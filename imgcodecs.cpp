#include "imgcodecs.h"

// Image
Mat Image_IMRead(const char* filename, int flags) {
    cv::Mat img = cv::imread(filename, flags);
    return new cv::Mat(img);
}

bool Image_IMWrite(const char* filename, Mat img, IntVector params) {
    std::vector<int> compression_params;
    for(int* i = params.val; *i; ++i) {
        compression_params.push_back(*i);
    }
    return cv::imwrite(filename, *img, compression_params);
}

struct ByteArray Image_IMEncode(const char* fileExt, Mat img) {
    std::vector<uchar> data;
    cv::imencode(fileExt, *img, data);
    return toByteArray(reinterpret_cast<const char*>(&data[0]), data.size());
}
