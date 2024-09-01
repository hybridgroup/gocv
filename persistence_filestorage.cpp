#include <string.h>
#include "persistence.h"


FileStorage FileStorage_Create(void) {
    return new cv::FileStorage();
}

FileStorage FileStorage_CreateWithParams(const char* filename, int flags, const char* encoding) {
    return new cv::FileStorage(filename, flags, encoding);
}

const char *FileStorage_ElName(FileStorage fs) {
    char* str = new char[fs->elname.length()+1];
    strcpy(str, fs->elname.c_str()); 
    return str;
}
int FileStorage_State(FileStorage fs) {
    return fs->state;
}

void FileStorage_Close(FileStorage fs) {
    fs->release();
    delete fs;
}

void FileStorage_EndWriteStruct(FileStorage fs) {
    fs->endWriteStruct();
}

int FileStorage_GetFormat(FileStorage fs){
    return fs->getFormat();
}

bool FileStorage_IsOpened(FileStorage fs) {
    return fs->isOpened();
}

bool FileStorage_Open(FileStorage fs, const char* filename, int flags, const char* encoding) {
    return fs->open(filename, flags, encoding);
}

void FileStorage_Release(FileStorage fs) {
    fs->release();
    delete fs;
}

const char* FileStorage_ReleaseAndGetString(FileStorage fs) {
    cv::String s = fs->releaseAndGetString();

    char* str = new char[s.length()+1];
    strcpy(str, s.c_str()); 
    return str;
}

void FileStorage_StartWriteStruct(FileStorage fs, const char* name, int flags, const char* typeName){
    fs->startWriteStruct(name, flags, typeName);
}

void FileStorage_WriteMat(FileStorage fs, const char* name, Mat val){
    fs->write(name, *val);
}

void FileStorage_WriteString(FileStorage fs, const char* name, const char* val) {
    fs->write(name, val);
}

void FileStorage_WriteStringArray(FileStorage fs, const char* name, const char** val, size_t len) {
    std::vector<cv::String> vals;

    for(int i = 0; i < len; i++) {
        vals.push_back(val[i]);
    }

    fs->write(name, vals);
}

void FileStorage_WriteDouble(FileStorage fs, const char* name, double val){
    fs->write(name, val);
}

void FileStorage_WriteInt(FileStorage fs, const char* name, int val){
    fs->write(name, val);
}

void FileStorage_WriteComment(FileStorage fs, const char* comment, bool append){
    fs->writeComment(comment, append);
}
void FileStorage_WriteRaw(FileStorage fs, const char* fmt, const void* vec, size_t len){
    fs->writeRaw(fmt, vec, len);
}

FileNode FileStorage_GetFirstTopLevelNode(FileStorage fs) {
    cv::FileNode node = fs->getFirstTopLevelNode();

    FileNode fn = new cv::FileNode(node);
    return fn;
}

FileNode FileStorage_GetNode(FileStorage fs, const char* nodename) {

    cv::FileNode node = (*fs)[nodename];

    FileNode fn = new cv::FileNode(node);
    return fn;
}

FileNode FileStorage_Root(FileStorage fs, int streamidx) {
    cv::FileNode node = fs->root(streamidx);

    FileNode fn = new cv::FileNode(node);
    return fn;
}