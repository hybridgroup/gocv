#ifndef _OPENCV3_CONTRIB_SIGNAL_H_
#define _OPENCV3_CONTRIB_SIGNAL_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/signal/signal_resample.hpp>
extern "C" {
#endif

#include "../core.h"

void Signal_ResampleSignal(Mat inputSignal, Mat outSignal, int inFreq, int outFreq);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_CONTRIB_SIGNAL_H_
