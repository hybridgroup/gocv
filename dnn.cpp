//go:build !gocv_specific_modules || (gocv_specific_modules && gocv_dnn)

#include "dnn.h"

Net Net_ReadNet(const char* model, const char* config) {
    try {
        Net n = new cv::dnn::Net(cv::dnn::readNet(model, config));
        return n;
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

Net Net_ReadNetBytes(const char* framework, struct ByteArray model, struct ByteArray config) {
    try {
        std::vector<uchar> modelv(model.data, model.data + model.length);
        std::vector<uchar> configv(config.data, config.data + config.length);
        Net n = new cv::dnn::Net(cv::dnn::readNet(framework, modelv, configv));
        return n;
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

Net Net_ReadNetFromCaffe(const char* prototxt, const char* caffeModel) {
    try {
        Net n = new cv::dnn::Net(cv::dnn::readNetFromCaffe(prototxt, caffeModel));
        return n;
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

Net Net_ReadNetFromCaffeBytes(struct ByteArray prototxt, struct ByteArray caffeModel) {
    try {
        Net n = new cv::dnn::Net(cv::dnn::readNetFromCaffe(prototxt.data, prototxt.length,
            caffeModel.data, caffeModel.length));
        return n;
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

Net Net_ReadNetFromTensorflow(const char* model) {
    try {
        Net n = new cv::dnn::Net(cv::dnn::readNetFromTensorflow(model));
        return n;
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

Net Net_ReadNetFromTensorflowBytes(struct ByteArray model) {
    try {
        Net n = new cv::dnn::Net(cv::dnn::readNetFromTensorflow(model.data, model.length));
        return n;
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

Net Net_ReadNetFromTorch(const char* model) {
    try {
        Net n = new cv::dnn::Net(cv::dnn::readNetFromTorch(model));
        return n;
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

Net Net_ReadNetFromONNX(const char* model) {
    try {
        Net n = new cv::dnn::Net(cv::dnn::readNetFromONNX(model));
        return n;
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

Net Net_ReadNetFromONNXBytes(struct ByteArray model) {
    try {
        Net n = new cv::dnn::Net(cv::dnn::readNetFromONNX(model.data, model.length));
        return n;
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

void Net_Close(Net net) {
    delete net;
}

bool Net_Empty(Net net) {
    try {
        return net->empty();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return false;
    }
}

void Net_SetInput(Net net, Mat blob, const char* name) {
    try {
        net->setInput(*blob, name);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

Mat Net_Forward(Net net, const char* outputName) {
    try {
        return new cv::Mat(net->forward(outputName));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return new cv::Mat();
    }
}

void Net_ForwardLayers(Net net, struct Mats* outputBlobs, struct CStrings outBlobNames) {
    try {
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
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void Net_SetPreferableBackend(Net net, int backend) {
    try {
        net->setPreferableBackend(backend);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void Net_SetPreferableTarget(Net net, int target) {
    try {
        net->setPreferableTarget(target);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

int64_t Net_GetPerfProfile(Net net) {
    try {
        std::vector<double> layersTimes;
        return net->getPerfProfile(layersTimes);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0;
    }
}

void Net_GetUnconnectedOutLayers(Net net, IntVector* res) {
    try {
        std::vector< int > cids(net->getUnconnectedOutLayers());
        int* ids = new int[cids.size()];
        
        for (size_t i = 0; i < cids.size(); ++i) {
            ids[i] = cids[i];
        }
    
        res->length = cids.size();
        res->val = ids;
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void Net_GetLayerNames(Net net, CStrings* names) {
    try {
        std::vector< cv::String > cstrs(net->getLayerNames());
        const char **strs = new const char*[cstrs.size()];
    
        for (size_t i = 0; i < cstrs.size(); ++i) {
            strs[i] = cstrs[i].c_str();
        }
    
        names->length = cstrs.size();
        names->strs = strs;
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

struct Rect Net_BlobRectToImageRect(struct Rect rect, Size originalSize, double scalefactor, Size size, Scalar mean, bool swapRB,
                    int ddepth, int dataLayout, int paddingMode, Scalar borderValue) {
    try {
        cv::Scalar sf(scalefactor);
        cv::Size sz(size.width, size.height);
        cv::Scalar cm(mean.val1, mean.val2, mean.val3, mean.val4);
        cv::dnn::DataLayout dl = static_cast<cv::dnn::DataLayout>(dataLayout);
        cv::dnn::ImagePaddingMode pm = static_cast<cv::dnn::ImagePaddingMode>(paddingMode);
        cv::Scalar bv(borderValue.val1, borderValue.val2, borderValue.val3, borderValue.val4);
        cv::dnn::Image2BlobParams params = cv::dnn::Image2BlobParams(sf, sz, cm, swapRB, ddepth, dl, pm, bv);
    
        cv::Rect bRect = params.blobRectToImageRect(cv::Rect(rect.x, rect.y, rect.width, rect.height), cv::Size(originalSize.width, originalSize.height));
        Rect r = {bRect.x, bRect.y, bRect.width, bRect.height};
        return r;
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        Rect r = {0, 0, 0, 0};
        return r;
    }
}

struct Rects Net_BlobRectsToImageRects(struct Rects rects, Size originalSize, double scalefactor, Size size, Scalar mean, bool swapRB,
                    int ddepth, int dataLayout, int paddingMode, Scalar borderValue) {
    try {
        std::vector<cv::Rect> _cRects;
        for (int i = 0; i < rects.length; ++i) {
            _cRects.push_back(cv::Rect(
                rects.rects[i].x,
                rects.rects[i].y,
                rects.rects[i].width,
                rects.rects[i].height
            ));
        }
    
        cv::Scalar sf(scalefactor);
        cv::Size sz(size.width, size.height);
        cv::Scalar cm(mean.val1, mean.val2, mean.val3, mean.val4);
        cv::dnn::DataLayout dl = static_cast<cv::dnn::DataLayout>(dataLayout);
        cv::dnn::ImagePaddingMode pm = static_cast<cv::dnn::ImagePaddingMode>(paddingMode);
        cv::Scalar bv(borderValue.val1, borderValue.val2, borderValue.val3, borderValue.val4);
        cv::dnn::Image2BlobParams params = cv::dnn::Image2BlobParams(sf, sz, cm, swapRB, ddepth, dl, pm, bv);
    
        std::vector<cv::Rect> detected;
        params.blobRectsToImageRects(_cRects, detected, cv::Size(originalSize.width, originalSize.height));
        Rect* drects = new Rect[detected.size()];
    
        for (size_t i = 0; i < detected.size(); ++i) {
            Rect r = {detected[i].x, detected[i].y, detected[i].width, detected[i].height};
            drects[i] = r;
        }
    
        Rects ret = {drects, (int)detected.size()};
        return ret;
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        Rects ret = {NULL, 0};
        return ret;
    }
}

Mat Net_BlobFromImage(Mat image, double scalefactor, Size size, Scalar mean, bool swapRB,
                      bool crop) {
    try {
        cv::Size sz(size.width, size.height);
        cv::Scalar cm(mean.val1, mean.val2, mean.val3, mean.val4);
        // use the default target ddepth here.
        return new cv::Mat(cv::dnn::blobFromImage(*image, scalefactor, sz, cm, swapRB, crop));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return new cv::Mat();
    }
}

Mat Net_BlobFromImageWithParams(Mat image, double scalefactor, Size size, Scalar mean, bool swapRB,
                      int ddepth, int dataLayout, int paddingMode, Scalar borderValue) {
    try {
        cv::Scalar sf(scalefactor);
        cv::Size sz(size.width, size.height);
        cv::Scalar cm(mean.val1, mean.val2, mean.val3, mean.val4);
        cv::dnn::DataLayout dl = static_cast<cv::dnn::DataLayout>(dataLayout);
        cv::dnn::ImagePaddingMode pm = static_cast<cv::dnn::ImagePaddingMode>(paddingMode);
        cv::Scalar bv(borderValue.val1, borderValue.val2, borderValue.val3, borderValue.val4);
        cv::dnn::Image2BlobParams params = cv::dnn::Image2BlobParams(sf, sz, cm, swapRB, ddepth, dl, pm, bv);
    
        return new cv::Mat(cv::dnn::blobFromImageWithParams(*image, params));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return new cv::Mat();
    }
}

void Net_BlobFromImages(struct Mats images, Mat blob, double scalefactor, Size size,
                       Scalar mean, bool swapRB, bool crop, int ddepth) {
    try {
        std::vector<cv::Mat> imgs;
    
        for (int i = 0; i < images.length; ++i) {
            imgs.push_back(*images.mats[i]);
        }
    
        cv::Size sz(size.width, size.height);
        cv::Scalar cm = cv::Scalar(mean.val1, mean.val2, mean.val3, mean.val4);
    
        // ignore the passed in ddepth, just use default.
        cv::dnn::blobFromImages(imgs, *blob, scalefactor, sz, cm, swapRB, crop);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void Net_BlobFromImagesWithParams(struct Mats images, Mat blob, double scalefactor, Size size,
                       Scalar mean, bool swapRB, int ddepth, int dataLayout, int paddingMode, Scalar borderValue) {
    try {
        std::vector<cv::Mat> imgs;
    
        for (int i = 0; i < images.length; ++i) {
            imgs.push_back(*images.mats[i]);
        }
    
        cv::Scalar sf(scalefactor);
        cv::Size sz(size.width, size.height);
        cv::Scalar cm(mean.val1, mean.val2, mean.val3, mean.val4);
        cv::dnn::DataLayout dl = static_cast<cv::dnn::DataLayout>(dataLayout);
        cv::dnn::ImagePaddingMode pm = static_cast<cv::dnn::ImagePaddingMode>(paddingMode);
        cv::Scalar bv(borderValue.val1, borderValue.val2, borderValue.val3, borderValue.val4);
        cv::dnn::Image2BlobParams params = cv::dnn::Image2BlobParams(sf, sz, cm, swapRB, ddepth, dl, pm, bv);
    
        cv::dnn::blobFromImagesWithParams(imgs, *blob, params);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void Net_ImagesFromBlob(Mat blob_, struct Mats* images_) {
    try {
        std::vector<cv::Mat> imgs;
        cv::dnn::imagesFromBlob(*blob_, imgs);
        images_->mats = new Mat[imgs.size()];
    
        for (size_t i = 0; i < imgs.size(); ++i) {
            images_->mats[i] = new cv::Mat(imgs[i]);
        }
        images_->length = (int) imgs.size();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
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
    try {
        return new cv::Ptr<cv::dnn::Layer>(net->getLayer(layerid));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

void Layer_Close(Layer layer) {
    delete layer;
}

int Layer_InputNameToIndex(Layer layer, const char* name) {
    try {
        return (*layer)->inputNameToIndex(name);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return -1;
    }
}

int Layer_OutputNameToIndex(Layer layer, const char* name) {
    try {
        return (*layer)->outputNameToIndex(name);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return -1;
    }
}

const char* Layer_GetName(Layer layer) {
    try {
        return (*layer)->name.c_str();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return "";
    }
}

const char* Layer_GetType(Layer layer) {
    try {
        return (*layer)->type.c_str();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return "";
    }
}

void NMSBoxes(struct Rects bboxes, FloatVector scores, float score_threshold, float nms_threshold, IntVector* indices) {
    try {
        std::vector<cv::Rect> _bboxes;

        for (int i = 0; i < bboxes.length; ++i) {
            _bboxes.push_back(cv::Rect(
                bboxes.rects[i].x,
                bboxes.rects[i].y,
                bboxes.rects[i].width,
                bboxes.rects[i].height
            ));
        }
    
        std::vector<float> _scores;
    
        float* f;
        int i;
        for (i = 0, f = scores.val; i < scores.length; ++f, ++i) {
            _scores.push_back(*f);
        }
    
        std::vector<int> _indices(indices->length);
    
        cv::dnn::NMSBoxes(_bboxes, _scores, score_threshold, nms_threshold, _indices, 1.f, 0);
    
        int* ptr = new int[_indices.size()];
    
        for (size_t i=0; i<_indices.size(); ++i) {
            ptr[i] = _indices[i];
        }
    
        indices->length = _indices.size();
        indices->val = ptr;
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void NMSBoxesWithParams(struct Rects bboxes, FloatVector scores, const float score_threshold, const float nms_threshold, IntVector* indices, const float eta, const int top_k) {
    try {
        std::vector<cv::Rect> _bboxes;

        for (int i = 0; i < bboxes.length; ++i) {
            _bboxes.push_back(cv::Rect(
                bboxes.rects[i].x,
                bboxes.rects[i].y,
                bboxes.rects[i].width,
                bboxes.rects[i].height
            ));
        }
    
        std::vector<float> _scores;
    
        float* f;
        int i;
        for (i = 0, f = scores.val; i < scores.length; ++f, ++i) {
            _scores.push_back(*f);
        }
    
        std::vector<int> _indices(indices->length);
    
        cv::dnn::NMSBoxes(_bboxes, _scores, score_threshold, nms_threshold, _indices, eta, top_k);
    
        int* ptr = new int[_indices.size()];
    
        for (size_t i=0; i<_indices.size(); ++i) {
            ptr[i] = _indices[i];
        }
    
        indices->length = _indices.size();
        indices->val = ptr;
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}