#ifndef _OPENCV3_CORE_H_
#define _OPENCV3_CORE_H_

#include <stdint.h>
#include <stdbool.h>

// Wrapper for std::vector<string>
typedef struct CStrings {
    const char** strs;
    int length;
} CStrings;

typedef struct ByteArray {
    char* data;
    int length;
} ByteArray;

// Wrapper for std::vector<int>
typedef struct IntVector {
    int* val;
    int length;
} IntVector;

// Wrapper for std::vector<float>
typedef struct FloatVector {
    float* val;
    int length;
} FloatVector;

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
extern "C" {
#endif

typedef struct RawData {
    int width;
    int height;
    struct ByteArray data;
} RawData;

// Wrapper for an individual cv::Point2f
typedef struct Point2f {
    float x;
    float y;
} Point2f;

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

// Wrapper for an individual cv::RotatedRect
typedef struct RotatedRect {
    Contour pts;
    Rect boundingRect;
    Point center;
    Size size;
    double angle;
} RotatedRect;

// Wrapper for an individual cv::cvScalar
typedef struct Scalar {
    double val1;
    double val2;
    double val3;
    double val4;
} Scalar;

// Wrapper for a individual cv::KeyPoint
typedef struct KeyPoint {
    double x;
    double y;
    double size;
    double angle;
    double response;
    int octave;
    int classID;
} KeyPoint;

// Wrapper for the vector of KeyPoint struct aka std::vector<KeyPoint>
typedef struct KeyPoints {
    KeyPoint* keypoints;
    int length;
} KeyPoints;

// Wrapper for SimpleBlobDetectorParams aka SimpleBlobDetector::Params
typedef struct SimpleBlobDetectorParams {
    unsigned char   blobColor;
    bool    filterByArea;
    bool    filterByCircularity;
    bool    filterByColor;
    bool    filterByConvexity;
    bool    filterByInertia;
    float   maxArea;
    float   maxCircularity;
    float   maxConvexity;
    float   maxInertiaRatio;
    float   maxThreshold;
    float   minArea;
    float   minCircularity;
    float   minConvexity;
    float   minDistBetweenBlobs;
    float   minInertiaRatio;
    size_t  minRepeatability;
    float   minThreshold;
    float   thresholdStep;
} SimpleBlobDetectorParams;

// Wrapper for an individual cv::DMatch
typedef struct DMatch {
    int queryIdx;
    int trainIdx;
    int imgIdx;
    float distance;
} DMatch;

// Wrapper for the vector of DMatch struct aka std::vector<DMatch>
typedef struct DMatches {
    DMatch* dmatches;
    int length;
} DMatches;

// Wrapper for the vector vector of DMatch struct aka std::vector<std::vector<DMatch>>
typedef struct MultiDMatches {
    DMatches* dmatches;
    int length;
} MultiDMatches;

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

// Wrapper for the vector of Mat aka std::vector<Mat>
typedef struct Mats {
    Mat* mats;
    int length;
} Mats;

Mat Mats_get(struct Mats mats, int i);
struct DMatches MultiDMatches_get(struct MultiDMatches mds, int index);

struct ByteArray toByteArray(const char* buf, int len);
void ByteArray_Release(struct ByteArray buf);

void Contours_Close(struct Contours cs);
void KeyPoints_Close(struct KeyPoints ks);
void Rects_Close(struct Rects rs);
void Mats_Close(struct Mats mats);
void Point_Close(struct Point p);
void Points_Close(struct Points ps);
void DMatches_Close(struct DMatches ds);
void MultiDMatches_Close(struct MultiDMatches mds);

Mat Mat_New();
Mat Mat_NewWithSize(int rows, int cols, int type);
Mat Mat_NewFromScalar(const Scalar ar, int type);
Mat Mat_NewWithSizeFromScalar(const Scalar ar, int rows, int cols, int type);
Mat Mat_NewFromBytes(int rows, int cols, int type, struct ByteArray buf);
Mat Mat_FromPtr(Mat m, int rows, int cols, int type, int prows, int pcols);
void Mat_Close(Mat m);
int Mat_Empty(Mat m);
Mat Mat_Clone(Mat m);
void Mat_CopyTo(Mat m, Mat dst);
int Mat_Total(Mat m);
void Mat_Size(Mat m, IntVector* res);
void Mat_CopyToWithMask(Mat m, Mat dst, Mat mask);
void Mat_ConvertTo(Mat m, Mat dst, int type);
struct ByteArray Mat_ToBytes(Mat m);
struct ByteArray Mat_DataPtr(Mat m);
Mat Mat_Region(Mat m, Rect r);
Mat Mat_Reshape(Mat m, int cn, int rows);
void Mat_PatchNaNs(Mat m);
Mat Mat_ConvertFp16(Mat m);
Scalar Mat_Mean(Mat m);
Mat Mat_Sqrt(Mat m);
int Mat_Rows(Mat m);
int Mat_Cols(Mat m);
int Mat_Channels(Mat m);
int Mat_Type(Mat m);
int Mat_Step(Mat m);

uint8_t Mat_GetUChar(Mat m, int row, int col);
uint8_t Mat_GetUChar3(Mat m, int x, int y, int z);
int8_t Mat_GetSChar(Mat m, int row, int col);
int8_t Mat_GetSChar3(Mat m, int x, int y, int z);
int16_t Mat_GetShort(Mat m, int row, int col);
int16_t Mat_GetShort3(Mat m, int x, int y, int z);
int32_t Mat_GetInt(Mat m, int row, int col);
int32_t Mat_GetInt3(Mat m, int x, int y, int z);
float Mat_GetFloat(Mat m, int row, int col);
float Mat_GetFloat3(Mat m, int x, int y, int z);
double Mat_GetDouble(Mat m, int row, int col);
double Mat_GetDouble3(Mat m, int x, int y, int z);

void Mat_SetTo(Mat m, Scalar value);
void Mat_SetUChar(Mat m, int row, int col, uint8_t val);
void Mat_SetUChar3(Mat m, int x, int y, int z, uint8_t val);
void Mat_SetSChar(Mat m, int row, int col, int8_t val);
void Mat_SetSChar3(Mat m, int x, int y, int z, int8_t val);
void Mat_SetShort(Mat m, int row, int col, int16_t val);
void Mat_SetShort3(Mat m, int x, int y, int z, int16_t val);
void Mat_SetInt(Mat m, int row, int col, int32_t val);
void Mat_SetInt3(Mat m, int x, int y, int z, int32_t val);
void Mat_SetFloat(Mat m, int row, int col, float val);
void Mat_SetFloat3(Mat m, int x, int y, int z, float val);
void Mat_SetDouble(Mat m, int row, int col, double val);
void Mat_SetDouble3(Mat m, int x, int y, int z, double val);

void Mat_AddUChar(Mat m, uint8_t val);
void Mat_SubtractUChar(Mat m, uint8_t val);
void Mat_MultiplyUChar(Mat m, uint8_t val);
void Mat_DivideUChar(Mat m, uint8_t val);
void Mat_AddFloat(Mat m, float val);
void Mat_SubtractFloat(Mat m, float val);
void Mat_MultiplyFloat(Mat m, float val);
void Mat_DivideFloat(Mat m, float val);

void LUT(Mat src, Mat lut, Mat dst);

void Mat_AbsDiff(Mat src1, Mat src2, Mat dst);
void Mat_Add(Mat src1, Mat src2, Mat dst);
void Mat_AddWeighted(Mat src1, double alpha, Mat src2, double beta, double gamma, Mat dst);
void Mat_BitwiseAnd(Mat src1, Mat src2, Mat dst);
void Mat_BitwiseAndWithMask(Mat src1, Mat src2, Mat dst, Mat mask);
void Mat_BitwiseNot(Mat src1, Mat dst);
void Mat_BitwiseNotWithMask(Mat src1, Mat dst, Mat mask);
void Mat_BitwiseOr(Mat src1, Mat src2, Mat dst);
void Mat_BitwiseOrWithMask(Mat src1, Mat src2, Mat dst, Mat mask);
void Mat_BitwiseXor(Mat src1, Mat src2, Mat dst);
void Mat_BitwiseXorWithMask(Mat src1, Mat src2, Mat dst, Mat mask);
void Mat_Compare(Mat src1, Mat src2, Mat dst, int ct);
void Mat_BatchDistance(Mat src1, Mat src2, Mat dist, int dtype, Mat nidx, int normType, int K,
                       Mat mask, int update, bool crosscheck);
int Mat_BorderInterpolate(int p, int len, int borderType);
void Mat_CalcCovarMatrix(Mat samples, Mat covar, Mat mean, int flags, int ctype);
void Mat_CartToPolar(Mat x, Mat y, Mat magnitude, Mat angle, bool angleInDegrees);
bool Mat_CheckRange(Mat m);
void Mat_CompleteSymm(Mat m, bool lowerToUpper);
void Mat_ConvertScaleAbs(Mat src, Mat dst, double alpha, double beta);
void Mat_CopyMakeBorder(Mat src, Mat dst, int top, int bottom, int left, int right, int borderType,
                        Scalar value);
int Mat_CountNonZero(Mat src);
void Mat_DCT(Mat src, Mat dst, int flags);
double Mat_Determinant(Mat m);
void Mat_DFT(Mat m, Mat dst, int flags);
void Mat_Divide(Mat src1, Mat src2, Mat dst);
bool Mat_Eigen(Mat src, Mat eigenvalues, Mat eigenvectors);
void Mat_EigenNonSymmetric(Mat src, Mat eigenvalues, Mat eigenvectors);
void Mat_Exp(Mat src, Mat dst);
void Mat_ExtractChannel(Mat src, Mat dst, int coi);
void Mat_FindNonZero(Mat src, Mat idx);
void Mat_Flip(Mat src, Mat dst, int flipCode);
void Mat_Gemm(Mat src1, Mat src2, double alpha, Mat src3, double beta, Mat dst, int flags);
int Mat_GetOptimalDFTSize(int vecsize);
void Mat_Hconcat(Mat src1, Mat src2, Mat dst);
void Mat_Vconcat(Mat src1, Mat src2, Mat dst);
void Rotate(Mat src, Mat dst, int rotationCode);
void Mat_Idct(Mat src, Mat dst, int flags);
void Mat_Idft(Mat src, Mat dst, int flags, int nonzeroRows);
void Mat_InRange(Mat src, Mat lowerb, Mat upperb, Mat dst);
void Mat_InRangeWithScalar(Mat src, const Scalar lowerb, const Scalar upperb, Mat dst);
void Mat_InsertChannel(Mat src, Mat dst, int coi);
double Mat_Invert(Mat src, Mat dst, int flags);
void Mat_Log(Mat src, Mat dst);
void Mat_Magnitude(Mat x, Mat y, Mat magnitude);
void Mat_Max(Mat src1, Mat src2, Mat dst);
void Mat_MeanStdDev(Mat src, Mat dstMean, Mat dstStdDev);
void Mat_Merge(struct Mats mats, Mat dst);
void Mat_Min(Mat src1, Mat src2, Mat dst);
void Mat_MinMaxIdx(Mat m, double* minVal, double* maxVal, int* minIdx, int* maxIdx);
void Mat_MinMaxLoc(Mat m, double* minVal, double* maxVal, Point* minLoc, Point* maxLoc);
void Mat_MulSpectrums(Mat a, Mat b, Mat c, int flags);
void Mat_Multiply(Mat src1, Mat src2, Mat dst);
void Mat_Subtract(Mat src1, Mat src2, Mat dst);
void Mat_Normalize(Mat src, Mat dst, double alpha, double beta, int typ);
double Norm(Mat src1, int normType);
void Mat_PerspectiveTransform(Mat src, Mat dst, Mat tm);
bool Mat_Solve(Mat src1, Mat src2, Mat dst, int flags);
int Mat_SolveCubic(Mat coeffs, Mat roots);
double Mat_SolvePoly(Mat coeffs, Mat roots, int maxIters);
void Mat_Reduce(Mat src, Mat dst, int dim, int rType, int dType);
void Mat_Repeat(Mat src, int nY, int nX, Mat dst);
void Mat_ScaleAdd(Mat src1, double alpha, Mat src2, Mat dst);
void Mat_Sort(Mat src, Mat dst, int flags);
void Mat_SortIdx(Mat src, Mat dst, int flags);
void Mat_Split(Mat src, struct Mats* mats);
void Mat_Subtract(Mat src1, Mat src2, Mat dst);
Scalar Mat_Trace(Mat src);
void Mat_Transform(Mat src, Mat dst, Mat tm);
void Mat_Transpose(Mat src, Mat dst);
void Mat_PolarToCart(Mat magnitude, Mat degree, Mat x, Mat y, bool angleInDegrees);
void Mat_Pow(Mat src, double power, Mat dst);
void Mat_Phase(Mat x, Mat y, Mat angle, bool angleInDegrees);
Scalar Mat_Sum(Mat src1);

TermCriteria TermCriteria_New(int typ, int maxCount, double epsilon);

int64_t GetCVTickCount();
double GetTickFrequency();

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_CORE_H_
