#ifndef _OPENCV3_CORE_H_
#define _OPENCV3_CORE_H_

typedef struct String {
  const char* str;
  int length;
} String;

typedef struct ByteArray{
  char *data;
  int length;
} ByteArray;

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
extern "C" {
#endif

typedef struct RawData {
  int width;
  int height;
  struct ByteArray data;
} RawData;

// Wrapper for an individual cv::cvPoint
typedef struct Point {
  int x;
  int y;
} Point;

// Wrapper for an individual cv::cvRect
typedef struct Rect {
  int x;
  int y;
  int width;
  int height;
} Rect;

// Wrapper for the vector of Rect struct aka std::vector<Rect>
typedef struct Rects {
  Rect* rects;
  int length;
} Rects;

// Wrapper for an individual cv::cvSize
typedef struct Size {
  int width;
  int height;
} Size;

// Wrapper for an individual cv::cvScalar
typedef struct Scalar {
  double val1;
  double val2;
  double val3;
  double val4;
} Scalar;
  
#ifdef __cplusplus
typedef cv::Mat* Mat;
#else
typedef void* Mat;
#endif
  
struct ByteArray toByteArray(const char* buf, int len);
void ByteArray_Release(struct ByteArray buf);

void Rects_Close(struct Rects rs);

Mat Mat_New();
Mat Mat_NewWithSize(int rows, int cols, int type);
void Mat_Close(Mat m);
int Mat_Empty(Mat m);
Mat Mat_Region(Mat m, Rect r);
int Mat_Rows(Mat m);
int Mat_Cols(Mat m);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_CORE_H_
