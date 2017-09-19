# Go OpenCV3

Go bindings for the [OpenCV 3](http://opencv.org/) computer vision package.

Supports the latest OpenCV v3.3

Uses C-style wrapper around the OpenCV 3 C++ classes to avoid having to deal with applying SWIG to a huge existing codebase.

The mappings are intended to match as close as possible to the original OpenCV project structure.

For example, the [OpenCV `videoio` module](https://github.com/opencv/opencv/tree/master/modules/videoio) wrappers can be found in this project in the `videoio.*` files.

Based on concepts & code from the blog post https://medium.com/@peterleyssens/using-opencv-3-from-golang-5510c312a3c and the repo at https://github.com/sensorbee/opencv
