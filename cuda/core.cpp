#include "../core.h"
#include "core.h"
#include <string.h>

void GpuRects_Close(struct Rects rs) {
    delete[] rs.rects;
}
