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