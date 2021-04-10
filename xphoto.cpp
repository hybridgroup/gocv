#include "xphoto.h"

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
