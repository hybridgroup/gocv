#include "pvl.h"

// Face
Face Face_New()
{
    return new cv::pvl::Face();
}

void Face_Delete(Face face)
{
    delete face;
}

Rect Face_GetRect(Face face)
{
    cv::Rect faceRect = face->get<cv::Rect>(cv::pvl::Face::FACE_RECT);

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

struct Faces FaceDetector_DetectFaceRect(FaceDetector fd, Mat img)
{
    // TODO: do conversion in imgproc
    cv::Mat grayedFrame;
    cv::cvtColor(*img, grayedFrame, cv::COLOR_BGR2GRAY);

    std::vector<cv::pvl::Face> faces;
    fd->detectFaceRect(grayedFrame, faces);

    Face* fs = new Face[faces.size()];
    for (size_t i = 0; i < faces.size(); ++i) {
        cv::Rect faceRect = faces[i].get<cv::Rect>(cv::pvl::Face::FACE_RECT);

        Face f = Face_New();
        f->setFaceRectInfo(faceRect);
        // TODO: copy all the other face info...

        fs[i] = f;
    }
    Faces ret = {fs, (int)faces.size()};
    return ret;
}