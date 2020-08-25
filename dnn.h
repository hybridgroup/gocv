#ifndef _OPENCV3_DNN_H_
#define _OPENCV3_DNN_H_

#include <stdbool.h>

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/dnn.hpp>
extern "C" {
#endif

#include "core.h"

#ifdef __cplusplus
typedef cv::dnn::Net* Net;
typedef cv::Ptr<cv::dnn::Layer>* Layer;
#else
typedef void* Net;
typedef void* Layer;
#endif

Net Net_ReadNet(const char* model, const char* config);
Net Net_ReadNetBytes(const char* framework, struct ByteArray model, struct ByteArray config);
Net Net_ReadNetFromCaffe(const char* prototxt, const char* caffeModel);
Net Net_ReadNetFromCaffeBytes(struct ByteArray prototxt, struct ByteArray caffeModel);
Net Net_ReadNetFromTensorflow(const char* model);
Net Net_ReadNetFromTensorflowBytes(struct ByteArray model);
Net Net_ReadNetFromTorch(const char* model);
Mat Net_BlobFromImage(Mat image, double scalefactor, Size size, Scalar mean, bool swapRB,
                      bool crop);
void Net_BlobFromImages(struct Mats images, Mat blob,  double scalefactor, Size size, 
                        Scalar mean, bool swapRB, bool crop, int ddepth);
void Net_ImagesFromBlob(Mat blob_, struct Mats* images_);
void Net_Close(Net net);
bool Net_Empty(Net net);
void Net_SetInput(Net net, Mat blob, const char* name);
Mat Net_Forward(Net net, const char* outputName);
void Net_ForwardLayers(Net net, struct Mats* outputBlobs, struct CStrings outBlobNames);
void Net_SetPreferableBackend(Net net, int backend);
void Net_SetPreferableTarget(Net net, int target);
int64_t Net_GetPerfProfile(Net net);
void Net_GetUnconnectedOutLayers(Net net, IntVector* res);
void Net_GetLayerNames(Net net, CStrings* names);

Mat Net_GetBlobChannel(Mat blob, int imgidx, int chnidx);
Scalar Net_GetBlobSize(Mat blob);

Layer Net_GetLayer(Net net, int layerid);
void Layer_Close(Layer layer);
int Layer_InputNameToIndex(Layer layer, const char* name);
int Layer_OutputNameToIndex(Layer layer, const char* name);
const char* Layer_GetName(Layer layer);
const char* Layer_GetType(Layer layer);

void NMSBoxes(struct Rects bboxes, FloatVector scores, float score_threshold, float nms_threshold, IntVector* indices);
void NMSBoxesWithParams(struct Rects bboxes, FloatVector scores, const float score_threshold, const float nms_threshold, IntVector* indices, const float eta, const int top_k);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_DNN_H_
