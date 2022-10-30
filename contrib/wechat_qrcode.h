#ifndef _OPENCV4_WECHAT_QRCODE_H_
#define _OPENCV4_WECHAT_QRCODE_H_

#ifdef __cplusplus

#include <opencv2/opencv.hpp>
#include <opencv2/wechat_qrcode.hpp>

extern "C" {
#endif

#include "../core.h"

#ifdef __cplusplus
typedef cv::Ptr<cv::wechat_qrcode::WeChatQRCode> *WeChatQRCode;
typedef std::vector<std::string> *StringsVector;
#else
typedef void* WeChatQRCode;
typedef void* StringsVector;
#endif

WeChatQRCode NewWeChatQRCode(const char *detector_prototxt_path, const char *detector_caffe_model_path,
                             const char *super_resolution_prototxt_path, const char *super_resolution_caffe_model_path);
CStrings WeChatQRCode_DetectAndDecode(WeChatQRCode wq, Mat img, struct Mats *points, StringsVector codes);
StringsVector NewStringsVector();
void WeChatQRCode_CStrings_Close(struct CStrings cstrs);
void WeChatQRCode_Mats_Close(struct Mats mats);
void WeChatQRCode_Mats_to(struct Mats mats, int i, Mat dst);


#ifdef __cplusplus
}
#endif

#endif //_OPENCV4_WECHAT_QRCODE_H_