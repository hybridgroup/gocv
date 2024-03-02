//
// Created by rainy on 2024/3/2.
//
#include "exception.h"

int ErrorCallbackProxy(int status, const char *func_name,
                       const char *err_msg, const char *file_name,
                       int line, void *callback){
    ((ErrorCallback) callback)(status, func_name, err_msg, file_name, line, nullptr);
    return status;
}

void registerErrorCallback(ErrorCallback callback){
    cv::redirectError(ErrorCallbackProxy, (void*)callback);
}

int CvException_GetCode(CvException exception) {
    return exception->code;
}

const char *CvException_GetErr(CvException exception) {
    return exception->err.c_str();
}

const char *CvException_GetFunc(CvException exception) {
    return exception->func.c_str();
}

const char *CvException_GetFile(CvException exception) {
    return exception->file.c_str();
}

int CvException_GetLine(CvException exception) {
    return exception->line;
}

void CvException_Close(CvException exception){
    delete exception;
}