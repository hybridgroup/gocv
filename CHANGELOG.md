0.20.0
---
* **build**
    * Use Go 1.12.x for build
    * Update to OpenCV 4.1.0
* **cuda**
    * Initial cuda implementation
* **docs**
    * Fix the command to install xquartz via brew/cask
* **features2d**
    * Add support for SimpleBlobDetectorParams (#434)
    * Added FastFeatureDetectorWithParams
* **imgproc**
    * Added function call to cv::morphologyDefaultBorderValue
* **test**
    * Increase test coverage for FP16BlobFromImage()
* **video**
    * Added calcOpticalFlowPyrLKWithParams
    * Addition of MOG2/KNN constructor with options

0.19.0
---
* **build**
    * Adds Dockerfile. Updates Makefile and README.
    * make maintainer tag same as dockerhub organization name
    * make sure to run tests for non-free contrib algorithms
    * update Appveyor build to use Go 1.12
* **calib3d**
    * add func InitUndistortRectifyMap (#405)
* **cmd**
    * correct formatting of code in example
* **core**
    * Added Bitwise Operations With Masks
    * update to OpenCV4.0.1
* **dnn**
    * add new backend and target types for NVIDIA and FPGA
    * Added blobFromImages in ROADMAP.md (#403)
    * Implement dnn methods for loading in-memory models.
* **docker**
    * update Dockerfile to use OpenCV 4.0.1
* **docs**
    * update ROADMAP from recent contributions
* **examples**
    * Fixing filename in caffe-classifier example
* **imgproc**
    * Add 'MinEnclosingCircle' function
    * added BoxPoints function and BorderIsolated const
    * Added Connected Components
    * Added the HoughLinesPointSet function.
    * Implement CLAHE to imgproc
* **openvino**
    * remove lib no longer included during non-FPGA installations
* **test**
    * Add len(kp) == 232 to TestMSER, seems this is necessary for MacOS for some reason.

0.18.0
---
* **build**
    * add OPENCV_GENERATE_PKGCONFIG flag to generate pkg-config file
    * Add required curl package to the RPM and DEBS
    * correct name for zip directory used for code download
    * Removing linking against face contrib module
    * update CI to use 4.0.0 release
    * update Makefile and Windows build command file to OpenCV 4.0.0
    * use opencv4 file for pkg-config
* **core**
    * add ScaleAdd() method to Mat
* **docs**
    * replace OpenCV 3.4.3 references with OpenCV 4
    * update macOS installation info to refer to new OpenCV 4.0 brew
    * Updated function documentation with information about errors.
* **examples**
    * Improve accuracy in hand gesture sample
* **features2d**
    * update drawKeypoints() to use new stricter enum
* **openvino**
    * changes to accommodate release 2018R4
* **profile**
    * add build tag matprofile to allow for conditional inclusion of custom profile
    * Add Mat profile wrapper in other areas of the library.
    * Add MatProfile.
    * Add MatProfileTest.
    * move MatProfile tests into separate test file so they only run when custom profiler active
* **test**
    * Close images in tests.
    * More Closes in tests.
    * test that we are using 4.0.x version now
* **videoio**
    * Return the right type and error when opening VideoCapture fails

0.17.0
---
* **build** 
    * Update Makefile
    * update version of OpenCV used to 3.4.3
    * use link to OpenCV 3.4.3 for Windows builds
* **core** 
    * add mulSpectrums wrapper
    * add PolarToCart() method to Mat
    * add Reduce() method to Mat
    * add Repeat() method to Mat
    * add Solve() method to Mat
    * add SolveCubic() method to Mat
    * add SolvePoly() method to Mat
    * add Sort() method to Mat
    * add SortIdx() method to Mat
    * add Trace() method to Mat
    * Added new MatType
    * Added Phase function
* **dnn** 
    * update test to match OpenCV 3.4.3 behavior
* **docs**
    * Add example of how to run individual test
    * adding instructions for installing pkgconfig for macOS
    * fixed GOPATH bug.
    * update ROADMAP from recent contributions
* **examples**
    * add condition to handle no circle found in circle detection example
* **imgcodecs**
    * Added IMEncodeWithParams function
* **imgproc**
    * Added Filter2D function
    * Added fitLine function
    * Added logPolar function
    * Added Remap function
    * Added SepFilter2D function
    * Added Sobel function
    * Added SpatialGradient function
* **xfeatures2d**
    * do not run SIFT test unless OpenCV was built using OPENCV_ENABLE_NONFREE
    * do not run SURF test unless OpenCV was built using OPENCV_ENABLE_NONFREE

0.16.0
---
* **build**
    * add make task for Raspbian install with ARM hardware optimizations
    * use all available cores to compile OpenCV on Windows as discussed in issue #275
    * download performance improvements for OpenCV installs on Windows
    * correct various errors and issues with OpenCV installs on Fedora and CentOS
* **core**
    * correct spelling error in constant to fix issue #269
    * implemented & added test for Mat.SetTo
    * improve Multiply() GoDoc and test showing Scalar() multiplication
    * mutator functions for Mat add, subtract, multiply, and divide for uint8 and float32 values.
* **dnn**
    * add FP16BlobFromImage() function to convert an image Mat to a half-float aka FP16 slice of bytes
* **docs**
    * fix a varible error in example code in README

0.15.0
---
* **build**
    * add max to make -j
    * improve path for Windows to use currently configured GOPATH
* **core**
    * Add Mat.DataPtr methods for direct access to OpenCV data
    * Avoid extra copy in Mat.ToBytes + code review feedback
* **dnn**
    * add test coverage for ParseNetBackend and ParseNetTarget
    * complete test coverage
* **docs**
    * minor cleanup of language for install
    * use chdir instead of cd in Windows instructions
* **examples**
    * add 'hello, video' example to repo
    * add HoughLinesP example
    * correct message on device close to match actual event
    * small change in display message for when file is input source
    * use DrawContours in motion detect example
* **imgproc**
    * Add MinAreaRect() function
* **test**
    * filling test coverage gaps
* **videoio**
    * add test coverage for OpenVideoCapture

0.14.0
---
* **build**
    * Add -lopencv_calib3d341 to the linker
    * auto-confirm on package installs from make deps command
    * display PowerShell download status for OpenCV files
    * obtain caffe test config file from new location in Travis build
    * remove VS only dependencies from OpenCV build, copy caffe test config file from new location
    * return back to GoCV directory after OpenCV install
    * update for release of OpenCV v3.4.2
    * use PowerShell for scripted OpenCV install for Windows
    * win32 version number has not changed yet
* **calib3d**
    * Add Calibrate for Fisheye model(WIP)
* **core**
    * add GetTickCount function
    * add GetTickFrequency function
    * add Size() and FromPtr() methods to Mat
    * add Total method to Mat
    * Added RotateFlag type
    * correct CopyTo to use pointer to Mat as destination
    * functions converting Image to Mat
    * rename implementation to avoid conflicts with Windows
    * stricter use of reflect.SliceHeader
* **dnn**
    * add backend/device options to caffe and tensorflow DNN examples
    * add Close to Layer
    * add first version of dnn-pose-detection example
    * add further comments to object detection/tracking DNN example
    * add GetPerfProfile function to Net
    * add initial Layer implementation alongside enhancements to Net
    * add InputNameToIndex to Layer
    * add new functions allowing DNN backends such as OpenVINO
    * additional refactoring and comments in dnn-pose-detection example
    * cleanup DNN face detection example
    * correct const for device targets to be called Target
    * correct test that expected init slice with blank entries
    * do not init slice with blank entries, since added via append
    * further cleanup of DNN face detection example
    * make dnn-pose-detection example use Go channels for async operation
    * refactoring and additional comments for object detection/tracking DNN example
    * refine comment in header for style transfer example
    * working style transfer example
    * added ForwardLayers() to accomodate models with multiple output layers
* **docs**
    * add scripted Windows install info to README
    * Added a sample gocv workflow contributing guideline
    * mention docker image in README.
    * mention work in progress on Android
    * simplify and add missing step in Linux installation in README
    * update contributing instructions to match latest version
    * update ROADMAP from recent calib3d module contribution
    * update ROADMAP from recent imgproc histogram contribution
* **examples**
    * cleanup header for caffe dnn classifier
    * show how to use either Caffe or Tensorflow for DNN object detection
    * further improve dnn samples
    * rearrange and add comments to dnn style transfer example
    * remove old copy of pose detector
    * remove unused example
* **features2d**
    * free memory allocation bug for C.KeyPoints as pointed out by @tzununbekov
    * Adding opencv::drawKeypoints() support
* **imgproc**
    * add equalizeHist function
    * Added opencv::calcHist implementation
* **openvino**
    * add needed environment config to execute examples
    * further details in README explaining how to use
    * remove opencv contrib references as they are not included in OpenVINO
* **videoio**
    * Add OpenVideoCapture
    * Use gocv.VideoCaptureFile if string is specified for device.

0.13.0
---
* **build**
    * Add cgo directives to contrib
    * contrib subpackage also needs cpp 11 or greater for a warning free build on Linux
    * Deprecate env scripts and update README
    * Don't set --std=c++1z on non-macOS
    * Remove CGO vars from CI and correct Windows cgo directives
    * Support pkg-config via cgo directives
    * we actually do need cpp 11 or greater for a warning free build on Linux
* **docs**
    * add a Github issue template to project
    * provide specific examples of using custom environment
* **imgproc**
    * add HoughLinesPWithParams() function
* **openvino**
    * add build tag specific to openvino
    * add roadmap info
    * add smoke test for ie

0.12.0
---
* **build**
    * convert to CRLF
    * Enable verbosity for travisCI
    * Further improvements to Makefile
* **core**
    * Add Rotate, VConcat
    * Adding InScalarRange and NewMatFromScalarWithSize functions
    * Changed NewMatFromScalarWithSize to NewMatWithSizeFromScalar
    * implement CheckRange(), Determinant(), EigenNonSymmetric(), Min(), and MinMaxIdx() functions
    * implement PerspectiveTransform() and Sqrt() functions
    * implement Transform() and Transpose() functions
    * Make toByteArray safe for empty byte slices
    * Renamed InScalarRange to InRangeWithScalar
* **docs**
    * nicer error if we can't read haarcascade_frontalface_default
    * correct some ROADMAP links
    * Fix example command.
    * Fix executable name in help text.
    * update ROADMAP from recent contributions
* **imgproc** 
    * add BoxFilter and SqBoxFilter functions
    * Fix the hack to convert C arrays to Go slices.
* **videoio** 
    * Add isColor to VideoWriterFile
    * Check numerical parameters for gocv.VideoWriterFile
    * CodecString()
* **features2d** 
    * add BFMatcher
* **img_hash** 
    * Add contrib/img_hash module
    * add GoDocs for new img_hash module
    * Add img-similarity as an example for img_hash
* **openvino** 
    * adds support for Intel OpenVINO toolkit PVL
    * starting experimental work on OpenVINO IE
    * update README files for Intel OpenVINO toolkit support
    * WIP on IE can load an IR network

0.11.0
---
* **build**
    * Add astyle config
    * Astyle cpp/h files
    * remove duplication in Makefile for astyle
* **core**
    * Add GetVecfAt() function to Mat
    * Add GetVeciAt() function to Mat
    * Add Mat.ToImage()
    * add MeanStdDev() method to Mat
    * add more functions
    * Compare Mat Type directly
    * further cleanup for GoDocs and enforce type for convariance operations
    * Make borderType in CopyMakeBorder be type BorderType
    * Mat Type() should return MatType
    * remove unused convenience functions
    * use Mat* to indicate when a Mat is mutable aka an output parameter
* **dnn**
    * add a ssd sample and a GetBlobChannel helper
    * added another helper func and a pose detection demo
* **docs**
    * add some additional detail about adding OpenCV functions to GoCV
    * updates to contribution guidelines
    * fill out complete list of needed imgproc functions for sections that have work started
    * indicate that missing imgproc functions need implementation
    * mention the WithParams patterns to be used for functions with default params
    * update README for the Mat* based API changes
    * update ROADMAP for recent changes especially awesome recent core contributions from @berak
* **examples**
    * Fix tf-classifier example
    * move new DNN advanced examples into separate folders
    * Update doc for the face contrib package
    * Update links in caffe-classifier demo
    * WIP on hand gestures tracking example
* **highgui**
    * fix constant in NewWindow
* **imgproc**
    * Add Ellipse() and FillPoly() functions
    * Add HoughCirclesWithParams() func
    * correct output Mat to for ConvexHull()
    * rename param being used for Mat image to be modified
* **tracking**
    * add support for TrackerMIL, TrackerBoosting, TrackerMedianFlow, TrackerTLD, TrackerKCF, TrackerMOSSE, TrackerCSRT trackers
    * removed mutitracker, added Csrt, rebased
    * update GoDocs and minor renaming based on gometalint output

0.10.0
---
* **build** 
    * install unzip before build
    * overwrite when unzipping file to install Tensorflow test model
    * use -DCPU_DISPATCH= flag for build to avoid problem with disabled AVX on Windows
    * update unzipped file when installing Tensorflow test model
* **core** 
    * add Compare() and CountNonZero() functions
    * add getter/setter using optional params for multi-dimensional Mat using row/col/channel
    * Add mat subtract function
    * add new toRectangle function to DRY up conversion from CRects to []image.Rectangle
    * add split subtract sum wrappers
    * Add toCPoints() helper function
    * Added Mat.CopyToWithMask() per #47
    * added Pow() method
    * BatchDistance BorderInterpolate CalcCovarMatrix CartToPolar
    * CompleteSymm ConvertScaleAbs CopyMakeBorder Dct
    * divide, multiply
    * Eigen Exp ExtractChannels
    * operations on a 3d Mat are not same as a 2d multichannel Mat
    * resolve merge conflict with duplicate Subtract() function
    * run gofmt on core tests
    * Updated type for Mat.GetUCharAt() and Mat.SetUCharAt() to reflect uint8 instead of int8
* **docs** 
    * update ROADMAP of completed functions in core from recent contributions
* **env** 
    * check loading resources
    * Add distribution detection to deps rule
    * Add needed environment variables for Linux
* **highgui** 
    * add some missing test coverage on WaitKey()
* **imgproc** 
    * Add adaptive threshold function
    * Add pyrDown and pyrUp functions
    * Expose DrawContours()
    * Expose WarpPerspective and GetPerspectiveTransform
    * implement ConvexHull() and ConvexityDefects() functions
* **opencv** 
    * update to OpenCV version 3.4.1

0.9.0
---
* **bugfix** 
    * correct several errors in size parameter ordering
* **build**
    * add missing opencv_face lib reference to env.sh
    * Support for non-brew installs of opencv on Darwin
* **core**
    * add Channels() method to Mat
    * add ConvertTo() and NewMatFromBytes() functions
    * add Type() method to Mat
    * implement ConvertFp16() function
* **dnn** 
    * use correct size for blob used for Caffe/Tensorflow tests
* **docs** 
    * Update copyright date and Apache 2.0 license to include full text
* **examples** 
    * cleanup mjpeg streamer code
    * cleanup motion detector comments
    * correct use of defer in loop
    * use correct size for blob used for Caffe/Tensorflow examples
* **imgproc**
    * Add cv::approxPolyDP() bindings.
    * Add cv::arcLength() bindings.
    * Add cv::matchTemplate() bindings.
    * correct comment and link for Blur function
    * correct docs for BilateralFilter()

0.8.0
---
* **core**
    * add ColorMapFunctions and their test
    * add Mat ToBytes
    * add Reshape and MinMaxLoc functions
    * also delete points
    * fix mistake in the norm function by taking NormType instead of int as parameter
    * SetDoubleAt func and his test
    * SetFloatAt func and his test
    * SetIntAt func and his test
    * SetSCharAt func and his test
    * SetShortAt func and his test
    * SetUCharAt fun and his test
    * use correct delete operator for array of new, eliminates a bunch of memory leaks
* **dnn**
    * add support for loading Tensorflow models
    * adjust test for Caffe now that we are auto-cropping blob
    * first pass at adding Caffe support
    * go back to older function signature to avoid version conflicts with Intel CV SDK
    * properly close DNN Net class
    * use approx. value from test result to account forr windows precision differences
* **features2d**
    * implement GFTTDetector, KAZE, and MSER algorithms
    * modify MSER test for Windows results
* **highgui**
    * un-deprecate WaitKey function needed for CLI apps
* **imgcodec**
    * add fileExt type
* **imgproc**
    * add the norm wrapper and use it in test for WarpAffine and WarpAffineWithParams
    * GetRotationMatrix2D, WarpAffine and WarpAffineWithParams
    * use NormL2 in wrap affine
* **pvl**
    * add support for FaceRecognizer
    * complete wrappers for all missing FaceDetector functions
    * update instructions to match R3 of Intel CV SDK
* **docs**
    * add more detail about exactly which functions are not yet implememented in the modules that are marked as 'Work Started'
    * add refernece to Tensorflow example, and also suggest brew upgrade for MacOS
    * improve ROADMAP to help would-be contributors know where to get started
    * in the readme, explain compiling to a static library
    * remove many godoc warnings by improving function descriptions
    * update all OpenCV 3.3.1 references to v3.4.0
    * update CGO_LDFLAGS references to match latest requirements
    * update contribution guidelines to try to make it more inviting
* **examples**
    * add Caffe classifier example
    * add Tensorflow classifier example
    * fixed closing window in examples in infinite loop
    * fixed format of the examples with gofmt
* **test**
    * add helper function for test : floatEquals
    * add some attiribution from test function
    * display OpenCV version in case that test fails
    * add round function to allow for floating point accuracy differences due to GPU usage.
* **build**
    * improve search for already installed OpenCV on MacOS
    * update Appveyor build to Opencv 3.4.0
    * update to Opencv 3.4.0

0.7.0
---
* **core**
    * correct Merge implementation
* **docs**
    * change wording and formatting for roadmap
    * update roadmap for a more complete list of OpenCV functionality
    * sequence docs in README in same way as the web site, aka by OS
    * show in README that some work was done on contrib face module
* **face**
    * LBPH facerecognizer bindings
* **highgui**
    * complete implementation for remaining API functions
* **imgcodecs**
    * add IMDecode function
* **imgproc**
    * elaborate on HoughLines & HoughLinesP tests to fetch a few individual results
* **objdetect**
    * add GroupRectangles function
* **xfeatures2d**
    * add SIFT and SURF algorithms from OpenCV contrib
    * improve description for OpenCV contrib
    * run tests from OpenCV contrib

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