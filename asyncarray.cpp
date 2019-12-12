// +build openvino

#include <string.h>
#include "asyncarray.h"


// AsyncArray_New creates a new empty AsyncArray
AsyncArray AsyncArray_New() {
    return new cv::AsyncArray();
}

// AsyncArray_Close deletes an existing AsyncArray
void AsyncArray_Close(AsyncArray a) {
    delete a;
}

const char* AsyncArray_GetAsync(AsyncArray async_out,Mat out) {
    try {
       async_out->get(*out);
    } catch(cv::Exception ex) {
        return ex.err.c_str();
    }
    return "";
}

AsyncArray Net_forwardAsync(Net net, const char* outputName) {
    return new cv::AsyncArray(net->forwardAsync(outputName));
}
