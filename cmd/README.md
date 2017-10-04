# GoCV Commands

GoCV comes with various useful command line utilities, that are also examples of how to use the package.

## Captest

Tests to verify you can capture video from a connected webcam.

Example: go run main.go captest camera-id=0

## Capwindow

Capture video from a connected webcam and display the video in a Window.

Example: go run main.go capwindow camera-id=0

## Faceblur

Captures video from a connected camera, then uses the CascadeClassifier to detect faces, blurs them using a Gaussian blur, then displays the blurred video in a window.

Example: go run main.go faceblur camera-id=0 clf-file="demo.xml"

## Facedetect

Captures video from a connected camera, then uses the CascadeClassifier to detect faces, and draw a rectangle around each of them, before displaying them within a Window

Example: go run main.go facedetect camera-id=0 clf-file="demo.xml"

## MJPEG-Streamer

Opens a video capture device, then streams MJPEG from it that you can view in any browser.

Example: go run main.go mstreamer camera-id=0 host="127.0.0.1:8080"

## Saveimage

Capture a single frame from a connected webcam, then save it to an image file on disk.

Example: go run main.go saveimage camera-id=0 image-file="image.jpg"

## Savevideo

Capture video from a connected camera, and save 100 frames worth to a video file on disk.

Example: go run main.go savevideo camera-id=0 video-file="video.mjpeg"

## Showimage

Open an image file from disk, then display it in a window.

Example: go run main.go showimage image-file="image.jpg"

## Version

Displays the current version of OpenCV that is being used by GoCV.

Example: go run main.go version
