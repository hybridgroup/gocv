0.40.0
---
* **all**
    - Add support for OpenCV 4.11.0
    - Update license year to 2025
- **bugfix**
    - Correct signature for FindHomography, since dst is actually target format for the operation, not something changed by the operation
- **core**
    - Added new Mat constructor and tests
    - Added NewMatFromPoint2fVector function
- **videoio**
    - Change type of VideoCaptureProperties to int32
    - Added VideoWriterFile with API and API Params
- **docker**
    - Alpine static improved (#1243)
- **make**
    - Improve Linux distro dtection code, and add specific dependendcies for Ubunutu 24.04
- **examples**
    - GStreamer VideoWriter example
- **build**
    - Update homebrew when running macOS tests
    - Update GH actions in linux build to latest versions

0.39.0
---
- **bugfix**
  - Fixed MinGW link typo in README.md.
  - Fixed function name typo (#1236).
- **core**
  - Added `FaceDetectorYN` example.
  - Query I/O API backends registry (#1237).
  - Face detector YN face recognizer SF (#1232).
- **cuda**
  - Added `createLookUpTable` and `split` functions.
  - Added missing CUDA `imgproc` standalone functions.
  - Added `XXXWithStream` standalone functions.
- **docker**
  - Added Dockerfile for container that can perform static builds of your own GoCV project (HighGUI not supported by static builds).
  - Added Dockerfile example showing how to build using static linking.
  - Added static build examples.
  - Updated version for `gocv-static-builder` image.
- **build**
  - Updated workflow for Docker builds to latest versions.
  - Used most recent NVIDIA CUDA base image.
  - Used static FFmpeg build for static OpenCV build.
  - Downgraded static build to Go 1.22 to avoid [Go issue #68976](https://github.com/golang/go/issues/68976).
  - Modified Dockerfile used for static builds to build own versions for static linking.
  - Corrected build options for OpenCV on arm64 for static builds.
  - Separated build tags and linker flags for arm64 and amd64 targets.
  - Modified LDFLAGS for correct static compilation.
  - Removed extra file to clear space for static build/standard dev build images.
  - Added options for Linux arm64 and separate Darwin builds.
- **cgo**
  - Changed tags for static OpenCV to `opencvstatic`.
  - Added options for Linux arm64 and separate Darwin builds.
  - Removed unneeded lib from link for Linux/arm64.
  - Modified LDFLAGS for correct static compilation.
- **docs**
  - Added missing `videoio` functions to ROADMAP.
  - Updated ROADMAP with missing `objdetect` functions for DNN faces, and moved `aruco` under `objdetect` module.
  - Simplified the YOLO example.

0.38.0
---
* **bugfix**
    * aruco: correct test from latest OpenCV update
    * exclude freetype.cpp file from being included in windows build
    * feat(demosaicing): release mat when conversion to bayer encounters invalid pattern
    * imgproc HomographyMethod const values typo fixed
* **build**
    * add macOS build for GH actions CI
    * adjust make and docker build files to build freetype support
    * correct ONNX DNN tests on Linux/macOS
    * move download for GOTURN models into testdata subdirectory
    * remove Caffe tests
    * run DNN tests on Windows
    * make: add task to run all cuda tests
    * make: build non-free modules when building opencv with cuda support
    * skip tests on macOS that are not passing due to OpenCV different results on macOS. See https://forum.opencv.org/t/match-template-different-results-on-mac-m1/10026 and other similar issues.
    * update all generated docker images to Go 1.23.1
* **examples**
    * add asciicam video to ascii in your terminal
    * add object detection example using YOLOv8
* **core**
    * add Closed() function to Mat
    * add OpenCV types for half-float values
    * add TransposeND() function
    * persistance implement Filestorage roadmap (#1208)
    * RotatedRect type constructors
* **dnn**
    * add BlobFromImageWithParams() and BlobFromImagesWithParams() functions
    * add BlobRectToImageRect() and BlobRectsToImageRects() functions
    * allow ReadNet() function to only pass model file, and remove tests for Caffe
* **features2d**
    * SIFT with params (#1186)
* **highgui**
    * added window pollkey function (#1198)
    * added window WaitKeyEx support (#1195)
    * Window set mouse callback (#1220)
* **imgcodecs**
    * added immultiread support
* **imgproc**
    * feat(imgproc): demosaicing wrapper
    * add HomographyMethodRHO HomographyMethod added
* **objdetect**
    * change QRCodeDetector signature to avoid pointer to slice
* **video**
    * added TrackerGOTURN (see roadmap)
* **videoio**
    * Capture from device and file with HW acceleration
* **cuda**
    * add Closed() function to Mat/GpuMat
    * add DeviceSupports function
    * add implementations for AddWeighted and CopyMakeBorder functions
    * add Merge and Transpose functions
    * add support for convertFp16 function
    * add tests for demosaicing
    * feat(imgproc): demosaicing wrapper
    * correct go fmt error
* **contrib/face**
    * added face recognizer interface (#1211)
    * BasicFaceRecognizer + EigenFaceRecognizer + FisherFaceRecognizer (#1213)
    * extra setters and getters for LBPHFaceRecognizer (#1194)
* **contrib/freetype**
    * imported freetype code by lz1998 from PR 873

0.37.0
---
* **all**
    * Add support for OpenCV 4.10.0

0.36.1
---
* **bugfix**
    * Correct error in CUDA function signature
* **test**
    * correct CUDA tests
* **docker**
    * add test image for CUDA 12

0.36.0
---
* **all**
    * Add support for OpenCV 4.9.0
    * update Go to version 1.22
    * update minimum go version to 1.21
* **bugfix**
    * aruco: correct test from latest OpenCV update
* **build**
    * add GH action for Windows
    * remove appveyor
    * adjusted Makefile to build for debian bookworm
* **core**
    * Add additional signature for MinMaxLoc.
    * add color conversion alias
    * add Mahalanobis(), Inv(), Row(), amd Col() functions
    * add MulTransposed() function
    * add PCABackProject() and PCAProject() functions
    * add PSNR() function
    * add SVBackSubst() and SVDecomp() functions
* **calib3d**
    * add FisheyeCalibrate, FisheyeDistortPoints, and CheckChessboard functions
    * Add func comments and update readme
    * add Rodrigues function
    * add SolvePnP function
    * Add more smoke tests
    * Initial commit of more stereo bindings
* **feature2d**
    * Add interface for `Feature2D` algorithms
    * Asserting some algorithms conform to `Feature2D`
    * Prepend "Feature2D" prefix to component interfaces of Feature2D
* **imgproc**
    * add CreateHanningWindow()
    * add EMD()
    * Add float version of BoxPoints and MinAreaRect
    * Add new binding for cv::Erode.
* **videoio**
    * add Retrieve function
* **contrib/xfeatures2d**
    * Add BriefDescriptorExtractor to xfeatures2d (#1114)
    * add NewSURFWithParams func
    * Add separate "Compute" bindings for detection algorithms (#1117)
* **cuda/core**
    * ADD Cuda MultiplyWithStream (#1142)

0.35.0
---
* **all**
    * Add support for OpenCV 4.8.1
    * correct Go formatting
* **features2d**
    * Add Match method for BFMatcher
* **build**
    * remove extra files from GH actions runner so GPU images builds have enough temp file space to run correctly
* **make**
    * for build_raspi added conditional cmake build for 64 and 32bit platforms
    * remove ENABLE_VFPV3=ON and add WITH_TBB=ON from 64bit build.
    * added sudo_pre_install_clean to raspberry pi and jetson installs
    * change sudo_pre_install_clean to support cleanup on 64bit architechtures (arm and x86)

0.34.0
---
* **all**
    * Add support for OpenCV 4.8.0
    * Add support for Go 1.21
* **build**
    * update all builds to use OpenCV 4.8.0
* **core**
    * Adds support for PCACompute
* **docker**
    * add dockerfile for OpenCV static build
* **make**
    * Leave one processor free instead of using all of them when building


0.33.0
---
* **bugfix**
    * Remove opencv2/aruco.hpp include
* **all**
    * build performance tests with all OpenCV builds
* **build**
    * build and push Ubuntu 22.04 base image with OpenCV 4.7.0
    * docker images with opencv
    * docker production images with opencv 4.7.0
    * Docker push to GHCR
* **core**
    * Add ReduceArgMax and ReduceArgMin
* **dnn**
    * improved NMSBoxes code
* **docker**
    * add dockerfile for Ubuntu 22.04 OpenCV base image
    * updates to migrate to GHCR
* **examples**
    * Deallocate Mats in feature-matching example.
    * Fix G108 (CWE-200) and G114 (CWE-676)
    * Fix G304 (CWE-22) and G307 (CWE-703)
    * Fix G304 (CWE-22) and G307 (CWE-703)
    * Missed #nosec tag
* **make**
    * Ubuntu Jammy (22) opencv build support.


0.32.0
---
* **all**
    * update to OpenCV 4.7.0
* **core**
    * Add the number of thread setter and getter
* **calib3d**
    * add EstimateAffinePartial2DWithParams()
* **imgcodecs**
    * Add IMDecodeIntoMat to reduce heap allocations (#1035)
* **imgproc**
    * add matchShapes function support
* **objdetect**
    * move aruco from contrib and also refactor/update to match current OpenCV API
* **photo**
    * add inpaint function
* **video**
    * cv::KalmanFilter bindings.
* **cuda**
    * add support for cuda::TemplateMatching
* **docker**
    * update all dockerfiles for OpenCV 4.7.0/GoCV 0.32.0
    * multiplatform for both amd64 and arm64
    * install libjpeg-turbo into docker image
    * add Ubunutu 18.04 and 20.04 prebuilt OpenCV images
    * add dockerfile for older version of CUDA for those who cannot upgrade
* **ci**
    * remove circleci
    * correct actions that trigger build
* **make**
    * change download path for OpenCV release tag
* **windows**
    * Update win_build_opencv.cmd
* **docs**
    * correct docs on building docker
    * update ROADMAP
    * typo in comment
    * update comments style with gofmt
* **openvino**
    * Add openvino Dockerfile
    * Fix OpenvinoVersion dangling pointer
    * Update env.sh and README.md for 2022.1

0.31.0
---
* **all**
    * update to OpenCV 4.6.0
* **build**
    * Switch to Github Actions for Linux CI build
    * Use go -tags static when verifying static build
* **core**
    * Add Mat.ElemSize (#964)
    * avoid index out of range panic in NewPointsVectorFromPoints
* **video**
    * add findTransformECC function
* **contrib/ximgproc**
    * add PeiLinNormalization() function
    * add anisotropicDiffusion() function
    * implement edgePreservingFilter()
    * implement niBlackThreshold and thinning filters

0.30.0
---
* **all**
    * update to OpenCV 4.5.5
* **build**
    * add install_nonfree make task to build all opencv_contrib modules
    * correct download location for onnx test file
    * Update Makefile for missing version changes
* **core**
    * correct how memory is being allocated for Eye(), Zeros(), and Ones() to address issue #930
* **calib3d** 
    * Adding support for estimateAffine2DWithParams (#924)
* **imgproc**
    * Add DrawContoursWithParams function
* **photo**
    * Add bindings for fastNlMeansDenoising and fastNlMeansDenoisingColored
    * add detailEnhance function
    * add EdgePreservingFilter function
    * add PencilSketch function
    * add stylization function
* **docs**
    * add godoc comments for FastNlMeansDenoising functions
    * update README with info on latest mingw-w64 t use for Windows builds
    * dnn pose detect examples correct the order of the argument variable name
* **examples**
    * Fixed memory leaks in the motion detection example
* **openvino**
    * Update env.sh and README.md
* **windows**
    * use mingw-w64 8.1.0 for protobuf compile
* **contrib**
    * add cv::wechat_qrcode::WeChatQRCode (#949)
    * Update cgo_static.go

0.29.0
---
* **all**
    * update to OpenCV 4.5.4
* **build**
    * add static build ability on windows
    * use tbb for all builds for CPU accelerated operations
* **cuda**
    * implement a bunch of per-element operations
    * add get/set/reset device functions
    * add NewGpuMatWithSize() to preallocate device memory
    * Reshape() returns a new GpuMat with the changed data
    * correct use of Stream by adding WaitForCompletion() and passing pre-allocated GpuMats
* **docs**
    * update ROADMAP from recent contributions
* **videoio**
    * Fix open video capture with api test (#895)
* **calib3d**
    * added EstimateAffine2D
    * findChessboardCornersSB
* **aruco**
    * added many functions as part of initial implementation

0.28.0
---
* **all**
    * update to OpenCV 4.5.3
    * make task and build tag for static build of OpenCV/GoCV on Linux
    * add Makefile tasks for OpenCV install on Nvidia Jetson
    * add gotest for more colorful test output running tests from containers
* **build**
    * correcting output format for code coverage report
    * enforce rule that all Go code is correctly formatted
    * remove codecov
* **core**
    * add NewPointVectorFromMat() and NewPoint2fVectorFromMat() functions
    * Fix possible MatProfile race by ordering remove before free.
* **cuda**
    * add core functions for GpuMat like Cols(), Rows(), and Type()
    * initial implementation for the Flip function
* **docs**
    * update ROADMAP from recent contributions
* **examples**
    * correct list of examples and fix comment
* **features2d**
    * Add NewORBWithParams
* **tracking**
    * change MOSSE to KCF
* **highgui**
    * Add function CreateTrackbarWithValue to Window type.
* **imgcodec**
    * optimize IMEncode avoiding multiple data copies.
* **imgproc**
    * Add CircleWithParams function
    * Add DilateWithParams() function (#827)
    * Add EllipseWithParams function
    * Add FillPolyWithParams function
    * Add PointPolygonTest function
    * Add RectangleWithParams function
* **photo**
    * add MergeMertens, AlignMTB and Denoising function (#848)
* **xphoto**
    * Add Xphoto contrib (#844)

0.27.0
---
* **all**
    * update to OpenCV 4.5.2
* **core**
    * add Append() to PointsVector/PointVector
    * add cv::RNG
    * add implementation for Point2fVector
    * add rand functions
    * add test coverage for PointsVector
    * create new PointsVector/PointVector wrappers to avoid repetitive memory copying for seeming innocent operations involving slices of image.Point
    * test coverage for Point2f
    * use PointVector for everything that we can to speed up pipeline when passing around Point vectors
    * use enum instead of int for Invert Method
* **cuda**
    * adding HoughLinesDetector and HoughSegmentDetector
    * adding tests for the CannyEdgeDetector
    * some refactoring of the API
    * adding dockerfiles for OpenCV 4.5.2 with CUDA 11.2
    * add GaussianFilter
    * correct signature and test for Threshold
    * implement SobelFilter
    * move arithm module functions into correct location
    * rename files to get rid of so many cudas
    * add abs function implementation
* **dnn**
    * increase test coverage
* **docker**
    * make all Dockerfiles names/tags more consistent
* **docs**
    * add CUDA functions that need implementation to ROADMAP
    * remove invalid sections and add some missing functions from ROADMAP
* **imgproc**
    * Add FindContoursWithParams function
    * Add ToImageYUV and ToImageYUVWithParams
* **make**
    * add make task to show changelog for next release
* **wechat_qrcode**
    * disable module in Windows due to linker error

0.26.0
---
* **all**
    * update to OpenCV 4.5.1
* **core**
    * add Matrix initializers: eye, ones, zeros (#758)
    * add multidimensional mat creation
    * add ndim mat constructor
    * added accumulators
    * added norm call with two mats (#600)
    * keep a reference to a []byte that backs a Mat. (#755)
    * remove guard for DataPtrUint8 since any Mat can be treated an Uint8
    * add Mat IsContinuous() function, and ensure that any Mat data pointers used to create Go slices only apply to continuous Mats
    * fix buffer size for Go strings for 32-bit operating systems
* **build**
    * bring back codecov.io
* **calib3d**
    * correctly close mat after test
* **dnn**
    * add ReadNetFromONNX and ReadNetFromONNXBytes (#760)
    * increase test coverage
* **docker**
    * dockerfiles for opencv gpu builds
* **docs**
    * corrected links to CUDA and OpenVINO
    * list all unimplemented functions in photo module
    * replace GoDocs with pkg docs
    * update ROADMAP from recent contributions
* **imgproc**
    * add test coverage for GetTextSizeWithBaseline()
    * close all Mats even those based on memory slices
    * close Mat to avoid memory leak in ToImage()
    * refactoring of ToImage and ImageToMatXX functions
* **openvino**
    * fix dldt repo in makefile for openvino
* **os**
    * adding gcc-c++ package to rpm deps
* **photo**
    * add SeamlessClone function
* **profile**
    * add created mats in Split and ForwardLayers to profile (#780)

0.25.0
---
* **all**
    * update to opencv release 4.5.0
* **build** 
    * add file dependencies needed for DNN tests
    * add verbose output for tests on CircleCI
    * also run unit tests on non-free algorithms. YMMV.
    * fix build with cuda
    * remove Travis and switch to CircleCI using Docker based builds
    * update CI builds to Go 1.15
* **core**
    * add mixChannels() method to Mat (#746)
    * Add toGoStrings helper
    * support ConvertToWithParams method
* **dnn**
    * Add NMSBoxes function (#736)
    * Added ability to load Torch file. Tested features for extracting 128d vectors
    * fix using wrong type for unconnectedlayertype
    * use default ddepth for conversions to blob from image as recommended by @berak
* **docker** 
    * use separate dockerfile for opencv to avoid massive rebuild
* **docs**
    * add recent contributions to ROADMAP and also add cuda functions still in need of implementation
    * display CircleCI badge in README
    * minor improvements to CUDA docs in READMEs
* **features2d**
    * add FlannBasedMatcher
    * add drawmatches (#720)
    * fix memory leak in SIFT
* **highgui**
    * refactored ROI methods
* **imgproc**
    * Add option to return baseline with GetTextSizeWithBaseline
* **objdetect** 
    * Add QRCode DetectAndDecodeMulti
* **videoio**
    * Add video capture properties and set preferred api backend (#739)
    * fix needed as discussed in golang/go issue #32479

0.24.0
---
* **all**
    * update Makefile and READMEChange constants and corresponding function signatures to have the correct types (#689)
    * replace master branch terminology with release
    * update to OpenCV 4.4.0
* **calib3d**
    * add FindHomography()
    * add function EstimateAffinePartial2D()
    * add GetAffineTransform() and GetAffineTransform2f()
    * add UndistortPoints(), FisheyeUndistortPoints() and EstimateNewCameraMatrixForUndistortRectify()
* **core**
    * add MultiplyWithParams
* **docs**
    * add recent contributions to ROADMAP
    * create CODE_OF_CONDUCT.md
    * update copyright year
* **features2d**
    * close returned Mat from SIFT algorithm
    * fix issue 707 with DrawKeyPoints
    * SIFT patent now expired so is part of main OpenCV modules
* **imgproc**
    * change struct to remove GNU old-style field designator extension warning

0.23.0
---
* **build**
    * update Makefile and README
    * update to use go1.14
* **calib3d**
    * add draw chessboard
* **core**
    * fix memory leak in Mat.Size() and Mat.Split() (#580)
* **cuda**
    * add build support
    * add cuda backend/target
    * add support for:
        * cv::cuda::CannyEdgeDetector
        * cv::cuda::CascadeClassifier Class
        * cv::cuda::HOG Class
    * remove breaking case statement
* **dnn**
    * avoid parallel test runs
    * remove attempt at providing grayscale image blog conversion that uses mean adjustment
* **docker**
    * docker file last command change (#505)
* **docs**
    * add recent contributions to ROADMAP
* **imgproc**
    * add ErodeWithParams function
    * add getGaussianKernel function
    * add Go Point2f type and update GetPerspectiveTransform() (#589)
    * add PhaseCorrelate binding (#626)
    * added Polylines feature
    * do not free contours data until after we have drawn the needed contours
    * Threshold() should return a value (#620)
* **make**
    * added raspberry pi zero support to the makefile
* **opencv**
    * update to OpenCV 4.3.0
* **openvino**
    * add build support
* **windows**
    * add cmake flag for allocator stats counter type to avoid opencv issue #16398

0.22.0
---
* **bgsegm**
    * Add BackgroundSubtractorCNT
* **calib3d**
    * Added undistort function (#520)
* **core**
    * add functions (singular value decomposition, multiply between matrices, transpose matrix) (#559)
    * Add new funcs (#578)
    * add setIdentity() method to Mat
    * add String method (#552)
    * MatType: add missing constants
* **dnn**
    * Adding GetLayerNames()
    * respect the bit depth of the input image to set the expected output when converting an image to a blob
* **doc**
    * change opencv version 3.x to 4.x
* **docker**
    * use Go1.13.5 for image
* **imgcodecs**
    * Fix webp image decode error (#523)
imgcodecs: optimize copy of data used for IMDecode method
* **imgproc**
    * Add GetRectSubPix
    * Added ClipLine
    * Added InvertAffineTransform
    * Added LinearPolar function (#524)
    * correct ksize param used for MedianBlur unit test
    * Feature/put text with line type (#527)
    * FitEllipse
    * In FillPoly and DrawContours functions, remove func() wrap to avoid memory freed before calling opencv functions. (#543)
* **objdetect**
    * Add support QR codes
* **opencv**
    * update to OpenCV 4.2.0 release
* **openvino**
    * Add openvino async
* **test**
    * Tolerate imprecise result in SolvePoly
    * Tolerate imprecision in TestHoughLines

0.21.0
---
* **build**
    * added go clean --cache to clean target, see issue 458
* **core**
    * Add KMeans function
    * added MeanWithMask function for Mats (#487)
    * Fix possible resource leak
* **cuda**
    * added cudaoptflow
    * added NewGpuMatFromMat which creates a GpuMat from a Mat
    * Support for CUDA Image Warping (#494)
* **dnn**
    * add BlobFromImages (#467)
    * add ImagesFromBlob (#468)
* **docs**
    * update ROADMAP with all recent contributions. Thank you!
* **examples**
    * face detection from image url by using IMDecode (#499)
    * better format
* **imgproc**
    * Add calcBackProject
    * Add CompareHist
    * Add DistanceTransform and Watershed
    * Add GrabCut
    * Add Integral
    * Add MorphologyExWithParams
* **opencv**
    * update to version 4.1.2
* **openvino**
    * updates needed for 2019 R3
* **videoio**
    * Added ToCodec to convert FOURCC string to numeric representation (#485)

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