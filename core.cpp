#include "core.h"
#include <string.h>

// Mat_New creates a new empty Mat
Mat Mat_New() {
    return new cv::Mat();
}

// Mat_NewWithSize creates a new Mat with a specific size dimension and number of channels.
Mat Mat_NewWithSize(int rows, int cols, int type) {
    return new cv::Mat(rows, cols, type, 0.0);
}

// Mat_NewFromScalar creates a new Mat from a Scalar. Intended to be used
// for Mat comparison operation such as InRange.
Mat Mat_NewFromScalar(Scalar ar, int type) {
    cv::Scalar c = cv::Scalar(ar.val1, ar.val2, ar.val3, ar.val4);
    return new cv::Mat(1, 1, type, c);
}

// Mat_NewWithSizeFromScalar creates a new Mat from a Scalar with a specific size dimension and number of channels
Mat Mat_NewWithSizeFromScalar(Scalar ar, int rows, int cols, int type) {
    cv::Scalar c = cv::Scalar(ar.val1, ar.val2, ar.val3, ar.val4);
    return new cv::Mat(rows, cols, type, c);
}

Mat Mat_NewFromBytes(int rows, int cols, int type, struct ByteArray buf) {
    return new cv::Mat(rows, cols, type, buf.data);
}

Mat Mat_FromPtr(Mat m, int rows, int cols, int type, int prow, int pcol) {
    return new cv::Mat(rows, cols, type, m->ptr(prow, pcol));
}

// Mat_Close deletes an existing Mat
void Mat_Close(Mat m) {
    delete m;
}

// Mat_Empty tests if a Mat is empty
int Mat_Empty(Mat m) {
    return m->empty();
}

// Mat_Clone returns a clone of this Mat
Mat Mat_Clone(Mat m) {
    return new cv::Mat(m->clone());
}

// Mat_CopyTo copies this Mat to another Mat.
void Mat_CopyTo(Mat m, Mat dst) {
    m->copyTo(*dst);
}

// Mat_CopyToWithMask copies this Mat to another Mat while applying the mask
void Mat_CopyToWithMask(Mat m, Mat dst, Mat mask) {
    m->copyTo(*dst, *mask);
}

void Mat_ConvertTo(Mat m, Mat dst, int type) {
    m->convertTo(*dst, type);
}

// Mat_ToBytes returns the bytes representation of the underlying data.
struct ByteArray Mat_ToBytes(Mat m) {
    return toByteArray(reinterpret_cast<const char*>(m->data), m->total() * m->elemSize());
}

struct ByteArray Mat_DataPtr(Mat m) {
    return ByteArray {reinterpret_cast<char*>(m->data), static_cast<int>(m->total() * m->elemSize())};
}

// Mat_Region returns a Mat of a region of another Mat
Mat Mat_Region(Mat m, Rect r) {
    return new cv::Mat(*m, cv::Rect(r.x, r.y, r.width, r.height));
}

Mat Mat_Reshape(Mat m, int cn, int rows) {
    return new cv::Mat(m->reshape(cn, rows));
}

void Mat_PatchNaNs(Mat m) {
    cv::patchNaNs(*m);
}

Mat Mat_ConvertFp16(Mat m) {
    Mat dst = new cv::Mat();
    cv::convertFp16(*m, *dst);
    return dst;
}

Mat Mat_Sqrt(Mat m) {
    Mat dst = new cv::Mat();
    cv::sqrt(*m, *dst);
    return dst;
}

// Mat_Mean calculates the mean value M of array elements, independently for each channel, and return it as Scalar vector
// TODO pass second paramter with mask
Scalar Mat_Mean(Mat m) {
    cv::Scalar c = cv::mean(*m);
    Scalar scal = Scalar();
    scal.val1 = c.val[0];
    scal.val2 = c.val[1];
    scal.val3 = c.val[2];
    scal.val4 = c.val[3];
    return scal;
}

void LUT(Mat src, Mat lut, Mat dst) {
    cv::LUT(*src, *lut, *dst);
}

// Mat_Rows returns how many rows in this Mat.
int Mat_Rows(Mat m) {
    return m->rows;
}

// Mat_Cols returns how many columns in this Mat.
int Mat_Cols(Mat m) {
    return m->cols;
}

// Mat_Channels returns how many channels in this Mat.
int Mat_Channels(Mat m) {
    return m->channels();
}

// Mat_Type returns the type from this Mat.
int Mat_Type(Mat m) {
    return m->type();
}

// Mat_Step returns the number of bytes each matrix row occupies.
int Mat_Step(Mat m) {
    return m->step;
}

int Mat_Total(Mat m) {
    return m->total();
}

void Mat_Size(Mat m, IntVector* res) {
    cv::MatSize ms(m->size);
    int* ids = new int[ms.dims()];

    for (size_t i = 0; i < ms.dims(); ++i) {
        ids[i] = ms[i];
    }

    res->length = ms.dims();
    res->val = ids;
    return;
}

// Mat_GetUChar returns a specific row/col value from this Mat expecting
// each element to contain a schar aka CV_8U.
uint8_t Mat_GetUChar(Mat m, int row, int col) {
    return m->at<uchar>(row, col);
}

uint8_t Mat_GetUChar3(Mat m, int x, int y, int z) {
    return m->at<uchar>(x, y, z);
}

// Mat_GetSChar returns a specific row/col value from this Mat expecting
// each element to contain a schar aka CV_8S.
int8_t Mat_GetSChar(Mat m, int row, int col) {
    return m->at<schar>(row, col);
}

int8_t Mat_GetSChar3(Mat m, int x, int y, int z) {
    return m->at<schar>(x, y, z);
}

// Mat_GetShort returns a specific row/col value from this Mat expecting
// each element to contain a short aka CV_16S.
int16_t Mat_GetShort(Mat m, int row, int col) {
    return m->at<short>(row, col);
}

int16_t Mat_GetShort3(Mat m, int x, int y, int z) {
    return m->at<short>(x, y, z);
}

// Mat_GetInt returns a specific row/col value from this Mat expecting
// each element to contain an int aka CV_32S.
int32_t Mat_GetInt(Mat m, int row, int col) {
    return m->at<int>(row, col);
}

int32_t Mat_GetInt3(Mat m, int x, int y, int z) {
    return m->at<int>(x, y, z);
}

// Mat_GetFloat returns a specific row/col value from this Mat expecting
// each element to contain a float aka CV_32F.
float Mat_GetFloat(Mat m, int row, int col) {
    return m->at<float>(row, col);
}

float Mat_GetFloat3(Mat m, int x, int y, int z) {
    return m->at<float>(x, y, z);
}

// Mat_GetDouble returns a specific row/col value from this Mat expecting
// each element to contain a double aka CV_64F.
double Mat_GetDouble(Mat m, int row, int col) {
    return m->at<double>(row, col);
}

double Mat_GetDouble3(Mat m, int x, int y, int z) {
    return m->at<double>(x, y, z);
}

void Mat_SetTo(Mat m, Scalar value) {
    cv::Scalar c_value(value.val1, value.val2, value.val3, value.val4);
    m->setTo(c_value);
}

// Mat_SetUChar set a specific row/col value from this Mat expecting
// each element to contain a schar aka CV_8U.
void Mat_SetUChar(Mat m, int row, int col, uint8_t val) {
    m->at<uchar>(row, col) = val;
}

void Mat_SetUChar3(Mat m, int x, int y, int z, uint8_t val) {
    m->at<uchar>(x, y, z) = val;
}

// Mat_SetSChar set a specific row/col value from this Mat expecting
// each element to contain a schar aka CV_8S.
void Mat_SetSChar(Mat m, int row, int col, int8_t val) {
    m->at<schar>(row, col) = val;
}

void Mat_SetSChar3(Mat m, int x, int y, int z, int8_t val) {
    m->at<schar>(x, y, z) = val;
}

// Mat_SetShort set a specific row/col value from this Mat expecting
// each element to contain a short aka CV_16S.
void Mat_SetShort(Mat m, int row, int col, int16_t val) {
    m->at<short>(row, col) = val;
}

void Mat_SetShort3(Mat m, int x, int y, int z, int16_t val) {
    m->at<short>(x, y, z) = val;
}

// Mat_SetInt set a specific row/col value from this Mat expecting
// each element to contain an int aka CV_32S.
void Mat_SetInt(Mat m, int row, int col, int32_t val) {
    m->at<int>(row, col) = val;
}

void Mat_SetInt3(Mat m, int x, int y, int z, int32_t val) {
    m->at<int>(x, y, z) = val;
}

// Mat_SetFloat set a specific row/col value from this Mat expecting
// each element to contain a float aka CV_32F.
void Mat_SetFloat(Mat m, int row, int col, float val) {
    m->at<float>(row, col) = val;
}

void Mat_SetFloat3(Mat m, int x, int y, int z, float val) {
    m->at<float>(x, y, z) = val;
}

// Mat_SetDouble set a specific row/col value from this Mat expecting
// each element to contain a double aka CV_64F.
void Mat_SetDouble(Mat m, int row, int col, double val) {
    m->at<double>(row, col) = val;
}

void Mat_SetDouble3(Mat m, int x, int y, int z, double val) {
    m->at<double>(x, y, z) = val;
}

void Mat_AddUChar(Mat m, uint8_t val) {
    *m += val;
}

void Mat_SubtractUChar(Mat m, uint8_t val) {
    *m -= val;
}

void Mat_MultiplyUChar(Mat m, uint8_t val) {
    *m *= val;
}

void Mat_DivideUChar(Mat m, uint8_t val) {
    *m /= val;
}

void Mat_AddFloat(Mat m, float val) {
    *m += val;
}

void Mat_SubtractFloat(Mat m, float val) {
    *m -= val;
}

void Mat_MultiplyFloat(Mat m, float val) {
    *m *= val;
}

void Mat_DivideFloat(Mat m, float val) {
    *m /= val;
}

void Mat_AbsDiff(Mat src1, Mat src2, Mat dst) {
    cv::absdiff(*src1, *src2, *dst);
}

void Mat_Add(Mat src1, Mat src2, Mat dst) {
    cv::add(*src1, *src2, *dst);
}

void Mat_AddWeighted(Mat src1, double alpha, Mat src2, double beta, double gamma, Mat dst) {
    cv::addWeighted(*src1, alpha, *src2, beta, gamma, *dst);
}

void Mat_BitwiseAnd(Mat src1, Mat src2, Mat dst) {
    cv::bitwise_and(*src1, *src2, *dst);
}

void Mat_BitwiseAndWithMask(Mat src1, Mat src2, Mat dst, Mat mask){
    cv::bitwise_and(*src1, *src2, *dst, *mask);
}

void Mat_BitwiseNot(Mat src1, Mat dst) {
    cv::bitwise_not(*src1, *dst);
}

void Mat_BitwiseNotWithMask(Mat src1, Mat dst, Mat mask) {
    cv::bitwise_not(*src1, *dst, *mask);
}

void Mat_BitwiseOr(Mat src1, Mat src2, Mat dst) {
    cv::bitwise_or(*src1, *src2, *dst);
}

void Mat_BitwiseOrWithMask(Mat src1, Mat src2, Mat dst, Mat mask) {
    cv::bitwise_or(*src1, *src2, *dst, *mask);
}

void Mat_BitwiseXor(Mat src1, Mat src2, Mat dst) {
    cv::bitwise_xor(*src1, *src2, *dst);
}

void Mat_BitwiseXorWithMask(Mat src1, Mat src2, Mat dst, Mat mask) {
    cv::bitwise_xor(*src1, *src2, *dst, *mask);
}

void Mat_BatchDistance(Mat src1, Mat src2, Mat dist, int dtype, Mat nidx, int normType, int K,
                       Mat mask, int update, bool crosscheck) {
    cv::batchDistance(*src1, *src2, *dist, dtype, *nidx, normType, K, *mask, update, crosscheck);
}

int Mat_BorderInterpolate(int p, int len, int borderType) {
    return cv::borderInterpolate(p, len, borderType);
}

void  Mat_CalcCovarMatrix(Mat samples, Mat covar, Mat mean, int flags, int ctype) {
    cv::calcCovarMatrix(*samples, *covar, *mean, flags, ctype);
}

void  Mat_CartToPolar(Mat x, Mat y, Mat magnitude, Mat angle, bool angleInDegrees) {
    cv::cartToPolar(*x, *y, *magnitude, *angle, angleInDegrees);
}

bool Mat_CheckRange(Mat m) {
    return cv::checkRange(*m);
}

void Mat_Compare(Mat src1, Mat src2, Mat dst, int ct) {
    cv::compare(*src1, *src2, *dst, ct);
}

void Mat_Compute(Mat a, Mat w, Mat u, Mat vt, int flags) {
    cv::SVD::compute(*a, *w, *u, *vt, flags);
}

int Mat_CountNonZero(Mat src) {
    return cv::countNonZero(*src);
}


void Mat_CompleteSymm(Mat m, bool lowerToUpper) {
    cv::completeSymm(*m, lowerToUpper);
}

void Mat_ConvertScaleAbs(Mat src, Mat dst, double alpha, double beta) {
    cv::convertScaleAbs(*src, *dst, alpha, beta);
}

void Mat_CopyMakeBorder(Mat src, Mat dst, int top, int bottom, int left, int right, int borderType,
                        Scalar value) {
    cv::Scalar c_value(value.val1, value.val2, value.val3, value.val4);
    cv::copyMakeBorder(*src, *dst, top, bottom, left, right, borderType, c_value);
}

void Mat_DCT(Mat src, Mat dst, int flags) {
    cv::dct(*src, *dst, flags);
}

double Mat_Determinant(Mat m) {
    return cv::determinant(*m);
}

void Mat_DFT(Mat m, Mat dst, int flags) {
    cv::dft(*m, *dst, flags);
}

void Mat_Divide(Mat src1, Mat src2, Mat dst) {
    cv::divide(*src1, *src2, *dst);
}

bool Mat_Eigen(Mat src, Mat eigenvalues, Mat eigenvectors) {
    return cv::eigen(*src, *eigenvalues, *eigenvectors);
}

void Mat_EigenNonSymmetric(Mat src, Mat eigenvalues, Mat eigenvectors) {
    cv::eigenNonSymmetric(*src, *eigenvalues, *eigenvectors);
}

void Mat_Exp(Mat src, Mat dst) {
    cv::exp(*src, *dst);
}

void Mat_ExtractChannel(Mat src, Mat dst, int coi) {
    cv::extractChannel(*src, *dst, coi);
}

void Mat_FindNonZero(Mat src, Mat idx) {
    cv::findNonZero(*src, *idx);
}

void Mat_Flip(Mat src, Mat dst, int flipCode) {
    cv::flip(*src, *dst, flipCode);
}

void Mat_Gemm(Mat src1, Mat src2, double alpha, Mat src3, double beta, Mat dst, int flags) {
    cv::gemm(*src1, *src2, alpha, *src3, beta, *dst, flags);
}

int Mat_GetOptimalDFTSize(int vecsize) {
    return cv::getOptimalDFTSize(vecsize);
}

void Mat_Hconcat(Mat src1, Mat src2, Mat dst) {
    cv::hconcat(*src1, *src2, *dst);
}

void Mat_Vconcat(Mat src1, Mat src2, Mat dst) {
    cv::vconcat(*src1, *src2, *dst);
}

void Rotate(Mat src, Mat dst, int rotateCode) {
    cv::rotate(*src, *dst, rotateCode);
}

void Mat_Idct(Mat src, Mat dst, int flags) {
    cv::idct(*src, *dst, flags);
}

void Mat_Idft(Mat src, Mat dst, int flags, int nonzeroRows) {
    cv::idft(*src, *dst, flags, nonzeroRows);
}

void Mat_InRange(Mat src, Mat lowerb, Mat upperb, Mat dst) {
    cv::inRange(*src, *lowerb, *upperb, *dst);
}

void Mat_InRangeWithScalar(Mat src, Scalar lowerb, Scalar upperb, Mat dst) {
    cv::Scalar lb = cv::Scalar(lowerb.val1, lowerb.val2, lowerb.val3, lowerb.val4);
    cv::Scalar ub = cv::Scalar(upperb.val1, upperb.val2, upperb.val3, upperb.val4);
    cv::inRange(*src, lb, ub, *dst);
}

void Mat_InsertChannel(Mat src, Mat dst, int coi) {
    cv::insertChannel(*src, *dst, coi);
}

double Mat_Invert(Mat src, Mat dst, int flags) {
    double ret = cv::invert(*src, *dst, flags);
    return ret;
}

void Mat_Log(Mat src, Mat dst) {
    cv::log(*src, *dst);
}

void Mat_Magnitude(Mat x, Mat y, Mat magnitude) {
    cv::magnitude(*x, *y, *magnitude);
}

void Mat_Max(Mat src1, Mat src2, Mat dst) {
    cv::max(*src1, *src2, *dst);
}

void Mat_MeanStdDev(Mat src, Mat dstMean, Mat dstStdDev) {
    cv::meanStdDev(*src, *dstMean, *dstStdDev);
}

void Mat_Merge(struct Mats mats, Mat dst) {
    std::vector<cv::Mat> images;

    for (int i = 0; i < mats.length; ++i) {
        images.push_back(*mats.mats[i]);
    }

    cv::merge(images, *dst);
}

void Mat_Min(Mat src1, Mat src2, Mat dst) {
    cv::min(*src1, *src2, *dst);
}

void Mat_MinMaxIdx(Mat m, double* minVal, double* maxVal, int* minIdx, int* maxIdx) {
    cv::minMaxIdx(*m, minVal, maxVal, minIdx, maxIdx);
}

void Mat_MinMaxLoc(Mat m, double* minVal, double* maxVal, Point* minLoc, Point* maxLoc) {
    cv::Point cMinLoc;
    cv::Point cMaxLoc;
    cv::minMaxLoc(*m, minVal, maxVal, &cMinLoc, &cMaxLoc);

    minLoc->x = cMinLoc.x;
    minLoc->y = cMinLoc.y;
    maxLoc->x = cMaxLoc.x;
    maxLoc->y = cMaxLoc.y;
}

void Mat_MulSpectrums(Mat a, Mat b, Mat c, int flags) {
    cv::mulSpectrums(*a, *b, *c, flags);
}

void Mat_Multiply(Mat src1, Mat src2, Mat dst) {
    cv::multiply(*src1, *src2, *dst);
}

void Mat_Normalize(Mat src, Mat dst, double alpha, double beta, int typ) {
    cv::normalize(*src, *dst, alpha, beta, typ);
}

double Norm(Mat src1, int normType) {
    return cv::norm(*src1, normType);
}

void Mat_PerspectiveTransform(Mat src, Mat dst, Mat tm) {
    cv::perspectiveTransform(*src, *dst, *tm);
}

bool Mat_Solve(Mat src1, Mat src2, Mat dst, int flags) {
    return cv::solve(*src1, *src2, *dst, flags);
}

int Mat_SolveCubic(Mat coeffs, Mat roots) {
    return cv::solveCubic(*coeffs, *roots);
}

double Mat_SolvePoly(Mat coeffs, Mat roots, int maxIters) {
    return cv::solvePoly(*coeffs, *roots, maxIters);
}

void Mat_Reduce(Mat src, Mat dst, int dim, int rType, int dType) {
    cv::reduce(*src, *dst, dim, rType, dType);
}

void Mat_Repeat(Mat src, int nY, int nX, Mat dst) {
    cv::repeat(*src, nY, nX, *dst);
}

void Mat_ScaleAdd(Mat src1, double alpha, Mat src2, Mat dst) {
    cv::scaleAdd(*src1, alpha, *src2, *dst);
}

void Mat_Sort(Mat src, Mat dst, int flags) {
    cv::sort(*src, *dst, flags);
}

void Mat_SortIdx(Mat src, Mat dst, int flags) {
    cv::sortIdx(*src, *dst, flags);
}

void Mat_Split(Mat src, struct Mats* mats) {
    std::vector<cv::Mat> channels;
    cv::split(*src, channels);
    mats->mats = new Mat[channels.size()];

    for (size_t i = 0; i < channels.size(); ++i) {
        mats->mats[i] = new cv::Mat(channels[i]);
    }

    mats->length = (int)channels.size();
}

void Mat_Subtract(Mat src1, Mat src2, Mat dst) {
    cv::subtract(*src1, *src2, *dst);
}

Scalar Mat_Trace(Mat src) {
    cv::Scalar c = cv::trace(*src);
    Scalar scal = Scalar();
    scal.val1 = c.val[0];
    scal.val2 = c.val[1];
    scal.val3 = c.val[2];
    scal.val4 = c.val[3];
    return scal;
}

void Mat_Transform(Mat src, Mat dst, Mat tm) {
    cv::transform(*src, *dst, *tm);
}

void Mat_Transpose(Mat src, Mat dst) {
    cv::transpose(*src, *dst);
}

void Mat_PolarToCart(Mat magnitude, Mat degree, Mat x, Mat y, bool angleInDegrees) {
    cv::polarToCart(*magnitude, *degree, *x, *y, angleInDegrees);
}

void Mat_Pow(Mat src, double power, Mat dst) {
    cv::pow(*src, power, *dst);
}

void Mat_Phase(Mat x, Mat y, Mat angle, bool angleInDegrees) {
	cv::phase(*x, *y, *angle, angleInDegrees);
}


Scalar Mat_Sum(Mat src) {
    cv::Scalar c = cv::sum(*src);
    Scalar scal = Scalar();
    scal.val1 = c.val[0];
    scal.val2 = c.val[1];
    scal.val3 = c.val[2];
    scal.val4 = c.val[3];
    return scal;
}

// TermCriteria_New creates a new TermCriteria
TermCriteria TermCriteria_New(int typ, int maxCount, double epsilon) {
    return new cv::TermCriteria(typ, maxCount, epsilon);
}

void Contours_Close(struct Contours cs) {
    for (int i = 0; i < cs.length; i++) {
        Points_Close(cs.contours[i]);
    }

    delete[] cs.contours;
}

void KeyPoints_Close(struct KeyPoints ks) {
    delete[] ks.keypoints;
}

void Points_Close(Points ps) {
    for (size_t i = 0; i < ps.length; i++) {
        Point_Close(ps.points[i]);
    }

    delete[] ps.points;
}

void Point_Close(Point p) {}

void Rects_Close(struct Rects rs) {
    delete[] rs.rects;
}

void DMatches_Close(struct DMatches ds) {
    delete[] ds.dmatches;
}

void MultiDMatches_Close(struct MultiDMatches mds) {
    for (size_t i = 0; i < mds.length; i++) {
        DMatches_Close(mds.dmatches[i]);
    }

    delete[] mds.dmatches;
}

struct DMatches MultiDMatches_get(struct MultiDMatches mds, int index) {
    return mds.dmatches[index];
}

// since it is next to impossible to iterate over mats.mats on the cgo side
Mat Mats_get(struct Mats mats, int i) {
    return mats.mats[i];
}

void Mats_Close(struct Mats mats) {
    delete[] mats.mats;
}

void ByteArray_Release(struct ByteArray buf) {
    delete[] buf.data;
}

struct ByteArray toByteArray(const char* buf, int len) {
    ByteArray ret = {new char[len], len};
    memcpy(ret.data, buf, len);
    return ret;
}

int64 GetCVTickCount() {
    return cv::getTickCount();
}

double GetTickFrequency() {
    return cv::getTickFrequency();
}

// --------------------add by yoda.guo--------------------------------
// xy = [x y]
// XY = [[x  y  1  0]
//       [y -x  0  1]]
cv::Mat stitch(const cv::Mat xy)
{
    int M = xy.rows;
    cv::Mat x = xy.col(0);
    cv::Mat y = xy.col(1);
    cv::Mat zeros = cv::Mat::zeros(M, 1, CV_32FC1);
    cv::Mat ones = cv::Mat::ones(M, 1, CV_32FC1);

    cv::Mat XY(2 * M, 4, CV_32FC1);

    cv::Mat XY_up = XY.rowRange(0, M);
    x.copyTo(XY_up.col(0));
    y.copyTo(XY_up.col(1));
    ones.copyTo(XY_up.col(2));
    zeros.copyTo(XY_up.col(3));

    cv::Mat XY_down = XY.rowRange(M, 2 * M);
    y.copyTo(XY_down.col(0));
    cv::Mat _x = x * -1;
    _x.copyTo(XY_down.col(1));
    zeros.copyTo(XY_down.col(2));
    ones.copyTo(XY_down.col(3));

    return std::move(XY);
}

int rank(const cv::Mat mat)
{
    cv::Mat1d w; // singular values
    cv::SVD::compute(mat, w);
    return cv::countNonZero(w > 0.f);
}

cv::Mat tformfwd(const cv::Mat trans, const cv::Mat uv)
{
    cv::Mat UV = cv::Mat::ones(uv.rows, 3, uv.type());
    uv.copyTo(UV.colRange(0, 2));
    cv::Mat XY = UV * trans;
    cv::Mat xy;
    XY.colRange(0, 2).copyTo(xy);
    return std::move(xy);
}

// 2-norm of matrix = largest singular value of mat.
double norm(const cv::Mat mat)
{
    cv::Mat1d w; // singular values
    cv::SVD::compute(mat, w);
    double min_v, max_v;
    cv::minMaxIdx(w, &min_v, &max_v);
    return max_v;
}

cv::Mat nonReflectiveSimilarity(const cv::Mat uv, const cv::Mat xy, int option_k = 2)
{
    int K = option_k;
    int M = xy.rows;
    cv::Mat X = stitch(xy);
    cv::Mat U = cv::Mat(uv.t()).reshape(1, 2 * M);

    // We know that X * r = U
    CV_Assert(::rank(X) >= 2 * K); // unique solution
    cv::Mat r(4, 1, CV_32FC1);
    cv::solve(X, U, r, cv::DECOMP_SVD);

    const float * pdata = r.ptr<float>(0);
    cv::Mat tinv = (cv::Mat_<float>(3, 3) <<
                    pdata[0], -pdata[1], 0,
                    pdata[1],  pdata[0], 0,
                    pdata[2],  pdata[3], 1);

    return std::move(tinv.inv(cv::DECOMP_SVD));
}

cv::Mat findSimilarity(const cv::Mat uv, const cv::Mat xy, int option_k = 2)
{
    /// solve for trans1
    cv::Mat trans1 = nonReflectiveSimilarity(uv, xy, option_k);
    /// solve for trans2
    // manually reflect the xy data across the Y-axis
    cv::Mat xyR;
    xy.copyTo(xyR);
    xyR.col(0) *= -1;
    cv::Mat trans2 = nonReflectiveSimilarity(uv, xyR, option_k);
    // manually reflect the tfrom to undo the reflection done on xyR.
    trans2.col(0) *= -1;
    // print(trans1, "trans1");
    // print(trans2, "trans2");

    // Figure out if trans1 or trans2 is better
    cv::Mat xy1 = tformfwd(trans1, uv);
    double norm1 = ::norm(xy1 - xy);

    cv::Mat xy2 = tformfwd(trans2, uv);
    double norm2 = ::norm(xy2 - xy);

    if (norm1 <= norm2)
        return std::move(trans1);
    else
        return std::move(trans2);
}

Mat Mat_SimilarityTransform(Mat src, Mat dst)
{
    cv::Mat trans = findSimilarity(*src, *dst);
    return new cv::Mat(std::move(trans));
}

//----insightface
/*
cv::Mat meanAxis0(const cv::Mat &src)
{
    int num = src.rows;
    int dim = src.cols;

    // x1 y1
    // x2 y2

    cv::Mat output(1,dim,CV_32F);
    for(int i = 0 ; i <  dim; i ++)
    {
        float sum = 0 ;
        for(int j = 0 ; j < num ; j++)
        {
            sum += src.at<float>(j,i);
        }
        output.at<float>(0,i) = sum/num;
    }
    return output;
}

cv::Mat elementwiseMinus(const cv::Mat &A,const cv::Mat &B)
{
    cv::Mat output(A.rows,A.cols,A.type());

    assert(B.cols == A.cols);
    if(B.cols == A.cols)
    {
        for(int i = 0 ; i <  A.rows; i ++)
        {
            for(int j = 0 ; j < B.cols; j++)
            {
               output.at<float>(i,j) = A.at<float>(i,j) - B.at<float>(0,j);
            }
        }
    }
    return output;
}

cv::Mat varAxis0(const cv::Mat &src)
{
    cv::Mat temp_ = elementwiseMinus(src,meanAxis0(src));
    cv::multiply(temp_ ,temp_ ,temp_ );
    return meanAxis0(temp_);
}

int MatrixRank(cv::Mat M)
{
    cv::Mat w, u, vt;
    cv::SVD::compute(M, w, u, vt);
    cv::Mat1b nonZeroSingularValues = w > 0.0001;
    int rank = cv::countNonZero(nonZeroSingularValues);
    return rank;
}

cv::Mat similarTransform(cv::Mat src,cv::Mat dst) {
    int num = src.rows;
    int dim = src.cols;
    cv::Mat src_mean = meanAxis0(src);
    cv::Mat dst_mean = meanAxis0(dst);
    cv::Mat src_demean = elementwiseMinus(src, src_mean);
    cv::Mat dst_demean = elementwiseMinus(dst, dst_mean);
    cv::Mat A = (dst_demean.t() * src_demean) / static_cast<float>(num);
    cv::Mat d(dim, 1, CV_32F);
    d.setTo(1.0f);
    if (cv::determinant(A) < 0) {
        d.at<float>(dim - 1, 0) = -1;
    }
    cv::Mat T = cv::Mat::eye(dim + 1, dim + 1, CV_32F);
    cv::Mat U, S, V;
    cv::SVD::compute(A, S,U, V);

        // the SVD function in opencv differ from scipy .
    int rank = MatrixRank(A);
    if (rank == 0) {
        assert(rank == 0);
    } else if (rank == dim - 1) {
        if (cv::determinant(U) * cv::determinant(V) > 0) {
            T.rowRange(0, dim).colRange(0, dim) = U * V;
        } else {
//           s = d[dim - 1]
//           d[dim - 1] = -1
//           T[:dim, :dim] = np.dot(U, np.dot(np.diag(d), V))
//           d[dim - 1] = s
            int s = d.at<float>(dim - 1, 0) = -1;
            d.at<float>(dim - 1, 0) = -1;

            T.rowRange(0, dim).colRange(0, dim) = U * V;
            cv::Mat diag_ = cv::Mat::diag(d);
            cv::Mat twp = diag_*V; //np.dot(np.diag(d), V.T)
            cv::Mat B = cv::Mat::zeros(3, 3, CV_8UC1);
            cv::Mat C = B.diag(0);
            T.rowRange(0, dim).colRange(0, dim) = U* twp;
            d.at<float>(dim - 1, 0) = s;
        }
    }
    else{
        cv::Mat diag_ = cv::Mat::diag(d);
        cv::Mat twp = diag_*V.t(); //np.dot(np.diag(d), V.T)
        cv::Mat res = U* twp; // U
        T.rowRange(0, dim).colRange(0, dim) = -U.t()* twp;
    }
    cv::Mat var_ = varAxis0(src_demean);
    float val = cv::sum(var_).val[0];
    cv::Mat res;
    cv::multiply(d,S,res);
    float scale =  1.0/val*cv::sum(res).val[0];
    T.rowRange(0, dim).colRange(0, dim) = - T.rowRange(0, dim).colRange(0, dim).t();
    cv::Mat  temp1 = T.rowRange(0, dim).colRange(0, dim); // T[:dim, :dim]
    cv::Mat  temp2 = src_mean.t(); //src_mean.T
    cv::Mat  temp3 = temp1*temp2; // np.dot(T[:dim, :dim], src_mean.T)
    cv::Mat temp4 = scale*temp3;
    T.rowRange(0, dim).colRange(dim, dim+1)=  -(temp4 - dst_mean.t()) ;
    T.rowRange(0, dim).colRange(0, dim) *= scale;
    return T;
}

Mat Mat_SimilarityTransform(Mat src, Mat dst)
{
    cv::Mat trans = similarTransform(*src, *dst);
    return new cv::Mat(std::move(trans));
}
*/
