#include "pvl.h"

// Face
Face Face_New()
{
    return new cv::pvl::Face();
}

void Face_Close(Face face)
{
    delete face;
}

void Face_CopyTo(Face src, Face dst)
{
    cv::Rect faceRect = src->get<cv::Rect>(cv::pvl::Face::FACE_RECT);
    int ripAngle = src->get<int>(cv::pvl::Face::RIP_ANGLE);
    int ropAngle = src->get<int>(cv::pvl::Face::ROP_ANGLE);
    int confidence = src->get<int>(cv::pvl::Face::FACE_RECT_CONFIDENCE);
    int trackingID = src->get<int>(cv::pvl::Face::TRACKING_ID);

    dst->setFaceRectInfo(faceRect, ripAngle, ropAngle, confidence, trackingID);

    // TODO: copy eye, mouth, blink, and smile info...
}

Rect Face_GetRect(Face face)
{
    cv::Rect faceRect = face->get<cv::Rect>(cv::pvl::Face::FACE_RECT);

    Rect r = {faceRect.x, faceRect.y, faceRect.width, faceRect.height};
    return r;
}

// FaceDetector
FaceDetector FaceDetector_New() 
{
    return cv::pvl::FaceDetector::create();
}

void FaceDetector_Close(FaceDetector f) 
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
        Face f = Face_New();
        Face_CopyTo(&faces[i], f);

        fs[i] = f;
    }
    Faces ret = {fs, (int)faces.size()};
    return ret;
}
