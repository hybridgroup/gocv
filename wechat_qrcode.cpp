#include "wechat_qrcode.h"

WeChatQRCode NewWeChatQRCode(const char *detector_prototxt_path,
                             const char *detector_caffe_model_path,
                             const char *super_resolution_prototxt_path,
                             const char *super_resolution_caffe_model_path) {
    return new cv::Ptr<cv::wechat_qrcode::WeChatQRCode>(
            cv::makePtr<cv::wechat_qrcode::WeChatQRCode>(detector_prototxt_path, detector_caffe_model_path,
                                                         super_resolution_prototxt_path,
                                                         super_resolution_caffe_model_path));
}

StringsVector NewStringsVector() {
    return new std::vector<std::string>;
}

CStrings WeChatQRCode_DetectAndDecode(WeChatQRCode wq, Mat img, struct Mats *points, StringsVector codes) {
    std::vector <cv::Mat> Points;
    *codes = ((*wq)->detectAndDecode(*img, Points));
    CStrings results;

    points->mats = new Mat[Points.size()];

    for (size_t i = 0; i < Points.size(); ++i) {
        points->mats[i] = new cv::Mat(Points[i]);
    }
    points->length = (int) Points.size();


    const char **decodes = new const char *[codes->size()];


    for (size_t i = 0; i < codes->size(); ++i) {
        decodes[i] = (*codes)[i].c_str();
    }
    (&results)->length = codes->size();
    (&results)->strs = decodes;
    return results;
}