
#ifndef _OPENCV3_TRACKING_H_
#define _OPENCV3_TRACKING_H_

#include "../core.h"

#ifdef __cplusplus
#include <opencv2/tracking/tracker.hpp>

extern "C" {
#endif


#ifdef __cplusplus
typedef cv::Ptr<cv::Tracker>* Tracker;
typedef cv::Ptr<cv::TrackerMIL>* TrackerMIL;
typedef cv::Ptr<cv::TrackerBoosting>* TrackerBoosting;
typedef cv::Ptr<cv::TrackerMedianFlow>* TrackerMedianFlow;
typedef cv::Ptr<cv::TrackerTLD>* TrackerTLD;
typedef cv::Ptr<cv::TrackerKCF>* TrackerKCF;
typedef cv::Ptr<cv::TrackerMOSSE>* TrackerMOSSE;
typedef cv::Ptr<cv::TrackerCSRT>* TrackerCSRT;
#else
typedef void* Tracker;
typedef void* TrackerMIL;
typedef void* TrackerBoosting;
typedef void* TrackerMedianFlow;
typedef void* TrackerTLD;
typedef void* TrackerKCF;
typedef void* TrackerMOSSE;
typedef void* TrackerCSRT;
#endif

bool Tracker_Init(Tracker self, Mat image, Rect boundingBox);
bool Tracker_Update(Tracker self, Mat image, Rect* boundingBox);

TrackerMIL TrackerMIL_Create();
void TrackerMIL_Close(TrackerMIL self);

TrackerBoosting TrackerBoosting_Create();
void TrackerBoosting_Close(TrackerBoosting self);

TrackerMedianFlow TrackerMedianFlow_Create();
void TrackerMedianFlow_Close(TrackerMedianFlow self);

TrackerTLD TrackerTLD_Create();
void TrackerTLD_Close(TrackerTLD self);

TrackerKCF TrackerKCF_Create();
void TrackerKCF_Close(TrackerKCF self);

TrackerMOSSE TrackerMOSSE_Create();
void TrackerMOSSE_Close(TrackerMOSSE self);

TrackerCSRT TrackerCSRT_Create();
void TrackerCSRT_Close(TrackerCSRT self);


#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_TRACKING_H_
