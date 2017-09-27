#ifndef _OPENCV3_PVL_H_
#define _OPENCV3_PVL_H_

#include <stdbool.h>

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/pvl.hpp>
extern "C" {
#endif

#include "../core.h"

#ifdef __cplusplus
typedef cv::pvl::Face* Face;
typedef cv::Ptr<cv::pvl::FaceDetector> FaceDetector;
#else
typedef void* Face;
typedef void* FaceDetector;
#endif

// Wrapper for the vector of Face struct aka std::vector<Face>
typedef struct Faces {
    Face* faces;
    int length;
} Faces;

// Face
Face Face_New();
void Face_Delete(Face f);
void Face_CopyTo(Face src, Face dst);
Rect Face_GetRect(Face f);

// Faces
struct Faces Faces_New();
void Faces_Delete(struct Faces fs);

// FaceDetector
FaceDetector FaceDetector_New();
void FaceDetector_Delete(FaceDetector f);
void FaceDetector_SetTrackingModeEnabled(FaceDetector f, bool enabled);
struct Faces FaceDetector_DetectFaceRect(FaceDetector f, Mat img);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_PVL_H_
