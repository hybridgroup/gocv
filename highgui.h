#ifndef _OPENCV3_HIGHGUI_H_
#define _OPENCV3_HIGHGUI_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
extern "C" {
#endif

#include "core.h"

// Window
void Window_New(const char* winname, int flags);
//void Window_Delete(const char* winname);
int Window_WaitKey(int);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_HIGHGUI_H_
