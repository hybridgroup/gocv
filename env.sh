uname_val="$(uname)"
if [[ "$uname_val" == "Darwin" ]]; then
  export CGO_CPPFLAGS="-I/usr/local/Cellar/opencv/3.3.0_3/include -I/usr/local/Cellar/opencv/3.3.0_3/include/opencv2"
  export CGO_CXXFLAGS="--std=c++1z -stdlib=libc++"
  export CGO_LDFLAGS="-L/usr/local/Cellar/opencv/3.3.0_3/lib -lopencv_core -lopencv_videoio -lopencv_imgproc -lopencv_highgui -lopencv_imgcodecs -lopencv_objdetect -lopencv_calib3d"
  echo "Environment variables configured for OSX"
elif [[ "$uname_val" == "Linux" ]]; then
        if [[ -f /etc/pacman.conf ]]; then
                export CGO_CPPFLAGS="-I/usr/include"
                export CGO_CXXFLAGS="--std=c++1z"
                export CGO_LDFLAGS="-L/lib64 -lopencv_core -lopencv_videoio -lopencv_imgproc -lopencv_highgui -lopencv_imgcodecs -lopencv_objdetect -lopencv_calib3d"
        else
                export CGO_CPPFLAGS="-I/usr/local/include"
                export CGO_CXXFLAGS="--std=c++1z"
                export CGO_LDFLAGS="-L/usr/local/lib -lopencv_core -lopencv_videoio -lopencv_imgproc -lopencv_highgui -lopencv_imgcodecs -lopencv_objdetect -lopencv_calib3d"
        fi
  echo "Environment variables configured for Linux"
else
  echo "Unknown platform '$uname_val'!"
fi
