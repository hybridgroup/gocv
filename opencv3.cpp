#include "opencv3.h"

#include <string.h>

MatVec3b MatVec3b_New() {
  return new cv::Mat_<cv::Vec3b>();
}

struct ByteArray MatVec3b_ToJpegData(MatVec3b m, int quality){
  std::vector<int> param(2);
  param[0] = CV_IMWRITE_JPEG_QUALITY;
  param[1] = quality;
  std::vector<uchar> data;
  cv::imencode(".jpg", *m, data, param);
  return toByteArray(reinterpret_cast<const char*>(&data[0]), data.size());
}

void MatVec3b_Delete(MatVec3b m) {
  delete m;
}

void MatVec3b_CopyTo(MatVec3b src, MatVec3b dst) {
  src->copyTo(*dst);
}

int MatVec3b_Empty(MatVec3b m) {
  return m->empty();
}

struct RawData MatVec3b_ToRawData(MatVec3b m) {
  int width = m->cols;
  int height = m->rows;
  int size = width * height * 3;
  char* data = reinterpret_cast<char*>(m->data);
  ByteArray byteData = {data, size};
  RawData raw = {width, height, byteData};
  return raw;
}

MatVec3b RawData_ToMatVec3b(struct RawData r) {
  int rows = r.height;
  int cols = r.width;
  cv::Mat_<cv::Vec3b>* mat = new cv::Mat_<cv::Vec3b>(rows, cols);
  unsigned char* data = reinterpret_cast<unsigned char*>(r.data.data);
  mat->data = data;
  return mat;
}

void MatVec4b_Delete(MatVec4b m) {
  delete m;
}

struct RawData MatVec4b_ToRawData(MatVec4b m) {
  int width = m->cols;
  int height = m->rows;
  int size = width * height * 4;
  char* data = reinterpret_cast<char*>(m->data);
  ByteArray byteData = {data, size};
  RawData raw = {width, height, byteData};
  return raw;
}

MatVec4b RawData_ToMatVec4b(struct RawData r) {
  int rows = r.height;
  int cols = r.width;
  cv::Mat_<cv::Vec4b>* mat = new cv::Mat_<cv::Vec4b>(rows, cols);
  unsigned char* data = reinterpret_cast<unsigned char*>(r.data.data);
  mat->data = data;
  return mat;
}

VideoCapture VideoCapture_New() {
  return new cv::VideoCapture();
}

void VideoCapture_Delete(VideoCapture v) {
  delete v;
}

int VideoCapture_Open(VideoCapture v, const char* uri) {
  return v->open(uri);
}

int VideoCapture_OpenDevice(VideoCapture v, int device) {
  return v->open(device);
}

void VideoCapture_Release(VideoCapture v) {
  v->release();
}

void VideoCapture_Set(VideoCapture v, int prop, int param) {
  v->set(prop, param);
}

int VideoCapture_IsOpened(VideoCapture v) {
  return v->isOpened();
}

int VideoCapture_Read(VideoCapture v, MatVec3b buf) {
  return v->read(*buf);
}

void VideoCapture_Grab(VideoCapture v, int skip) {
  for (int i =0; i < skip; i++) {
    v->grab();
  }
}

VideoWriter VideoWriter_New() {
  return new cv::VideoWriter();
}

void VideoWriter_Delete(VideoWriter vw) {
  delete vw;
}

void VideoWriter_Open(VideoWriter vw, const char* name, double fps, int width,
    int height) {
  vw->open(name, CV_FOURCC('M', 'J', 'P', 'G'), fps, cv::Size(width, height), true);
}

void VideoWriter_OpenWithMat(VideoWriter vw, const char* name, double fps,
    MatVec3b img) {
  vw->open(name, CV_FOURCC('M', 'J', 'P', 'G'), fps, img->size(), true);
}

int VideoWriter_IsOpened(VideoWriter vw) {
  return vw->isOpened();
}

void VideoWriter_Write(VideoWriter vw, MatVec3b img) {
  *vw << *img;
}

CascadeClassifier CascadeClassifier_New() {
  return new cv::CascadeClassifier();
}

void CascadeClassifier_Delete(CascadeClassifier cs) {
  delete cs;
}

int CascadeClassifier_Load(CascadeClassifier cs, const char* name) {
  return cs->load(name);
}

struct Rects CascadeClassifier_DetectMultiScale(CascadeClassifier cs, MatVec3b img) {
  std::vector<cv::Rect> faces;
  cs->detectMultiScale(*img, faces); // TODO control default parameter
  Rect* rects = new Rect[faces.size()];
  for (size_t i = 0; i < faces.size(); ++i) {
    Rect r = {faces[i].x, faces[i].y, faces[i].width, faces[i].height};
    rects[i] = r;
  }
  Rects ret = {rects, (int)faces.size()};
  return ret;
}

void Rects_Delete(struct Rects rs) {
  delete rs.rects;
}

void DrawRectsToImage(MatVec3b img, struct Rects rects) {
  for (int i = 0; i < rects.length; ++i) {
    Rect r = rects.rects[i];
    cv::rectangle(*img, cv::Point(r.x, r.y), cv::Point(r.x+r.width, r.y+r.height),
      cv::Scalar(0, 200, 0), 3, CV_AA);
  }
}

MatVec4b LoadAlphaImg(const char* name) {
  cv::Mat_<cv::Vec4b> img = cv::imread(name, cv::IMREAD_UNCHANGED);
  return new cv::Mat_<cv::Vec4b>(img);
}

void MountAlphaImage(MatVec4b img, MatVec3b back, struct Rects rects) {
  cv::Mat_<cv::Vec4b> resized;
  for (int i = 0; i < rects.length; ++i) {
    Rect r = rects.rects[i];
    int col, row;
    if (r.width < r.height) {
      col = img->cols * r.height / img->rows;
      row = r.height;
    } else {
      col = r.width;
      row = img->rows * r.width / img->cols;
    }
    int ltx = r.x + r.width * 0.5 - col * 0.5;
    int lty = r.y + r.height * 0.5 - row * 0.5;
    std::vector<cv::Point2f> tgtPt;
    tgtPt.push_back(cv::Point2f(ltx, lty));
    tgtPt.push_back(cv::Point2f(ltx+col, lty));
    tgtPt.push_back(cv::Point2f(ltx+col, lty+row));
    tgtPt.push_back(cv::Point2f(ltx, lty+row));

    cv::Mat img_rgb, img_aaa, img_backa;
    std::vector<cv::Mat> planes_rgba, planes_rgb, planes_aaa, planes_backa;
    int maxVal = pow(2, 8 * back->elemSize1()) - 1;

    std::vector<cv::Point2f> srcPt;
    srcPt.push_back(cv::Point2f(0, 0));
    srcPt.push_back(cv::Point2f(img->cols-1, 0));
    srcPt.push_back(cv::Point2f(img->cols-1, img->rows-1));
    srcPt.push_back(cv::Point2f(0, img->rows-1));
    cv::Mat mat = cv::getPerspectiveTransform(srcPt, tgtPt);

    cv::Mat alpha0(back->rows, back->cols, img->type());
    alpha0 = cv::Scalar::all(0);
    cv::warpPerspective(*img, alpha0, mat, alpha0.size(), cv::INTER_CUBIC,
      cv::BORDER_TRANSPARENT);

    cv::split(alpha0, planes_rgba);

    planes_rgb.push_back(planes_rgba[0]);
    planes_rgb.push_back(planes_rgba[1]);
    planes_rgb.push_back(planes_rgba[2]);
    merge(planes_rgb, img_rgb);

    planes_aaa.push_back(planes_rgba[3]);
    planes_aaa.push_back(planes_rgba[3]);
    planes_aaa.push_back(planes_rgba[3]);
    merge(planes_aaa, img_aaa);

    planes_backa.push_back(maxVal - planes_rgba[3]);
    planes_backa.push_back(maxVal - planes_rgba[3]);
    planes_backa.push_back(maxVal - planes_rgba[3]);
    merge(planes_backa, img_backa);

    *back = img_rgb.mul(img_aaa, 1.0/(float)maxVal)
      + back->mul(img_backa, 1.0/(float)maxVal);
  }
}
