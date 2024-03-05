#include "version.h"

const char* openCVVersion() {
    return CV_VERSION;
}

const char* getBuildInfo(){
    return cv::getBuildInformation().c_str();
}
