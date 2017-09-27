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
  
  #ifdef __cplusplus
  typedef cv::Mat* Mat;
  #else
  typedef void* Mat;
  #endif
  
  struct ByteArray toByteArray(const char* buf, int len);
  void ByteArray_Release(struct ByteArray buf);
  
  void Rects_Close(struct Rects rs);
  void DrawRectsToImage(Mat img, struct Rects rects);
  
  Mat Mat_New();
  void Mat_Close(Mat m);
  int Mat_Empty(Mat m);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_CORE_H_
