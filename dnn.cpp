#include "dnn.h"

Net Net_ReadNet(const char* model, const char* config) {
    Net n = new cv::dnn::Net(cv::dnn::readNet(model, config));
    return n;
}

Net Net_ReadNetBytes(const char* framework, struct ByteArray model, struct ByteArray config) {
    std::vector<uchar> modelv(model.data, model.data + model.length);
    std::vector<uchar> configv(config.data, config.data + config.length);
    Net n = new cv::dnn::Net(cv::dnn::readNet(framework, modelv, configv));
    return n;
}

Net Net_ReadNetFromCaffe(const char* prototxt, const char* caffeModel) {
    Net n = new cv::dnn::Net(cv::dnn::readNetFromCaffe(prototxt, caffeModel));
    return n;
}

Net Net_ReadNetFromCaffeBytes(struct ByteArray prototxt, struct ByteArray caffeModel) {
    Net n = new cv::dnn::Net(cv::dnn::readNetFromCaffe(prototxt.data, prototxt.length,
                            caffeModel.data, caffeModel.length));
    return n;
}

Net Net_ReadNetFromTensorflow(const char* model) {
    Net n = new cv::dnn::Net(cv::dnn::readNetFromTensorflow(model));
    return n;
}

Net Net_ReadNetFromTensorflowBytes(struct ByteArray model) {
    Net n = new cv::dnn::Net(cv::dnn::readNetFromTensorflow(model.data, model.length));
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

void Net_ForwardLayers(Net net, struct Mats* outputBlobs, struct CStrings outBlobNames) {
    std::vector< cv::Mat > blobs;

    std::vector< cv::String > names;
    for (int i = 0; i < outBlobNames.length; ++i) {
        names.push_back(cv::String(outBlobNames.strs[i]));
    }
    net->forward(blobs, names);

    // copy blobs into outputBlobs
    outputBlobs->mats = new Mat[blobs.size()];

    for (size_t i = 0; i < blobs.size(); ++i) {
        outputBlobs->mats[i] = new cv::Mat(blobs[i]);
    }

    outputBlobs->length = (int)blobs.size();
}

void Net_SetPreferableBackend(Net net, int backend) {
    net->setPreferableBackend(backend);
}

void Net_SetPreferableTarget(Net net, int target) {
    net->setPreferableTarget(target);
}

int64_t Net_GetPerfProfile(Net net) {
    std::vector<double> layersTimes;
    return net->getPerfProfile(layersTimes);
}

void Net_GetUnconnectedOutLayers(Net net, IntVector* res) {
    std::vector< int > cids(net->getUnconnectedOutLayers());
    int* ids = new int[cids.size()];
    
    for (size_t i = 0; i < cids.size(); ++i) {
        ids[i] = cids[i];
    }

    res->length = cids.size();
    res->val = ids;
    return;
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

Layer Net_GetLayer(Net net, int layerid) {
    return new cv::Ptr<cv::dnn::Layer>(net->getLayer(layerid));
}

void Layer_Close(Layer layer) {
    delete layer;
}

int Layer_InputNameToIndex(Layer layer, const char* name) {
    return (*layer)->inputNameToIndex(name);
}

int Layer_OutputNameToIndex(Layer layer, const char* name) {
    return (*layer)->outputNameToIndex(name);
}

const char* Layer_GetName(Layer layer) {
    return (*layer)->name.c_str();
}

const char* Layer_GetType(Layer layer) {
    return (*layer)->type.c_str();
}
