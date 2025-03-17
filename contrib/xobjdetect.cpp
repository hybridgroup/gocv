#include "xobjdetect.h"

WBDetector WBDetector_Create(){
    try {
        return new cv::Ptr<cv::xobjdetect::WBDetector>(cv::xobjdetect::WBDetector::create());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

void WBDetector_Close(WBDetector det){
    delete det;
}

WBDetector_detect_result WBDetector_Detect(WBDetector det, Mat img){
    try {
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
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        WBDetector_detect_result result;
        return result;
    }
}

void WBDetector_Read(WBDetector det, FileNode node){
    try {
        (*det)->read(*node);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void WBDetector_Train(WBDetector det, const char* pos_samples, const char* neg_imgs){
    try {
        (*det)->train(pos_samples, neg_imgs);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void WBDetector_Write(WBDetector det, FileStorage fs){
    try {
        (*det)->write(*fs);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}
