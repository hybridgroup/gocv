0.6.0
---
* **core**
    * Add cv::LUT binding
* **examples** 
    * do not try to go fullscreen, since does not work on OSX
* **features2d** 
    * add AKAZE algorithm
    * add BRISK algorithm
    * add FastFeatureDetector algorithm
    * implement AgastFeatureDetector algorithm
    * implement ORB algorithm
    * implement SimpleBlobDetector algorithm
* **osx**
    * Fix to get the OpenCV path with "brew info".
* **highgui** 
    * use new Window with thread lock, and deprecate WaitKey() in favor of Window.WaitKey()
    * use Window.WaitKey() in tests
* **imgproc** 
    * add tests for HoughCircles
* **pvl**
    * use correct Ptr referencing
* **video** 
    * use smart Ptr for Algorithms thanks to @alalek
    * use unsafe.Pointer for Algorithm    
    * move tests to single file now that they all pass

0.5.0
---
* **core**
    * add TermCriteria for iterative algorithms
* **imgproc**
    * add CornerSubPix() and GoodFeaturesToTrack() for corner detection
* **objdetect**
    * add DetectMultiScaleWithParams() for HOGDescriptor
    * add DetectMultiScaleWithParams() to allow override of defaults for CascadeClassifier
* **video**
    * add CalcOpticalFlowFarneback() for Farneback optical flow calculations
    * add CalcOpticalFlowPyrLK() for Lucas-Kanade optical flow calculations
* **videoio**
    * use temp directory for Windows test compat.
* **build**
    * enable Appveyor build w/cache
* **osx**
    * update env path to always match installed OpenCV from Homebrew

0.4.0
---
* **core**
    * Added cv::mean binding with single argument
    * fix the write-strings warning
    * return temp pointer fix
* **examples**
    * add counter example
    * add motion-detect command
    * correct counter
    * remove redundant cast and other small cleanup
    * set motion detect example to fullscreen
    * use MOG2 for continous motion detection, instead of simplistic first frame only
* **highgui**
    * ability to better control the fullscreen window
* **imgproc**
    * add BorderType param type for GaussianBlur
    * add BoundingRect() function
    * add ContourArea() function
    * add FindContours() function along with associated data types
    * add Laplacian and Scharr functions
    * add Moments() function
    * add Threshold function
* **pvl**
    * add needed lib for linker missing in README
* **test**
    * slightly more permissive version test
* **videoio**
    * Add image compression flags for gocv.IMWrite
    * Fixed possible looping out of compression parameters length
    * Make dedicated function to run cv::imwrite with compression parameters

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