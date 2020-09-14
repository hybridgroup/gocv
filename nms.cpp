#include "nms.h"

void NMSBoxes(struct Rects bboxes, FloatVector scores, const float score_threshold, const float nms_threshold, IntVector indices) {
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

    std::vector<int> _indices;

    for (int i = 0, *v = indices.val; i < indices.length; ++v, ++i) {
        _indices.push_back(*v);
    }

    cv::dnn::NMSBoxes(_bboxes, _scores, score_threshold, nms_threshold, _indices, 1.f, 0);
}

void NMSBoxesWithParams(struct Rects bboxes, FloatVector scores, const float score_threshold, const float nms_threshold, IntVector indices, const float eta, const int top_k) {
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

    std::vector<int> _indices;

    for (int i = 0, *v = indices.val; i < indices.length; ++v, ++i) {
        _indices.push_back(*v);
    }

    cv::dnn::NMSBoxes(_bboxes, _scores, score_threshold, nms_threshold, _indices, eta, top_k);
}