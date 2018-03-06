
#ifndef _OPENCV3_MY_H_
#define _OPENCV3_MY_H_

#include "../core.h"

#ifdef __cplusplus
#include <opencv2/tracking/tracker.hpp>

extern "C" {
#endif


#ifdef __cplusplus
typedef cv::Ptr<cv::Tracker>* Tracker;
typedef cv::Ptr<cv::TrackerMIL>* TrackerMil;
typedef cv::Ptr<cv::TrackerBoosting>* TrackerBoosting;
typedef cv::Ptr<cv::TrackerMedianFlow>* TrackerMedianFlow;
typedef cv::Ptr<cv::TrackerTLD>* TrackerTld;
typedef cv::Ptr<cv::TrackerKCF>* TrackerKcf;
typedef cv::Ptr<cv::TrackerMOSSE>* TrackerMosse;
typedef cv::Ptr<cv::TrackerCSRT>* TrackerCsrt;
#else
typedef void* Tracker;
typedef void* TrackerMil;
typedef void* TrackerBoosting;
typedef void* TrackerMedianFlow;
typedef void* TrackerTld;
typedef void* TrackerKcf;
typedef void* TrackerMosse;
typedef void* TrackerCsrt;
#endif

bool Tracker_Init(Tracker self, Mat image, Rect boundingBox);
bool Tracker_Update(Tracker self, Mat image, Rect *boundingBox);


TrackerMil TrackerMil_Create();
void TrackerMil_Close(TrackerMil self);


TrackerBoosting TrackerBoosting_Create();
void TrackerBoosting_Close(TrackerBoosting self);


TrackerMedianFlow TrackerMedianFlow_Create();
void TrackerMedianFlow_Close(TrackerMedianFlow self);


TrackerTld TrackerTld_Create();
void TrackerTld_Close(TrackerTld self);


TrackerKcf TrackerKcf_Create();
void TrackerKcf_Close(TrackerKcf self);


TrackerMosse TrackerMosse_Create();
void TrackerMosse_Close(TrackerMosse self);


TrackerCsrt TrackerCsrt_Create();
void TrackerCsrt_Close(TrackerCsrt self);


#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_MY_H_
