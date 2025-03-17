#include <stdexcept>
#include "videoio.h"

// VideoWriter
VideoCapture VideoCapture_New() {
    try {
        return new cv::VideoCapture();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return NULL;
    }
}

void VideoCapture_Close(VideoCapture v) {
    delete v;
}

bool VideoCapture_Open(VideoCapture v, const char* uri) {
    try {
        return v->open(uri);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return false;
    }
}

bool VideoCapture_OpenWithAPI(VideoCapture v, const char* uri, int apiPreference) {
    try {
        return v->open(uri, apiPreference);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return false;
    }
}

bool VideoCapture_OpenWithAPIParams(VideoCapture v, const char* uri, int apiPreference, int *paramsv, int paramsc) {
    try {
        std::vector< int > params;

        for( int i = 0; i< paramsc; i++) {
            params.push_back(paramsv[i]);
        }
    
        return v->open(cv::String(uri), apiPreference, params);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return false;
    }
}

bool VideoCapture_OpenDevice(VideoCapture v, int device) {
    try {
        return v->open(device);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return false;
    }
}

bool VideoCapture_OpenDeviceWithAPI(VideoCapture v, int device, int apiPreference) {
    try {
        return v->open(device, apiPreference);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return false;
    }
}

bool VideoCapture_OpenDeviceWithAPIParams(VideoCapture v, int device, int apiPreference, int *paramsv, int paramsc) {
    try {
        std::vector< int > params;

        for( int i = 0; i< paramsc; i++) {
            params.push_back(paramsv[i]);
        }
    
        return v->open(device, apiPreference, params);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return false;
    }
}


OpenCVResult VideoCapture_Set(VideoCapture v, int prop, double param) {
    try {
        v->set(prop, param);
		OpenCVResult result = {0, NULL};
		return result;
	} catch(const cv::Exception& e) {
		OpenCVResult result = {e.code, e.what()};
		return result;
    }
}

double VideoCapture_Get(VideoCapture v, int prop) {
    try {
        return v->get(prop);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return 0.0;
    }
}

int VideoCapture_IsOpened(VideoCapture v) {
    try {
        return v->isOpened();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return 0;
    }
}

int VideoCapture_Read(VideoCapture v, Mat buf) {
    try {
        return v->read(*buf);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return 0;
    }
}

OpenCVResult VideoCapture_Grab(VideoCapture v, int skip) {
    try {
        for (int i = 0; i < skip; i++) {
            v->grab();
        }
		OpenCVResult result = {0, NULL};
		return result;
	} catch(const cv::Exception& e) {
		OpenCVResult result = {e.code, e.what()};
		return result;
    }
}

int VideoCapture_Retrieve(VideoCapture v, Mat buf) {
    try {
        return v->retrieve(*buf);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return 0;
    }
}

// VideoWriter
VideoWriter VideoWriter_New() {
    try {
        return new cv::VideoWriter();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return NULL;
    }
}

void VideoWriter_Close(VideoWriter vw) {
    delete vw;
}

OpenCVResult VideoWriter_Open(VideoWriter vw, const char* name, const char* codec, double fps, int width, int height, bool isColor) {
    try {
        int codecCode = cv::VideoWriter::fourcc(codec[0], codec[1], codec[2], codec[3]);
        vw->open(name, codecCode, fps, cv::Size(width, height), isColor);
		OpenCVResult result = {0, NULL};
		return result;
	} catch(const cv::Exception& e) {
		OpenCVResult result = {e.code, e.what()};
		return result;
    }
}

OpenCVResult VideoWriter_OpenWithAPI(VideoWriter vw, const char* name, int apiPreference, const char* codec, double fps, int width, int height, bool isColor) {
    try {
        int codecCode = cv::VideoWriter::fourcc(codec[0], codec[1], codec[2], codec[3]);
        vw->open(name, apiPreference, codecCode, fps, cv::Size(width, height), isColor);
		OpenCVResult result = {0, NULL};
		return result;
	} catch(const cv::Exception& e) {
		OpenCVResult result = {e.code, e.what()};
		return result;
    }
}

OpenCVResult VideoWriter_OpenWithAPIParams(VideoWriter vw, const char* name, int apiPreference, const char* codec, double fps, int width, int height, IntVector params) {
    try {
        std::vector<int> cpp_params;

        for(int i = 0; i < params.length; i++) {
            cpp_params.push_back(params.val[i]);
        }
    
        int codecCode = cv::VideoWriter::fourcc(codec[0], codec[1], codec[2], codec[3]);
        vw->open(name, apiPreference, codecCode, fps, cv::Size(width, height), cpp_params);
		OpenCVResult result = {0, NULL};
		return result;
	} catch(const cv::Exception& e) {
		OpenCVResult result = {e.code, e.what()};
		return result;
    }
}

int VideoWriter_IsOpened(VideoWriter vw) {
    try {
        return vw->isOpened();
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return 0;
    }
}

OpenCVResult VideoWriter_Write(VideoWriter vw, Mat img) {
    try {
        *vw << *img;
		OpenCVResult result = {0, NULL};
		return result;
	} catch(const cv::Exception& e) {
		OpenCVResult result = {e.code, e.what()};
		return result;
    }
}

const char* Videoio_Registry_GetBackendName(int api) {
    try {
        cv::String name;
        name = cv::videoio_registry::getBackendName((cv::VideoCaptureAPIs)(api));
        return strdup(name.c_str());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return "";
    }
}

IntVector Videio_Registry_GetBackends() {
    try {
        IntVector c_backs;
        std::vector<cv::VideoCaptureAPIs> backs = cv::videoio_registry::getBackends();
    
        c_backs.val = new int[backs.size()];
        c_backs.length = backs.size();
    
        for(int i = 0; i < c_backs.length; i++) {
            c_backs.val[i] = backs[i];
        }
    
        return c_backs;
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        IntVector empty;
		return empty;
    }
}

const char* Videoio_Registry_GetCameraBackendPluginVersion(int api, int* version_ABI, int* version_API) {
    try {
        std::string desc = cv::videoio_registry::getCameraBackendPluginVersion((cv::VideoCaptureAPIs)(api), *version_ABI, *version_API);

        return strdup(desc.c_str());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return "";
    }
}

IntVector Videoio_Registry_GetCameraBackends() {
    try {
        IntVector c_backs;
        std::vector<cv::VideoCaptureAPIs> backs = cv::videoio_registry::getCameraBackends();
    
        c_backs.val = new int[backs.size()];
        c_backs.length = backs.size();
    
        for(int i = 0; i < c_backs.length; i++) {
            c_backs.val[i] = backs[i];
        }
    
        return c_backs;
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        IntVector empty;
		return empty;
    }
}

const char* Videoio_Registry_GetStreamBackendPluginVersion(int api, int* version_ABI, int* version_API){
    try {
        std::string desc = cv::videoio_registry::getStreamBackendPluginVersion((cv::VideoCaptureAPIs)(api), *version_ABI, *version_API);

        return strdup(desc.c_str());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return "";
    } 
}

IntVector Videoio_Registry_GetStreamBackends() {
    try {
        IntVector c_backs;

        std::vector<cv::VideoCaptureAPIs> backs = cv::videoio_registry::getStreamBackends();
    
        c_backs.val = new int[backs.size()];
        c_backs.length = backs.size();
    
        for(int i = 0; i < c_backs.length; i++) {
            c_backs.val[i] = backs[i];
        }
    
        return c_backs;
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        IntVector empty;
		return empty;
    }
}

const char* Videoio_Registry_GetWriterBackendPluginVersion(int api, int* version_ABI, int* version_API){
    try {
        std::string desc = cv::videoio_registry::getWriterBackendPluginVersion((cv::VideoCaptureAPIs)(api), *version_ABI, *version_API);

        return strdup(desc.c_str());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return "";
    }
}

IntVector Videoio_Registry_GetWriterBackends() {
    try {
        IntVector c_backs;

        std::vector<cv::VideoCaptureAPIs> backs = cv::videoio_registry::getWriterBackends();
    
        c_backs.val = new int[backs.size()];
        c_backs.length = backs.size();
    
        for(int i = 0; i < c_backs.length; i++) {
            c_backs.val[i] = backs[i];
        }
    
        return c_backs;
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        IntVector empty;
		return empty;
    }
}

bool Videoio_Registry_HasBackend(int api) {
    try {
        return cv::videoio_registry::hasBackend((cv::VideoCaptureAPIs)(api));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return false;
    }
}

bool Videoio_Registry_IsBackendBuiltIn(int api) {
    try {
        return cv::videoio_registry::isBackendBuiltIn((cv::VideoCaptureAPIs)(api));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return false;
    }
}
