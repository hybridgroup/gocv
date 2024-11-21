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

// Mat_NewWithSizes creates a new Mat with specific dimension sizes and number of channels.
Mat Mat_NewWithSizes(struct IntVector sizes, int type) {
	std::vector<int> sizess;
    for (int i = 0; i < sizes.length; ++i) {
        sizess.push_back(sizes.val[i]);
    }
    return new cv::Mat(sizess, type);
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

// Mat_NewWithSizesFromScalar creates multidimensional Mat from a scalar
Mat Mat_NewWithSizesFromScalar(IntVector sizes, int type, Scalar ar) {
    std::vector<int> _sizes;
    for (int i = 0, *v = sizes.val; i < sizes.length; ++v, ++i) {
        _sizes.push_back(*v);
    }

    cv::Scalar c = cv::Scalar(ar.val1, ar.val2, ar.val3, ar.val4);
    return new cv::Mat(_sizes, type, c);
}

// Mat_NewWithSizesFromBytes creates multidimensional Mat from a bytes
Mat Mat_NewWithSizesFromBytes(IntVector sizes, int type, struct ByteArray buf) {
    std::vector<int> _sizes;
    for (int i = 0, *v = sizes.val; i < sizes.length; ++v, ++i) {
        _sizes.push_back(*v);
    }

    return new cv::Mat(_sizes, type, buf.data);
}

Mat Mat_NewFromPoint2fVector(Point2fVector pfv, bool copy_data) {
    return new cv::Mat(*pfv, copy_data);
}

Mat Eye(int rows, int cols, int type) {
    cv::Mat* mat = new cv::Mat(rows, cols, type);
    *mat = cv::Mat::eye(rows, cols, type);
    return mat;
}

Mat Zeros(int rows, int cols, int type) {
    cv::Mat* mat = new cv::Mat(rows, cols, type);
    *mat = cv::Mat::zeros(rows, cols, type);
    return mat;
}

Mat Ones(int rows, int cols, int type) {
    cv::Mat* mat = new cv::Mat(rows, cols, type);
    *mat = cv::Mat::ones(rows, cols, type);
    return mat;
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

// Mat_IsContinuous tests if a Mat is continuous
bool Mat_IsContinuous(Mat m) {
    return m->isContinuous();
}

void Mat_Inv(Mat m) {
    m->inv();
}

Mat Mat_Col(Mat m, int c) {
    return new cv::Mat(m->col(c));
}

Mat Mat_Row(Mat m, int r) {
    return new cv::Mat(m->row(r));
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

void Mat_ConvertToWithParams(Mat m, Mat dst, int type, float alpha, float beta) {
    m->convertTo(*dst, type, alpha, beta);
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
Scalar Mat_Mean(Mat m) {
    cv::Scalar c = cv::mean(*m);
    Scalar scal = Scalar();
    scal.val1 = c.val[0];
    scal.val2 = c.val[1];
    scal.val3 = c.val[2];
    scal.val4 = c.val[3];
    return scal;
}

// Mat_MeanWithMask calculates the mean value M of array elements,
// independently for each channel, and returns it as Scalar vector
// while applying the mask.

Scalar Mat_MeanWithMask(Mat m, Mat mask){
    cv::Scalar c = cv::mean(*m, *mask);
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

int Mat_ElemSize(Mat m){
    return m->elemSize();
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

Mat Mat_MultiplyMatrix(Mat x, Mat y) {
    return new cv::Mat((*x) * (*y));
}

Mat Mat_T(Mat x) {
    return new cv::Mat(x->t());
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

void Mat_PCABackProject(Mat data, Mat mean, Mat eigenvectors, Mat result) {
    cv::PCABackProject(*data, *mean, *eigenvectors, *result);
}

void Mat_PCACompute(Mat src, Mat mean, Mat eigenvectors, Mat eigenvalues, int maxComponents) {
    cv::PCACompute(*src, *mean, *eigenvectors, *eigenvalues, maxComponents);
}

void Mat_PCAProject(Mat data, Mat mean, Mat eigenvectors, Mat result) {
    cv::PCAProject(*data, *mean, *eigenvectors, *result);
}

double PSNR(Mat src1, Mat src2) {
    return cv::PSNR(*src1, *src2);
}

void SVBackSubst(Mat w, Mat u, Mat vt, Mat rhs, Mat dst) {
    cv::SVBackSubst(*w, *u, *vt, *rhs, *dst);
}

void SVDecomp(Mat src, Mat w, Mat u, Mat vt) {
    cv::SVDecomp(*src, *w, *u, *vt);
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

double KMeans(Mat data, int k, Mat bestLabels, TermCriteria criteria, int attempts, int flags, Mat centers) {
    double ret = cv::kmeans(*data, k, *bestLabels, *criteria, attempts, flags, *centers);
    return ret;
}

double KMeansPoints(PointVector points, int k, Mat bestLabels, TermCriteria criteria, int attempts, int flags, Mat centers) {
    std::vector<cv::Point2f> pts;
    copyPointVectorToPoint2fVector(points, &pts);
    double ret = cv::kmeans(pts, k, *bestLabels, *criteria, attempts, flags, *centers);
    return ret;
}

void Mat_Log(Mat src, Mat dst) {
    cv::log(*src, *dst);
}

void Mat_Magnitude(Mat x, Mat y, Mat magnitude) {
    cv::magnitude(*x, *y, *magnitude);
}

double Mat_Mahalanobis(Mat v1, Mat v2, Mat icovar) {
    return cv::Mahalanobis(*v1, *v2, *icovar);
}

void MulTransposed(Mat src, Mat dest, bool ata) {
    cv::mulTransposed(*src, *dest, ata);
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

void Mat_MinMaxLocWithMask(Mat m, double* minVal, double* maxVal, Point* minLoc, Point* maxLoc, Mat mask) {
    cv::Point cMinLoc;
    cv::Point cMaxLoc;
    cv::minMaxLoc(*m, minVal, maxVal, &cMinLoc, &cMaxLoc, *mask);

    minLoc->x = cMinLoc.x;
    minLoc->y = cMinLoc.y;
    maxLoc->x = cMaxLoc.x;
    maxLoc->y = cMaxLoc.y;
}

void Mat_MixChannels(struct Mats src, struct Mats dst, struct IntVector fromTo) {
    std::vector<cv::Mat> srcMats;

    for (int i = 0; i < src.length; ++i) {
        srcMats.push_back(*src.mats[i]);
    }

    std::vector<cv::Mat> dstMats;

    for (int i = 0; i < dst.length; ++i) {
        dstMats.push_back(*dst.mats[i]);
    }

    std::vector<int> fromTos;

    for (int i = 0; i < fromTo.length; ++i) {
        fromTos.push_back(fromTo.val[i]);
    }

    cv::mixChannels(srcMats, dstMats, fromTos);
}

void Mat_MulSpectrums(Mat a, Mat b, Mat c, int flags) {
    cv::mulSpectrums(*a, *b, *c, flags);
}

void Mat_Multiply(Mat src1, Mat src2, Mat dst) {
    cv::multiply(*src1, *src2, *dst);
}

void Mat_MultiplyWithParams(Mat src1, Mat src2, Mat dst, double scale, int dtype) {
    cv::multiply(*src1, *src2, *dst, scale, dtype);
}

void Mat_Normalize(Mat src, Mat dst, double alpha, double beta, int typ) {
    cv::normalize(*src, *dst, alpha, beta, typ);
}

double Norm(Mat src1, int normType) {
    return cv::norm(*src1, normType);
}

double NormWithMats(Mat src1, Mat src2, int normType) {
    return cv::norm(*src1, *src2, normType);
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

void Mat_ReduceArgMax(Mat src, Mat dst, int axis, bool lastIndex) {
    cv::reduceArgMax(*src, *dst, axis, lastIndex);
}

void Mat_ReduceArgMin(Mat src, Mat dst, int axis, bool lastIndex) {
    cv::reduceArgMin(*src, *dst, axis, lastIndex);
}


void Mat_Repeat(Mat src, int nY, int nX, Mat dst) {
    cv::repeat(*src, nY, nX, *dst);
}

void Mat_ScaleAdd(Mat src1, double alpha, Mat src2, Mat dst) {
    cv::scaleAdd(*src1, alpha, *src2, *dst);
}

void Mat_SetIdentity(Mat src, double scalar) {
    cv::setIdentity(*src, scalar);
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

void Mat_TransposeND(Mat src, struct IntVector order, Mat dst) {
    std::vector<int> _order;
    for (int i = 0, *v = order.val; i < order.length; ++v, ++i) {
        _order.push_back(*v);
    }

    cv::transposeND(*src, _order, *dst);
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

void CStrings_Close(struct CStrings cstrs) {
    for ( int i = 0; i < cstrs.length; i++ ) {
        delete [] cstrs.strs[i];
    }
    delete [] cstrs.strs;
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

void Points2f_Close(Points2f ps) {
    for (size_t i = 0; i < ps.length; i++) {
        Point2f_Close(ps.points[i]);
    }

    delete[] ps.points;
}

void Point2f_Close(Point2f p) {}

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

Mat Mat_rowRange(Mat m,int startrow,int endrow) {
    return new cv::Mat(m->rowRange(startrow,endrow));
}

Mat Mat_colRange(Mat m,int startrow,int endrow) {
    return new cv::Mat(m->colRange(startrow,endrow));
}

PointVector PointVector_New() {
    return new std::vector< cv::Point >;
}

PointVector PointVector_NewFromPoints(Contour points) {
    std::vector<cv::Point>* cntr = new std::vector<cv::Point>;

    for (size_t i = 0; i < points.length; i++) {
        cntr->push_back(cv::Point(points.points[i].x, points.points[i].y));
    }

    return cntr;
}

PointVector PointVector_NewFromMat(Mat mat) {
    std::vector<cv::Point>* pts = new std::vector<cv::Point>;
    *pts = (std::vector<cv::Point>) *mat;
    return pts;
}

Point PointVector_At(PointVector pv, int idx) {
    cv::Point p = pv->at(idx);
    return Point{.x = p.x, .y = p.y};
}

void PointVector_Append(PointVector pv, Point p) {
    pv->push_back(cv::Point(p.x, p.y));
}

int PointVector_Size(PointVector p) {
    return p->size();
}

void PointVector_Close(PointVector p) {
    p->clear();
    delete p;
}

PointsVector PointsVector_New() {
    return new std::vector< std::vector< cv::Point > >;
}

PointsVector PointsVector_NewFromPoints(Contours points) {
    std::vector< std::vector< cv::Point > >* pv = new std::vector< std::vector< cv::Point > >;

    for (size_t i = 0; i < points.length; i++) {
        Contour contour = points.contours[i];

        std::vector<cv::Point> cntr;

        for (size_t i = 0; i < contour.length; i++) {
            cntr.push_back(cv::Point(contour.points[i].x, contour.points[i].y));
        }

        pv->push_back(cntr);
    }

    return pv;
}

int PointsVector_Size(PointsVector ps) {
    return ps->size();
}

PointVector PointsVector_At(PointsVector ps, int idx) {
    std::vector< cv::Point >* p = &(ps->at(idx));
    return p;
}

void PointsVector_Append(PointsVector psv, PointVector pv) {
    psv->push_back(*pv);
}

void PointsVector_Close(PointsVector ps) {
    ps->clear();
    delete ps;
}

Point2fVector Point2fVector_New() {
    return new std::vector< cv::Point2f >;
}

Point2fVector Point2fVector_NewFromPoints(Contour2f points) {
    std::vector<cv::Point2f>* cntr = new std::vector<cv::Point2f>;

    for (size_t i = 0; i < points.length; i++) {
        cntr->push_back(cv::Point2f(points.points[i].x, points.points[i].y));
    }

    return cntr;
}

Point2fVector Point2fVector_NewFromMat(Mat mat) {
    std::vector<cv::Point2f>* pts = new std::vector<cv::Point2f>;
    *pts = (std::vector<cv::Point2f>) *mat;
    return pts;
}

Point2f Point2fVector_At(Point2fVector pfv, int idx) {
    cv::Point2f p = pfv->at(idx);
    return Point2f{.x = p.x, .y = p.y};
}

int Point2fVector_Size(Point2fVector pfv) {
    return pfv->size();
}

void Point2fVector_Close(Point2fVector pv) {
    pv->clear();
    delete pv;
}


void IntVector_Close(struct IntVector ivec) {
    delete[] ivec.val;
}

RNG TheRNG() {
    return &cv::theRNG();
}

void SetRNGSeed(int seed) {
    cv::setRNGSeed(seed);
}

void RNG_Fill(RNG rng, Mat mat, int distType, double a, double b, bool saturateRange) {
    rng->fill(*mat, distType, a, b, saturateRange);
}

double RNG_Gaussian(RNG rng, double sigma) {
    return rng->gaussian(sigma);
}

unsigned int RNG_Next(RNG rng) {
    return rng->next();
}

void RandN(Mat mat, Scalar mean, Scalar stddev) {
    cv::Scalar m = cv::Scalar(mean.val1, mean.val2, mean.val3, mean.val4);
    cv::Scalar s = cv::Scalar(stddev.val1, stddev.val2, stddev.val3, stddev.val4);
    cv::randn(*mat, m, s);
}

void RandShuffle(Mat mat) {
    cv::randShuffle(*mat);
}

void RandShuffleWithParams(Mat mat, double iterFactor, RNG rng) {
    cv::randShuffle(*mat, iterFactor, rng);
}

void RandU(Mat mat, Scalar low, Scalar high) {
    cv::Scalar l = cv::Scalar(low.val1, low.val2, low.val3, low.val4);
    cv::Scalar h = cv::Scalar(high.val1, high.val2, high.val3, high.val4);
    cv::randn(*mat, l, h);
}

void copyPointVectorToPoint2fVector(PointVector src, Point2fVector dest) {
    for (size_t i = 0; i < src->size(); i++) {
        dest->push_back(cv::Point2f(src->at(i).x, src->at(i).y));
    }
}

void StdByteVectorInitialize(void* data) {
    new (data) std::vector<uchar>();
}

void StdByteVectorFree(void *data) {
    reinterpret_cast<std::vector<uchar> *>(data)->~vector<uchar>();
}

size_t StdByteVectorLen(void *data) {
    return reinterpret_cast<std::vector<uchar> *>(data)->size();
}

uint8_t* StdByteVectorData(void *data) {
    return reinterpret_cast<std::vector<uchar> *>(data)->data();
}

Points2fVector Points2fVector_New(){
    return new std::vector< std::vector< cv::Point2f > >;
}

Points2fVector Points2fVector_NewFromPoints(Contours2f points) {
    Points2fVector pv = Points2fVector_New();
    for(size_t i = 0;i<points.length;i++){
        Contour2f contour2f = points.contours[i];
        Point2fVector cntr = Point2fVector_NewFromPoints(contour2f);
        Points2fVector_Append(pv, cntr);
    }

    return pv;
}

int Points2fVector_Size(Points2fVector ps) {
    return ps->size();
}

Point2fVector Points2fVector_At(Points2fVector ps, int idx) {
    return &(ps->at(idx));
}

void Points2fVector_Append(Points2fVector psv, Point2fVector pv) {
    psv->push_back(*pv);
}

void Points2fVector_Close(Points2fVector ps) {
    ps->clear();
    delete ps;
}

Point3fVector Point3fVector_New() {
    return new std::vector< cv::Point3f >;
}


Point3fVector Point3fVector_NewFromPoints(Contour3f points) {
    std::vector<cv::Point3f> *cntr = new std::vector<cv::Point3f>;
    for(size_t i = 0;i< points.length;i++) {
        cntr->push_back(cv::Point3f(
            points.points[i].x,
            points.points[i].y,
            points.points[i].z
        ));
    }

    return cntr;
}

Point3fVector Point3fVector_NewFromMat(Mat mat) {
    std::vector<cv::Point3f> *pts = new std::vector<cv::Point3f>;
    *pts = (std::vector<cv::Point3f>) *mat;
    return pts;
}

Point3f Point3fVector_At(Point3fVector pfv, int idx) {
    cv::Point3f p = pfv->at(idx);
    return Point3f{
        .x = p.x,
        .y = p.y,
        .z = p.z
    };
}

void Point3fVector_Append(Point3fVector pfv, Point3f point) {
    pfv->push_back(cv::Point3f(point.x, point.y, point.z));
}

int Point3fVector_Size(Point3fVector pfv) {
    return pfv->size();
}

void Point3fVector_Close(Point3fVector pv) {
    pv->clear();
    delete pv;
}

Points3fVector Points3fVector_New(){
    return new std::vector< std::vector< cv::Point3f > >;
}

Points3fVector Points3fVector_NewFromPoints(Contours3f points) {
    Points3fVector pv = Points3fVector_New();
    for(size_t i = 0;i<points.length;i++){
        Contour3f contour3f = points.contours[i];
        Point3fVector cntr = Point3fVector_NewFromPoints(contour3f);
        Points3fVector_Append(pv, cntr);
    }

    return pv;
}

int Points3fVector_Size(Points3fVector ps) {
    return ps->size();
}

Point3fVector Points3fVector_At(Points3fVector ps, int idx) {
    return &(ps->at(idx));
}

void Points3fVector_Append(Points3fVector psv, Point3fVector pv) {
    psv->push_back(*pv);
}

void Points3fVector_Close(Points3fVector ps) {
    ps->clear();
    delete ps;
}

void SetNumThreads(int n) {
    cv::setNumThreads(n);
}

int GetNumThreads() {
    return cv::getNumThreads();
}

struct RotatedRect RotatedRect_Create(struct Point2f center, int width, int height, float angle){

    cv::Point2f cvpoint2f = cv::Point2f(center.x, center.y);
    cv::Size2f cvsize2f = cv::Size2f(width, height);

    cv::RotatedRect cvrect = cv::RotatedRect(cvpoint2f, cvsize2f, angle);

    Point* rpts = new Point[4];
    cv::Point2f* pts4 = new cv::Point2f[4];
    cvrect.points(pts4);

    for (size_t j = 0; j < 4; j++) {
        Point pt = {int(lroundf(pts4[j].x)), int(lroundf(pts4[j].y))};
        rpts[j] = pt;
    }

    delete[] pts4;

    cv::Rect bRect = cvrect.boundingRect();
    Rect r = {bRect.x, bRect.y, bRect.width, bRect.height};
    Point centrpt = {int(lroundf(cvrect.center.x)), int(lroundf(cvrect.center.y))};
    Size szsz = {int(lroundf(cvrect.size.width)), int(lroundf(cvrect.size.height))};

    RotatedRect retrect = {(Contour){rpts, 4}, r, centrpt, szsz, cvrect.angle};
    return retrect;
}

struct RotatedRect2f RotatedRect2f_Create(struct Point2f center, float width, float height, float angle){

    cv::Point2f cvpoint2f = cv::Point2f(center.x, center.y);
    cv::Size2f cvsize2f = cv::Size2f(width, height);

    cv::RotatedRect cvrect = cv::RotatedRect(cvpoint2f, cvsize2f, angle);

    Point2f* rpts = new Point2f[4];
    cv::Point2f* pts4 = new cv::Point2f[4];
    cvrect.points(pts4);

    for (size_t j = 0; j < 4; j++) {
        Point2f pt = {pts4[j].x, pts4[j].y};
        rpts[j] = pt;
    }

    delete[] pts4;

    cv::Rect bRect = cvrect.boundingRect();
    Rect r = {bRect.x, bRect.y, bRect.width, bRect.height};
    Point2f centrpt = {cvrect.center.x, cvrect.center.y};
    Size2f szsz = {cvrect.size.width, cvrect.size.height};

    RotatedRect2f retrect = {(Contour2f){rpts, 4}, r, centrpt, szsz, cvrect.angle};
    return retrect;
}
