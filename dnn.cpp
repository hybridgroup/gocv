#include "dnn.h"

Net Net_ReadNetFromCaffe(const char* prototxt, const char* caffeModel) {
    try {
        return new cv::dnn::Net(cv::dnn::readNetFromCaffe(prototxt, caffeModel));
    } catch(...) {}
    return 0;
}

Net Net_ReadNetFromDarknet(const char* prototxt, const char* model) {
    try {
        return new cv::dnn::Net(cv::dnn::readNetFromDarknet(prototxt, model));
    } catch(...) {}
    return 0;
}

Net Net_ReadNetFromTensorflow(const char* model) {
    try {
        return new cv::dnn::Net(cv::dnn::readNetFromTensorflow(model));
    } catch(...) {}
    return 0;
}

Net Net_ReadNetFromTensorflowProto(const char* model, const char* prototxt) {
    try {
        return new cv::dnn::Net(cv::dnn::readNetFromTensorflow(model, prototxt));
    } catch(...) {}
    return 0;
}

Net Net_ReadNetFromTorch(const char* model) {
    try {
        return new cv::dnn::Net(cv::dnn::readNetFromTorch(model));
    } catch(...) {}
    return 0;
}

void Net_Close(Net net) {
    delete net;
}

bool Net_Empty(Net net) {
    return net->empty();
}

void Net_SetInput(Net net, Mat blob, const char* name) {
    net->setInput(*blob, name);
}

Mat Net_Forward(Net net, const char* outputName) {
    return new cv::Mat(net->forward(outputName));
}

Mat Net_BlobFromImage(Mat image, double scalefactor, Size size, Scalar mean, bool swapRB, bool crop) {
    cv::Size sz(size.width, size.height);
    cv::Scalar cm = cv::Scalar(mean.val1, mean.val2, mean.val3, mean.val4);

    // TODO: handle different version signatures of this function v2 vs v3.
    return new cv::Mat(cv::dnn::blobFromImage(*image, scalefactor, sz, cm, swapRB));
}
