#include "giflib.h"
#include "gif_lib.h"
#include <stdbool.h>

struct giflib_decoder_struct {
    GifFileType* gif;
    const cv::Mat* mat;
    ptrdiff_t read_index;
    GifByteType* pixels;
    size_t pixel_len;
    int prev_frame_disposal;
    int prev_frame_delay_time;
    int prev_frame_left;
    int prev_frame_top;
    int prev_frame_width;
    int prev_frame_height;
    uint8_t bg_green;
    uint8_t bg_red;
    uint8_t bg_blue;
    uint8_t bg_alpha;
    bool have_read_first_frame;
    bool seek_clear_extensions;
};

// this structure will help save us work of "reversing" a palette
// we will bit-crush a RGB value and use it to look up one of these
// entries, which if present, prevents us from searching for the
// nearest palette entry for that color
typedef struct {
    uint8_t index;
    uint8_t present;
} encoder_palette_lookup;

struct giflib_encoder_struct {
    GifFileType* gif;
    uint8_t* dst;
    size_t dst_len;
    ptrdiff_t dst_offset;

    // palette lookup is a computational-saving structure to convert
    // (reduced-depth) RGB values into the frame's 256-entry palette
    encoder_palette_lookup* palette_lookup;

    GifByteType* pixels;
    size_t pixel_len;

    ColorMapObject* frame_color_map;
    ColorMapObject* prev_frame_color_map;

    int prev_frame_disposal;

    uint8_t* prev_frame_bgra;

    bool have_written_first_frame;

    // keep track of all of the things we've allocated
    // we could technically just stuff all of these into a vector
    // of void*s but it might be interesting to build a pool
    // for these later, so it makes sense to keep them separated
    // n.b. that even if we do that, giflib still uses hella mallocs
    // when building the decoder, so it would only save us on the encoder
    std::vector<ExtensionBlock*> extension_blocks;
    std::vector<GifByteType*> gif_bytes;
    std::vector<ColorMapObject*> color_maps;
    std::vector<GifColorType*> colors;
    std::vector<SavedImage*> saved_images;
};

opencv_mat opencv_mat_create_from_data(int width, int height, int type, void* data, size_t data_len)
{
    size_t total_size = width * height * CV_ELEM_SIZE(type);
    if (total_size > data_len) {
        return NULL;
    }
    auto mat = new cv::Mat(height, width, type, data);
    mat->datalimit = (uint8_t*)data + data_len;
    return mat;
}

void opencv_mat_release(opencv_mat mat)
{
    auto m = static_cast<cv::Mat*>(mat);
    delete m;
}

int opencv_type_depth(int type)
{
    return CV_ELEM_SIZE1(type) * 8;
}

int opencv_type_convert_depth(int t, int depth)
{
    return CV_MAKETYPE(depth, CV_MAT_CN(t));
}

int decode_func(GifFileType* gif, GifByteType* buf, int len)
{
    auto d = static_cast<giflib_decoder>(gif->UserData);
    size_t buf_len = d->mat->total() - d->read_index;
    size_t read_len = (buf_len > len) ? len : buf_len;
    memmove(buf, d->mat->data + d->read_index, read_len);
    d->read_index += read_len;
    return read_len;
}

giflib_decoder giflib_decoder_create(const opencv_mat buf)
{
    giflib_decoder d = new struct giflib_decoder_struct();
    memset(d, 0, sizeof(struct giflib_decoder_struct));
    d->mat = static_cast<const cv::Mat*>(buf);

    int error = 0;
    GifFileType* gif = DGifOpen(d, decode_func, &error);
    if (error) {
        delete d;
        return NULL;
    }
    d->gif = gif;

    return d;
}

int giflib_decoder_get_width(const giflib_decoder d)
{
    return d->gif->SWidth;
}

int giflib_decoder_get_height(const giflib_decoder d)
{
    return d->gif->SHeight;
}

int giflib_decoder_get_num_frames(const giflib_decoder d)
{
    return d->gif->ImageCount;
}

int giflib_decoder_get_frame_width(const giflib_decoder d)
{
    return d->gif->Image.Width;
}

int giflib_decoder_get_frame_height(const giflib_decoder d)
{
    return d->gif->Image.Height;
}

int giflib_decoder_get_prev_frame_delay(const giflib_decoder d)
{
    return d->prev_frame_delay_time;
}

void giflib_decoder_release(giflib_decoder d)
{
    if (d->pixels) {
        free(d->pixels);
    }
    int error = 0;
    DGifCloseFile(d->gif, &error);
    delete d;
}

static bool giflib_decoder_read_extensions(giflib_decoder d)
{
    GifByteType* ExtData;
    int ExtFunction;

    if (DGifGetExtension(d->gif, &ExtFunction, &ExtData) == GIF_ERROR) {
        return false;
    }

    // XXX filter out everything but GRAPHICS_EXT_FUNC_CODE
    if (ExtData != NULL) {
        int res = GifAddExtensionBlock(&d->gif->ExtensionBlockCount,
                                       &d->gif->ExtensionBlocks,
                                       ExtFunction,
                                       ExtData[0],
                                       &ExtData[1]);
        if (res == GIF_ERROR) {
            return false;
        }
    }

    while (ExtData != NULL) {
        if (DGifGetExtensionNext(d->gif, &ExtData) == GIF_ERROR) {
            return false;
        }

        if (ExtData != NULL) {
            int res = GifAddExtensionBlock(&d->gif->ExtensionBlockCount,
                                           &d->gif->ExtensionBlocks,
                                           CONTINUE_EXT_FUNC_CODE,
                                           ExtData[0],
                                           &ExtData[1]);
            if (res == GIF_ERROR) {
                return false;
            }
        }
    }

    return true;
}

static bool giflib_get_frame_gcb(GifFileType* gif, GraphicsControlBlock* gcb)
{
    gcb->DisposalMode = DISPOSAL_UNSPECIFIED;
    gcb->UserInputFlag = false;
    gcb->DelayTime = 0;
    gcb->TransparentColor = NO_TRANSPARENT_COLOR;

    bool success = true;
    for (int i = 0; i < gif->ExtensionBlockCount; i++) {
        ExtensionBlock* b = &gif->ExtensionBlocks[i];
        if (b->Function == GRAPHICS_EXT_FUNC_CODE) {
            int res = DGifExtensionToGCB(b->ByteCount, b->Bytes, gcb);
            success = res == GIF_OK;
        }
    }

    return success;
}

static bool giflib_set_frame_gcb(GifFileType* gif, const GraphicsControlBlock* gcb)
{
    bool success = true;
    for (int i = 0; i < gif->ExtensionBlockCount; i++) {
        ExtensionBlock* b = &gif->ExtensionBlocks[i];
        if (b->Function == GRAPHICS_EXT_FUNC_CODE) {
            int res = EGifGCBToExtension(gcb, b->Bytes);
            success = res == GIF_OK;
        }
    }

    return success;
}

static giflib_decoder_frame_state giflib_decoder_seek_next_frame(giflib_decoder d)
{
    GifRecordType RecordType;

    if (d->seek_clear_extensions) {
        GifFreeExtensions(&d->gif->ExtensionBlockCount, &d->gif->ExtensionBlocks);
        d->seek_clear_extensions = false;
    }

    do {
        if (DGifGetRecordType(d->gif, &RecordType) == GIF_ERROR) {
            return giflib_decoder_error;
        }

        switch (RecordType) {
        case IMAGE_DESC_RECORD_TYPE:
            // we are now at the next frame, so quit
            return giflib_decoder_have_next_frame;

        case EXTENSION_RECORD_TYPE:
            if (!giflib_decoder_read_extensions(d)) {
                return giflib_decoder_error;
            }
            break;

        case TERMINATE_RECORD_TYPE:
            break;

        default:
            break;
        }
    } while (RecordType != TERMINATE_RECORD_TYPE);

    return giflib_decoder_eof;
}

// get just the header without attempting to read its pixel data
// this sets the image properties on d->gif->Image
// includes dimensions and frame origin coordinates, color map
giflib_decoder_frame_state giflib_decoder_decode_frame_header(giflib_decoder d)
{
    giflib_decoder_frame_state seek = giflib_decoder_seek_next_frame(d);

    if (seek == giflib_decoder_eof || seek == giflib_decoder_error) {
        return seek;
    }

    if (DGifGetImageHeader(d->gif) == GIF_ERROR) {
        return giflib_decoder_error;
    }

    return giflib_decoder_have_next_frame;
}

static bool giflib_decoder_render_frame(giflib_decoder d, GraphicsControlBlock* gcb, opencv_mat mat)
{
    auto cvMat = static_cast<cv::Mat*>(mat);
    GifImageDesc desc = d->gif->Image;
    int transparency_index = gcb->TransparentColor;

    if (desc.Width < 0) {
        fprintf(stderr, "encountered error, gif frame width less than 0\n");
        return false;
    }

    if (desc.Height < 0) {
        fprintf(stderr, "encountered error, gif frame height less than 0\n");
        return false;
    }

    int frame_left = desc.Left;
    int frame_top = desc.Top;
    int frame_width = desc.Width;
    int frame_height = desc.Height;

    int buf_width = cvMat->cols;
    int buf_height = cvMat->rows;

    // calculate the out-of-bounds skip lengths
    // for whatever reason, gifs allow frames to draw outside of the viewport
    // we can't just cap these values because we have to skip the (unrenderable!) raster bits
    int skip_left = (frame_left < 0) ? -frame_left : 0;
    int skip_top = (frame_top < 0) ? -frame_top : 0;
    int skip_right =
      (frame_left + frame_width > buf_width) ? (frame_left + frame_width - buf_width) : 0;
    int skip_bottom =
      (frame_top + frame_height > buf_height) ? (frame_top + frame_height - buf_height) : 0;

    ColorMapObject* globalColorMap = d->gif->SColorMap;
    ColorMapObject* frameColorMap = desc.ColorMap;
    ColorMapObject* colorMap = frameColorMap ? frameColorMap : globalColorMap;

    if (!colorMap) {
        fprintf(stderr, "encountered error, gif frame has no color map\n");
        return false;
    }

    if (!d->have_read_first_frame) {
        // first frame -- draw the background
        for (size_t y = 0; y < buf_height; y++) {
            uint8_t* dst = cvMat->data + y * cvMat->step;
            for (size_t x = 0; x < buf_width; x++) {
                *dst++ = d->bg_blue;
                *dst++ = d->bg_green;
                *dst++ = d->bg_red;
                *dst++ = d->bg_alpha;
            }
        }
    }

    if (d->have_read_first_frame) {
        if (d->prev_frame_disposal == DISPOSE_BACKGROUND) {
            // draw over the previous frame with the BG color
            int prev_frame_left = d->prev_frame_left;
            int prev_frame_top = d->prev_frame_top;
            int prev_frame_width = d->prev_frame_width;
            int prev_frame_height = d->prev_frame_height;

            if (prev_frame_left < 0) {
                // "subtract" the width that hangs off the left edge
                prev_frame_width += prev_frame_left;
                prev_frame_left = 0;
            }

            if (prev_frame_top < 0) {
                // do same subtracting for height off top edge
                prev_frame_height += prev_frame_top;
                prev_frame_top = 0;
            }

            if (prev_frame_left + prev_frame_width > buf_width) {
                // cap width to keep frame within right edge
                prev_frame_width = buf_width - prev_frame_left;
            }

            if (prev_frame_top + prev_frame_height > buf_height) {
                // do same cap to keep frame within bottom edge
                prev_frame_height = buf_height - prev_frame_top;
            }

            // if either of these is true, we'll just do nothing in the loop
            // we could bail out of here somehow, seems easiest to do it this way
            prev_frame_height = (prev_frame_height < 0) ? 0 : prev_frame_height;
            prev_frame_width = (prev_frame_width < 0) ? 0 : prev_frame_width;

            for (int y = prev_frame_top; y < prev_frame_top + prev_frame_height; y++) {
                uint8_t* dst = cvMat->data + y * cvMat->step + (prev_frame_left * 4);
                for (int x = prev_frame_left; x < prev_frame_left + prev_frame_width; x++) {
                    *dst++ = d->bg_blue;
                    *dst++ = d->bg_green;
                    *dst++ = d->bg_red;
                    *dst++ = d->bg_alpha;
                }
            }
        }
        else if (d->prev_frame_disposal == DISPOSE_PREVIOUS) {
            // TODO or maybe not to do
            // should we at least log this happened so that we know this exists?
            // tldr this crazy method requires you to walk back across all previous
            //    frames until you reach one with DISPOSAL_DO_NOT
            //    and "undraw them", most likely would be done by building a temp
            //    buffer when first one is encountered
        }
    }

    // TODO handle interlaced gifs?

    // TODO if top > 0 or left > 0, we could actually just return an ROI
    // of the pixel buffer and then resize just the ROI frame
    // we would then have to rescale the origin coordinates of that frame
    // when encoding back to gif, so that the resized frame is drawn to the
    // correct location
    int pixel_index = 0;

    // skip entire rows at the top if frame_top < 0
    // start by skipping the raster bits -- we're skipping full rows here
    pixel_index += (skip_top * desc.Width);
    // now reduce how far we iterate by subtracting how many rows we skipped
    // if we were supposed to start at y = -2 and go for 5 rows, then instead
    // start at y = 0 and go for 3 rows
    frame_height -= skip_top;
    // move the top of the frame over by how far we skipped
    frame_top += skip_top;

    // do similar thing for left-side skip as with top-side
    // here we only skip by some columns, not an entire row
    frame_width -= skip_left;
    frame_left += skip_left;

    // right-side skip requires shortening the loop iter and moving raster pointer
    // here we just shorten loop, move raster at bottom of row loop
    frame_width -= skip_right;

    // bottom skip is simple, we just reduce # of rows we do
    frame_height -= skip_bottom;

    for (int y = frame_top; y < frame_top + frame_height; y++) {
        // draw a single row of pixels in this iteration

        // do actual column skipping here
        pixel_index += skip_left;

        uint8_t* dst = cvMat->data + y * cvMat->step + (frame_left * 4);
        for (int x = frame_left; x < frame_left + frame_width; x++) {
            // draw a single pixel in this iteration
            GifByteType palette_index = d->pixels[pixel_index++];
            if (palette_index == transparency_index) {
                // TODO: don't hardcode 4 channels (8UC4) here
                dst += 4;
                continue;
            }
            *dst++ = colorMap->Colors[palette_index].Blue;
            *dst++ = colorMap->Colors[palette_index].Green;
            *dst++ = colorMap->Colors[palette_index].Red;
            *dst++ = 255;
        }

        pixel_index += skip_right;
    }

    // because we turn partial frames into full frames, we need to ensure that a transparency color
    // is defined, so that the encoder can use it (we convert partial frames to full frames with
    // a lot of transparency)

    // let's check if we have a partial frame and whether no transparency is defined
    bool have_partial_frame = false;
    have_partial_frame |= frame_height < buf_height;
    have_partial_frame |= frame_width < buf_width;
    have_partial_frame |= frame_left != 0;
    have_partial_frame |= frame_top != 0;

    if (have_partial_frame && transparency_index == -1) {
        // make sure our pseudo partial frame impl can use transparency
        // if no color is set, force the palette to have one
        // evict the last color and make it transparency instead
        gcb->TransparentColor = colorMap->ColorCount - 1;
        giflib_set_frame_gcb(d->gif, gcb);
    }

    return true;
}

giflib_decoder_frame_state giflib_decoder_skip_frame(giflib_decoder d)
{
    giflib_decoder_frame_state seek = giflib_decoder_decode_frame_header(d);

    if (seek != giflib_decoder_have_next_frame) {
        return seek;
    }

    GifByteType* block;
    while (true) {
        if (DGifGetCodeNext(d->gif, &block) == GIF_ERROR) {
            return giflib_decoder_error;
        }

        if (block == NULL) {
            break;
        }
    }

    return giflib_decoder_have_next_frame;
}

static int interlace_offset[] = {0, 4, 2, 1};
static int interlace_jumps[] = {8, 8, 4, 2};

// decode the full frame and write it into mat
// decode_frame_header *must* be called before this function
bool giflib_decoder_decode_frame(giflib_decoder d, opencv_mat mat)
{
    GifImageDesc desc = d->gif->Image;

    if (desc.Width <= 0) {
        fprintf(stderr, "encountered error, gif frame has negative or zero width\n");
        return false;
    }

    if (desc.Height <= 0) {
        fprintf(stderr, "encountered error, gif frame has negative or zero height\n");
        return false;
    }

    if (desc.Width > (INT_MAX / desc.Height)) {
        fprintf(stderr, "encountered error, gif frame is too wide\n");
        return false;
    }

    // since we aren't actually writing into mat, we don't check for large
    // dimensions here. it is up to the caller to do that after reading the
    // header

    size_t image_size = desc.Width * desc.Height;

    if (image_size > (SIZE_MAX / sizeof(GifPixelType))) {
        fprintf(stderr, "encountered error, gif frame is too large\n");
        return false;
    }

    if (image_size > d->pixel_len) {
        // only realloc if we need to size up
        // no point in shrinking, we'll free when decode has finished
        d->pixel_len = image_size;
        d->pixels = (GifByteType*)(realloc(d->pixels, d->pixel_len * sizeof(GifPixelType)));
    }

    if (d->pixels == NULL) {
        fprintf(stderr, "encountered error, gif pixel buffer failed to allocate\n");
        return false;
    }

    if (desc.Interlace) {
        for (int i = 0; i < sizeof(interlace_offset) / sizeof(int); i++) {
            for (int j = interlace_offset[i]; j < desc.Height; j += interlace_jumps[i]) {
                int res = DGifGetLine(d->gif, d->pixels + j * desc.Width, desc.Width);
                if (res == GIF_ERROR) {
                    fprintf(stderr, "encountered error, could not rasterize gif line\n");
                    return false;
                }
            }
        }
    }
    else {
        int res = DGifGetLine(d->gif, d->pixels, image_size);
        if (res == GIF_ERROR) {
            fprintf(stderr, "encountered error, could not rasterize gif\n");
            return false;
        }
    }

    GraphicsControlBlock gcb;
    giflib_get_frame_gcb(d->gif, &gcb);

    if (!d->have_read_first_frame) {
        bool have_transparency = (gcb.TransparentColor != NO_TRANSPARENT_COLOR);
        if (have_transparency) {
            d->bg_red = d->bg_green = d->bg_blue = d->bg_alpha = 0;
        }
        else if (d->gif->SColorMap && d->gif->SColorMap->Colors) {
            d->bg_red = d->gif->SColorMap->Colors[d->gif->SBackGroundColor].Red;
            d->bg_green = d->gif->SColorMap->Colors[d->gif->SBackGroundColor].Green;
            d->bg_blue = d->gif->SColorMap->Colors[d->gif->SBackGroundColor].Blue;
            d->bg_alpha = 255;
        }
        else {
            d->bg_red = d->bg_green = d->bg_blue = d->bg_alpha = 255;
        }
    }

    if (!giflib_decoder_render_frame(d, &gcb, mat)) {
        return false;
    }

    d->prev_frame_disposal = gcb.DisposalMode;
    d->prev_frame_delay_time = gcb.DelayTime;
    d->prev_frame_left = d->gif->Image.Left;
    d->prev_frame_top = d->gif->Image.Top;
    d->prev_frame_width = d->gif->Image.Width;
    d->prev_frame_height = d->gif->Image.Height;
    d->have_read_first_frame = true;
    d->seek_clear_extensions = true;

    return true;
}

ExtensionBlock* giflib_encoder_allocate_extension_blocks(giflib_encoder e, size_t count)
{
    ExtensionBlock* blocks = (ExtensionBlock*)(malloc(count * sizeof(ExtensionBlock)));
    e->extension_blocks.push_back(blocks);
    return blocks;
}

GifByteType* giflib_encoder_allocate_gif_bytes(giflib_encoder e, size_t count)
{
    GifByteType* bytes = (GifByteType*)(malloc(count * sizeof(GifByteType)));
    e->gif_bytes.push_back(bytes);
    return bytes;
}

ColorMapObject* giflib_encoder_allocate_color_maps(giflib_encoder e, size_t count)
{
    ColorMapObject* color_maps = (ColorMapObject*)(malloc(count * sizeof(ColorMapObject)));
    e->color_maps.push_back(color_maps);
    return color_maps;
}

GifColorType* giflib_encoder_allocate_colors(giflib_encoder e, size_t count)
{
    GifColorType* colors = (GifColorType*)(malloc(count * sizeof(GifColorType)));
    e->colors.push_back(colors);
    return colors;
}

SavedImage* giflib_encoder_allocate_saved_images(giflib_encoder e, size_t count)
{
    SavedImage* saved_images = (SavedImage*)(malloc(count * sizeof(SavedImage)));
    e->saved_images.push_back(saved_images);
    return saved_images;
}

int encode_func(GifFileType* gif, const GifByteType* buf, int len)
{
    giflib_encoder e = static_cast<giflib_encoder>(gif->UserData);
    if (e->dst_offset + len > e->dst_len) {
        return 0;
    }
    memcpy(e->dst + e->dst_offset, &buf[0], len);
    e->dst_offset += len;
    return len;
}

giflib_encoder giflib_encoder_create(void* buf, size_t buf_len)
{
    giflib_encoder e = new struct giflib_encoder_struct();
    memset(e, 0, sizeof(struct giflib_encoder_struct));
    e->dst = (uint8_t*)(buf);
    e->dst_len = buf_len;

    int error = 0;
    GifFileType* gif_out = EGifOpen(e, encode_func, &error);
    if (error) {
        fprintf(stderr, "encountered error opening gif, %d\n", error);
        delete e;
        return NULL;
    }
    e->gif = gif_out;

    // set up palette lookup table. we need 2^15 entries because we will be
    // using bit-crushed RGB values, 5 bits each. this is a reasonable compromise
    // between fidelity and computation/storage
    e->palette_lookup =
      (encoder_palette_lookup*)(malloc((1 << 15) * sizeof(encoder_palette_lookup)));

    return e;
}

// this function should be called just once when we know the global dimensions
bool giflib_encoder_init(giflib_encoder e, const giflib_decoder d, int width, int height)
{
    // all gifs will output as gif89
    EGifSetGifVersion(e->gif, true);
    e->gif->SWidth = width;
    e->gif->SHeight = height;

    e->prev_frame_bgra = (uint8_t*)(malloc(width * height * 4));

    // preserve # of palette entries and aspect ratio of original gif
    e->gif->SColorResolution = d->gif->SColorResolution;
    e->gif->AspectByte = d->gif->AspectByte;

    // copy global color palette, if any
    if (d->gif->SColorMap) {
        e->gif->SColorMap = giflib_encoder_allocate_color_maps(e, 1);
        memmove(e->gif->SColorMap, d->gif->SColorMap, sizeof(ColorMapObject));
        e->gif->SColorMap->Colors =
          giflib_encoder_allocate_colors(e, e->gif->SColorMap->ColorCount);
        memmove(e->gif->SColorMap->Colors,
                d->gif->SColorMap->Colors,
                e->gif->SColorMap->ColorCount * sizeof(GifColorType));
    }

    int res = EGifPutScreenDesc(e->gif,
                                e->gif->SWidth,
                                e->gif->SHeight,
                                e->gif->SColorResolution,
                                e->gif->SBackGroundColor,
                                e->gif->SColorMap);
    if (res == GIF_ERROR) {
        return false;
    }

    return true;
}

static bool giflib_encoder_setup_frame(giflib_encoder e, const giflib_decoder d)
{
    // initialize frame with input gif's frame metadata
    // this includes, amongst other things, inter-frame delays
    GifImageDesc* im_in = &d->gif->Image;
    GifImageDesc* im_out = &e->gif->Image;

    // XXX we're just going to copy here, but this probably isn't right since
    // the decoder doesn't handle interlacing correctly. might be worthwhile to
    // just set this to false always (or enhance the decoder)
    im_out->Interlace = im_in->Interlace;

    // prepare frame local palette, if any
    e->frame_color_map = NULL;
    if (im_in->ColorMap) {
        e->frame_color_map = giflib_encoder_allocate_color_maps(e, 1);
        memmove(e->frame_color_map, im_in->ColorMap, sizeof(ColorMapObject));
        // copy all of the RGB color values from input frame palette to output frame palette
        e->frame_color_map->Colors =
          giflib_encoder_allocate_colors(e, e->frame_color_map->ColorCount);
        memmove(e->frame_color_map->Colors,
                im_in->ColorMap->Colors,
                e->frame_color_map->ColorCount * sizeof(GifColorType));
    }

    // copy extension blocks specific to this frame
    // this sets up the frame delay as well as which palette entry is transparent, if any
    e->gif->ExtensionBlockCount = d->gif->ExtensionBlockCount;
    e->gif->ExtensionBlocks = NULL;
    if (e->gif->ExtensionBlockCount > 0) {
        // TODO here and in global extension blocks, we should filter out worthless blocks
        // we're only really interested in ExtensionBlock.Function = GRAPHICS_EXT_FUNC_CODE
        // other values like COMMENT_ and PLAINTEXT_ are not essential to viewing the image
        e->gif->ExtensionBlocks =
          giflib_encoder_allocate_extension_blocks(e, e->gif->ExtensionBlockCount);
        for (int i = 0; i < e->gif->ExtensionBlockCount; i++) {
            ExtensionBlock* eb_in = &(d->gif->ExtensionBlocks[i]);
            ExtensionBlock* eb_out = &(e->gif->ExtensionBlocks[i]);
            eb_out->ByteCount = eb_in->ByteCount;
            eb_out->Function = eb_in->Function;
            eb_out->Bytes = giflib_encoder_allocate_gif_bytes(e, eb_out->ByteCount);
            memmove(eb_out->Bytes, eb_in->Bytes, eb_out->ByteCount);
        }
    }

    return true;
}

// TODO this probably should be the euclidean distance
// the manhattan distance will still be "good enough"
// euclidean requires calculating pow(2) and sqrt()?
static inline int rgb_distance(int r0, int g0, int b0, int r1, int g1, int b1)
{
    int dist = 0;
    dist += (r0 > r1) ? r0 - r1 : r1 - r0;
    dist += (g0 > g1) ? g0 - g1 : g1 - g0;
    dist += (b0 > b1) ? b0 - b1 : b1 - b0;
    return dist;
}

static bool giflib_encoder_render_frame(giflib_encoder e,
                                        const giflib_decoder d,
                                        const opencv_mat opaque_frame)
{
    GifFileType* gif_out = e->gif;
    auto frame = static_cast<const cv::Mat*>(opaque_frame);

    // basic bounds checking - would this frame be wider than the global gif width?
    // if we do partial frames, we'll need to change this to account for top/left
    if (frame->cols > gif_out->SWidth) {
        fprintf(stderr, "encountered error, gif frame wider than gif global width\n");
        return false;
    }

    if (frame->rows > gif_out->SHeight) {
        fprintf(stderr, "encountered error, gif frame taller than gif global height\n");
        return false;
    }

    GifImageDesc* im_out = &gif_out->Image;
    // TODO some day consider making partial frames/make these not 0
    im_out->Left = 0;
    im_out->Top = 0;
    im_out->Width = frame->cols;
    im_out->Height = frame->rows;

    int image_size = im_out->Width * im_out->Height;

    if (image_size > e->pixel_len) {
        // only realloc if we need to size up
        e->pixel_len = image_size;
        e->pixels = (GifByteType*)(realloc(e->pixels, e->pixel_len * sizeof(GifPixelType)));
    }

    ColorMapObject* global_color_map = e->gif->SColorMap;
    ColorMapObject* frame_color_map = e->frame_color_map;
    ColorMapObject* color_map = frame_color_map ? frame_color_map : global_color_map;

    if (!color_map) {
        fprintf(stderr, "encountered error, gif frame has no color map\n");
        return false;
    }

    // prepare our palette lookup table. if we used the same (byte-equal) palette table last
    // frame, we can just reuse it this frame. otherwise we need to clear the lookup out
    bool clear_palette_lookup = true;
    // on the first frame, we will always clear
    if (e->have_written_first_frame) {
        ColorMapObject* last_color_map = e->prev_frame_color_map;
        if (last_color_map && last_color_map->ColorCount == color_map->ColorCount) {
            int cmp = memcmp(last_color_map->Colors,
                             color_map->Colors,
                             color_map->ColorCount * sizeof(GifColorType));
            clear_palette_lookup = (cmp != 0);
        }
    }

    if (clear_palette_lookup) {
        memset(e->palette_lookup, 0, (1 << 15) * sizeof(encoder_palette_lookup));
    }

    GraphicsControlBlock gcb;
    giflib_get_frame_gcb(e->gif, &gcb);
    int transparency_index = gcb.TransparentColor;
    bool have_transparency = (transparency_index != NO_TRANSPARENT_COLOR);

    // decide whether we can use transparency against the previous frame
    bool prev_frame_valid = e->have_written_first_frame &&
      (e->prev_frame_disposal == DISPOSAL_UNSPECIFIED || e->prev_frame_disposal == DISPOSE_DO_NOT);

    // convenience names for these dimensions
    int frame_left = im_out->Left;
    int frame_top = im_out->Top;
    int frame_width = im_out->Width;
    int frame_height = im_out->Height;

    GifByteType* raster_out = e->pixels;

    int raster_index = 0;
    for (int y = frame_top; y < frame_top + frame_height; y++) {
        uint8_t* src = frame->data + y * frame->step + (frame_left * 4);
        for (int x = frame_left; x < frame_left + frame_width; x++) {
            uint32_t B = *src++;
            uint32_t G = *src++;
            uint32_t R = *src++;
            uint32_t A = *src++;

            // TODO come up with what this threshold value should be
            // probably ought to be a lot smaller, but greater than 0
            // for now we just pick halfway
            if (A < 128 && have_transparency) {
                // this composite frame pixel is actually transparent
                // what this means is that the background color must be transparent
                // AND this frame pixel must be transparent
                // for now we'll just assume bg is transparent since otherwise decoder
                // could not have generated this frame pixel with a low opacity
                *raster_out++ = transparency_index;
                continue;
            }

            uint32_t crushed = ((R >> 3) << 10) | ((G >> 3) << 5) | ((B >> 3));
            int least_dist = INT_MAX;
            int best_color = 0;
            if (!(e->palette_lookup[crushed].present)) {
                // calculate the best palette entry based on the midpoint of the crushed colors
                // what this means is that we drop the crushed bits (& 0xf8)
                // and then OR the highest-order crushed bit back in, which is approx midpoint
                uint32_t R_center = (R & 0xf8) | 4;
                uint32_t G_center = (G & 0xf8) | 4;
                uint32_t B_center = (B & 0xf8) | 4;

                // we're calculating the best, so keep track of which palette entry has least
                // distance
                int count = color_map->ColorCount;
                for (int i = 0; i < count; i++) {
                    if (i == transparency_index) {
                        // this index doesn't point to an actual color
                        continue;
                    }
                    int dist = rgb_distance(R_center,
                                            G_center,
                                            B_center,
                                            color_map->Colors[i].Red,
                                            color_map->Colors[i].Green,
                                            color_map->Colors[i].Blue);
                    if (dist < least_dist) {
                        least_dist = dist;
                        best_color = i;
                    }
                }
                e->palette_lookup[crushed].present = 1;
                e->palette_lookup[crushed].index = best_color;
            }
            else {
                best_color = e->palette_lookup[crushed].index;
                least_dist = rgb_distance(R,
                                          G,
                                          B,
                                          color_map->Colors[best_color].Red,
                                          color_map->Colors[best_color].Green,
                                          color_map->Colors[best_color].Blue);
            }

            // now that we for sure know which palette entry to pick, we have one more test
            // to perform. it's possible that the best color for this pixel is actually
            // the color of this pixel in the previous frame. if that's true, we'll just
            // choose the transparency color, which will compress better on average
            // (plus it improves color range of image)
            if (prev_frame_valid && have_transparency) {
                ptrdiff_t frame_index = 4 * ((y * e->gif->SWidth) + x);
                uint32_t last_B = e->prev_frame_bgra[frame_index];
                uint32_t last_G = e->prev_frame_bgra[frame_index + 1];
                uint32_t last_R = e->prev_frame_bgra[frame_index + 2];
                int dist = rgb_distance(R, G, B, last_R, last_G, last_B);
                if (dist < least_dist) {
                    least_dist = dist;
                    best_color = transparency_index;
                }
            }

            *raster_out++ = best_color;
        }
    }

    // XXX change this if we do partial frames (only copy over some)
    memcpy(e->prev_frame_bgra, frame->data, 4 * e->gif->SWidth * e->gif->SHeight);

    e->prev_frame_color_map = color_map;
    e->prev_frame_disposal = gcb.DisposalMode;

    return true;
}

static int giflib_encoder_write_extensions(giflib_encoder e)
{
    if (e->gif->ExtensionBlocks) {
        ExtensionBlock* ep;

        for (int i = 0; i < e->gif->ExtensionBlockCount; i++) {
            ep = &e->gif->ExtensionBlocks[i];
            if (ep->Function != CONTINUE_EXT_FUNC_CODE) {
                if (EGifPutExtensionLeader(e->gif, ep->Function) == GIF_ERROR) {
                    return false;
                }
            }
            if (EGifPutExtensionBlock(e->gif, ep->ByteCount, ep->Bytes) == GIF_ERROR) {
                return false;
            }
            if (i == e->gif->ExtensionBlockCount - 1 ||
                (ep + 1)->Function != CONTINUE_EXT_FUNC_CODE) {
                if (EGifPutExtensionTrailer(e->gif) == GIF_ERROR) {
                    return false;
                }
            }
        }
    }

    return true;
}

bool giflib_encoder_encode_frame(giflib_encoder e,
                                 const giflib_decoder d,
                                 const opencv_mat opaque_frame)
{
    giflib_encoder_setup_frame(e, d);
    giflib_encoder_render_frame(e, d, opaque_frame);

    GifImageDesc* im_out = &e->gif->Image;
    int frame_height = im_out->Height;
    int frame_width = im_out->Width;

    int res = giflib_encoder_write_extensions(e);
    if (res == GIF_ERROR) {
        return false;
    }

    res = EGifPutImageDesc(e->gif,
                           im_out->Left,
                           im_out->Top,
                           im_out->Width,
                           im_out->Height,
                           im_out->Interlace,
                           e->frame_color_map);
    if (res == GIF_ERROR) {
        return false;
    }

    if (im_out->Interlace) {
        /* Need to perform 4 passes on the images: */
        for (int i = 0; i < 4; i++) {
            for (int j = interlace_offset[i]; j < frame_height; j += interlace_jumps[i]) {
                res = EGifPutLine(e->gif, e->pixels + j * frame_width, frame_width);
                if (res == GIF_ERROR) {
                    fprintf(stderr, "encountered error, could not serialize gif line\n");
                    return false;
                }
            }
        }
    }
    else {
        for (int i = 0; i < frame_height; i++) {
            res = EGifPutLine(e->gif, e->pixels + i * frame_width, frame_width);
            if (res == GIF_ERROR) {
                return false;
            }
        }
    }

    e->have_written_first_frame = true;

    return true;
}

bool giflib_encoder_flush(giflib_encoder e, const giflib_decoder d)
{
    // XXX we need to pull these trailing blocks on d
    // does decoder's state machine allow that?

    // set up "trailing" extension blocks, which appear after all the frames
    // brian note: what do these do? do we actually need them?
    e->gif->ExtensionBlockCount = d->gif->ExtensionBlockCount;
    e->gif->ExtensionBlocks = NULL;
    if (e->gif->ExtensionBlockCount > 0) {
        e->gif->ExtensionBlocks =
          giflib_encoder_allocate_extension_blocks(e, e->gif->ExtensionBlockCount);
        for (int i = 0; i < e->gif->ExtensionBlockCount; i++) {
            ExtensionBlock* eb = &(e->gif->ExtensionBlocks[i]);
            eb->ByteCount = d->gif->ExtensionBlocks[i].ByteCount;
            eb->Function = d->gif->ExtensionBlocks[i].Function;
            eb->Bytes = giflib_encoder_allocate_gif_bytes(e, eb->ByteCount);
            memmove(eb->Bytes, d->gif->ExtensionBlocks[i].Bytes, eb->ByteCount);
        }
    }

    int res = giflib_encoder_write_extensions(e);
    if (res == GIF_ERROR) {
        return false;
    }

    if (EGifCloseFile(e->gif, NULL) == GIF_ERROR) {
        return false;
    }

    e->gif = NULL;

    return true;
}

void giflib_encoder_release(giflib_encoder e)
{
    // don't free dst -- we're borrowing it

    if (e->prev_frame_bgra) {
        free(e->prev_frame_bgra);
    }

    if (e->palette_lookup) {
        free(e->palette_lookup);
    }

    if (e->pixels) {
        free(e->pixels);
    }

    for (std::vector<ExtensionBlock*>::iterator it = e->extension_blocks.begin();
         it != e->extension_blocks.end();
         ++it) {
        free(*it);
    }
    e->extension_blocks.clear();

    for (std::vector<GifByteType*>::iterator it = e->gif_bytes.begin(); it != e->gif_bytes.end();
         ++it) {
        free(*it);
    }
    e->gif_bytes.clear();

    ColorMapObject* gif_scolor = NULL;
    ColorMapObject* gif_last_color = NULL;
    if (e->gif) {
        gif_scolor = e->gif->SColorMap;
        gif_last_color = e->gif->Image.ColorMap;
    }
    for (std::vector<ColorMapObject*>::iterator it = e->color_maps.begin();
         it != e->color_maps.end();
         ++it) {
        if (gif_scolor && gif_scolor == *it) {
            // this is extremely unlikely to happen, but giflib transitions from
            // borrowing this ptr to owning it, and it's possible that it will try
            // to free ours in certain circumstances
            // so swap its ptr if it matches one we own
            e->gif->SColorMap = NULL;
        }
        if (gif_last_color && gif_last_color == *it) {
            // this seemingly will never happen, but given how strange last case is,
            // check for it anyway
            e->gif->Image.ColorMap = NULL;
        }
        free(*it);
    }
    e->color_maps.clear();

    for (std::vector<GifColorType*>::iterator it = e->colors.begin(); it != e->colors.end(); ++it) {
        free(*it);
    }
    e->colors.clear();

    for (std::vector<SavedImage*>::iterator it = e->saved_images.begin();
         it != e->saved_images.end();
         ++it) {
        free(*it);
    }
    e->saved_images.clear();

    if (e->gif) {
        // we most likely won't actually call this since Spew() does it
        // but in exceptional cases we'll need it for cleanup
        int error_code = 0;
        EGifCloseFile(e->gif, &error_code);
        if (error_code) {
            fprintf(stderr, "encountered error closing gif, %d\n", error_code);
        }
    }

    delete e;
}

int giflib_encoder_get_output_length(giflib_encoder e)
{
    return e->dst_offset;
}
