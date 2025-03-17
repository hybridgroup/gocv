#include <stdlib.h>
#include "imgcodecs.h"

// Image
Mat Image_IMRead(const char* filename, int flags) {
    try {
        cv::Mat img = cv::imread(filename, flags);
        return new cv::Mat(img);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return new cv::Mat();
    }
}

Mats Image_IMReadMulti(const char* filename, int flags) {
    try {
        std::vector<cv::Mat> dst;
        Mats m = Mats();
    
        bool b = cv::imreadmulti(filename, dst, flags);
        if (b) {
            m.mats = new Mat[dst.size()];
            for (size_t i = 0; i < dst.size(); ++i) {
                m.mats[i] = new cv::Mat(dst[i]);
            }
            m.length = (int)dst.size();        
        }
    
        return m;
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return Mats();
    }
}

Mats Image_IMReadMulti_WithParams(const char* filename, int start, int count, int flags) {
    try {
        std::vector<cv::Mat> dst;
        auto m = Mats();
        
        auto b = cv::imreadmulti(filename, dst, start, count, flags);
        if (b) {
            m.mats = new Mat[dst.size()];
            for (size_t i = 0; i < dst.size(); ++i) {
                m.mats[i] = new cv::Mat(dst[i]);
            }
            m.length = (int)dst.size();        
        }
    
        return m;
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return Mats();
    }
}


bool Image_IMWrite(const char* filename, Mat img) {
    try {
        return cv::imwrite(filename, *img);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return false;
    }
}

bool Image_IMWrite_WithParams(const char* filename, Mat img, IntVector params) {
    try {
        std::vector<int> compression_params;

        for (int i = 0, *v = params.val; i < params.length; ++v, ++i) {
            compression_params.push_back(*v);
        }
    
        return cv::imwrite(filename, *img, compression_params);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return false;
    }
}

OpenCVResult Image_IMEncode(const char* fileExt, Mat img, void* vector) {
    try {
        auto vectorPtr = reinterpret_cast<std::vector<uchar> *>(vector);
        cv::imencode(fileExt, *img, *vectorPtr);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

OpenCVResult Image_IMEncode_WithParams(const char* fileExt, Mat img, IntVector params, void* vector) {
    try {
        auto vectorPtr = reinterpret_cast<std::vector<uchar> *>(vector);
        std::vector<int> compression_params;
    
        for (int i = 0, *v = params.val; i < params.length; ++v, ++i) {
            compression_params.push_back(*v);
        }
    
        cv::imencode(fileExt, *img, *vectorPtr, compression_params);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}

Mat Image_IMDecode(ByteArray buf, int flags) {
    try {
        std::vector<uchar> data(buf.data, buf.data + buf.length);
        cv::Mat img = cv::imdecode(data, flags);
        return new cv::Mat(img);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
        return new cv::Mat();
    }
}

OpenCVResult Image_IMDecodeIntoMat(ByteArray buf, int flags, Mat dest) {
    try {
        std::vector<uchar> data(buf.data, buf.data + buf.length);
        cv::imdecode(data, flags, dest);
        return successResult();
    } catch(const cv::Exception& e) {
        return errorResult(e.code, e.what());
    }
}
