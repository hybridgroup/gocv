#include "face_detector.h"

// FaceDetector
FaceDetector FaceDetector_New() 
{
    return new cv::Ptr<cv::pvl::FaceDetector>(cv::pvl::FaceDetector::create());
}

void FaceDetector_Close(FaceDetector f) 
{
    delete f;
}

void FaceDetector_SetTrackingModeEnabled(FaceDetector f, bool enabled)
{
    (*f)->setTrackingModeEnabled(enabled);
    return;
}

struct Faces FaceDetector_DetectFaceRect(FaceDetector fd, Mat img)
{
    std::vector<cv::pvl::Face> faces;
    (*fd)->detectFaceRect(*img, faces);

    Face* fs = new Face[faces.size()];
    for (size_t i = 0; i < faces.size(); ++i) {
        Face f = Face_New();
        Face_CopyTo(&faces[i], f);

        fs[i] = f;
    }
    Faces ret = {fs, (int)faces.size()};
    return ret;
}

void FaceDetector_DetectEye(FaceDetector f, Mat img, Face face)
{
    (*f)->detectEye(*img, *face);
    return;
}

void FaceDetector_DetectMouth(FaceDetector f, Mat img, Face face)
{
    (*f)->detectMouth(*img, *face);
    return;
}

void FaceDetector_DetectSmile(FaceDetector f, Mat img, Face face)
{
    (*f)->detectSmile(*img, *face);
    return;
}

void FaceDetector_DetectBlink(FaceDetector f, Mat img, Face face)
{
    (*f)->detectBlink(*img, *face);
    return;
}
