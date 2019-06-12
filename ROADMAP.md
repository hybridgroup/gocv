# Roadmap

This is a list of all of the functionality areas within OpenCV, and OpenCV Contrib.

Any section listed with an "X" means that all of the relevant OpenCV functionality has been wrapped for use within GoCV.

Any section listed with **WORK STARTED** indicates that some work has been done, but not all functionality in that module has been completed. If there are any functions listed under a section marked **WORK STARTED**, it indicates that that function still requires a wrapper implemented.

And any section that is simply listed, indicates that so far, no work has been done on that module.

Your pull requests will be greatly appreciated!

## Modules list

- [ ] **core. Core functionality - WORK STARTED**
    - [ ] **Basic structures - WORK STARTED**
    - [ ] **Operations on arrays - WORK STARTED**. The following functions still need implementation:
        - [ ] [Mahalanobis](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga4493aee129179459cbfc6064f051aa7d)
        - [ ] [mixChannels](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga51d768c270a1cdd3497255017c4504be)
        - [ ] [mulTransposed](https://docs.opencv.org/master/d2/de8/group__core__array.html#gadc4e49f8f7a155044e3be1b9e3b270ab)
        - [ ] [PCABackProject](https://docs.opencv.org/master/d2/de8/group__core__array.html#gab26049f30ee8e94f7d69d82c124faafc)
        - [ ] [PCACompute](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga4e2073c7311f292a0648f04c37b73781)
        - [ ] [PCAProject](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga6b9fbc7b3a99ebfd441bbec0a6bc4f88)
        - [ ] [PSNR](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga07aaf34ae31d226b1b847d8bcff3698f)
        - [ ] [randn](https://docs.opencv.org/master/d2/de8/group__core__array.html#gaeff1f61e972d133a04ce3a5f81cf6808)
        - [ ] [randShuffle](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga6a789c8a5cb56c6dd62506179808f763)
        - [ ] [randu](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga1ba1026dca0807b27057ba6a49d258c0)
        - [ ] [scaleAdd](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga9e0845db4135f55dcf20227402f00d98)
        - [ ] [setIdentity](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga388d7575224a4a277ceb98ccaa327c99)
        - [ ] [setRNGSeed](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga757e657c037410d9e19e819569e7de0f)
        - [ ] [SVBackSubst](https://docs.opencv.org/master/d2/de8/group__core__array.html#gab4e620e6fc6c8a27bb2be3d50a840c0b)
        - [ ] [SVDecomp](https://docs.opencv.org/master/d2/de8/group__core__array.html#gab477b5b7b39b370bb03e75b19d2d5109)
        - [ ] [theRNG](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga75843061d150ad6564b5447e38e57722)
    - [ ] XML/YAML Persistence
    - [ ] Clustering
    - [ ] Utility and system functions and macros
    - [ ] OpenGL interoperability
    - [ ] Intel IPP Asynchronous C/C++ Converters
    - [ ] Optimization Algorithms
    - [ ] OpenCL support 

- [ ] **imgproc. Image processing - WORK STARTED**
    - [ ] **Image Filtering - WORK STARTED** The following functions still need implementation:
        - [ ] [buildPyramid](https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#gacfdda2bc1ac55e96de7e9f0bce7238c0)
        - [ ] [getDerivKernels](https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga6d6c23f7bd3f5836c31cfae994fc4aea)
        - [ ] [getGaborKernel](https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#gae84c92d248183bd92fa713ce51cc3599)
        - [ ] [getGaussianKernel](https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#gac05a120c1ae92a6060dd0db190a61afa)
        - [X] [morphologyDefaultBorderValue](https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga94756fad83d9d24d29c9bf478558c40a)
        - [ ] [morphologyExWithParams](https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga67493776e3ad1a3df63883829375201f)
        - [ ] [pyrMeanShiftFiltering](https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga9fabdce9543bd602445f5db3827e4cc0)
    
    - [ ] **Geometric Image Transformations - WORK STARTED** The following functions still need implementation:
        - [ ] [convertMaps](https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga9156732fa8f01be9ebd1a194f2728b7f)
        - [ ] [getAffineTransform](https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga8f6d378f9f8eebb5cb55cd3ae295a999)
        - [ ] [getDefaultNewCameraMatrix](https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga744529385e88ef7bc841cbe04b35bfbf)
        - [ ] [getRectSubPix](https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga77576d06075c1a4b6ba1a608850cd614)
        - [ ] [initUndistortRectifyMap](https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga7dfb72c9cf9780a347fbe3d1c47e5d5a)
        - [ ] [initWideAngleProjMap](https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#gaceb049ec48898d1dadd5b50c604429c8)
        - [ ] [invertAffineTransform](https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga57d3505a878a7e1a636645727ca08f51)
        - [ ] [linearPolar](https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#gaa38a6884ac8b6e0b9bed47939b5362f3)
        - [ ] [undistort](https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga69f2545a8b62a6b0fc2ee060dc30559d)
        - [ ] [undistortPoints](https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga55c716492470bfe86b0ee9bf3a1f0f7e)

    - [ ] **Miscellaneous Image Transformations - WORK STARTED** The following functions still need implementation:
        - [ ] [cvtColorTwoPlane](https://docs.opencv.org/master/d7/d1b/group__imgproc__misc.html#ga8e873314e72a1a6c0252375538fbf753)
        - [ ] [distanceTransform](https://docs.opencv.org/master/d7/d1b/group__imgproc__misc.html#ga8a0b7fdfcb7a13dde018988ba3a43042)
        - [ ] [floodFill](https://docs.opencv.org/master/d7/d1b/group__imgproc__misc.html#gaf1f55a048f8a45bc3383586e80b1f0d0)
        - [ ] [grabCut](https://docs.opencv.org/master/d7/d1b/group__imgproc__misc.html#ga909c1dda50efcbeaa3ce126be862b37f)
        - [ ] [integral](https://docs.opencv.org/master/d7/d1b/group__imgproc__misc.html#ga97b87bec26908237e8ba0f6e96d23e28)
        - [ ] [watershed](https://docs.opencv.org/master/d7/d1b/group__imgproc__misc.html#ga3267243e4d3f95165d55a618c65ac6e1)

    - [ ] **Drawing Functions - WORK STARTED** The following functions still need implementation:
        - [ ] [clipLine](https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#gaf483cb46ad6b049bc35ec67052ef1c2c)
        - [ ] [drawMarker](https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga482fa7b0f578fcdd8a174904592a6250)
        - [ ] [ellipse2Poly](https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga727a72a3f6a625a2ae035f957c61051f)
        - [ ] [fillConvexPoly](https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga906aae1606ea4ed2f27bec1537f6c5c2)
        - [ ] [getFontScaleFromHeight](https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga442ff925c1a957794a1309e0ed3ba2c3)
        - [ ] [polylines](https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga444cb8a2666320f47f09d5af08d91ffb)
    
    - [ ] ColorMaps in OpenCV
    - [ ] Planar Subdivision
    - [ ] **Histograms  - WORK STARTED** The following functions still need implementation:
        - [ ] [calcBackProject](https://docs.opencv.org/master/d6/dc7/group__imgproc__hist.html#ga3a0af640716b456c3d14af8aee12e3ca)
        - [ ] [calcHist](https://docs.opencv.org/master/d6/dc7/group__imgproc__hist.html#ga4b2b5fd75503ff9e6844cc4dcdaed35d)
        - [ ] [compareHist](https://docs.opencv.org/master/d6/dc7/group__imgproc__hist.html#gaf4190090efa5c47cb367cf97a9a519bd)
        - [ ] [EMD](https://docs.opencv.org/master/d6/dc7/group__imgproc__hist.html#ga902b8e60cc7075c8947345489221e0e0)
        - [ ] [wrapperEMD](https://docs.opencv.org/master/d6/dc7/group__imgproc__hist.html#ga31fdda0864e64ca6b9de252a2611758b)

    - [ ] **Structural Analysis and Shape Descriptors - WORK STARTED** The following functions still need implementation:
        - [ ] [fitEllipse](https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#gaf259efaad93098103d6c27b9e4900ffa)
        - [ ] [fitEllipseAMS](https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga69e90cda55c4e192a8caa0b99c3e4550)
        - [ ] [fitEllipseDirect](https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga6421884fd411923a74891998bbe9e813)
        - [ ] [HuMoments](https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#gab001db45c1f1af6cbdbe64df04c4e944)
        - [ ] [intersectConvexConvex](https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga8e840f3f3695613d32c052bec89e782c)
        - [ ] [isContourConvex](https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga8abf8010377b58cbc16db6734d92941b)
        - [ ] [matchShapes](https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#gaadc90cb16e2362c9bd6e7363e6e4c317)
        - [ ] [minEnclosingTriangle](https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga1513e72f6bbdfc370563664f71e0542f)
        - [ ] [pointPolygonTest](https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga1a539e8db2135af2566103705d7a5722)
        - [ ] [rotatedRectangleIntersection](https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga8740e7645628c59d238b0b22c2abe2d4)

    - [ ] Motion Analysis and Object Tracking
    - [ ] **Feature Detection - WORK STARTED** The following functions still need implementation:
        - [ ] [cornerEigenValsAndVecs](https://docs.opencv.org/master/dd/d1a/group__imgproc__feature.html#ga4055896d9ef77dd3cacf2c5f60e13f1c)
        - [ ] [cornerHarris](https://docs.opencv.org/master/dd/d1a/group__imgproc__feature.html#gac1fc3598018010880e370e2f709b4345)
        - [ ] [cornerMinEigenVal](https://docs.opencv.org/master/dd/d1a/group__imgproc__feature.html#ga3dbce297c1feb859ee36707e1003e0a8)
        - [ ] [createLineSegmentDetector](https://docs.opencv.org/master/dd/d1a/group__imgproc__feature.html#ga6b2ad2353c337c42551b521a73eeae7d)
        - [ ] [preCornerDetect](https://docs.opencv.org/master/dd/d1a/group__imgproc__feature.html#gaa819f39b5c994871774081803ae22586)

    - [X] **Object Detection**

- [X] **imgcodecs. Image file reading and writing.**
- [X] **videoio. Video I/O**
- [X] **highgui. High-level GUI**
- [ ] **video. Video Analysis - WORK STARTED**
    - [X] **Motion Analysis**
    - [ ] **Object Tracking - WORK STARTED** The following functions still need implementation:
        - [ ] [buildOpticalFlowPyramid](https://docs.opencv.org/master/dc/d6b/group__video__track.html#ga86640c1c470f87b2660c096d2b22b2ce)
        - [ ] [estimateRigidTransform](https://docs.opencv.org/master/dc/d6b/group__video__track.html#ga762cbe5efd52cf078950196f3c616d48)
        - [ ] [findTransformECC](https://docs.opencv.org/master/dc/d6b/group__video__track.html#ga7ded46f9a55c0364c92ccd2019d43e3a)
        - [ ] [meanShift](https://docs.opencv.org/master/dc/d6b/group__video__track.html#ga7ded46f9a55c0364c92ccd2019d43e3a)
        - [ ] [CamShift](https://docs.opencv.org/master/dc/d6b/group__video__track.html#gaef2bd39c8356f423124f1fe7c44d54a1)
        - [ ] [DualTVL1OpticalFlow](https://docs.opencv.org/master/dc/d47/classcv_1_1DualTVL1OpticalFlow.html)
        - [ ] [FarnebackOpticalFlow](https://docs.opencv.org/master/de/d9e/classcv_1_1FarnebackOpticalFlow.html)
        - [ ] [KalmanFilter](https://docs.opencv.org/master/dd/d6a/classcv_1_1KalmanFilter.html)
        - [ ] [SparsePyrLKOpticalFlow](https://docs.opencv.org/master/d7/d08/classcv_1_1SparsePyrLKOpticalFlow.html)

- [ ] **calib3d. Camera Calibration and 3D Reconstruction - WORK STARTED**. The following functions still need implementation:
    - [ ] Camera Calibration
    - [ ] **Fisheye - WORK STARTED** The following functions still need implementation:
        - [ ] [calibrate](https://docs.opencv.org/3.4.1/db/d58/group__calib3d__fisheye.html#gad626a78de2b1dae7489e152a5a5a89e1)
        - [ ] [distortPoints](https://docs.opencv.org/3.4.1/db/d58/group__calib3d__fisheye.html#ga75d8877a98e38d0b29b6892c5f8d7765)
        - [ ] [estimateNewCameraMatrixForUndistortRectify](https://docs.opencv.org/3.4.1/db/d58/group__calib3d__fisheye.html#ga384940fdf04c03e362e94b6eb9b673c9)
        - [ ] [projectPoints](https://docs.opencv.org/3.4.1/db/d58/group__calib3d__fisheye.html#gab1ad1dc30c42ee1a50ce570019baf2c4)
        - [ ] [stereoCalibrate](https://docs.opencv.org/3.4.1/db/d58/group__calib3d__fisheye.html#gadbb3a6ca6429528ef302c784df47949b)
        - [ ] [stereoRectify](https://docs.opencv.org/3.4.1/db/d58/group__calib3d__fisheye.html#gac1af58774006689056b0f2ef1db55ecc)
        - [ ] [undistortPoints](https://docs.opencv.org/3.4.1/db/d58/group__calib3d__fisheye.html#gab738cdf90ceee97b2b52b0d0e7511541)

- [ ] **features2d. 2D Features Framework - WORK STARTED**
    - [X] **Feature Detection and Description**
    - [ ] **Descriptor Matchers - WORK STARTED** The following functions still need implementation:
        - [ ] [FlannBasedMatcher](https://docs.opencv.org/master/dc/de2/classcv_1_1FlannBasedMatcher.html)
    - [ ] **Drawing Function of Keypoints and Matches - WORK STARTED** The following function still needs implementation:
        - [ ] [drawMatches](https://docs.opencv.org/master/d4/d5d/group__features2d__draw.html#ga7421b3941617d7267e3f2311582f49e1)
    - [ ] Object Categorization
        - [ ] [BOWImgDescriptorExtractor](https://docs.opencv.org/master/d2/d6b/classcv_1_1BOWImgDescriptorExtractor.html)
        - [ ] [BOWKMeansTrainer](https://docs.opencv.org/master/d4/d72/classcv_1_1BOWKMeansTrainer.html)

- [X] **objdetect. Object Detection**
- [ ] **dnn. Deep Neural Network module - WORK STARTED** The following functions still need implementation:
    - [ ] [imagesFromBlob](https://docs.opencv.org/master/d6/d0f/group__dnn.html#ga4051b5fa2ed5f54b76c059a8625df9f5)
    - [ ] [blobFromImages](https://docs.opencv.org/master/d6/d0f/group__dnn.html#ga2b89ed84432e4395f5a1412c2926293c)
    - [ ] [NMSBoxes](https://docs.opencv.org/master/d6/d0f/group__dnn.html#ga9d118d70a1659af729d01b10233213ee)

- [ ] ml. Machine Learning
- [ ] flann. Clustering and Search in Multi-Dimensional Spaces
- [ ] photo. Computational Photography
- [ ] stitching. Images stitching
- [ ] cudaarithm. Operations on Matrices
- [ ] cudabgsegm. Background Segmentation
- [ ] cudacodec. Video Encoding/Decoding
- [ ] cudafeatures2d. Feature Detection and Description
- [ ] cudafilters. Image Filtering
- [ ] cudaimgproc. Image Processing
- [ ] cudalegacy. Legacy support
- [ ] cudaobjdetect. Object Detection
- [ ] cudaoptflow. Optical Flow
- [ ] cudastereo. Stereo Correspondence
- [ ] cudawarping. Image Warping
- [ ] cudev. Device layer
- [ ] shape. Shape Distance and Matching
- [ ] superres. Super Resolution
- [ ] videostab. Video Stabilization
- [ ] viz. 3D Visualizer

## Contrib modules list

- [ ] aruco. ArUco Marker Detection
- [ ] bgsegm. Improved Background-Foreground Segmentation Methods
- [ ] bioinspired. Biologically inspired vision models and derivated tools
- [ ] ccalib. Custom Calibration Pattern for 3D reconstruction
- [ ] cnn_3dobj. 3D object recognition and pose estimation API
- [ ] cvv. GUI for Interactive Visual Debugging of Computer Vision Programs
- [ ] datasets. Framework for working with different datasets
- [ ] dnn_modern. Deep Learning Modern Module
- [ ] dpm. Deformable Part-based Models
- [ ] **face. Face Recognition - WORK STARTED**
- [ ] freetype. Drawing UTF-8 strings with freetype/harfbuzz
- [ ] fuzzy. Image processing based on fuzzy mathematics
- [ ] hdf. Hierarchical Data Format I/O routines
- [X] **img_hash. The module brings implementations of different image hashing algorithms.**
- [ ] line_descriptor. Binary descriptors for lines extracted from an image
- [ ] matlab. MATLAB Bridge
- [ ] optflow. Optical Flow Algorithms
- [ ] phase_unwrapping. Phase Unwrapping API
- [ ] plot. Plot function for Mat data
- [ ] reg. Image Registration
- [ ] rgbd. RGB-Depth Processing
- [ ] saliency. Saliency API
- [ ] sfm. Structure From Motion
- [ ] stereo. Stereo Correspondance Algorithms
- [ ] structured_light. Structured Light API
- [ ] surface_matching. Surface Matching
- [ ] text. Scene Text Detection and Recognition
- [ ] **tracking. Tracking API - WORK STARTED**
- [ ] **xfeatures2d. Extra 2D Features Framework - WORK STARTED**
- [ ] ximgproc. Extended Image Processing
- [ ] xobjdetect. Extended object detection
- [ ] xphoto. Additional photo processing algorithms
