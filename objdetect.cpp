#include "objdetect.h"

// CascadeClassifier

CascadeClassifier CascadeClassifier_New() {
    return new cv::CascadeClassifier();
}

void CascadeClassifier_Close(CascadeClassifier cs) {
    delete cs;
}

int CascadeClassifier_Load(CascadeClassifier cs, const char* name) {
    return cs->load(name);
}

struct Rects CascadeClassifier_DetectMultiScale(CascadeClassifier cs, Mat img) {
    std::vector<cv::Rect> detected;
    cs->detectMultiScale(*img, detected); // uses all default parameters
    Rect* rects = new Rect[detected.size()];

    for (size_t i = 0; i < detected.size(); ++i) {
        Rect r = {detected[i].x, detected[i].y, detected[i].width, detected[i].height};
        rects[i] = r;
    }

    Rects ret = {rects, (int)detected.size()};
    return ret;
}

struct Rects CascadeClassifier_DetectMultiScaleWithParams(CascadeClassifier cs, Mat img,
        double scale, int minNeighbors, int flags, Size minSize, Size maxSize) {

    cv::Size minSz(minSize.width, minSize.height);
    cv::Size maxSz(maxSize.width, maxSize.height);

    std::vector<cv::Rect> detected;
    cs->detectMultiScale(*img, detected, scale, minNeighbors, flags, minSz, maxSz);
    Rect* rects = new Rect[detected.size()];

    for (size_t i = 0; i < detected.size(); ++i) {
        Rect r = {detected[i].x, detected[i].y, detected[i].width, detected[i].height};
        rects[i] = r;
    }

    Rects ret = {rects, (int)detected.size()};
    return ret;
}

// HOGDescriptor

HOGDescriptor HOGDescriptor_New() {
    return new cv::HOGDescriptor();
}

void HOGDescriptor_Close(HOGDescriptor hog) {
    delete hog;
}

int HOGDescriptor_Load(HOGDescriptor hog, const char* name) {
    return hog->load(name);
}

struct Rects HOGDescriptor_DetectMultiScale(HOGDescriptor hog, Mat img) {
    std::vector<cv::Rect> detected;
    hog->detectMultiScale(*img, detected);
    Rect* rects = new Rect[detected.size()];

    for (size_t i = 0; i < detected.size(); ++i) {
        Rect r = {detected[i].x, detected[i].y, detected[i].width, detected[i].height};
        rects[i] = r;
    }

    Rects ret = {rects, (int)detected.size()};
    return ret;
}

struct Rects HOGDescriptor_DetectMultiScaleWithParams(HOGDescriptor hog, Mat img,
        double hitThresh, Size winStride, Size padding, double scale, double finalThresh,
        bool useMeanshiftGrouping) {

    cv::Size wSz(winStride.width, winStride.height);
    cv::Size pSz(padding.width, padding.height);

    std::vector<cv::Rect> detected;
    hog->detectMultiScale(*img, detected, hitThresh, wSz, pSz, scale, finalThresh,
                          useMeanshiftGrouping);
    Rect* rects = new Rect[detected.size()];

    for (size_t i = 0; i < detected.size(); ++i) {
        Rect r = {detected[i].x, detected[i].y, detected[i].width, detected[i].height};
        rects[i] = r;
    }

    Rects ret = {rects, (int)detected.size()};
    return ret;
}

Mat HOG_GetDefaultPeopleDetector() {
    return new cv::Mat(cv::HOGDescriptor::getDefaultPeopleDetector());
}

void HOGDescriptor_SetSVMDetector(HOGDescriptor hog, Mat det) {
    hog->setSVMDetector(*det);
}

struct Rects GroupRectangles(struct Rects rects, int groupThreshold, double eps) {
    std::vector<cv::Rect> vRect;

    for (int i = 0; i < rects.length; ++i) {
        cv::Rect r = cv::Rect(rects.rects[i].x, rects.rects[i].y, rects.rects[i].width,
                              rects.rects[i].height);
        vRect.push_back(r);
    }

    cv::groupRectangles(vRect, groupThreshold, eps);

    Rect* results = new Rect[vRect.size()];

    for (size_t i = 0; i < vRect.size(); ++i) {
        Rect r = {vRect[i].x, vRect[i].y, vRect[i].width, vRect[i].height};
        results[i] = r;
    }

    Rects ret = {results, (int)vRect.size()};
    return ret;
}

// QRCodeDetector

QRCodeDetector QRCodeDetector_New() {
    return new cv::QRCodeDetector();
}

void QRCodeDetector_Close(QRCodeDetector qr) {
    delete qr;
}

const char* QRCodeDetector_DetectAndDecode(QRCodeDetector qr, Mat input,Mat points,Mat straight_qrcode) {
  cv::String *str = new cv::String(qr->detectAndDecode(*input,*points,*straight_qrcode)); 
  return str->c_str();
}

bool QRCodeDetector_Detect(QRCodeDetector qr, Mat input,Mat points) {
  return qr->detect(*input,*points); 
}

const char* QRCodeDetector_Decode(QRCodeDetector qr, Mat input,Mat inputPoints,Mat straight_qrcode) {
  cv::String *str = new cv::String(qr->detectAndDecode(*input,*inputPoints,*straight_qrcode)); 
  return str->c_str();
}

bool QRCodeDetector_DetectMulti(QRCodeDetector qr, Mat input, Mat points) {
  return qr->detectMulti(*input,*points);
}

bool QRCodeDetector_DetectAndDecodeMulti(QRCodeDetector qr, Mat input, CStrings* decoded, Mat points, struct Mats* qrCodes) {
  std::vector<cv::String> decodedCodes;
  std::vector<cv::Mat> straightQrCodes;
  bool res = qr->detectAndDecodeMulti(*input, decodedCodes, *points, straightQrCodes);
  if (!res) {
    return res;
  }

  qrCodes->mats = new Mat[straightQrCodes.size()];
  qrCodes->length = straightQrCodes.size();
  for (size_t i = 0; i < straightQrCodes.size(); i++) {
     qrCodes->mats[i] = new cv::Mat(straightQrCodes[i]);
  }

  const char **strs = new const char*[decodedCodes.size()];
  for (size_t i = 0; i < decodedCodes.size(); ++i) {
      strs[i] = decodedCodes[i].c_str();
  }
  decoded->length = decodedCodes.size();
  decoded->strs = strs;
  return res;
}

FaceDetectorYN FaceDetectorYN_Create(const char* model, const char* config, Size size) {
    cv::String smodel = cv::String(model);
    cv::String sconfig = cv::String(config);
    cv::Size   ssize = cv::Size(size.width, size.height);

    return new cv::Ptr<cv::FaceDetectorYN>(cv::FaceDetectorYN::create(smodel, sconfig, ssize));
}

FaceDetectorYN FaceDetectorYN_Create_WithParams(const char* model, const char* config, Size size, float score_threshold, float nms_threshold, int top_k, int backend_id, int target_id) {
    cv::String smodel = cv::String(model);
    cv::String sconfig = cv::String(config);
    cv::Size   ssize = cv::Size(size.width, size.height);

    return new cv::Ptr<cv::FaceDetectorYN>(cv::FaceDetectorYN::create(smodel, sconfig, ssize, score_threshold, nms_threshold, top_k, backend_id, target_id));
}

FaceDetectorYN FaceDetectorYN_Create_FromBytes(const char* framework, void* bufferModel, int model_size, void* bufferConfig, int config_size, Size size) {
    cv::String sframework = cv::String(framework);
    cv::Size   ssize = cv::Size(size.width, size.height);
    
    std::vector<uchar> bufferModelV;
    std::vector<uchar> bufferConfigV;

    uchar* bmv = (uchar*)bufferModel;
    uchar* bcv = (uchar*)bufferConfig;


    for(int i = 0; i < model_size; i ++) {
        bufferModelV.push_back(bmv[i]);
    }
   for(int i = 0; i < config_size; i ++) {
        bufferConfigV.push_back(bcv[i]);
    }

    return new cv::Ptr<cv::FaceDetectorYN>(cv::FaceDetectorYN::create(sframework, bufferModelV, bufferConfigV, ssize));
}

FaceDetectorYN FaceDetectorYN_Create_FromBytes_WithParams(const char* framework, void* bufferModel, int model_size, void* bufferConfig, int config_size, Size size, float score_threshold, float nms_threshold, int top_k, int backend_id, int target_id) {
    cv::String sframework = cv::String(framework);
    cv::Size   ssize = cv::Size(size.width, size.height);
    
    std::vector<uchar> bufferModelV;
    std::vector<uchar> bufferConfigV;

    uchar* bmv = (uchar*)bufferModel;
    uchar* bcv = (uchar*)bufferConfig;


    for(int i = 0; i < model_size; i ++) {
        bufferModelV.push_back(bmv[i]);
    }
   for(int i = 0; i < config_size; i ++) {
        bufferConfigV.push_back(bcv[i]);
    }

    return new cv::Ptr<cv::FaceDetectorYN>(cv::FaceDetectorYN::create(sframework, bufferModelV, bufferConfigV, ssize, score_threshold, nms_threshold, top_k, backend_id, target_id));
}

void FaceDetectorYN_Close(FaceDetectorYN fd) {
    delete fd;
}

int FaceDetectorYN_Detect(FaceDetectorYN fd, Mat image, Mat faces) {
    return (*fd)->detect(*image, *faces);
}

Size FaceDetectorYN_GetInputSize(FaceDetectorYN fd) {
    Size sz;

    cv::Size cvsz = (*fd)->getInputSize();

    sz.width = cvsz.width;
    sz.height = cvsz.height;

    return sz;
}

float FaceDetectorYN_GetNMSThreshold(FaceDetectorYN fd) {
    return (*fd)->getNMSThreshold();
}

float FaceDetectorYN_GetScoreThreshold(FaceDetectorYN fd) {
    return (*fd)->getScoreThreshold();
}

int FaceDetectorYN_GetTopK(FaceDetectorYN fd) {
    return (*fd)->getTopK();
}

void FaceDetectorYN_SetInputSize(FaceDetectorYN fd, Size input_size){
    cv::Size isz(input_size.width, input_size.height);
    (*fd)->setInputSize(isz);
}

void FaceDetectorYN_SetNMSThreshold(FaceDetectorYN fd, float nms_threshold){
    (*fd)->setNMSThreshold(nms_threshold);
}

void FaceDetectorYN_SetScoreThreshold(FaceDetectorYN fd, float score_threshold){
    (*fd)->setScoreThreshold(score_threshold);
}

void FaceDetectorYN_SetTopK(FaceDetectorYN fd, int top_k){
    (*fd)->setTopK(top_k);
}

FaceRecognizerSF FaceRecognizerSF_Create(const char* model, const char* config) {
    return FaceRecognizerSF_Create_WithParams(model, config, 0, 0);
}

FaceRecognizerSF FaceRecognizerSF_Create_WithParams(const char* model, const char* config, int backend_id, int target_id) {
    cv::Ptr<cv::FaceRecognizerSF>* p = new cv::Ptr<cv::FaceRecognizerSF>(cv::FaceRecognizerSF::create(model, config, backend_id, target_id));
    return p;
}

void FaceRecognizerSF_Close(FaceRecognizerSF fr) {
    delete fr;
}

void FaceRecognizerSF_AlignCrop(FaceRecognizerSF fr, Mat src_img, Mat face_box, Mat aligned_img) {
    (*fr)->alignCrop(*src_img, *face_box, *aligned_img);
}

void FaceRecognizerSF_Feature(FaceRecognizerSF fr, Mat aligned_img, Mat face_feature) {
    (*fr)->feature(*aligned_img, *face_feature);
}

float FaceRecognizerSF_Match(FaceRecognizerSF fr, Mat face_feature1, Mat face_feature2) {
    return FaceRecognizerSF_Match_WithParams(fr, face_feature1, face_feature2, 0);
}

float FaceRecognizerSF_Match_WithParams(FaceRecognizerSF fr, Mat face_feature1, Mat face_feature2, int dis_type) {
    double rv = (*fr)->match(*face_feature1, *face_feature2, dis_type);
    return (float)rv;
}