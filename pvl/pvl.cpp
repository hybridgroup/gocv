#include "pvl.h"

// Face
Face Face_New()
{
    return new cv::pvl::Face();
}

void Face_Delete(Face f)
{
    delete f;
}

Rect Face_GetRect(Face f)
{
    cv::Rect faceRect = f->get<cv::Rect>(cv::pvl::Face::FACE_RECT);
    Rect r = {faceRect.x, faceRect.y, faceRect.width, faceRect.height};
    return r;
}

void Faces_Delete(struct Faces fs) {
    delete fs.faces;
}

// FaceDetector
FaceDetector FaceDetector_New() 
{
    return cv::pvl::FaceDetector::create();
}

void FaceDetector_Delete(FaceDetector f) 
{
    delete f;
}

void FaceDetector_SetTrackingModeEnabled(FaceDetector f, bool enabled)
{
    f->setTrackingModeEnabled(enabled);
    return;
}

struct Faces FaceDetector_DetectFaceRect(FaceDetector f, Mat img)
{
    std::vector<cv::pvl::Face> faces;
    cv::Mat grayedFrame;
    cv::cvtColor(*img, grayedFrame, cv::COLOR_BGR2GRAY);
    f->detectFaceRect(grayedFrame, faces);

    Face* fs = new Face[faces.size()];
    for (size_t i = 0; i < faces.size(); ++i) {
        fs[i] = &faces[i];
    }
    Faces ret = {fs, (int)faces.size()};
    return ret;
}