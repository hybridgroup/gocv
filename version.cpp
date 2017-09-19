#include "version.h"
#include <sstream>
#include <string>
#include <cstring>

using namespace std;

char *openCVVersion() {
    ostringstream os;
    os << CV_VERSION;
    return const_cast<char*>(os.str().c_str());
}
