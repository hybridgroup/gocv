uname_val="$(uname)"
if [[ "$uname_val" == "Darwin" ]]; then
  CVPATH=$(brew info opencv | sed -n "4p" | sed -e "s/ (.*//g")
  export CGO_CPPFLAGS="-I$CVPATH/include -I$CVPATH/include/opencv2"
  export CGO_CXXFLAGS="--std=c++1z -stdlib=libc++"
  export CGO_LDFLAGS="-L$CVPATH/lib -lopencv_core -lopencv_videoio -lopencv_imgproc -lopencv_highgui -lopencv_imgcodecs -lopencv_objdetect -lopencv_features2d -lopencv_video"
  echo "Environment variables configured for OSX"
elif [[ "$uname_val" == "Linux" ]]; then
        if [[ -f /etc/pacman.conf ]]; then
                export CGO_CPPFLAGS="-I/usr/include"
                export CGO_CXXFLAGS="--std=c++1z"
                export CGO_LDFLAGS="-L/lib64 -lopencv_core -lopencv_videoio -lopencv_imgproc -lopencv_highgui -lopencv_imgcodecs -lopencv_objdetect -lopencv_features2d -lopencv_video"
        else
                export CGO_CPPFLAGS="-I/usr/local/include"
                export CGO_CXXFLAGS="--std=c++1z"
                export CGO_LDFLAGS="-L/usr/local/lib -lopencv_core -lopencv_videoio -lopencv_imgproc -lopencv_highgui -lopencv_imgcodecs -lopencv_objdetect -lopencv_features2d -lopencv_video"
        fi
  echo "Environment variables configured for Linux"
else
  echo "Unknown platform '$uname_val'!"
fi
