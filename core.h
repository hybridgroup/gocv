#ifndef _OPENCV3_CORE_H_
#define _OPENCV3_CORE_H_

#include <stdint.h>

typedef struct String {
  const char* str;
  int length;
} String;

typedef struct ByteArray{
  char *data;
  int length;
} ByteArray;

// Wrapper for std::vector<int>
typedef struct IntVector {
  int *val;
  int length;
} IntVector;

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

// Wrapper for the vector of Point structs aka std::vector<Point>
typedef struct Points {
  Point* points;
  int length;
} Points;

// Contour is alias for Points
typedef Points Contour;

// Wrapper for the vector of Points vectors aka std::vector< std::vector<Point> >
typedef struct Contours {
  Contour* contours;
  int length;
} Contours;

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

// Wrapper for an individual cv::Moment
typedef struct Moment {
  double m00;
  double m10;
  double m01;
  double m20;
  double m11;
  double m02;
  double m30;
  double m21;
  double m12;
  double m03;

  double mu20;
  double mu11;
  double mu02;
  double mu30;
  double mu21;
  double mu12;
  double mu03;

  double nu20;
  double nu11;
  double nu02;
  double nu30;
  double nu21;
  double nu12;
  double nu03;
} Moment;

#ifdef __cplusplus
typedef cv::Mat* Mat;
typedef cv::TermCriteria* TermCriteria;
#else
typedef void* Mat;
typedef void* TermCriteria;
#endif
  
struct ByteArray toByteArray(const char* buf, int len);
void ByteArray_Release(struct ByteArray buf);

void Contours_Close(struct Contours cs);
void Rects_Close(struct Rects rs);

Mat Mat_New();
Mat Mat_NewWithSize(int rows, int cols, int type);
Mat Mat_NewFromScalar(const Scalar ar, int type);
void Mat_Close(Mat m);
int Mat_Empty(Mat m);
Mat Mat_Clone(Mat m);
void Mat_CopyTo(Mat m, Mat dst);
Mat Mat_Region(Mat m, Rect r);
Scalar Mat_Mean(Mat m);
int Mat_Rows(Mat m);
int Mat_Cols(Mat m);
uint8_t Mat_GetUChar(Mat m, int row, int col);
int8_t Mat_GetSChar(Mat m, int row, int col);
int16_t Mat_GetShort(Mat m, int row, int col);
int32_t Mat_GetInt(Mat m, int row, int col);
float Mat_GetFloat(Mat m, int row, int col);
double Mat_GetDouble(Mat m, int row, int col);

void Mat_AbsDiff(Mat src1, Mat src2, Mat dst);
void Mat_Add(Mat src1, Mat src2, Mat dst);
void Mat_AddWeighted(Mat src1, double alpha, Mat src2, double beta, double gamma, Mat dst);
void Mat_BitwiseAnd(Mat src1, Mat src2, Mat dst);
void Mat_BitwiseNot(Mat src1, Mat dst);
void Mat_BitwiseOr(Mat src1, Mat src2, Mat dst);
void Mat_BitwiseXor(Mat src1, Mat src2, Mat dst);
void Mat_InRange(Mat src, Mat lowerb, Mat upperb, Mat dst);
int Mat_GetOptimalDFTSize(int vecsize);
void Mat_DFT(Mat m, Mat dst);
void Mat_Merge(Mat m, size_t count, Mat dst);
void Mat_Normalize(Mat src, Mat dst, double alpha, double beta, int typ);

TermCriteria TermCriteria_New(int typ, int maxCount, double epsilon);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_CORE_H_
