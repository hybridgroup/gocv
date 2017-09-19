#ifndef _OPENCV3_H_
#define _OPENCV3_H_

#include "convert.h"

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
extern "C" {
#endif

typedef struct RawData {
  int width;
  int height;
  struct ByteArray data;
} RawData;

typedef struct Rect {
  int x;
  int y;
  int width;
  int height;
} Rect;

typedef struct Rects {
  Rect* rects;
  int length;
} Rects;

#ifdef __cplusplus
typedef cv::Mat_<cv::Vec3b>* MatVec3b;
typedef cv::Mat_<cv::Vec4b>* MatVec4b;
typedef cv::CascadeClassifier* CascadeClassifier;
#else
typedef void* MatVec3b;
typedef void* MatVec4b;
typedef void* CascadeClassifier;
#endif

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

CascadeClassifier CascadeClassifier_New();
void CascadeClassifier_Delete(CascadeClassifier cs);
int CascadeClassifier_Load(CascadeClassifier cs, const char* name);
struct Rects CascadeClassifier_DetectMultiScale(CascadeClassifier cs, MatVec3b img);
void Rects_Delete(struct Rects rs);
void DrawRectsToImage(MatVec3b img, struct Rects rects);
MatVec4b LoadAlphaImg(const char* name);
void MountAlphaImage(MatVec4b img, MatVec3b back, struct Rects rects);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_H_