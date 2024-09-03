#include "face.h"

bool FaceRecognizer_Empty(FaceRecognizer fr) {
    return (*fr)->empty();
}

void FaceRecognizer_Train(FaceRecognizer fr, Mats mats, IntVector labels_in) {
    std::vector<int> labels;

    for (int i = 0, *v = labels_in.val; i < labels_in.length; ++v, ++i) {
        labels.push_back(*v);
    }

    std::vector<cv::Mat> images;

    for (int i = 0; i < mats.length; ++i) {
        images.push_back(*mats.mats[i]);
    }

    (*fr)->train(images, labels);

    return;
}

void FaceRecognizer_Update(FaceRecognizer fr, Mats mats, IntVector labels_in) {
    std::vector<int> labels;

    for (int i = 0, *v = labels_in.val; i < labels_in.length; ++v, ++i) {
        labels.push_back(*v);
    }

    std::vector<cv::Mat> images;

    for (int i = 0; i < mats.length; ++i) {
        images.push_back(*mats.mats[i]);
    }

    (*fr)->update(images, labels);

    return;
}

int FaceRecognizer_Predict(FaceRecognizer fr, Mat sample) {
    int label;
    label = (*fr)->predict(*sample);

    return label;
}

struct PredictResponse FaceRecognizer_PredictExtended(FaceRecognizer fr, Mat sample) {
    struct PredictResponse response;
    int label;
    double confidence;

    (*fr)->predict(*sample, label, confidence);
    response.label = label;
    response.confidence = confidence;

    return response;
}

double FaceRecognizer_GetThreshold(FaceRecognizer fr){
    return (*fr)->getThreshold();
}

void FaceRecognizer_SetThreshold(FaceRecognizer fr, double threshold) {
    (*fr)->setThreshold(threshold);
    return;
}

void FaceRecognizer_SaveFile(FaceRecognizer fr, const char*  filename) {
    (*fr)->write(filename);
    return;
}

void FaceRecognizer_LoadFile(FaceRecognizer fr, const char*  filename) {
    (*fr)->read(filename);
    return;
}

void BasicFaceRecognizer_Train(BasicFaceRecognizer fr, Mats mats, IntVector labels_in){
    std::vector<int> labels;

    for (int i = 0, *v = labels_in.val; i < labels_in.length; ++v, ++i) {
        labels.push_back(*v);
    }

    std::vector<cv::Mat> images;

    for (int i = 0; i < mats.length; ++i) {
        images.push_back(*mats.mats[i]);
    }

    (*fr)->train(images, labels);

    return;
}

void BasicFaceRecognizer_Update(BasicFaceRecognizer fr, Mats mats, IntVector labels_in){
    std::vector<int> labels;
    
    for (int i = 0, *v = labels_in.val; i < labels_in.length; ++v, ++i) {
        labels.push_back(*v);
    }

    std::vector<cv::Mat> images;

    for (int i = 0; i < mats.length; ++i) {
        images.push_back(*mats.mats[i]);
    }
    (*fr)->update(images, labels);
}

Mat BasicFaceRecognizer_getEigenValues(BasicFaceRecognizer fr){
    return new cv::Mat((*fr)->getEigenValues());
}

Mat BasicFaceRecognizer_getEigenVectors(BasicFaceRecognizer fr){
    return new cv::Mat((*fr)->getEigenVectors());
}

Mat BasicFaceRecognizer_getLabels(BasicFaceRecognizer fr){
    return new cv::Mat((*fr)->getLabels());
}

Mat BasicFaceRecognizer_getMean(BasicFaceRecognizer fr){
    return new cv::Mat((*fr)->getMean());
}

int BasicFaceRecognizer_getNumComponents(BasicFaceRecognizer fr) {
    return (*fr)->getNumComponents();
}

Mats BasicFaceRecognizer_getProjections(BasicFaceRecognizer fr) {
    Mats mats;

    std::vector<cv::Mat> vec = (*fr)->getProjections();

    mats.length = (int)vec.size();
    mats.mats = new Mat[vec.size()];

    for(size_t i = 0; i < vec.size(); i++) {
        mats.mats[i] = new cv::Mat(vec[i]);
    }
    return mats;
}

void BasicFaceRecognizer_setNumComponents(BasicFaceRecognizer fr, int val){
    (*fr)->setNumComponents(val);
}	

void BasicFaceRecognizer_SaveFile(BasicFaceRecognizer fr, const char*  filename){
    (*fr)->write(filename);
}

void BasicFaceRecognizer_LoadFile(BasicFaceRecognizer fr, const char*  filename){
    (*fr)->read(filename);
}




LBPHFaceRecognizer CreateLBPHFaceRecognizer() {
    return new cv::Ptr<cv::face::LBPHFaceRecognizer>(cv::face::LBPHFaceRecognizer::create());
}

void LBPHFaceRecognizer_SetRadius(LBPHFaceRecognizer fr, int radius) {
    (*fr)->setRadius(radius);
    return;
}

void LBPHFaceRecognizer_SetNeighbors(LBPHFaceRecognizer fr, int neighbors) {
    (*fr)->setNeighbors(neighbors);
    return;
}

int LBPHFaceRecognizer_GetNeighbors(LBPHFaceRecognizer fr) {
    int n;

    n = (*fr)->getNeighbors();
    return n;
}

void LBPHFaceRecognizer_SetGridX(LBPHFaceRecognizer fr, int x) {
    (*fr)->setGridX(x);
    return;
}

void LBPHFaceRecognizer_SetGridY(LBPHFaceRecognizer fr, int y) {
    (*fr)->setGridY(y);
    return;
}

int LBPHFaceRecognizer_GetGridX(LBPHFaceRecognizer fr) {
    int n = (*fr)->getGridX();
    return n;
}

int LBPHFaceRecognizer_GetGridY(LBPHFaceRecognizer fr) {
    int n = (*fr)->getGridY();
    return n;
}

void LBPHFaceRecognizer_Close(LBPHFaceRecognizer fr) {
    delete fr;
}


FisherFaceRecognizer FisherFaceRecognizer_Create(void) {
    return new cv::Ptr<cv::face::FisherFaceRecognizer>(cv::face::FisherFaceRecognizer::create());
}

FisherFaceRecognizer FisherFaceRecognizer_CreateWithParams(int num_components, float threshold) {
    return new cv::Ptr<cv::face::FisherFaceRecognizer>(cv::face::FisherFaceRecognizer::create(num_components, threshold));
}

void FisherFaceRecognizer_Close(FisherFaceRecognizer fr) {
    delete fr;
}


EigenFaceRecognizer EigenFaceRecognizer_Create(void) {
    return new cv::Ptr<cv::face::EigenFaceRecognizer>(cv::face::EigenFaceRecognizer::create());
}

EigenFaceRecognizer EigenFaceRecognizer_CreateWithParams(int num_components, float threshold) {
    return new cv::Ptr<cv::face::EigenFaceRecognizer>(cv::face::EigenFaceRecognizer::create(num_components, threshold));
}

void EigenFaceRecognizer_Close(EigenFaceRecognizer fr) {
    delete fr;
}