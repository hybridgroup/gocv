#include "dnn.h"

Net Net_ReadNet(const char* model, const char* config) {
    Net n = new cv::dnn::Net(cv::dnn::readNet(model, config));
    return n;
}

Net Net_ReadNetFromCaffe(const char* prototxt, const char* caffeModel) {
    Net n = new cv::dnn::Net(cv::dnn::readNetFromCaffe(prototxt, caffeModel));
    return n;
}

Net Net_ReadNetFromTensorflow(const char* model) {
    Net n = new cv::dnn::Net(cv::dnn::readNetFromTensorflow(model));
    return n;
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

void Net_ForwardLayers(Net net, struct Mats outputBlobs, struct CStrings outBlobNames) {
    std::vector< cv::Mat > blobs;
    for (int i = 0; i < outputBlobs.length; ++i) {
        blobs.push_back(*outputBlobs.mats[i]);
    }

    std::vector< cv::String > names;
    for (int i = 0; i < outBlobNames.length; ++i) {
        names.push_back(cv::String(outBlobNames.strs[i]));
    }
    net->forward(blobs, names);
}

void Net_SetPreferableBackend(Net net, int backend) {
    net->setPreferableBackend(backend);
}

void Net_SetPreferableTarget(Net net, int target) {
    net->setPreferableTarget(target);
}

Mat Net_BlobFromImage(Mat image, double scalefactor, Size size, Scalar mean, bool swapRB,
                      bool crop) {
    cv::Size sz(size.width, size.height);
    cv::Scalar cm = cv::Scalar(mean.val1, mean.val2, mean.val3, mean.val4);

    // TODO: handle different version signatures of this function v2 vs v3.
    return new cv::Mat(cv::dnn::blobFromImage(*image, scalefactor, sz, cm, swapRB, crop));
}

Mat Net_GetBlobChannel(Mat blob, int imgidx, int chnidx) {
    size_t w = blob->size[3];
    size_t h = blob->size[2];
    return new cv::Mat(h, w, CV_32F, blob->ptr<float>(imgidx, chnidx));
}

Scalar Net_GetBlobSize(Mat blob) {
    Scalar scal = Scalar();
    scal.val1 = blob->size[0];
    scal.val2 = blob->size[1];
    scal.val3 = blob->size[2];
    scal.val4 = blob->size[3];
    return scal;
}
