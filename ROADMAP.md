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
        - [x] [mixChannels](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga51d768c270a1cdd3497255017c4504be)
        - [ ] [mulTransposed](https://docs.opencv.org/master/d2/de8/group__core__array.html#gadc4e49f8f7a155044e3be1b9e3b270ab)
        - [ ] [PCABackProject](https://docs.opencv.org/master/d2/de8/group__core__array.html#gab26049f30ee8e94f7d69d82c124faafc)
        - [ ] [PCACompute](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga4e2073c7311f292a0648f04c37b73781)
        - [ ] [PCAProject](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga6b9fbc7b3a99ebfd441bbec0a6bc4f88)
        - [ ] [PSNR](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga07aaf34ae31d226b1b847d8bcff3698f)
        - [ ] [randn](https://docs.opencv.org/master/d2/de8/group__core__array.html#gaeff1f61e972d133a04ce3a5f81cf6808)
        - [ ] [randShuffle](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga6a789c8a5cb56c6dd62506179808f763)
        - [ ] [randu](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga1ba1026dca0807b27057ba6a49d258c0)
        - [ ] [setRNGSeed](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga757e657c037410d9e19e819569e7de0f)
        - [ ] [SVBackSubst](https://docs.opencv.org/master/d2/de8/group__core__array.html#gab4e620e6fc6c8a27bb2be3d50a840c0b)
        - [ ] [SVDecomp](https://docs.opencv.org/master/d2/de8/group__core__array.html#gab477b5b7b39b370bb03e75b19d2d5109)
        - [ ] [theRNG](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga75843061d150ad6564b5447e38e57722)
    - [ ] XML/YAML Persistence
    - [ ] **Clustering - WORK STARTED**. The following functions still need implementation:
        - [ ] [partition](https://docs.opencv.org/master/d5/d38/group__core__cluster.html#ga2037c989e69b499c1aa271419f3a9b34)

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
        - [ ] [morphologyExWithParams](https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga67493776e3ad1a3df63883829375201f)
        - [ ] [pyrMeanShiftFiltering](https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga9fabdce9543bd602445f5db3827e4cc0)
    
    - [ ] **Geometric Image Transformations - WORK STARTED** The following functions still need implementation:
        - [ ] [convertMaps](https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga9156732fa8f01be9ebd1a194f2728b7f)
        - [ ] [getDefaultNewCameraMatrix](https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga744529385e88ef7bc841cbe04b35bfbf)
        - [ ] [initUndistortRectifyMap](https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga7dfb72c9cf9780a347fbe3d1c47e5d5a)
        - [ ] [initWideAngleProjMap](https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#gaceb049ec48898d1dadd5b50c604429c8)
        - [ ] [undistort](https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga69f2545a8b62a6b0fc2ee060dc30559d)

    - [ ] **Miscellaneous Image Transformations - WORK STARTED** The following functions still need implementation:
        - [ ] [cvtColorTwoPlane](https://docs.opencv.org/master/d7/d1b/group__imgproc__misc.html#ga8e873314e72a1a6c0252375538fbf753)
        - [ ] [floodFill](https://docs.opencv.org/master/d7/d1b/group__imgproc__misc.html#gaf1f55a048f8a45bc3383586e80b1f0d0)

    - [ ] **Drawing Functions - WORK STARTED** The following functions still need implementation:
        - [ ] [drawMarker](https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga482fa7b0f578fcdd8a174904592a6250)
        - [ ] [ellipse2Poly](https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga727a72a3f6a625a2ae035f957c61051f)
        - [ ] [fillConvexPoly](https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga906aae1606ea4ed2f27bec1537f6c5c2)
        - [ ] [getFontScaleFromHeight](https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga442ff925c1a957794a1309e0ed3ba2c3)
    
    - [ ] ColorMaps in OpenCV
    - [ ] Planar Subdivision
    - [ ] **Histograms  - WORK STARTED** The following functions still need implementation:
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
    - [ ] **Camera Calibration - WORK STARTED** The following functions still need implementation:
        - [ ] [calibrateCamera](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [calibrateCameraRO](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [calibrateHandEye](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [calibrationMatrixValues](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [checkChessboard](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [composeRT](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [computeCorrespondEpilines](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [convertPointsFromHomogeneous](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [convertPointsHomogeneous](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [convertPointsToHomogeneous](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [correctMatches](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [decomposeEssentialMat](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [decomposeHomographyMat](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [decomposeProjectionMatrix](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [drawChessboardCorners](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [drawFrameAxes](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [estimateAffine2D](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [estimateAffine3D](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [filterHomographyDecompByVisibleRefpoints](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [filterSpeckles](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [find4QuadCornerSubpix](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [findChessboardCorners](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [findChessboardCornersSB](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [findCirclesGrid](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [findEssentialMat](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [findFundamentalMat](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [getDefaultNewCameraMatrix](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [getOptimalNewCameraMatrix](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [getValidDisparityROI](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [initCameraMatrix2D](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [initUndistortRectifyMap](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [initWideAngleProjMap](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [matMulDeriv](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [projectPoints](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [recoverPose](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [rectify3Collinear](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [reprojectImageTo3D](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [Rodrigues](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [RQDecomp3x3](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [sampsonDistance](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [solveP3P](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [solvePnP](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [solvePnPGeneric](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [solvePnPRansac](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [solvePnPRefineLM](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [solvePnPRefineVVS](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [stereoCalibrate](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [stereoRectify](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [stereoRectifyUncalibrated](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [triangulatePoints](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [validateDisparity](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)

    - [ ] **Fisheye - WORK STARTED** The following functions still need implementation:
        - [ ] [calibrate](https://docs.opencv.org/master/db/d58/group__calib3d__fisheye.html#gad626a78de2b1dae7489e152a5a5a89e1)
        - [ ] [distortPoints](https://docs.opencv.org/master/db/d58/group__calib3d__fisheye.html#ga75d8877a98e38d0b29b6892c5f8d7765)
        - [ ] [projectPoints](https://docs.opencv.org/master/db/d58/group__calib3d__fisheye.html#gab1ad1dc30c42ee1a50ce570019baf2c4)
        - [ ] [stereoCalibrate](https://docs.opencv.org/master/db/d58/group__calib3d__fisheye.html#gadbb3a6ca6429528ef302c784df47949b)
        - [ ] [stereoRectify](https://docs.opencv.org/master/db/d58/group__calib3d__fisheye.html#gac1af58774006689056b0f2ef1db55ecc)

- [ ] **features2d. 2D Features Framework - WORK STARTED**
    - [X] **Feature Detection and Description**
    - [ ] **Descriptor Matchers - WORK STARTED** The following functions still need implementation:
        - [X] [FlannBasedMatcher](https://docs.opencv.org/master/dc/de2/classcv_1_1FlannBasedMatcher.html)
    - [ ] **Drawing Function of Keypoints and Matches - WORK STARTED** The following function still needs implementation:
        - [X] [drawMatches](https://docs.opencv.org/master/d4/d5d/group__features2d__draw.html#ga7421b3941617d7267e3f2311582f49e1)
    - [ ] Object Categorization
        - [ ] [BOWImgDescriptorExtractor](https://docs.opencv.org/master/d2/d6b/classcv_1_1BOWImgDescriptorExtractor.html)
        - [ ] [BOWKMeansTrainer](https://docs.opencv.org/master/d4/d72/classcv_1_1BOWKMeansTrainer.html)

- [X] **objdetect. Object Detection**
- [ ] **dnn. Deep Neural Network module - WORK STARTED** The following functions still need implementation:
    - [X] [NMSBoxes](https://docs.opencv.org/master/d6/d0f/group__dnn.html#ga9d118d70a1659af729d01b10233213ee)

- [ ] ml. Machine Learning
- [ ] flann. Clustering and Search in Multi-Dimensional Spaces
- [ ] photo. Computational Photography
- [ ] stitching. Images stitching
- [ ] **cudaarithm. Operations on Matrices - WORK STARTED** The following functions still need implementation:
    - [ ] [cv::cuda::abs](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga54a72bd772494ab34d05406fd76df2b6)
    - [ ] [cv::cuda::absdiff](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#gac062b283cf46ee90f74a773d3382ab54)
    - [ ] [cv::cuda::add](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga5d9794bde97ed23d1c1485249074a8b1)
    - [ ] [cv::cuda::addWeighted](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga2cd14a684ea70c6ab2a63ee90ffe6201)
    - [ ] [cv::cuda::bitwise_and](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga78d7c1a013877abd4237fbfc4e13bd76)
    - [ ] [cv::cuda::bitwise_not](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#gae58159a2259ae1acc76b531c171cf06a)
    - [ ] [cv::cuda::bitwise_or](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#gafd098ee3e51c68daa793999c1da3dfb7)
    - [ ] [cv::cuda::bitwise_xor](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga3d95d4faafb099aacf18e8b915a4ad8d)
    - [ ] [cv::cuda::cartToPolar](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga82210c7d1c1d42e616e554bf75a53480)
    - [ ] [cv::cuda::compare](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga4d41cd679f4a83862a3de71a6057db54)
    - [ ] [cv::cuda::divide](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga124315aa226260841e25cc0b9ea99dc3)
    - [ ] [cv::cuda::exp](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#gac6e51541d3bb0a7a396128e4d5919b61)
    - [ ] [cv::cuda::log](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#gaae9c60739e2d1a977b4d3250a0be42ca)
    - [ ] [cv::cuda::lshift](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#gafd072accecb14c9adccdad45e3bf2300)
    - [ ] [cv::cuda::magnitude](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga3d17f4fcd79d7c01fadd217969009463)
    - [ ] [cv::cuda::magnitudeSqr](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga7613e382d257e150033d0ce4d6098f6a)
    - [ ] [cv::cuda::max](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#gadb5dd3d870f10c0866035755b929b1e7)
    - [ ] [cv::cuda::min](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga74f0b05a65b3d949c237abb5e6c60867)
    - [ ] [cv::cuda::multiply](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga497cc0615bf717e1e615143b56f00591)
    - [ ] [cv::cuda::phase](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga5b75ec01be06dcd6e27ada09a0d4656a)
    - [ ] [cv::cuda::polarToCart](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga01516a286a329c303c2db746513dd9df)
    - [ ] [cv::cuda::pow](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga82d04ef4bcc4dfa9bfbe76488007c6c4)
    - [ ] [cv::cuda::rshift](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga87af0b66358cc302676f35c1fd56c2ed)
    - [ ] [cv::cuda::sqr](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga8aae233da90ce0ffe309ab8004342acb)
    - [ ] [cv::cuda::sqrt](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga09303680cb1a5521a922b6d392028d8c)
    - [ ] [cv::cuda::subtract](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga6eab60fc250059e2fda79c5636bd067f)
    - [X] [cv::cuda::threshold](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga40f1c94ae9a9456df3cad48e3cb008e1)

- [X] **cudabgsegm. Background Segmentation**
- [ ] cudacodec. Video Encoding/Decoding
- [ ] cudafeatures2d. Feature Detection and Description
- [ ] cudafilters. Image Filtering
- [ ] **cudaimgproc. Image Processing - WORK STARTED** The following functions still need implementation:
    - [X] [cv::cuda::CannyEdgeDetector](https://docs.opencv.org/master/d0/d43/classcv_1_1cuda_1_1CannyEdgeDetector.html)
    - [ ] [cv::cuda::TemplateMatching](https://docs.opencv.org/master/d2/d58/classcv_1_1cuda_1_1TemplateMatching.html)
    - [ ] [cv::cuda::alphaComp](https://docs.opencv.org/master/db/d8c/group__cudaimgproc__color.html#ga08a698700458d9311390997b57fbf8dc)
    - [X] [cv::cuda::cvtColor](https://docs.opencv.org/master/db/d8c/group__cudaimgproc__color.html#ga48d0f208181d5ca370d8ff6b62cbe826)
    - [ ] [cv::cuda::demosaicing](https://docs.opencv.org/master/db/d8c/group__cudaimgproc__color.html#ga7fb153572b573ebd2d7610fcbe64166e)
    - [ ] [cv::cuda::gammaCorrection](https://docs.opencv.org/master/db/d8c/group__cudaimgproc__color.html#gaf4195a8409c3b8fbfa37295c2b2c4729)
    - [ ] [cv::cuda::swapChannels](https://docs.opencv.org/master/db/d8c/group__cudaimgproc__color.html#ga75a29cc4a97cde0d43ea066b01de927e)
    - [ ] [cv::cuda::calcHist](https://docs.opencv.org/master/d8/d0e/group__cudaimgproc__hist.html#gaaf3944106890947020bb4522a7619c26)
    - [ ] [cv::cuda::CLAHE](https://docs.opencv.org/master/db/d79/classcv_1_1cuda_1_1CLAHE.html)
    - [ ] [cv::cuda::equalizeHist](https://docs.opencv.org/master/d8/d0e/group__cudaimgproc__hist.html#ga2384be74bd2feba7e6c46815513f0060)
    - [ ] [cv::cuda::evenLevels](https://docs.opencv.org/master/d8/d0e/group__cudaimgproc__hist.html#ga2f2cbd21dc6d7367a7c4ee1a826f389d)
    - [ ] [cv::cuda::histEven](https://docs.opencv.org/master/d8/d0e/group__cudaimgproc__hist.html#gacd3b14279fb77a57a510cb8c89a1856f)
    - [ ] [cv::cuda::histRange](https://docs.opencv.org/master/d8/d0e/group__cudaimgproc__hist.html#ga87819085c1059186d9cdeacd92cea783)
    - [ ] [cv::cuda::HoughCirclesDetector](https://docs.opencv.org/master/da/d80/classcv_1_1cuda_1_1HoughCirclesDetector.html)
    - [ ] [cv::cuda::HoughLinesDetector](https://docs.opencv.org/master/d2/dcd/classcv_1_1cuda_1_1HoughLinesDetector.html)
    - [ ] [cv::cuda::HoughSegmentDetector](https://docs.opencv.org/master/d6/df9/classcv_1_1cuda_1_1HoughSegmentDetector.html)
    - [ ] [cv::cuda::createGoodFeaturesToTrackDetector](https://docs.opencv.org/master/dc/d6d/group__cudaimgproc__feature.html#ga478b474a598ece101f7e706fee2c8e91)
    - [ ] [cv::cuda::createHarrisCorner](https://docs.opencv.org/master/dc/d6d/group__cudaimgproc__feature.html#ga3e5878a803e9bba51added0c10101979)
    - [ ] [cv::cuda::createMinEigenValCorner](https://docs.opencv.org/master/dc/d6d/group__cudaimgproc__feature.html#ga7457fd4b53b025f990b1c1dd1b749915)
    - [ ] [cv::cuda::bilateralFilter](https://docs.opencv.org/master/d0/d05/group__cudaimgproc.html#ga6abeaecdd4e7edc0bd1393a04f4f20bd)
    - [ ] [cv::cuda::blendLinear](https://docs.opencv.org/master/d0/d05/group__cudaimgproc.html#ga4793607e5729bcc15b27ea33d9fe335e)
    - [ ] [cv::cuda::meanShiftFiltering](https://docs.opencv.org/master/d0/d05/group__cudaimgproc.html#gae13b3035bc6df0e512d876dbb8c00555)
    - [ ] [cv::cuda::meanShiftProc](https://docs.opencv.org/master/d0/d05/group__cudaimgproc.html#ga6039dc8ecbe2f912bc83fcc9b3bcca39)
    - [ ] [cv::cuda::meanShiftSegmentation](https://docs.opencv.org/master/d0/d05/group__cudaimgproc.html#ga70ed80533a448829dc48cf22b1845c16)

- [ ] cudalegacy. Legacy support
- [X] **cudaobjdetect. Object Detection**
- [ ] **cudaoptflow. Optical Flow - WORK STARTED** The following functions still need implementation:
    - [ ] [BroxOpticalFlow](https://docs.opencv.org/master/d7/d18/classcv_1_1cuda_1_1BroxOpticalFlow.html)
    - [ ] [DenseOpticalFlow](https://docs.opencv.org/master/d6/d4a/classcv_1_1cuda_1_1DenseOpticalFlow.html)
    - [ ] [DensePyrLKOpticalFlow](https://docs.opencv.org/master/d0/da4/classcv_1_1cuda_1_1DensePyrLKOpticalFlow.html)
    - [ ] [FarnebackOpticalFlow](https://docs.opencv.org/master/d9/d30/classcv_1_1cuda_1_1FarnebackOpticalFlow.html)
    - [ ] [NvidiaHWOpticalFlow](https://docs.opencv.org/master/d5/d26/classcv_1_1cuda_1_1NvidiaHWOpticalFlow.html)
    - [ ] [NvidiaOpticalFlow_1_0](https://docs.opencv.org/master/dc/d9d/classcv_1_1cuda_1_1NvidiaOpticalFlow__1__0.html)
    - [ ] [SparseOpticalFlow](https://docs.opencv.org/master/d5/dcf/classcv_1_1cuda_1_1SparseOpticalFlow.html)
    - [ ] **[SparsePyrLKOpticalFlow](https://docs.opencv.org/master/d7/d05/classcv_1_1cuda_1_1SparsePyrLKOpticalFlow.html) - WORK STARTED**

- [ ] cudastereo. Stereo Correspondence
- [X] **cudawarping. Image Warping**
- [ ] cudev. Device layer
- [ ] shape. Shape Distance and Matching
- [ ] superres. Super Resolution
- [ ] videostab. Video Stabilization
- [ ] viz. 3D Visualizer

## Contrib modules list

- [ ] aruco. ArUco Marker Detection
- [X] **bgsegm. Improved Background-Foreground Segmentation Methods - WORK STARTED**
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
