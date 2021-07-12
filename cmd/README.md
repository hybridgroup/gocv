# GoCV Commands

GoCV comes with various useful command line utilities, that are also examples of how to use the package.

## Basic Drawing

Demonstrates the basic drawing primitives available for drawing on images.

https://github.com/hybridgroup/gocv/blob/release/cmd/basic-drawing/main.go

## Caffe Classifier

Capture video from a connected webcam, then use the Caffe deep learning framework to classify whatever is in front of the camera.

https://github.com/hybridgroup/gocv/blob/release/cmd/caffe-classifier/main.go

## Capture test

Tests to verify you can capture video from a connected webcam.

https://github.com/hybridgroup/gocv/blob/release/cmd/captest/main.go

## Capture window

Capture video from a connected webcam and display the video in a Window.

https://github.com/hybridgroup/gocv/blob/release/cmd/capwindow/main.go

## Counter

Capture video from a pre-recorded file, and then count the number of detected objects that cross a user-definable vertical or horizontal line.

https://github.com/hybridgroup/gocv/blob/release/cmd/counter/main.go

## DNN Detection

Use a Deep Neural Network to detect and track objects or faces.

https://github.com/hybridgroup/gocv/blob/release/cmd/dnn-detection/main.go

## DNN Pose Detection

Use a Deep Neural Network trained using OpenPose to detect and track human body poses.

https://github.com/hybridgroup/gocv/blob/release/cmd/dnn-pose-detection/main.go

## DNN Style Transfer

Use a Deep Neural Network to perform real-time style transfer.

https://github.com/hybridgroup/gocv/blob/release/cmd/dnn-style-transfer/main.go

## Faceblur

Captures video from a connected camera, then uses the CascadeClassifier to detect faces, blurs them using a Gaussian blur, then displays the blurred video in a window.

https://github.com/hybridgroup/gocv/blob/release/cmd/faceblur/main.go

## Facedetect

Captures video from a connected camera, then uses the CascadeClassifier to detect faces, and draw a rectangle around each of them, before displaying them within a Window.

https://github.com/hybridgroup/gocv/blob/release/cmd/facedetect/main.go

## Facedetect from URL

Using a single image downloaded from a URL, it then uses the CascadeClassifier to detect faces, and draw a rectangle around each of them, before displaying them within a Window.

https://github.com/hybridgroup/gocv/blob/release/cmd/facedetect/main.go

## Feature Matching

Match features in an image using the SIFT algorithm.

https://github.com/hybridgroup/gocv/blob/release/cmd/feature-matching/main.go

## Find Chessboard

Find chessboard in an image using FindChessboardCorners.

https://github.com/hybridgroup/gocv/blob/release/cmd/find-circles/main.go

## Find Circles

Find circles in an image using the Hough transform.

https://github.com/hybridgroup/gocv/blob/release/cmd/find-circles/main.go

## Find Lines

Find lines in an image using the Hough transform.

https://github.com/hybridgroup/gocv/blob/release/cmd/find-lines/main.go

## Hand Gestures

Count the number of fingers being held up in front of the camera by looking for convexity defects.

https://github.com/hybridgroup/gocv/blob/release/cmd/hand-gestures/main.go

## Hello

The "hello world" of computer vision.

https://github.com/hybridgroup/gocv/blob/release/cmd/hello/main.go

## Image Similarity

Compute and compare perceptual hashes for a pair of images, with a variety of algorithms.

https://github.com/hybridgroup/gocv/blob/release/cmd/img-similarity/main.go

## MJPEG Streamer

Opens a video capture device, then streams MJPEG from it that you can view in any browser.

https://github.com/hybridgroup/gocv/blob/release/cmd/mjpeg-streamer/main.go

## Motion Detection

Opens a video capture device, then processes it looking for motion, human or otherwise.

https://github.com/hybridgroup/gocv/blob/release/cmd/motion-detect/main.go

## Save Image

Capture a single frame from a connected webcam, then save it to an image file on disk.

https://github.com/hybridgroup/gocv/blob/release/cmd/saveimage/main.go

## Save Video

Capture video from a connected camera, and save 100 frames worth to a video file on disk.

https://github.com/hybridgroup/gocv/blob/release/cmd/savevideo/main.go

## Show Image

Open an image file from disk, then display it in a window.

https://github.com/hybridgroup/gocv/blob/release/cmd/showimage/main.go

## SSD Face Detection

Advanced Deep Neural Network example that uses SSD classifier to detect faces from a connected camera.

https://github.com/hybridgroup/gocv/blob/release/cmd/ssd-facedetect/main.go

## Tensorflow Classifier

Capture video from a connected webcam, then use the Tensorflow machine learning framework to classify whatever is in front of the camera.

https://github.com/hybridgroup/gocv/blob/release/cmd/tf-classifier/main.go

## Tracking

Example of using Tracker from OpenCV Contrib to track any region of interest selected by the user using the TrackerMOSSE algorithm using the connected camera.

https://github.com/hybridgroup/gocv/blob/release/cmd/tracking/main.go

## Version

Displays the current version of OpenCV that is being used by GoCV.

https://github.com/hybridgroup/gocv/blob/release/cmd/version/main.go

## XPhoto

This example demonstrates a couple different uses of the XPhoto module. It can use the GrayworldWB class with BalanceWhite image to save an image file on disk. It can also use the Inpaint functions with inpaint algorithms type to save an image file on disk.

https://github.com/hybridgroup/gocv/blob/release/cmd/xphoto/main.go
