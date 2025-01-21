#include "xobjdetect.h"

WBDetector WBDetector_Create(){
    return new cv::Ptr<cv::xobjdetect::WBDetector>(cv::xobjdetect::WBDetector::create());
}

void WBDetector_Close(WBDetector det){
    delete det;
}

WBDetector_detect_result WBDetector_Detect(WBDetector det, Mat img){
    std::vector<cv::Rect> bb;
    std::vector<double> conf;
    WBDetector_detect_result result;

    (*det)->detect(*img, bb, conf);

    Rect* result_rects = (Rect*)malloc(bb.size()*sizeof(Rect));
    float* result_confs = (float*)malloc(conf.size()*sizeof(float));

    for(int i = 0; i < bb.size(); i ++) {
        result_rects[i].x = bb[i].x;
        result_rects[i].y = bb[i].y;
        result_rects[i].width = bb[i].width;
        result_rects[i].height = bb[i].height;
    }

    for(int i = 0; i < conf.size(); i ++){
        result_confs[i] = conf[i];
    }

    result.bboxes.rects = result_rects;
    result.bboxes.length = bb.size();

    result.confidences.val = result_confs;
    result.confidences.length = conf.size();

    return result;
}

void WBDetector_Read(WBDetector det, FileNode node){
    (*det)->read(*node);    
}

void WBDetector_Train(WBDetector det, const char* pos_samples, const char* neg_imgs){
    (*det)->train(pos_samples, neg_imgs);
}

void WBDetector_Write(WBDetector det, FileStorage fs){
    (*det)->write(*fs);
}
