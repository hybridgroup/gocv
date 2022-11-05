# Roadmap

This is a list of all of the functionality areas within OpenCV, and OpenCV Contrib.

Any section listed with an "X" means that all of the relevant OpenCV functionality has been wrapped for use within GoCV.

Any section listed with **WORK STARTED** indicates that some work has been done, but not all functionality in that module has been completed. If there are any functions listed under a section marked **WORK STARTED**, it indicates that that function still requires a wrapper implemented.

And any section that is simply listed, indicates that so far, no work has been done on that module.

Your pull requests will be greatly appreciated!

## Modules list

- [ ] **core. Core functionality - WORK STARTED**
    - [X] **Basic structures**
    - [ ] **Operations on arrays - WORK STARTED**. The following functions still need implementation:
        - [ ] [Mahalanobis](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga4493aee129179459cbfc6064f051aa7d)
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
        - [ ] [FileStorage](https://docs.opencv.org/master/da/d56/classcv_1_1FileStorage.html)

    - [ ] **Clustering - WORK STARTED**. The following functions still need implementation:
        - [ ] [partition](https://docs.opencv.org/master/d5/d38/group__core__cluster.html#ga2037c989e69b499c1aa271419f3a9b34)

    - [ ] Optimization Algorithms
        - [ ] [ConjGradSolver](https://docs.opencv.org/master/d0/d21/classcv_1_1ConjGradSolver.html)
        - [ ] [DownhillSolver](https://docs.opencv.org/master/d4/d43/classcv_1_1DownhillSolver.html)
        - [ ] [solveLP](https://docs.opencv.org/master/da/d01/group__core__optim.html#ga9a06d237a9d38ace891efa1ca1b5d00a)

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
        - [X] [matchShapes](https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#gaadc90cb16e2362c9bd6e7363e6e4c317)
        - [ ] [minEnclosingTriangle](https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga1513e72f6bbdfc370563664f71e0542f)
        - [ ] [rotatedRectangleIntersection](https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga8740e7645628c59d238b0b22c2abe2d4)

    - [ ] **Motion Analysis and Object Tracking - WORK STARTED** The following functions still need implementation:
        - [ ] [createHanningWindow](https://docs.opencv.org/master/d7/df3/group__imgproc__motion.html#ga80e5c3de52f6bab3a7c1e60e89308e1b)
        - [ ] [phaseCorrelate](https://docs.opencv.org/master/d7/df3/group__imgproc__motion.html#ga552420a2ace9ef3fb053cd630fdb4952)

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
        - [ ] [GOTURN](https://docs.opencv.org/master/d7/d4c/classcv_1_1TrackerGOTURN.html)

- [ ] **calib3d. Camera Calibration and 3D Reconstruction - WORK STARTED**. The following functions still need implementation:
    - [ ] **Camera Calibration - WORK STARTED** The following functions still need implementation:
        - [X] [calibrateCamera](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
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
        - [X] [estimateAffine2D](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [estimateAffine3D](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [filterHomographyDecompByVisibleRefpoints](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [filterSpeckles](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [ ] [find4QuadCornerSubpix](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [X] [findChessboardCorners](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
        - [X] [findChessboardCornersSB](https://docs.opencv.org/master/d9/d0c/group__calib3d.html)
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
    - [X] **Descriptor Matchers**
    - [X] **Drawing Function of Keypoints and Matches**
    - [ ] Object Categorization
        - [ ] [BOWImgDescriptorExtractor](https://docs.opencv.org/master/d2/d6b/classcv_1_1BOWImgDescriptorExtractor.html)
        - [ ] [BOWKMeansTrainer](https://docs.opencv.org/master/d4/d72/classcv_1_1BOWKMeansTrainer.html)

- [X] **objdetect. Object Detection**
- [X] **dnn. Deep Neural Network module**
- [ ] ml. Machine Learning
- [ ] flann. Clustering and Search in Multi-Dimensional Spaces
- [ ] **photo. Computational Photography - WORK STARTED** The following functions still need implementation:
    - [ ] [inpaint](https://docs.opencv.org/master/d7/d8b/group__photo__inpaint.html#gaedd30dfa0214fec4c88138b51d678085)
    - [ ] [denoise_TVL1](https://docs.opencv.org/master/d1/d79/group__photo__denoise.html#ga7602ed5ae17b7de40152b922227c4e4f)
    - [X] [fastNlMeansDenoising](https://docs.opencv.org/master/d1/d79/group__photo__denoise.html#ga4c6b0031f56ea3f98f768881279ffe93)
    - [X] [fastNlMeansDenoisingColored](https://docs.opencv.org/master/d1/d79/group__photo__denoise.html#ga03aa4189fc3e31dafd638d90de335617)
    - [X] [fastNlMeansDenoisingMulti](https://docs.opencv.org/master/d1/d79/group__photo__denoise.html#gaf4421bf068c4d632ea7f0aa38e0bf172)
    - [ ] [createCalibrateDebevec](https://docs.opencv.org/master/d6/df5/group__photo__hdr.html#ga7fed9707ad5f2cc0e633888867109f90)
    - [ ] [createCalibrateRobertson](https://docs.opencv.org/master/d6/df5/group__photo__hdr.html#gae77813a21cd351a596619e5ff013be5d)
    - [ ] [createMergeDebevec](https://docs.opencv.org/master/d6/df5/group__photo__hdr.html#gaa8eab36bc764abb2a225db7c945f87f9)
    - [ ] [createMergeRobertson](https://docs.opencv.org/master/d6/df5/group__photo__hdr.html#ga460d4a1df1a7e8cdcf7445bb87a8fb78)
    - [ ] [createTonemap](https://docs.opencv.org/master/d6/df5/group__photo__hdr.html#gabcbd653140b93a1fa87ccce94548cd0d)
    - [ ] [createTonemapDrago](https://docs.opencv.org/master/d6/df5/group__photo__hdr.html#ga72bf92bb6b8653ee4be650ac01cf50b6)
    - [ ] [createTonemapMantiuk](https://docs.opencv.org/master/d6/df5/group__photo__hdr.html#ga3b3f3bf083b7515802f039a6a70f2d21)
    - [ ] [createTonemapReinhard](https://docs.opencv.org/master/d6/df5/group__photo__hdr.html#gadabe7f6bf1fa96ad0fd644df9182c2fb)
    - [ ] [decolor](https://docs.opencv.org/master/d4/d32/group__photo__decolor.html#ga4864d4c007bda5dacdc5e9d4ed7e222c)
    - [X] [detailEnhance](https://docs.opencv.org/master/df/dac/group__photo__render.html#ga0de660cb6f371a464a74c7b651415975)
    - [X] [edgePreservingFilter](https://docs.opencv.org/master/df/dac/group__photo__render.html#gafaee2977597029bc8e35da6e67bd31f7)
    - [X] [pencilSketch](https://docs.opencv.org/master/df/dac/group__photo__render.html#gae5930dd822c713b36f8529b21ddebd0c)
    - [X] [stylization](https://docs.opencv.org/master/df/dac/group__photo__render.html#gacb0f7324017df153d7b5d095aed53206)

- [ ] stitching. Images stitching

## CUDA

- [ ] **core. - WORK STARTED** The following functions still need implementation:
    - [ ] [cv::cuda::convertFp16](https://docs.opencv.org/master/d8/d40/group__cudacore__init.html#gaa1c52258763197958eb9e6681917f723)
    - [ ] [cv::cuda::deviceSupports](https://docs.opencv.org/master/d8/d40/group__cudacore__init.html#ga170b10cc9af4aa8cce8c0afdb4b1d08c)
    - [X] [cv::cuda::getDevice](https://docs.opencv.org/master/d8/d40/group__cudacore__init.html#ga6ded4ed8e4fc483a9863d31f34ec9c0e)
    - [X] [cv::cuda::resetDevice](https://docs.opencv.org/master/d8/d40/group__cudacore__init.html#ga6153b6f461101374e655a54fc77e725e)
    - [X] [cv::cuda::setDevice](https://docs.opencv.org/master/d8/d40/group__cudacore__init.html#gaefa34186b185de47851836dba537828b)

- [ ] **cudaarithm. Operations on Matrices - WORK STARTED** The following functions still need implementation:
    - [ ] **core** The following functions still need implementation:
        - [ ] [cv::cuda::copyMakeBorder](https://docs.opencv.org/master/de/d09/group__cudaarithm__core.html#ga5368db7656eacf846b40089c98053a49)
        - [ ] [cv::cuda::createLookUpTable](https://docs.opencv.org/master/de/d09/group__cudaarithm__core.html#ga2d9d9780dea8c5cd85d3c19b7e01979c)
        - [ ] [cv::cuda::merge](https://docs.opencv.org/master/de/d09/group__cudaarithm__core.html#gaac939dc3b178ee92fb6e7078f342622c)
        - [ ] [cv::cuda::split](https://docs.opencv.org/master/de/d09/group__cudaarithm__core.html#gabe5013d55d4ff586b20393913726179e)
        - [ ] [cv::cuda::transpose](https://docs.opencv.org/master/de/d09/group__cudaarithm__core.html#ga327b71c3cb811a904ccf5fba37fc29f2)

    - [ ] **per-element operations - WORK STARTED** The following functions still need implementation:
        - [X] [cv::cuda::absdiff](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#gac062b283cf46ee90f74a773d3382ab54)
        - [X] [cv::cuda::add](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga5d9794bde97ed23d1c1485249074a8b1)
        - [ ] [cv::cuda::addWeighted](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga2cd14a684ea70c6ab2a63ee90ffe6201)
        - [X] [cv::cuda::bitwise_and](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga78d7c1a013877abd4237fbfc4e13bd76)
        - [X] [cv::cuda::bitwise_not](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#gae58159a2259ae1acc76b531c171cf06a)
        - [X] [cv::cuda::bitwise_or](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#gafd098ee3e51c68daa793999c1da3dfb7)
        - [X] [cv::cuda::bitwise_xor](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga3d95d4faafb099aacf18e8b915a4ad8d)
        - [ ] [cv::cuda::cartToPolar](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga82210c7d1c1d42e616e554bf75a53480)
        - [ ] [cv::cuda::compare](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga4d41cd679f4a83862a3de71a6057db54)
        - [X] [cv::cuda::divide](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga124315aa226260841e25cc0b9ea99dc3)
        - [X] [cv::cuda::exp](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#gac6e51541d3bb0a7a396128e4d5919b61)
        - [ ] [cv::cuda::inRange](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#gaf611ab6b1d85e951feb6f485b1ed9672)
        - [X] [cv::cuda::log](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#gaae9c60739e2d1a977b4d3250a0be42ca)
        - [ ] [cv::cuda::lshift](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#gafd072accecb14c9adccdad45e3bf2300)
        - [ ] [cv::cuda::magnitude](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga3d17f4fcd79d7c01fadd217969009463)
        - [ ] [cv::cuda::magnitudeSqr](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga7613e382d257e150033d0ce4d6098f6a)
        - [X] [cv::cuda::max](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#gadb5dd3d870f10c0866035755b929b1e7)
        - [X] [cv::cuda::min](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga74f0b05a65b3d949c237abb5e6c60867)
        - [X] [cv::cuda::multiply](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga497cc0615bf717e1e615143b56f00591)
        - [ ] [cv::cuda::phase](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga5b75ec01be06dcd6e27ada09a0d4656a)
        - [ ] [cv::cuda::polarToCart](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga01516a286a329c303c2db746513dd9df)
        - [ ] [cv::cuda::pow](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga82d04ef4bcc4dfa9bfbe76488007c6c4)
        - [ ] [cv::cuda::rshift](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga87af0b66358cc302676f35c1fd56c2ed)
        - [X] [cv::cuda::sqr](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga8aae233da90ce0ffe309ab8004342acb)
        - [X] [cv::cuda::sqrt](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga09303680cb1a5521a922b6d392028d8c)
        - [X] [cv::cuda::subtract](https://docs.opencv.org/master/d8/d34/group__cudaarithm__elem.html#ga6eab60fc250059e2fda79c5636bd067f)

    - [ ] **matrix reductions** The following functions still need implementation:
        - [ ] [cv::cuda::absSum](https://docs.opencv.org/master/d5/de6/group__cudaarithm__reduce.html#ga690fa79ba4426c53f7d2bebf3d37a32a)
        - [ ] [cv::cuda::calcAbsSum](https://docs.opencv.org/master/d5/de6/group__cudaarithm__reduce.html#ga15c403b76ab2c4d7ed0f5edc09891b7e)
        - [ ] [cv::cuda::calcNorm](https://docs.opencv.org/master/d5/de6/group__cudaarithm__reduce.html#ga39d2826990d29b7e4b69dbe02bdae2e1)
        - [ ] [cv::cuda::calcNormDiff](https://docs.opencv.org/master/d5/de6/group__cudaarithm__reduce.html#ga9be3d9a7b6c5760955f37d1039d01265)
        - [ ] [cv::cuda::calcSqrSum](https://docs.opencv.org/master/d5/de6/group__cudaarithm__reduce.html#gac998c83597f6c206c78cee16aa87946f)
        - [ ] [cv::cuda::calcSum](https://docs.opencv.org/master/d5/de6/group__cudaarithm__reduce.html#ga98a09144047f09f5cb1d6b6ea8e0856f)
        - [ ] [cv::cuda::countNonZero](https://docs.opencv.org/master/d5/de6/group__cudaarithm__reduce.html#ga98a09144047f09f5cb1d6b6ea8e0856f)
        - [ ] [cv::cuda::findMinMax](https://docs.opencv.org/master/d5/de6/group__cudaarithm__reduce.html#gae7f5f2aa9f65314470a76fccdff887f2)
        - [ ] [cv::cuda::findMinMaxLoc](https://docs.opencv.org/master/d5/de6/group__cudaarithm__reduce.html#ga93916bc473a62d215d1130fab84d090a)
        - [ ] [cv::cuda::integral](https://docs.opencv.org/master/d5/de6/group__cudaarithm__reduce.html#ga07e5104eba4bf45212ac9dbc5bf72ba6)
        - [ ] [cv::cuda::meanStdDev](https://docs.opencv.org/master/d5/de6/group__cudaarithm__reduce.html#ga990a4db4c6d7e8f0f3a6685ba48fbddc)
        - [ ] [cv::cuda::minMax](https://docs.opencv.org/master/d5/de6/group__cudaarithm__reduce.html#ga8d7de68c10717cf25e787e3c20d2dfee)
        - [ ] [cv::cuda::minMaxLoc](https://docs.opencv.org/master/d5/de6/group__cudaarithm__reduce.html#ga5cacbc2a2323c4eaa81e7390c5d9f530)
        - [ ] [cv::cuda::norm](https://docs.opencv.org/master/d5/de6/group__cudaarithm__reduce.html#ga6c01988a58d92126a7c60a4ab76d8324)
        - [ ] [cv::cuda::normalize](https://docs.opencv.org/master/d5/de6/group__cudaarithm__reduce.html#ga4da4738b9956a5baaa2f5f8c2fba438a)
        - [ ] [cv::cuda::rectStdDev](https://docs.opencv.org/master/d5/de6/group__cudaarithm__reduce.html#gac311484a4e57cab2ce2cfdc195fda7ee)
        - [ ] [cv::cuda::reduce](https://docs.opencv.org/master/d5/de6/group__cudaarithm__reduce.html#ga21d57f661db7be093caf2c4378be2007)
        - [ ] [cv::cuda::sqrIntegral](https://docs.opencv.org/master/d5/de6/group__cudaarithm__reduce.html#ga40c75196202706399a60bf6ba7a052ac)
        - [ ] [cv::cuda::sqlSum](https://docs.opencv.org/master/d5/de6/group__cudaarithm__reduce.html#ga056c804ebf5d2eb9f6f35e3dcb01524c)
        - [ ] [cv::cuda::sum](https://docs.opencv.org/master/d5/de6/group__cudaarithm__reduce.html#ga1f582844670199281e8012733b50c582)

    - [ ] **Operations on matrices** The following functions still need implementation:
        - [ ] [cv::cuda::createConvolution](https://docs.opencv.org/4.5.0/d9/d88/group__cudaarithm__arithm.html#ga2695e05ef624bf3ce03cfbda383a821d)
        - [ ] [cv::cuda::createDFT](https://docs.opencv.org/4.5.0/d9/d88/group__cudaarithm__arithm.html#ga0f72d063b73c8bb995678525eb076f10)
        - [ ] [cv::cuda::dft](https://docs.opencv.org/4.5.0/d9/d88/group__cudaarithm__arithm.html#gadea99cb15a715c983bcc2870d65a2e78)
        - [ ] [cv::cuda::gemm](https://docs.opencv.org/4.5.0/d9/d88/group__cudaarithm__arithm.html#ga42efe211d7a43bbc922da044c4f17130)
        - [ ] [cv::cuda::mulAndScaleSpectrums](https://docs.opencv.org/4.5.0/d9/d88/group__cudaarithm__arithm.html#ga5704c25b8be4f19da812e6d98c8ee464)
        - [ ] [cv::cuda::mulSpectrums](https://docs.opencv.org/4.5.0/d9/d88/group__cudaarithm__arithm.html#gab3e8900d67c4f59bdc137a0495206cd8)

- [X] **cudabgsegm. Background Segmentation**

- [ ] **cudacodec** Video Encoding/Decoding. The following functions still need implementation:
    - [ ] [cv::cuda::VideoReader](https://docs.opencv.org/master/db/ded/classcv_1_1cudacodec_1_1VideoReader.html)
    - [ ] [cv::cuda::VideoWriter](https://docs.opencv.org/master/df/dde/classcv_1_1cudacodec_1_1VideoWriter.html)

- [ ] **cudafeatures2d** Feature Detection and Description. The following functions still need implementation:
    - [ ] [cv::cuda::FastFeatureDetector](https://docs.opencv.org/master/d4/d6a/classcv_1_1cuda_1_1FastFeatureDetector.html)
    - [ ] [cv::cuda::ORB](https://docs.opencv.org/master/da/d44/classcv_1_1cuda_1_1ORB.html)

- [ ] **cudafilters. Image Filtering - WORK STARTED** The following functions still need implementation:
    - [ ] [cv::cuda::createBoxFilter](https://docs.opencv.org/master/dc/d66/group__cudafilters.html#ga3113b66e289bad7caef412e6e13ec2be)
    - [ ] [cv::cuda::createBoxMaxFilter](https://docs.opencv.org/master/dc/d66/group__cudafilters.html#gaaf4740c51128d23a37f6f1b22cee49e8)
    - [ ] [cv::cuda::createBoxMinFilter](https://docs.opencv.org/master/dc/d66/group__cudafilters.html#ga77fd36949bc8d92aabc120b4b1cfaafa)
    - [ ] [cv::cuda::createColumnSumFilter](https://docs.opencv.org/master/dc/d66/group__cudafilters.html#gac13bf7c41a34bfde2a7f33ad8caacfdf)
    - [ ] [cv::cuda::createDerivFilter](https://docs.opencv.org/master/dc/d66/group__cudafilters.html#ga14d76dc6982ce739c67198f52bc16ee1)
    - [ ] [cv::cuda::createLaplacianFilter](https://docs.opencv.org/master/dc/d66/group__cudafilters.html#ga53126e88bb7e6185dcd5628e28e42cd2)
    - [ ] [cv::cuda::createLinearFilter](https://docs.opencv.org/master/dc/d66/group__cudafilters.html#ga57cb1804ad9d1280bf86433858daabf9)
    - [ ] [cv::cuda::createMorphologyFilter](https://docs.opencv.org/master/dc/d66/group__cudafilters.html#gae58694e07be6bdbae126f36c75c08ee6)
    - [ ] [cv::cuda::createRowSumFilter](https://docs.opencv.org/master/dc/d66/group__cudafilters.html#gaf735de273ccb5072f3c27816fb97a53a)
    - [ ] [cv::cuda::createScharrFilter](https://docs.opencv.org/master/dc/d66/group__cudafilters.html#ga4ac8df158e5771ddb0bd5c9091188ce6)
    - [ ] [cv::cuda::createSeparableLinearFilter](https://docs.opencv.org/master/dc/d66/group__cudafilters.html#gaf7b79a9a92992044f328dad07a52c4bf)

- [ ] **cudaimgproc. Image Processing - WORK STARTED** The following functions still need implementation:
    - [ ] [cv::cuda::TemplateMatching](https://docs.opencv.org/master/d2/d58/classcv_1_1cuda_1_1TemplateMatching.html)
    - [ ] [cv::cuda::alphaComp](https://docs.opencv.org/master/db/d8c/group__cudaimgproc__color.html#ga08a698700458d9311390997b57fbf8dc)
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
    - [ ] [cv::cuda::createGoodFeaturesToTrackDetector](https://docs.opencv.org/master/dc/d6d/group__cudaimgproc__feature.html#ga478b474a598ece101f7e706fee2c8e91)
    - [ ] [cv::cuda::createHarrisCorner](https://docs.opencv.org/master/dc/d6d/group__cudaimgproc__feature.html#ga3e5878a803e9bba51added0c10101979)
    - [ ] [cv::cuda::createMinEigenValCorner](https://docs.opencv.org/master/dc/d6d/group__cudaimgproc__feature.html#ga7457fd4b53b025f990b1c1dd1b749915)
    - [ ] [cv::cuda::bilateralFilter](https://docs.opencv.org/master/d0/d05/group__cudaimgproc.html#ga6abeaecdd4e7edc0bd1393a04f4f20bd)
    - [ ] [cv::cuda::blendLinear](https://docs.opencv.org/master/d0/d05/group__cudaimgproc.html#ga4793607e5729bcc15b27ea33d9fe335e)
    - [ ] [cv::cuda::meanShiftFiltering](https://docs.opencv.org/master/d0/d05/group__cudaimgproc.html#gae13b3035bc6df0e512d876dbb8c00555)
    - [ ] [cv::cuda::meanShiftProc](https://docs.opencv.org/master/d0/d05/group__cudaimgproc.html#ga6039dc8ecbe2f912bc83fcc9b3bcca39)
    - [ ] [cv::cuda::meanShiftSegmentation](https://docs.opencv.org/master/d0/d05/group__cudaimgproc.html#ga70ed80533a448829dc48cf22b1845c16)

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

- [ ] **cudastereo** Stereo Correspondence
    - [ ] [cv::cuda::createDisparityBilateralFilter](https://docs.opencv.org/master/dd/d47/group__cudastereo.html#gaafb5f9902f7a9e74cb2cd4e680569590)
    - [ ] [cv::cuda::createStereoBeliefPropagation](https://docs.opencv.org/master/dd/d47/group__cudastereo.html#ga8d22dd80bdfb4e3d7d2ac09e8a07c22b)
    - [ ] [cv::cuda::createStereoBM](https://docs.opencv.org/master/dd/d47/group__cudastereo.html#ga77edc901350dd0a7f46ec5aca4138039)
    - [ ] [cv::cuda::createStereoConstantSpaceBP](https://docs.opencv.org/master/dd/d47/group__cudastereo.html#gaec3b49c7cf9f7701a6f549a227be4df2)
    - [ ] [cv::cuda::createStereoSGM](https://docs.opencv.org/master/dd/d47/group__cudastereo.html#gafb7e5284de5f488d664c3155acb12c93)
    - [ ] [cv::cuda::drawColorDisp](https://docs.opencv.org/master/dd/d47/group__cudastereo.html#ga469b23a77938dd7c06861e59cecc08c5)
    - [ ] [cv::cuda::reprojectImageTo3D](https://docs.opencv.org/master/dd/d47/group__cudastereo.html#gaff851e3932da0f3e74d1be1d8855f094)

- [X] **cudawarping. Image Warping**

## Contrib modules list

- [ ] alphamat. Alpha Matting
- [X] **aruco. ArUco Marker Detection - WORK STARTED**
- [ ] barcode. Barcode detecting and decoding methods
- [X] **bgsegm. Improved Background-Foreground Segmentation Methods - WORK STARTED**
- [ ] bioinspired. Biologically inspired vision models and derivated tools
- [ ] ccalib. Custom Calibration Pattern for 3D reconstruction
- [ ] cnn_3dobj. 3D object recognition and pose estimation API
- [ ] cvv. GUI for Interactive Visual Debugging of Computer Vision Programs
- [ ] datasets. Framework for working with different datasets
- [ ] dnn_modern. Deep Learning Modern Module
- [ ] dnn_objdetect. DNN used for object detection
- [ ] dnn_superres. DNN used for super resolution
- [ ] dpm. Deformable Part-based Models
- [ ] **face. Face Recognition - WORK STARTED**
- [ ] freetype. Drawing UTF-8 strings with freetype/harfbuzz
- [ ] fuzzy. Image processing based on fuzzy mathematics
- [ ] hdf. Hierarchical Data Format I/O routines
- [ ] hfs. Hierarchical Feature Selection for Efficient Image Segmentation
- [X] **img_hash. The module brings implementations of different image hashing algorithms.**
- [ ] intensity_transform. The module brings implementations of intensity transformation algorithms to adjust image contrast.
- [ ] line_descriptor. Binary descriptors for lines extracted from an image
- [ ] mcc. Macbeth Chart module
- [ ] optflow. Optical Flow Algorithms
- [ ] ovis. OGRE 3D Visualiser
- [ ] phase_unwrapping. Phase Unwrapping API
- [ ] plot. Plot function for Mat data
- [ ] quality. Image Quality Analysis (IQA) API
- [ ] rapid. silhouette based 3D object tracking
- [ ] reg. Image Registration
- [ ] rgbd. RGB-Depth Processing
- [ ] saliency. Saliency API
- [ ] sfm. Structure From Motion
- [ ] shape. Shape Distance and Matching
- [ ] stereo. Stereo Correspondance Algorithms
- [ ] structured_light. Structured Light API
- [ ] superres. Super Resolution
- [ ] surface_matching. Surface Matching
- [ ] text. Scene Text Detection and Recognition
- [ ] **tracking. Tracking API - WORK STARTED**
- [ ] videostab. Video Stabilization
- [ ] viz. 3D Visualizer
- [X] **wechat_qrcode. WeChat QR code detector for detecting and parsing QR code**
- [ ] **xfeatures2d. Extra 2D Features Framework - WORK STARTED**
- [ ] **ximgproc. Extended Image Processing - WORK STARTED**
- [ ] xobjdetect. Extended object detection
- [X] **xphoto. Additional photo processing algorithms**
