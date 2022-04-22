#ifndef _OPENCV_GIFLIB_H_
#define _OPENCV_GIFLIB_H_

#include <stdbool.h>

#ifdef __cplusplus
#include <opencv2/opencv.hpp>

extern "C" {
#endif

typedef struct giflib_decoder_struct* giflib_decoder;
typedef struct giflib_encoder_struct* giflib_encoder;

typedef enum {
    giflib_decoder_have_next_frame,
    giflib_decoder_eof,
    giflib_decoder_error,
} giflib_decoder_frame_state;

typedef void* opencv_mat;
typedef void* opencv_decoder;
typedef void* opencv_encoder;

opencv_mat opencv_mat_create_from_data(int width, int height, int type, void* data, size_t data_len);
void opencv_mat_release(opencv_mat mat);
int opencv_type_depth(int type);
int opencv_type_convert_depth(int type, int depth);

giflib_decoder giflib_decoder_create(const opencv_mat buf);
int giflib_decoder_get_width(const giflib_decoder d);
int giflib_decoder_get_height(const giflib_decoder d);
int giflib_decoder_get_num_frames(const giflib_decoder d);
int giflib_decoder_get_frame_width(const giflib_decoder d);
int giflib_decoder_get_frame_height(const giflib_decoder d);
int giflib_decoder_get_prev_frame_delay(const giflib_decoder d);
void giflib_decoder_release(giflib_decoder d);
giflib_decoder_frame_state giflib_decoder_decode_frame_header(giflib_decoder d);
bool giflib_decoder_decode_frame(giflib_decoder d, opencv_mat mat);
giflib_decoder_frame_state giflib_decoder_skip_frame(giflib_decoder d);

giflib_encoder giflib_encoder_create(void* buf, size_t buf_len);
bool giflib_encoder_init(giflib_encoder e, const giflib_decoder d, int width, int height);
bool giflib_encoder_encode_frame(giflib_encoder e, const giflib_decoder d, const opencv_mat frame);
bool giflib_encoder_flush(giflib_encoder e, const giflib_decoder d);
void giflib_encoder_release(giflib_encoder e);
int giflib_encoder_get_output_length(giflib_encoder e);

#ifdef __cplusplus
}
#endif

#endif
