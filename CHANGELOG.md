0.3.1
---
* **overall**
    * Update to use OpenCV 3.3.1

0.3.0
---
* **docs** 
    * Correct Windows build location from same @jpfarias fix to gocv-site
* **core**
    * Add Resize
    * Add Mat merge and Discrete Fourier Transform
    * Add CopyTo() and Normalize()
    * Implement various important Mat logical operations
* **video**
    * BackgroundSubtractorMOG2 algorithm now working
    * Add BackgroundSubtractorKNN algorithm from video module
* **videoio**
    * Add VideoCapture::get
* **imgproc**
    * Add BilateralFilter and MedianBlur
    * Additional drawing functions implemented
    * Add HoughCircles filter
    * Implement various morphological operations
* **highgui**
    * Add Trackbar support
* **objdetect**
    * Add HOGDescriptor
* **build** 
    * Remove race from test on Travis, since it causes CGo segfault in MOG2

0.2.0
---
* Switchover to custom domain for package import
* Yes, we have Windows

0.1.0
---
Initial release!

- [X] Video capture
- [X] GUI Window to display video
- [X] Image load/save
- [X] CascadeClassifier for object detection/face tracking/etc.
- [X] Installation instructions for Ubuntu
- [X] Installation instructions for OS X
- [X] Code example to use VideoWriter
- [X] Intel CV SDK PVL FaceTracker support
- [X] imgproc Image processing
- [X] Travis CI build
- [X] At least minimal test coverage for each OpenCV class
- [X] Implement more of imgproc Image processing