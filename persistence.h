#ifndef _OPENCV3_OBJDETECT_H_
#define _OPENCV3_OBJDETECT_H_

#include <stdbool.h>

#ifdef __cplusplus
#include <opencv2/opencv.hpp>
extern "C" {
#endif

#include "core.h"

#ifdef __cplusplus
typedef cv::FileStorage* FileStorage;
typedef cv::FileNode* FileNode;
#else
typedef void* FileStorage;
typedef void* FileNode;
#endif

// FileStorage
FileStorage FileStorage_Create(void);
FileStorage FileStorage_CreateWithParams(const char* filename, int flags, const char* encoding);
void FileStorage_Close(FileStorage fs);

const char *FileStorage_ElName(FileStorage fs);
int FileStorage_State(FileStorage fs);

void FileStorage_EndWriteStruct(FileStorage fs);
int FileStorage_GetFormat(FileStorage fs);
bool FileStorage_IsOpened(FileStorage fs);
bool FileStorage_Open(FileStorage fs, const char* filename, int flags, const char* encoding);
void FileStorage_Release(FileStorage fs);
const char* FileStorage_ReleaseAndGetString(FileStorage fs);
void FileStorage_StartWriteStruct(FileStorage fs, const char* name, int flags, const char* typeName);
void FileStorage_WriteMat(FileStorage fs, const char* name, Mat val);
void FileStorage_WriteString(FileStorage fs, const char* name, const char* val);
void FileStorage_WriteStringArray(FileStorage fs, const char* name, const char** val, size_t len);
void FileStorage_WriteDouble(FileStorage fs, const char* name, double val);
void FileStorage_WriteInt(FileStorage fs, const char* name, int val);
void FileStorage_WriteComment(FileStorage fs, const char* comment, bool append);
void FileStorage_WriteRaw(FileStorage fs, const char* fmt, const void* vec, size_t len);

FileNode FileStorage_GetFirstTopLevelNode(FileStorage fs);
FileNode FileStorage_GetNode(FileStorage fs, const char* nodename);
FileNode FileStorage_Root(FileStorage fs, int streamidx);

bool FileNode_Empty(FileNode fn);
bool FileNode_IsInt(FileNode fn);
bool FileNode_IsMap(FileNode fn);
bool FileNode_IsNamed(FileNode fn);
bool FileNode_IsNone(FileNode fn);
bool FileNode_IsReal(FileNode fn);
bool FileNode_IsSeq(FileNode fn);
bool FileNode_IsString(FileNode fn);
char** FileNode_Keys(FileNode fn);
size_t FileNode_KeysCount(FileNode fn);
void FileNode_KeysFree(char** keys, size_t len);
Mat FileNode_Mat(FileNode fn);
const char* FileNode_Name(FileNode fn);
float FileNode_Float(FileNode fn);
const char* FileNode_String(FileNode fn);
FileNode FileNode_Get(FileNode fn, int i); //FileNode operator[] (int i) const
FileNode FileNode_GetByName(FileNode fn, const char* nodename); //FileNode operator[] (const char *nodename) const
size_t FileNode_RawSize(FileNode fn);
void FileNode_ReadRaw(FileNode fn, const char* fmt, void *vec, size_t len);
void FileNode_SetValue(FileNode fn, int type, const void *value, int len);
size_t FileNode_Size(FileNode fn);
int FileNode_Type(FileNode fn);

void FileNode_Close(FileNode fn);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_OBJDETECT_H_
