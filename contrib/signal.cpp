#include "signal.h"

void Signal_ResampleSignal(Mat inputSignal, Mat outSignal, int inFreq, int outFreq) {
    cv::signal::resampleSignal(*inputSignal, *outSignal, inFreq, outFreq);
}

