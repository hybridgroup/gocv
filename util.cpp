#include "util.h"
#include <string.h>
#include <cstdlib>

void ByteArray_Release(struct ByteArray buf) {
  delete[] buf.data;
}

struct ByteArray toByteArray(const char* buf, int len) {
  ByteArray ret = {new char[len], len};
  memcpy(ret.data, buf, len);
  return ret;
}
