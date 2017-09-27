#include "core.h"
#include <string.h>

// Mat_New creates a new empty Mat
Mat Mat_New() {
    return new cv::Mat();
}

// Mat_Close deletes an existing Mat
void Mat_Close(Mat m) {
    delete m;
}

// Mat_Empty tests if a Mat is empty
int Mat_Empty(Mat m) {
    return m->empty();
}

void Rects_Close(struct Rects rs) {
    delete rs.rects;
}

void DrawRectsToImage(Mat img, struct Rects rects) {
    for (int i = 0; i < rects.length; ++i) {
        Rect r = rects.rects[i];
        cv::rectangle(*img, cv::Point(r.x, r.y), cv::Point(r.x+r.width, r.y+r.height),
            cv::Scalar(0, 200, 0), 3, CV_AA);
    }
}

void ByteArray_Release(struct ByteArray buf) {
  delete[] buf.data;
}

struct ByteArray toByteArray(const char* buf, int len) {
  ByteArray ret = {new char[len], len};
  memcpy(ret.data, buf, len);
  return ret;
}
