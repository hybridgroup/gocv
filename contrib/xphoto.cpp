#include "xphoto.h"

void Xphoto_ApplyChannelGains(Mat src, Mat dst, float gainB, float gainG, float gainR) {
    cv::xphoto::applyChannelGains(*src, *dst, gainB, gainG, gainR);
}

void Xphoto_Bm3dDenoising_Step(Mat src, Mat dststep1, Mat dststep2) {

    cv::xphoto::bm3dDenoising(
        *src, *dststep1, *dststep2,
        1, 4,
        16, 2500,
        400, 8,
        1, 2.0f,
        cv::NORM_L2, cv::xphoto::BM3D_STEPALL,
        cv::xphoto::HAAR
    );
}


void Xphoto_Bm3dDenoising_Step_WithParams(
    Mat src, Mat dststep1, Mat dststep2,
    float h, int templateWindowSize,
    int searchWindowSize, int blockMatchingStep1,
    int blockMatchingStep2, int groupSize,
    int slidingStep, float beta,
    int normType, int step,
    int transformType
) {

    cv::xphoto::bm3dDenoising(
        *src, *dststep1, *dststep2,
        h, templateWindowSize,
        searchWindowSize, blockMatchingStep1,
        blockMatchingStep2, groupSize,
        slidingStep, beta,
        normType, step,
        transformType
    );
}

void Xphoto_Bm3dDenoising(Mat src, Mat dst) {

    cv::xphoto::bm3dDenoising(*src, *dst,
                              1, 4,
                              16, 2500,
                              400, 8,
                              1, 2.0f,
                              cv::NORM_L2, cv::xphoto::BM3D_STEPALL,
                              cv::xphoto::HAAR
                             );

}

void Xphoto_Bm3dDenoising_WithParams(
    Mat src, Mat dst, float h, int templateWindowSize,
    int searchWindowSize, int blockMatchingStep1,
    int blockMatchingStep2, int groupSize,
    int slidingStep, float beta,
    int normType, int step,
    int transformType
) {

    cv::xphoto::bm3dDenoising(*src, *dst, h, templateWindowSize,
                              searchWindowSize, blockMatchingStep1,
                              blockMatchingStep2, groupSize,
                              slidingStep, beta,
                              normType, step,
                              transformType
                             );

}




// ----------------------- GrayworldWB -----------------------

GrayworldWB GrayworldWB_Create() {
    return new cv::Ptr<cv::xphoto::GrayworldWB>(cv::xphoto::createGrayworldWB());
}

void GrayworldWB_Close(GrayworldWB b) {
    delete b;
}

void GrayworldWB_SetSaturationThreshold(GrayworldWB b, float saturationThreshold) {
    (*b)->setSaturationThreshold(saturationThreshold);
}

float GrayworldWB_GetSaturationThreshold(GrayworldWB b) {
    return (*b)->getSaturationThreshold();
}

void GrayworldWB_BalanceWhite(GrayworldWB b, Mat src, Mat dst) {
    (*b)->balanceWhite(*src, *dst);
}

// ----------------------- LearningBasedWB -----------------------

LearningBasedWB LearningBasedWB_Create() {
    return new cv::Ptr<cv::xphoto::LearningBasedWB>(cv::xphoto::createLearningBasedWB());
}

LearningBasedWB LearningBasedWB_CreateWithParams(const char* pathmodel) {
    cv::String path(pathmodel);
    return new cv::Ptr<cv::xphoto::LearningBasedWB>(cv::xphoto::createLearningBasedWB(path));
}

void LearningBasedWB_Close(LearningBasedWB b) {
    delete b;
}

void LearningBasedWB_ExtractSimpleFeatures(LearningBasedWB b, Mat src, Mat dst) {
    (*b)->extractSimpleFeatures(*src, *dst);
}

int LearningBasedWB_GetHistBinNum(LearningBasedWB b)  {
    return (*b)->getHistBinNum();
}

int LearningBasedWB_GetRangeMaxVal(LearningBasedWB b)  {
    return (*b)->getRangeMaxVal();
}

float LearningBasedWB_GetSaturationThreshold(LearningBasedWB b)  {
    return (*b)->getSaturationThreshold();
}

void LearningBasedWB_SetHistBinNum(LearningBasedWB b, int val)  {
    (*b)->setHistBinNum(val);
}

void LearningBasedWB_SetRangeMaxVal(LearningBasedWB b, int val) {
    (*b)->setRangeMaxVal(val);
}

void LearningBasedWB_SetSaturationThreshold(LearningBasedWB b, float val) {
    (*b)->setSaturationThreshold(val);
}

void LearningBasedWB_BalanceWhite(LearningBasedWB b, Mat src, Mat dst) {
    (*b)->balanceWhite(*src, *dst);
}

// ----------------------- SimpleWB -----------------------


SimpleWB SimpleWB_Create() {
    return new cv::Ptr<cv::xphoto::SimpleWB>(cv::xphoto::createSimpleWB());
}

void SimpleWB_Close(SimpleWB b) {
    delete b;
}

//  Input image range maximum value.
float SimpleWB_GetInputMax(SimpleWB b) {
    return (*b)->getInputMax();
}

//  Input image range minimum value.
float SimpleWB_GetInputMin(SimpleWB b) {
    return (*b)->getInputMin();
}

//  Output image range maximum value.
float SimpleWB_GetOutputMax(SimpleWB b) {
    return (*b)->getOutputMax();
}

//  Output image range minimum value.
float SimpleWB_GetOutputMin(SimpleWB b) {
    return (*b)->getOutputMin();
}

//  Percent of top/bottom values to ignore.
float SimpleWB_GetP(SimpleWB b) {
    return (*b)->getP();
}

//  Input image range maximum value.
void SimpleWB_SetInputMax(SimpleWB b, float val) {
    return (*b)->setInputMax(val);
}

//  Input image range minimum value.
void SimpleWB_SetInputMin(SimpleWB b, float val) {
    return (*b)->setInputMin(val);
}

//  Output image range maximum value.
void SimpleWB_SetOutputMax(SimpleWB b, float val) {
    return (*b)->setOutputMax(val);
}

//  Output image range minimum value.
void SimpleWB_SetOutputMin(SimpleWB b, float val) {
    return (*b)->setOutputMin(val);
}

//  Percent of top/bottom values to ignore.
void SimpleWB_SetP(SimpleWB b, float val) {
    return (*b)->setP(val);
}

void SimpleWB_BalanceWhite(SimpleWB b, Mat src, Mat dst) {
    (*b)->balanceWhite(*src, *dst);
}


// -------------------- TonemapDurand --------------------

// Creates TonemapDurand object. More...
TonemapDurand TonemapDurand_Create() {
    return new cv::Ptr<cv::xphoto::TonemapDurand>(cv::xphoto::createTonemapDurand(1.0f, 4.0f, 1.0f,
            2.0f, 2.0f));
}

TonemapDurand TonemapDurand_CreateWithParams(float gamma, float contrast, float saturation,
        float sigma_color, float sigma_space) {
    return new cv::Ptr<cv::xphoto::TonemapDurand>(cv::xphoto::createTonemapDurand(gamma, contrast,
            saturation, sigma_color, sigma_space));
}

void TonemapDurand_Close(TonemapDurand b) {
    delete b;
}

float TonemapDurand_GetContrast(TonemapDurand b) {
    return (*b)->getContrast();
}
float TonemapDurand_GetSaturation(TonemapDurand b) {
    return (*b)->getSaturation();
}
float TonemapDurand_GetSigmaColor(TonemapDurand b) {
    return (*b)->getSigmaColor();
}
float TonemapDurand_GetSigmaSpace(TonemapDurand b) {
    return (*b)->getSigmaSpace();
}

void TonemapDurand_SetContrast(TonemapDurand b, float contrast) {
    return (*b)->setContrast(contrast);
}
void TonemapDurand_SetSaturation(TonemapDurand b, float saturation) {
    return (*b)->setSaturation(saturation);
}
void TonemapDurand_SetSigmaColor(TonemapDurand b, float sigma_color) {
    return (*b)->setSigmaColor(sigma_color);
}
void TonemapDurand_SetSigmaSpace(TonemapDurand b, float sigma_space) {
    return (*b)->setSigmaSpace(sigma_space);
}

float TonemapDurand_GetGamma(TonemapDurand b) {
    return (*b)->getGamma();
}

void TonemapDurand_SetGamma(TonemapDurand b, float gamma) {
    (*b)->setGamma(gamma);
}

void TonemapDurand_Process(TonemapDurand b, Mat src, Mat dst) {
    (*b)->process(*src, *dst);
}

// -------------------- cv::xphoto::Inpaint --------------------

void Inpaint(Mat src, Mat mask, Mat dst, int algorithmType) {
    cv::xphoto::inpaint(*src, *mask, *dst, algorithmType);
}

void OilPaintingWithParams(Mat src, Mat dst, int size, int dynRatio, int code) {
    cv::xphoto::oilPainting(*src, *dst, size, dynRatio, code);
}

void OilPainting(Mat src, Mat dst, int size, int dynRatio) {
    cv::xphoto::oilPainting(*src, *dst, size, dynRatio);
}
