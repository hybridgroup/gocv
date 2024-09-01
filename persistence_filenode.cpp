#include <string.h>
#include "persistence.h"

bool FileNode_Empty(FileNode fn) {
    return fn->empty();
}

bool FileNode_IsInt(FileNode fn){
    return fn->isInt();
}

bool FileNode_IsMap(FileNode fn){
    return fn->isMap();
}

bool FileNode_IsNamed(FileNode fn) {
    return fn->isNamed();
}

bool FileNode_IsNone(FileNode fn){
    return fn->isNone();
}

bool FileNode_IsReal(FileNode fn){
    return fn->isReal();
}

bool FileNode_IsSeq(FileNode fn) {
    return fn->isSeq();
}

bool FileNode_IsString(FileNode fn) {
    return fn->isString();
}

char** FileNode_Keys(FileNode fn) {

    std::vector<cv::String> keys = fn->keys();

    char** c_keys = new char*[keys.size()];

    for (int i = 0; i < keys.size(); i++) {
        char *c_key = new char[keys[i].length()+1];
        strcpy(c_key, keys[i].c_str());
        c_keys[i] = c_key;
    }

    return c_keys;
}

size_t FileNode_KeysCount(FileNode fn) {
    return fn->keys().size();
}


void FileNode_KeysFree(char** keys, size_t len) {
    for(int i = 0; i < len; i++) {
        delete keys[i];
    }
    delete keys;
}

Mat FileNode_Mat(FileNode fn) {
    return new cv::Mat(fn->mat());
}

const char* FileNode_Name(FileNode fn) {
    char* str = new char[fn->name().length()+1];
    strcpy(str, fn->name().c_str()); 
    return str;
}

float FileNode_Float(FileNode fn) {
    return (float)fn->real();
}

const char* FileNode_String(FileNode fn) {
    char* str = new char[fn->string().length()+1];
    strcpy(str, fn->string().c_str());
    return str;
}

FileNode FileNode_Get(FileNode fn, int i) {
    return new cv::FileNode((*fn)[i]);
}

FileNode FileNode_GetByName(FileNode fn, const char* nodename) {
    return new cv::FileNode((*fn)[nodename]);
}

size_t FileNode_RawSize(FileNode fn) {
    return fn->rawSize();
}

void FileNode_ReadRaw(FileNode fn, const char* fmt, void *vec, size_t len) {
    fn->readRaw(fmt, vec, len);
}
 
void FileNode_SetValue(FileNode fn, int type, const void *value, int len) {
    fn->setValue(type, value, len);
}

size_t FileNode_Size(FileNode fn) {
    return fn->size();
}

int FileNode_Type(FileNode fn) {
    return fn->type();
}

void FileNode_Close(FileNode fn){
    delete fn;
}
