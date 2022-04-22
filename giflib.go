package gocv

/*
#include <stddef.h>
#include "giflib.h"
*/
import "C"

import (
	"bytes"
	"errors"
	"io"
	"sync/atomic"
	"time"
	"unsafe"
)

var (
	ErrInvalidImage     = errors.New("unrecognized image format")
	ErrDecodingFailed   = errors.New("failed to decode image")
	ErrBufTooSmall      = errors.New("buffer too small to hold image")
	ErrFrameBufNoPixels = errors.New("Framebuffer contains no pixels")
	ErrSkipNotSupported = errors.New("skip operation not supported by this decoder")

	gif87Magic   = []byte("GIF87a")
	gif89Magic   = []byte("GIF89a")
	mp42Magic    = []byte("ftypmp42")
	mp4IsomMagic = []byte("ftypisom")
	pngMagic     = []byte{0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a}
)

type PixelType int

func (p PixelType) Depth() int {
	return int(C.opencv_type_depth(C.int(p)))
}

const (
	// Not available since we don't have defined value from C
	// OrientationTopLeft     = ImageOrientation(C.CV_IMAGE_ORIENTATION_TL)
	// OrientationTopRight    = ImageOrientation(C.CV_IMAGE_ORIENTATION_TR)
	// OrientationBottomRight = ImageOrientation(C.CV_IMAGE_ORIENTATION_BR)
	// OrientationBottomLeft  = ImageOrientation(C.CV_IMAGE_ORIENTATION_BL)
	// OrientationLeftTop     = ImageOrientation(C.CV_IMAGE_ORIENTATION_LT)
	// OrientationRightTop    = ImageOrientation(C.CV_IMAGE_ORIENTATION_RT)
	// OrientationRightBottom = ImageOrientation(C.CV_IMAGE_ORIENTATION_RB)
	// OrientationLeftBottom  = ImageOrientation(C.CV_IMAGE_ORIENTATION_LB)

	pngChunkSizeFieldLen = 4
	pngChunkTypeFieldLen = 4
	pngChunkAllFieldsLen = 12
)

var (
	pngActlChunkType = []byte{0x61, 0x63, 0x54, 0x4c}
	pngFctlChunkType = []byte{0x66, 0x63, 0x54, 0x4c}
	pngFdatChunkType = []byte{0x66, 0x64, 0x41, 0x54}
)

type ImageOrientation int

type ImageHeader struct {
	width     int
	height    int
	pixelType PixelType
	// orientation ImageOrientation
	numFrames int
}

func (h *ImageHeader) Width() int {
	return h.width
}

// Height returns the height of the image in number of pixels.
func (h *ImageHeader) Height() int {
	return h.height
}

// PixelType returns a PixelType describing the image's pixels.
func (h *ImageHeader) PixelType() PixelType {
	return h.pixelType
}

// Framebuffer contains an array of raw, decoded pixel data.
type Framebuffer struct {
	buf       []byte
	mat       C.opencv_mat
	width     int
	height    int
	pixelType PixelType
	duration  time.Duration
}

// Width returns the width of the contained pixel data in number of pixels. This may
// differ from the capacity of the framebuffer.
func (f *Framebuffer) Width() int {
	return f.width
}

// Height returns the height of the contained pixel data in number of pixels. This may
// differ from the capacity of the framebuffer.
func (f *Framebuffer) Height() int {
	return f.height
}

func (f *Framebuffer) resizeMat(width, height int, pixelType PixelType) error {
	if f.mat != nil {
		C.opencv_mat_release(f.mat)
		f.mat = nil
	}
	if pixelType.Depth() > 8 {
		pixelType = PixelType(C.opencv_type_convert_depth(C.int(pixelType), C.int(MatTypeCV8U)))
	}
	newMat := C.opencv_mat_create_from_data(C.int(width), C.int(height), C.int(pixelType), unsafe.Pointer(&f.buf[0]), C.size_t(len(f.buf)))
	if newMat == nil {
		return ErrBufTooSmall
	}
	f.mat = newMat
	f.width = width
	f.height = height
	f.pixelType = pixelType
	return nil
}

type Decoder interface {
	// Header returns basic image metadata from the image.
	// This is done lazily, reading only the first part of the image and not
	// a full decode.
	Header() (*ImageHeader, error)

	// Close releases any resources associated with the Decoder
	Close()

	// Description returns a string description of the image type, such as
	// "PNG"
	Description() string

	// Duration returns the duration of the content. This property is 0 for
	// static images and animated GIFs.
	Duration() time.Duration

	// DecodeTo fully decodes the image pixel data into f. Generally users should
	// prefer instead using the ImageOps object to decode images.
	DecodeTo(f *Framebuffer) error

	// SkipFrame skips a frame if the decoder supports multiple frames
	// and returns io.EOF if the last frame has been reached
	SkipFrame() error
}

// An Encoder compresses raw pixel data into a well-known image type.
type Encoder interface {
	// Encode encodes the pixel data in f into the dst provided to NewEncoder. Encode quality
	// options can be passed into opt, such as map[int]int{lilliput.JpegQuality: 80}
	Encode(f *Framebuffer, opt map[int]int) ([]byte, error)

	// Close releases any resources associated with the Encoder
	Close()
}

func isGIF(maybeGIF []byte) bool {
	return bytes.HasPrefix(maybeGIF, gif87Magic) || bytes.HasPrefix(maybeGIF, gif89Magic)
}

func isMP4(maybeMP4 []byte) bool {
	if len(maybeMP4) < 12 {
		return false
	}

	magic := maybeMP4[4:]
	return bytes.HasPrefix(magic, mp42Magic) || bytes.HasPrefix(magic, mp4IsomMagic)
}

// NewDecoder returns a Decoder which can be used to decode
// image data provided in buf. If the first few bytes of buf do not
// point to a valid magic string, an error will be returned.
func NewDecoder(buf []byte) (Decoder, error) {
	// Check buffer length before accessing it
	if len(buf) == 0 {
		return nil, ErrInvalidImage
	}

	return newGifDecoder(buf)
}

// NewEncoder returns an Encode which can be used to encode Framebuffer
// into compressed image data. ext should be a string like ".jpeg" or
// ".png". decodedBy is optional and can be the Decoder used to make
// the Framebuffer. dst is where an encoded image will be written.
func NewEncoder(ext string, decodedBy Decoder, dst []byte) (Encoder, error) {
	return newGifEncoder(decodedBy, dst)
}

type gifDecoder struct {
	decoder    C.giflib_decoder
	mat        C.opencv_mat
	buf        []byte
	frameIndex int
}

type gifEncoder struct {
	encoder    C.giflib_encoder
	decoder    C.giflib_decoder
	buf        []byte
	frameIndex int
	hasFlushed bool
}

const defaultMaxFrameDimension = 10000

var (
	gifMaxFrameDimension uint64

	ErrGifEncoderNeedsDecoder = errors.New("GIF encoder needs decoder used to create image")
)

// SetGIFMaxFrameDimension sets the largest GIF width/height that can be
// decoded
func SetGIFMaxFrameDimension(dim uint64) {
	// TODO we should investigate if this can be removed/become a mat check in decoder
	atomic.StoreUint64(&gifMaxFrameDimension, dim)
}

func newGifDecoder(buf []byte) (*gifDecoder, error) {
	mat := C.opencv_mat_create_from_data(C.int(len(buf)), 1, C.int(MatTypeCV8U), unsafe.Pointer(&buf[0]), C.size_t(len(buf)))

	if mat == nil {
		return nil, ErrBufTooSmall
	}

	decoder := C.giflib_decoder_create(mat)
	if decoder == nil {
		return nil, ErrInvalidImage
	}

	return &gifDecoder{
		decoder:    decoder,
		mat:        mat,
		buf:        buf,
		frameIndex: 0,
	}, nil
}

func (d *gifDecoder) Header() (*ImageHeader, error) {
	return &ImageHeader{
		width:     int(C.giflib_decoder_get_width(d.decoder)),
		height:    int(C.giflib_decoder_get_height(d.decoder)),
		pixelType: PixelType(MatTypeCV8UC4),
		// orientation: OrientationTopLeft,
		numFrames: int(C.giflib_decoder_get_num_frames(d.decoder)),
	}, nil
}

func (d *gifDecoder) FrameHeader() (*ImageHeader, error) {
	return &ImageHeader{
		width:     int(C.giflib_decoder_get_frame_width(d.decoder)),
		height:    int(C.giflib_decoder_get_frame_height(d.decoder)),
		pixelType: PixelType(MatTypeCV8U),
		// orientation: OrientationTopLeft,
		numFrames: 1,
	}, nil
}

func (d *gifDecoder) Close() {
	C.giflib_decoder_release(d.decoder)
	C.opencv_mat_release(d.mat)
	d.buf = nil
}

func (d *gifDecoder) Description() string {
	return "GIF"
}

func (d *gifDecoder) Duration() time.Duration {
	return time.Duration(0)
}

func (d *gifDecoder) DecodeTo(f *Framebuffer) error {
	h, err := d.Header()
	if err != nil {
		return err
	}

	err = f.resizeMat(h.Width(), h.Height(), h.PixelType())
	if err != nil {
		return err
	}

	nextFrameResult := int(C.giflib_decoder_decode_frame_header(d.decoder))
	if nextFrameResult == C.giflib_decoder_eof {
		return io.EOF
	}
	if nextFrameResult == C.giflib_decoder_error {
		return ErrInvalidImage
	}

	frameHeader, err := d.FrameHeader()
	if err != nil {
		return ErrInvalidImage
	}
	maxDim := int(atomic.LoadUint64(&gifMaxFrameDimension))
	if frameHeader.Width() > maxDim || frameHeader.Height() > maxDim {
		return ErrInvalidImage
	}

	ret := C.giflib_decoder_decode_frame(d.decoder, f.mat)
	if !ret {
		return ErrDecodingFailed
	}
	f.duration = time.Duration(C.giflib_decoder_get_prev_frame_delay(d.decoder)) * 10 * time.Millisecond
	d.frameIndex++
	return nil
}

func (d *gifDecoder) SkipFrame() error {
	nextFrameResult := int(C.giflib_decoder_skip_frame(d.decoder))

	if nextFrameResult == C.giflib_decoder_eof {
		return io.EOF
	}
	if nextFrameResult == C.giflib_decoder_error {
		return ErrInvalidImage
	}

	return nil
}

func newGifEncoder(decodedBy Decoder, buf []byte) (*gifEncoder, error) {
	// we must have a decoder since we can't build our own palettes
	// so if we don't get a gif decoder, bail out
	if decodedBy == nil {
		return nil, ErrGifEncoderNeedsDecoder
	}

	gifDecoder, ok := decodedBy.(*gifDecoder)
	if !ok {
		return nil, ErrGifEncoderNeedsDecoder
	}

	buf = buf[:1]
	enc := C.giflib_encoder_create(unsafe.Pointer(&buf[0]), C.size_t(cap(buf)))
	if enc == nil {
		return nil, ErrBufTooSmall
	}

	return &gifEncoder{
		encoder:    enc,
		decoder:    gifDecoder.decoder,
		buf:        buf,
		frameIndex: 0,
	}, nil
}

func (e *gifEncoder) Encode(f *Framebuffer, opt map[int]int) ([]byte, error) {
	if e.hasFlushed {
		return nil, io.EOF
	}

	if f == nil {
		ret := C.giflib_encoder_flush(e.encoder, e.decoder)
		if !ret {
			return nil, ErrInvalidImage
		}
		e.hasFlushed = true

		len := C.int(C.giflib_encoder_get_output_length(e.encoder))

		return e.buf[:len], nil
	}

	if e.frameIndex == 0 {
		// first run setup
		// TODO figure out actual gif width/height?
		C.giflib_encoder_init(e.encoder, e.decoder, C.int(f.Width()), C.int(f.Height()))
	}

	if !C.giflib_encoder_encode_frame(e.encoder, e.decoder, f.mat) {
		return nil, ErrInvalidImage
	}

	e.frameIndex++

	return nil, nil
}

func (e *gifEncoder) Close() {
	C.giflib_encoder_release(e.encoder)
}

func init() {
	SetGIFMaxFrameDimension(defaultMaxFrameDimension)
}
