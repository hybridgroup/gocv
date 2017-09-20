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
  typedef cv::Mat_<cv::Vec3b>* MatVec3b;
  typedef cv::Mat_<cv::Vec4b>* MatVec4b;
  #else
  typedef void* Mat;
  typedef void* MatVec3b;
  typedef void* MatVec4b;
  #endif
  
  struct ByteArray toByteArray(const char* buf, int len);
  void ByteArray_Release(struct ByteArray buf);
  
  void Rects_Delete(struct Rects rs);
  void DrawRectsToImage(Mat img, struct Rects rects);
  
  Mat Mat_New();
  void Mat_Delete(Mat m);
  int Mat_Empty(Mat m);

  MatVec3b MatVec3b_New();
  struct ByteArray MatVec3b_ToJpegData(MatVec3b m, int quality);
  void MatVec3b_Delete(MatVec3b m);
  void MatVec3b_CopyTo(MatVec3b src, MatVec3b dst);
  int MatVec3b_Empty(MatVec3b m);
  struct RawData MatVec3b_ToRawData(MatVec3b m);
  MatVec3b RawData_ToMatVec3b(struct RawData r);
  
  void MatVec4b_Delete(MatVec4b m);
  struct RawData MatVec4b_ToRawData(MatVec4b m);
  MatVec4b RawData_ToMatVec4b(struct RawData r);
  MatVec4b LoadAlphaImg(const char* name);
  void MountAlphaImage(MatVec4b img, MatVec3b back, struct Rects rects);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_CORE_H_
