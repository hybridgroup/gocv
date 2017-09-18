#ifndef _OPENCV3_UTIL_H_
#define _OPENCV3_UTIL_H_

#ifdef __cplusplus
extern "C" {
#endif

typedef struct String {
  const char* str;
  int length;
} String;
typedef struct ByteArray{
  char *data;
  int length;
} ByteArray;

struct ByteArray toByteArray(const char* buf, int len);
void ByteArray_Release(struct ByteArray buf);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_UTIL_H_