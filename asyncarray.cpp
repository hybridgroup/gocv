// +build openvino

#include <string.h>
#include "asyncarray.h"


// AsyncArray_New creates a new empty AsyncArray
AsyncArray AsyncArray_New() {
    try {
        return new cv::AsyncArray();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}

// AsyncArray_Close deletes an existing AsyncArray
void AsyncArray_Close(AsyncArray a) {
    delete a;
}

const char* AsyncArray_GetAsync(AsyncArray async_out,Mat out) {
    try {
       async_out->get(*out);
    } catch(const cv::Exception& ex) {
        return ex.what();
    }
    return "";
}

AsyncArray Net_forwardAsync(Net net, const char* outputName) {
    try {
        return new cv::AsyncArray(net->forwardAsync(outputName));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return NULL;
    }
}
