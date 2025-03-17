#include "face.h"

bool FaceRecognizer_Empty(FaceRecognizer fr) {
    try {
        return (*fr)->empty();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return false;
    }
}

void FaceRecognizer_Train(FaceRecognizer fr, Mats mats, IntVector labels_in) {
    try {
        std::vector<int> labels;

        for (int i = 0, *v = labels_in.val; i < labels_in.length; ++v, ++i) {
            labels.push_back(*v);
        }
    
        std::vector<cv::Mat> images;
    
        for (int i = 0; i < mats.length; ++i) {
            images.push_back(*mats.mats[i]);
        }
    
        (*fr)->train(images, labels);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void FaceRecognizer_Update(FaceRecognizer fr, Mats mats, IntVector labels_in) {
    try {
        std::vector<int> labels;

        for (int i = 0, *v = labels_in.val; i < labels_in.length; ++v, ++i) {
            labels.push_back(*v);
        }
    
        std::vector<cv::Mat> images;
    
        for (int i = 0; i < mats.length; ++i) {
            images.push_back(*mats.mats[i]);
        }
    
        (*fr)->update(images, labels);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

int FaceRecognizer_Predict(FaceRecognizer fr, Mat sample) {
    try {
        int label;
        label = (*fr)->predict(*sample);
    
        return label;
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0;
    }
}

struct PredictResponse FaceRecognizer_PredictExtended(FaceRecognizer fr, Mat sample) {
    try {
        struct PredictResponse response;
        int label;
        double confidence;
    
        (*fr)->predict(*sample, label, confidence);
        response.label = label;
        response.confidence = confidence;
    
        return response;
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        struct PredictResponse response;
        return response;
    }
}

double FaceRecognizer_GetThreshold(FaceRecognizer fr){
    try {
        return (*fr)->getThreshold();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0.0;
    }
}

void FaceRecognizer_SetThreshold(FaceRecognizer fr, double threshold) {
    try {
        (*fr)->setThreshold(threshold);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void FaceRecognizer_SaveFile(FaceRecognizer fr, const char*  filename) {
    try {
        (*fr)->write(filename);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void FaceRecognizer_LoadFile(FaceRecognizer fr, const char*  filename) {
    try {
        (*fr)->read(filename);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void BasicFaceRecognizer_Train(BasicFaceRecognizer fr, Mats mats, IntVector labels_in){
    try {
        std::vector<int> labels;

        for (int i = 0, *v = labels_in.val; i < labels_in.length; ++v, ++i) {
            labels.push_back(*v);
        }
    
        std::vector<cv::Mat> images;
    
        for (int i = 0; i < mats.length; ++i) {
            images.push_back(*mats.mats[i]);
        }
    
        (*fr)->train(images, labels);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void BasicFaceRecognizer_Update(BasicFaceRecognizer fr, Mats mats, IntVector labels_in){
    try {
        std::vector<int> labels;
    
        for (int i = 0, *v = labels_in.val; i < labels_in.length; ++v, ++i) {
            labels.push_back(*v);
        }
    
        std::vector<cv::Mat> images;
    
        for (int i = 0; i < mats.length; ++i) {
            images.push_back(*mats.mats[i]);
        }
        (*fr)->update(images, labels);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

Mat BasicFaceRecognizer_getEigenValues(BasicFaceRecognizer fr){
    try {
        return new cv::Mat((*fr)->getEigenValues());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return new cv::Mat();
    }
}

Mat BasicFaceRecognizer_getEigenVectors(BasicFaceRecognizer fr){
    try {
        return new cv::Mat((*fr)->getEigenVectors());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return new cv::Mat();
    }
}

Mat BasicFaceRecognizer_getLabels(BasicFaceRecognizer fr){
    try {
        return new cv::Mat((*fr)->getLabels());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return new cv::Mat();
    }
}

Mat BasicFaceRecognizer_getMean(BasicFaceRecognizer fr){
    try {
        return new cv::Mat((*fr)->getMean());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return new cv::Mat();
    }
}

int BasicFaceRecognizer_getNumComponents(BasicFaceRecognizer fr) {
    try {
        return (*fr)->getNumComponents();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0;
    }
}

Mats BasicFaceRecognizer_getProjections(BasicFaceRecognizer fr) {
    try {
        Mats mats;

        std::vector<cv::Mat> vec = (*fr)->getProjections();
    
        mats.length = (int)vec.size();
        mats.mats = new Mat[vec.size()];
    
        for(size_t i = 0; i < vec.size(); i++) {
            mats.mats[i] = new cv::Mat(vec[i]);
        }
        return mats;
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        Mats mats;
        return mats;
    }
}

void BasicFaceRecognizer_setNumComponents(BasicFaceRecognizer fr, int val){
    try {
        (*fr)->setNumComponents(val);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}	

void BasicFaceRecognizer_SaveFile(BasicFaceRecognizer fr, const char*  filename){
    try {
        (*fr)->write(filename);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void BasicFaceRecognizer_LoadFile(BasicFaceRecognizer fr, const char*  filename){
    try {
        (*fr)->read(filename);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

LBPHFaceRecognizer CreateLBPHFaceRecognizer() {
    try {
        return new cv::Ptr<cv::face::LBPHFaceRecognizer>(cv::face::LBPHFaceRecognizer::create());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

void LBPHFaceRecognizer_SetRadius(LBPHFaceRecognizer fr, int radius) {
    try {
        (*fr)->setRadius(radius);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void LBPHFaceRecognizer_SetNeighbors(LBPHFaceRecognizer fr, int neighbors) {
    try {
        (*fr)->setNeighbors(neighbors);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

int LBPHFaceRecognizer_GetNeighbors(LBPHFaceRecognizer fr) {
    try {
        return (*fr)->getNeighbors();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0;
    }
}

void LBPHFaceRecognizer_SetGridX(LBPHFaceRecognizer fr, int x) {
    try {
        (*fr)->setGridX(x);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

void LBPHFaceRecognizer_SetGridY(LBPHFaceRecognizer fr, int y) {
    try {
        (*fr)->setGridY(y);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
    }
}

int LBPHFaceRecognizer_GetGridX(LBPHFaceRecognizer fr) {
    try {
        return (*fr)->getGridX();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0;
    }
}

int LBPHFaceRecognizer_GetGridY(LBPHFaceRecognizer fr) {
    try {
        return (*fr)->getGridY();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return 0;
    }
}

void LBPHFaceRecognizer_Close(LBPHFaceRecognizer fr) {
    delete fr;
}

FisherFaceRecognizer FisherFaceRecognizer_Create(void) {
    try {
        return new cv::Ptr<cv::face::FisherFaceRecognizer>(cv::face::FisherFaceRecognizer::create());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

FisherFaceRecognizer FisherFaceRecognizer_CreateWithParams(int num_components, float threshold) {
    try {
        return new cv::Ptr<cv::face::FisherFaceRecognizer>(cv::face::FisherFaceRecognizer::create(num_components, threshold));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

void FisherFaceRecognizer_Close(FisherFaceRecognizer fr) {
    delete fr;
}

EigenFaceRecognizer EigenFaceRecognizer_Create(void) {
    try {
        return new cv::Ptr<cv::face::EigenFaceRecognizer>(cv::face::EigenFaceRecognizer::create());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

EigenFaceRecognizer EigenFaceRecognizer_CreateWithParams(int num_components, float threshold) {
    try {
        return new cv::Ptr<cv::face::EigenFaceRecognizer>(cv::face::EigenFaceRecognizer::create(num_components, threshold));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

void EigenFaceRecognizer_Close(EigenFaceRecognizer fr) {
    delete fr;
}
