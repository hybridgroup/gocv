//go:build !gocv_specific_modules || (gocv_specific_modules && gocv_contrib_xphoto)

#include "xphoto.h"

OpenCVResult Xphoto_ApplyChannelGains(Mat src, Mat dst, float gainB, float gainG, float gainR) {
    try {
        cv::xphoto::applyChannelGains(*src, *dst, gainB, gainG, gainR);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Xphoto_Bm3dDenoising_Step(Mat src, Mat dststep1, Mat dststep2) {
    try {
        cv::xphoto::bm3dDenoising(
            *src, *dststep1, *dststep2,
            1, 4,
            16, 2500,
            400, 8,
            1, 2.0f,
            cv::NORM_L2, cv::xphoto::BM3D_STEPALL,
            cv::xphoto::HAAR
        );
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}


OpenCVResult Xphoto_Bm3dDenoising_Step_WithParams(
    Mat src, Mat dststep1, Mat dststep2,
    float h, int templateWindowSize,
    int searchWindowSize, int blockMatchingStep1,
    int blockMatchingStep2, int groupSize,
    int slidingStep, float beta,
    int normType, int step,
    int transformType
) {
    try {
        cv::xphoto::bm3dDenoising(
            *src, *dststep1, *dststep2,
            h, templateWindowSize,
            searchWindowSize, blockMatchingStep1,
            blockMatchingStep2, groupSize,
            slidingStep, beta,
            normType, step,
            transformType
        );    
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Xphoto_Bm3dDenoising(Mat src, Mat dst) {
    try {
        cv::xphoto::bm3dDenoising(*src, *dst,
            1, 4,
            16, 2500,
            400, 8,
            1, 2.0f,
            cv::NORM_L2, cv::xphoto::BM3D_STEPALL,
            cv::xphoto::HAAR
           );
           OpenCVResult result = {0, NULL};
           return result;
    } catch(const cv::Exception& e) {
        OpenCVResult result = {e.code, e.what()};
        return result;
    }
}

OpenCVResult Xphoto_Bm3dDenoising_WithParams(
    Mat src, Mat dst, float h, int templateWindowSize,
    int searchWindowSize, int blockMatchingStep1,
    int blockMatchingStep2, int groupSize,
    int slidingStep, float beta,
    int normType, int step,
    int transformType
) {
    try {
        cv::xphoto::bm3dDenoising(*src, *dst, h, templateWindowSize,
            searchWindowSize, blockMatchingStep1,
            blockMatchingStep2, groupSize,
            slidingStep, beta,
            normType, step,
            transformType
           );
           OpenCVResult result = {0, NULL};
           return result;
    } catch(const cv::Exception& e) {
        OpenCVResult result = {e.code, e.what()};
        return result;
    }
}

// ----------------------- GrayworldWB -----------------------

GrayworldWB GrayworldWB_Create() {
    try {
        return new cv::Ptr<cv::xphoto::GrayworldWB>(cv::xphoto::createGrayworldWB());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

void GrayworldWB_Close(GrayworldWB b) {
    delete b;
}

void GrayworldWB_SetSaturationThreshold(GrayworldWB b, float saturationThreshold) {
    try {
        (*b)->setSaturationThreshold(saturationThreshold);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

float GrayworldWB_GetSaturationThreshold(GrayworldWB b) {
    try {
        return (*b)->getSaturationThreshold();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0.0;
    }
}

OpenCVResult GrayworldWB_BalanceWhite(GrayworldWB b, Mat src, Mat dst) {
    try {
        (*b)->balanceWhite(*src, *dst);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

// ----------------------- LearningBasedWB -----------------------

LearningBasedWB LearningBasedWB_Create() {
    try {
        return new cv::Ptr<cv::xphoto::LearningBasedWB>(cv::xphoto::createLearningBasedWB());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

LearningBasedWB LearningBasedWB_CreateWithParams(const char* pathmodel) {
    try {
        cv::String path(pathmodel);
        return new cv::Ptr<cv::xphoto::LearningBasedWB>(cv::xphoto::createLearningBasedWB(path));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

void LearningBasedWB_Close(LearningBasedWB b) {
    delete b;
}

OpenCVResult LearningBasedWB_ExtractSimpleFeatures(LearningBasedWB b, Mat src, Mat dst) {
    try {
        (*b)->extractSimpleFeatures(*src, *dst);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

int LearningBasedWB_GetHistBinNum(LearningBasedWB b)  {
    try {
        return (*b)->getHistBinNum();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0;
    }
}

int LearningBasedWB_GetRangeMaxVal(LearningBasedWB b)  {
    try {
        return (*b)->getRangeMaxVal();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0;
    }
}

float LearningBasedWB_GetSaturationThreshold(LearningBasedWB b)  {
    try {
        return (*b)->getSaturationThreshold();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0.0;
    }
}

void LearningBasedWB_SetHistBinNum(LearningBasedWB b, int val)  {
    try {
        (*b)->setHistBinNum(val);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void LearningBasedWB_SetRangeMaxVal(LearningBasedWB b, int val) {
    try {
        (*b)->setRangeMaxVal(val);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void LearningBasedWB_SetSaturationThreshold(LearningBasedWB b, float val) {
    try {
        (*b)->setSaturationThreshold(val);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

OpenCVResult LearningBasedWB_BalanceWhite(LearningBasedWB b, Mat src, Mat dst) {
    try {
        (*b)->balanceWhite(*src, *dst);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

// ----------------------- SimpleWB -----------------------


SimpleWB SimpleWB_Create() {
    try {
        return new cv::Ptr<cv::xphoto::SimpleWB>(cv::xphoto::createSimpleWB());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

void SimpleWB_Close(SimpleWB b) {
    delete b;
}

//  Input image range maximum value.
float SimpleWB_GetInputMax(SimpleWB b) {
    try {
        return (*b)->getInputMax();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0.0;
    }
}

//  Input image range minimum value.
float SimpleWB_GetInputMin(SimpleWB b) {
    try {
        return (*b)->getInputMin();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0.0;
    }
}

//  Output image range maximum value.
float SimpleWB_GetOutputMax(SimpleWB b) {
    try {
        return (*b)->getOutputMax();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0.0;
    }
}

//  Output image range minimum value.
float SimpleWB_GetOutputMin(SimpleWB b) {
    try {
        return (*b)->getOutputMin();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0.0;
    }
}

//  Percent of top/bottom values to ignore.
float SimpleWB_GetP(SimpleWB b) {
    try {
        return (*b)->getP();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0.0;
    }
}

//  Input image range maximum value.
void SimpleWB_SetInputMax(SimpleWB b, float val) {
    try {
        (*b)->setInputMax(val);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

//  Input image range minimum value.
void SimpleWB_SetInputMin(SimpleWB b, float val) {
    try {
        (*b)->setInputMin(val);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

//  Output image range maximum value.
void SimpleWB_SetOutputMax(SimpleWB b, float val) {
    try {
        (*b)->setOutputMax(val);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

//  Output image range minimum value.
void SimpleWB_SetOutputMin(SimpleWB b, float val) {
    try {
        (*b)->setOutputMin(val);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

//  Percent of top/bottom values to ignore.
void SimpleWB_SetP(SimpleWB b, float val) {
    try {
        (*b)->setP(val);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

OpenCVResult SimpleWB_BalanceWhite(SimpleWB b, Mat src, Mat dst) {
    try {
        (*b)->balanceWhite(*src, *dst);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

// -------------------- TonemapDurand --------------------

// Creates TonemapDurand object. More...
TonemapDurand TonemapDurand_Create() {
    try {
        return new cv::Ptr<cv::xphoto::TonemapDurand>(cv::xphoto::createTonemapDurand(1.0f, 4.0f, 1.0f, 2.0f, 2.0f));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

TonemapDurand TonemapDurand_CreateWithParams(float gamma, float contrast, float saturation,
        float sigma_color, float sigma_space) {

    try {
        return new cv::Ptr<cv::xphoto::TonemapDurand>(cv::xphoto::createTonemapDurand(gamma, contrast,
            saturation, sigma_color, sigma_space));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

void TonemapDurand_Close(TonemapDurand b) {
    delete b;
}

float TonemapDurand_GetContrast(TonemapDurand b) {
    try {
        return (*b)->getContrast();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0.0;
    }
}

float TonemapDurand_GetSaturation(TonemapDurand b) {
    try {
        return (*b)->getSaturation();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0.0;
    }
}

float TonemapDurand_GetSigmaColor(TonemapDurand b) {
    try {
        return (*b)->getSigmaColor();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0.0;
    }
}

float TonemapDurand_GetSigmaSpace(TonemapDurand b) {
    try {
        return (*b)->getSigmaSpace();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0.0;
    }
}

void TonemapDurand_SetContrast(TonemapDurand b, float contrast) {
    try {
        (*b)->setContrast(contrast);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void TonemapDurand_SetSaturation(TonemapDurand b, float saturation) {
    try {
        (*b)->setSaturation(saturation);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void TonemapDurand_SetSigmaColor(TonemapDurand b, float sigma_color) {
    try {
        (*b)->setSigmaColor(sigma_color);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void TonemapDurand_SetSigmaSpace(TonemapDurand b, float sigma_space) {
    try {
        (*b)->setSigmaSpace(sigma_space);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

float TonemapDurand_GetGamma(TonemapDurand b) {
    try {
        return (*b)->getGamma();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0.0;
    }
}

void TonemapDurand_SetGamma(TonemapDurand b, float gamma) {
    try {
        (*b)->setGamma(gamma);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

OpenCVResult TonemapDurand_Process(TonemapDurand b, Mat src, Mat dst) {
    try {
        (*b)->process(*src, *dst);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

// -------------------- cv::xphoto::Inpaint --------------------

OpenCVResult Inpaint(Mat src, Mat mask, Mat dst, int algorithmType) {
    try {
        cv::xphoto::inpaint(*src, *mask, *dst, algorithmType);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult OilPaintingWithParams(Mat src, Mat dst, int size, int dynRatio, int code) {
    try {
        cv::xphoto::oilPainting(*src, *dst, size, dynRatio, code);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult OilPainting(Mat src, Mat dst, int size, int dynRatio) {
    try {
        cv::xphoto::oilPainting(*src, *dst, size, dynRatio);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}
