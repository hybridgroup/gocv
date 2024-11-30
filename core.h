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

typedef struct Point3f {
    float x;
    float y;
    float z;
} Point3f;

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

// Wrapper for the vector of Point2f structs aka std::vector<Point2f>
typedef struct Points2f {
    Point2f* points;
    int length;
} Points2f;

typedef struct Points3f {
    Point3f *points;
    int length;
} Points3f;

// Contour is alias for Points
typedef Points Contour;


// Contour2f is alias for Points2f
typedef Points2f Contour2f;

typedef struct Contours2f {
    Contour2f *contours;
    int length;
} Contours2f;

// Contour3f is alias for Points3f
typedef Points3f Contour3f;

// Wrapper for the vector of Points3f vectors aka std::vector< std::vector<Point3f> >
typedef struct Contours3f {
    Contour3f *contours;
    int length;
} Contours3f;


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

// Wrapper for an individual cv::cvRect2f
typedef struct Rect2f {
    float x;
    float y;
    float width;
    float height;
} Rect2f;

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

// Wrapper for an individual cv::cvSize
typedef struct Size2f {
    float width;
    float height;
} Size2f;

// Wrapper for an individual cv::RotatedRect
typedef struct RotatedRect {
    Points pts;
    Rect boundingRect;
    Point center;
    Size size;
    double angle;
} RotatedRect;

// Wrapper for an individual cv::RotatedRect2f
typedef struct RotatedRect2f {
    Points2f pts;
    Rect boundingRect;
    Point2f center;
    Size2f size;
    double angle;
} RotatedRect2f;

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
typedef cv::RNG* RNG;
typedef std::vector< cv::Point >* PointVector;
typedef std::vector< std::vector< cv::Point > >* PointsVector;
typedef std::vector< cv::Point2f >* Point2fVector;
typedef std::vector< std::vector< cv::Point2f> >* Points2fVector;
typedef std::vector< cv::Point3f >* Point3fVector;
typedef std::vector< std::vector< cv::Point3f > >* Points3fVector;
typedef cv::RotatedRect* RotatedRectT;
#else
typedef void* Mat;
typedef void* TermCriteria;
typedef void* RNG;
typedef void* PointVector;
typedef void* PointsVector;
typedef void* Point2fVector;
typedef void* Points2fVector;
typedef void* Point3fVector;
typedef void* Points3fVector;
typedef void* RotatedRectT;
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
void Point2f_Close(struct Point2f p);
void Points2f_Close(struct Points2f ps);
void DMatches_Close(struct DMatches ds);
void MultiDMatches_Close(struct MultiDMatches mds);

Mat Mat_New();
Mat Mat_NewWithSize(int rows, int cols, int type);
Mat Mat_NewWithSizes(struct IntVector sizes, int type);
Mat Mat_NewWithSizesFromScalar(IntVector sizes, int type, Scalar ar);
Mat Mat_NewWithSizesFromBytes(IntVector sizes, int type, struct ByteArray buf);
Mat Mat_NewFromScalar(const Scalar ar, int type);
Mat Mat_NewWithSizeFromScalar(const Scalar ar, int rows, int cols, int type);
Mat Mat_NewFromBytes(int rows, int cols, int type, struct ByteArray buf);
Mat Mat_NewFromPoint2fVector(Point2fVector pfv, bool copy_data);
Mat Mat_NewFromPointVector(PointVector pv, bool copy_data);
Mat Mat_FromPtr(Mat m, int rows, int cols, int type, int prows, int pcols);
void Mat_Close(Mat m);
int Mat_Empty(Mat m);
bool Mat_IsContinuous(Mat m);
void Mat_Inv(Mat m);
Mat Mat_Col(Mat m, int c);
Mat Mat_Row(Mat m, int r);
Mat Mat_Clone(Mat m);
void Mat_CopyTo(Mat m, Mat dst);
int Mat_Total(Mat m);
void Mat_Size(Mat m, IntVector* res);
void Mat_CopyToWithMask(Mat m, Mat dst, Mat mask);
void Mat_ConvertTo(Mat m, Mat dst, int type);
void Mat_ConvertToWithParams(Mat m, Mat dst, int type, float alpha, float beta);
struct ByteArray Mat_ToBytes(Mat m);
struct ByteArray Mat_DataPtr(Mat m);
Mat Mat_Region(Mat m, Rect r);
Mat Mat_Reshape(Mat m, int cn, int rows);
void Mat_PatchNaNs(Mat m);
Mat Mat_ConvertFp16(Mat m);
Scalar Mat_Mean(Mat m);
Scalar Mat_MeanWithMask(Mat m, Mat mask);
Mat Mat_Sqrt(Mat m);
int Mat_Rows(Mat m);
int Mat_Cols(Mat m);
int Mat_Channels(Mat m);
int Mat_Type(Mat m);
int Mat_Step(Mat m);
int Mat_ElemSize(Mat m);
Mat Eye(int rows, int cols, int type);
Mat Zeros(int rows, int cols, int type);
Mat Ones(int rows, int cols, int type);

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
Mat Mat_MultiplyMatrix(Mat x, Mat y);

Mat Mat_T(Mat x);

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
void Mat_PCABackProject(Mat data, Mat mean, Mat eigenvectors, Mat result);
void Mat_PCACompute(Mat src, Mat mean, Mat eigenvectors, Mat eigenvalues, int maxComponents);
void Mat_PCAProject(Mat data, Mat mean, Mat eigenvectors, Mat result);
double PSNR(Mat src1, Mat src2);
void SVBackSubst(Mat w, Mat u, Mat vt, Mat rhs, Mat dst);
void SVDecomp(Mat src, Mat w, Mat u, Mat vt);
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
double KMeans(Mat data, int k, Mat bestLabels, TermCriteria criteria, int attempts, int flags, Mat centers);
double KMeansPoints(PointVector pts, int k, Mat bestLabels, TermCriteria criteria, int attempts, int flags, Mat centers);
void Mat_Log(Mat src, Mat dst);
void Mat_Magnitude(Mat x, Mat y, Mat magnitude);
double Mat_Mahalanobis(Mat v1, Mat v2, Mat icovar);
void MulTransposed(Mat src, Mat dest, bool ata);
void Mat_Max(Mat src1, Mat src2, Mat dst);
void Mat_MeanStdDev(Mat src, Mat dstMean, Mat dstStdDev);
void Mat_Merge(struct Mats mats, Mat dst);
void Mat_Min(Mat src1, Mat src2, Mat dst);
void Mat_MinMaxIdx(Mat m, double* minVal, double* maxVal, int* minIdx, int* maxIdx);
void Mat_MinMaxLoc(Mat m, double* minVal, double* maxVal, Point* minLoc, Point* maxLoc);
void Mat_MinMaxLocWithMask(Mat m, double* minVal, double* maxVal, Point* minLoc, Point* maxLoc, Mat mask);
void Mat_MixChannels(struct Mats src, struct Mats dst, struct IntVector fromTo);
void Mat_MulSpectrums(Mat a, Mat b, Mat c, int flags);
void Mat_Multiply(Mat src1, Mat src2, Mat dst);
void Mat_MultiplyWithParams(Mat src1, Mat src2, Mat dst, double scale, int dtype);
void Mat_Subtract(Mat src1, Mat src2, Mat dst);
void Mat_Normalize(Mat src, Mat dst, double alpha, double beta, int typ);
double Norm(Mat src1, int normType);
double NormWithMats(Mat src1, Mat src2, int normType);
void Mat_PerspectiveTransform(Mat src, Mat dst, Mat tm);
bool Mat_Solve(Mat src1, Mat src2, Mat dst, int flags);
int Mat_SolveCubic(Mat coeffs, Mat roots);
double Mat_SolvePoly(Mat coeffs, Mat roots, int maxIters);
void Mat_Reduce(Mat src, Mat dst, int dim, int rType, int dType);
void Mat_ReduceArgMax(Mat src, Mat dst, int axis, bool lastIndex);
void Mat_ReduceArgMin(Mat src, Mat dst, int axis, bool lastIndex);
void Mat_Repeat(Mat src, int nY, int nX, Mat dst);
void Mat_ScaleAdd(Mat src1, double alpha, Mat src2, Mat dst);
void Mat_SetIdentity(Mat src, double scalar);
void Mat_Sort(Mat src, Mat dst, int flags);
void Mat_SortIdx(Mat src, Mat dst, int flags);
void Mat_Split(Mat src, struct Mats* mats);
void Mat_Subtract(Mat src1, Mat src2, Mat dst);
Scalar Mat_Trace(Mat src);
void Mat_Transform(Mat src, Mat dst, Mat tm);
void Mat_Transpose(Mat src, Mat dst);
void Mat_TransposeND(Mat src, struct IntVector order, Mat dst);
void Mat_PolarToCart(Mat magnitude, Mat degree, Mat x, Mat y, bool angleInDegrees);
void Mat_Pow(Mat src, double power, Mat dst);
void Mat_Phase(Mat x, Mat y, Mat angle, bool angleInDegrees);
Scalar Mat_Sum(Mat src1);

TermCriteria TermCriteria_New(int typ, int maxCount, double epsilon);

int64_t GetCVTickCount();
double GetTickFrequency();

Mat Mat_rowRange(Mat m,int startrow,int endrow);
Mat Mat_colRange(Mat m,int startrow,int endrow);

PointVector PointVector_New();
PointVector PointVector_NewFromPoints(Contour points);
PointVector PointVector_NewFromMat(Mat mat);
Point PointVector_At(PointVector pv, int idx);
void PointVector_Append(PointVector pv, Point p);
int PointVector_Size(PointVector pv);
void PointVector_Close(PointVector pv);

PointsVector PointsVector_New();
PointsVector PointsVector_NewFromPoints(Contours points);
PointVector PointsVector_At(PointsVector psv, int idx);
void PointsVector_Append(PointsVector psv, PointVector pv);
int PointsVector_Size(PointsVector psv);
void PointsVector_Close(PointsVector psv);

Point2fVector Point2fVector_New();
void Point2fVector_Close(Point2fVector pfv);
Point2fVector Point2fVector_NewFromPoints(Contour2f pts);
Point2fVector Point2fVector_NewFromMat(Mat mat);
Point2f Point2fVector_At(Point2fVector pfv, int idx);
int Point2fVector_Size(Point2fVector pfv);

void IntVector_Close(struct IntVector ivec);

void CStrings_Close(struct CStrings cstrs);

RNG TheRNG();

void SetRNGSeed(int seed);

void RNG_Fill(RNG rng, Mat mat, int distType, double a, double b, bool saturateRange);

double RNG_Gaussian(RNG rng, double sigma);

unsigned int RNG_Next(RNG rng);

void RandN(Mat mat, Scalar mean, Scalar stddev);

void RandShuffle(Mat mat);

void RandShuffleWithParams(Mat mat, double iterFactor, RNG rng);

void RandU(Mat mat, Scalar low, Scalar high);

void copyPointVectorToPoint2fVector(PointVector src, Point2fVector dest);

void StdByteVectorInitialize(void* data);
void StdByteVectorFree(void *data);
size_t StdByteVectorLen(void *data);
uint8_t* StdByteVectorData(void *data);

Points2fVector Points2fVector_New();
Points2fVector Points2fVector_NewFromPoints(Contours2f points);
int Points2fVector_Size(Points2fVector ps);
Point2fVector Points2fVector_At(Points2fVector ps, int idx);
void Points2fVector_Append(Points2fVector psv, Point2fVector pv);
void Points2fVector_Close(Points2fVector ps);

Point3fVector Point3fVector_New();
Point3fVector Point3fVector_NewFromPoints(Contour3f points);
Point3fVector Point3fVector_NewFromMat(Mat mat);
void Point3fVector_Append(Point3fVector pfv, Point3f point);
Point3f Point3fVector_At(Point3fVector pfv, int idx);
int Point3fVector_Size(Point3fVector pfv);
void Point3fVector_Close(Point3fVector pv);
Points3fVector Points3fVector_New();
Points3fVector Points3fVector_NewFromPoints(Contours3f points);
int Points3fVector_Size(Points3fVector ps);
Point3fVector Points3fVector_At(Points3fVector ps, int idx);
void Points3fVector_Append(Points3fVector psv, Point3fVector pv);
void Points3fVector_Close(Points3fVector ps);

void SetNumThreads(int n);
int GetNumThreads();


struct RotatedRect RotatedRect_Create(struct Point2f center, int width, int height, float angle);
struct RotatedRect2f RotatedRect2f_Create(struct Point2f center, float width, float height, float angle);


#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_CORE_H_
