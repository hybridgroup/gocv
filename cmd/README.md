# GoCV Commands

GoCV comes with various useful command line utilities, that are also examples of how to use the package.

## Caffe Classifier

Capture video from a connected webcam, then use the Caffe deep learning framework to classify whatever is in front of the camera.

https://github.com/hybridgroup/gocv/blob/master/cmd/caffe-classifier/main.go

## Captest

Tests to verify you can capture video from a connected webcam.

https://github.com/hybridgroup/gocv/blob/master/cmd/captest/main.go

## Capwindow

Capture video from a connected webcam and display the video in a Window.

https://github.com/hybridgroup/gocv/blob/master/cmd/capwindow/main.go

## Counter

Capture video from a pre-recorded file, and then count the number of detected objects that cross a user-definable vertical or horizontal line.

https://github.com/hybridgroup/gocv/blob/master/cmd/counter/main.go

## Faceblur

Captures video from a connected camera, then uses the CascadeClassifier to detect faces, blurs them using a Gaussian blur, then displays the blurred video in a window.

https://github.com/hybridgroup/gocv/blob/master/cmd/faceblur/main.go

## Facedetect

Captures video from a connected camera, then uses the CascadeClassifier to detect faces, and draw a rectangle around each of them, before displaying them within a Window.

https://github.com/hybridgroup/gocv/blob/master/cmd/facedetect/main.go

## Find-circles

Find circles in an image using the Hough transform.

https://github.com/hybridgroup/gocv/blob/master/cmd/find-circles/main.go

## Hand-gestures

Count the number of fingers being held up in front of the camera by looking for convexity defects.

https://github.com/hybridgroup/gocv/blob/master/cmd/hand-gestures/main.go

## Img-similarity

Compute and compare perceptual hashes for a pair of images, with a variety of algorithms.

https://github.com/hybridgroup/gocv/blob/master/cmd/img-similarity/main.go

## MJPEG-Streamer

Opens a video capture device, then streams MJPEG from it that you can view in any browser.

https://github.com/hybridgroup/gocv/blob/master/cmd/mjpeg-streamer/main.go

## Motion-detect

Opens a video capture device, then processes it looking for motion, human or otherwise.

https://github.com/hybridgroup/gocv/blob/master/cmd/motion-detect/main.go

## Pose

Advanced Deep Neural Network example does pose detection on an image.

https://github.com/hybridgroup/gocv/blob/master/cmd/pose/main.go

## Saveimage

Capture a single frame from a connected webcam, then save it to an image file on disk.

https://github.com/hybridgroup/gocv/blob/master/cmd/saveimage/main.go

## Savevideo

Capture video from a connected camera, and save 100 frames worth to a video file on disk.

https://github.com/hybridgroup/gocv/blob/master/cmd/savevideo/main.go

## Showimage

Open an image file from disk, then display it in a window.

https://github.com/hybridgroup/gocv/blob/master/cmd/showimage/main.go

## SSD Facedetect

Advanced Deep Neural Network example that uses SSD classifier to detect faces from a connected camera.

https://github.com/hybridgroup/gocv/blob/master/cmd/ssd-facedetect/main.go

## TF Classifier

Capture video from a connected webcam, then use the Tensorflow machine learning framework to classify whatever is in front of the camera.

https://github.com/hybridgroup/gocv/blob/master/cmd/tf-classifier/main.go

## Tracking

Example of using Tracker from OpenCV Contrib to track any region of interest selected by the user using the TrackerMOSSE algorithm using the connected camera.

https://github.com/hybridgroup/gocv/blob/master/cmd/tracking/main.go

## Version

Displays the current version of OpenCV that is being used by GoCV.

https://github.com/hybridgroup/gocv/blob/master/cmd/version/main.go
