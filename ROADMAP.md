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
        - [ ] [mulSpectrums](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga3ab38646463c59bf0ce962a9d51db64f)
        - [ ] [mulTransposed](https://docs.opencv.org/master/d2/de8/group__core__array.html#gadc4e49f8f7a155044e3be1b9e3b270ab)
        - [ ] [PCABackProject](https://docs.opencv.org/master/d2/de8/group__core__array.html#gab26049f30ee8e94f7d69d82c124faafc)
        - [ ] [PCACompute](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga4e2073c7311f292a0648f04c37b73781)
        - [ ] [PCAProject](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga6b9fbc7b3a99ebfd441bbec0a6bc4f88)
        - [ ] [phase](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga9db9ca9b4d81c3bde5677b8f64dc0137)
        - [ ] [polarToCart](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga581ff9d44201de2dd1b40a50db93d665)
        - [ ] [PSNR](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga07aaf34ae31d226b1b847d8bcff3698f)
        - [ ] [randn](https://docs.opencv.org/master/d2/de8/group__core__array.html#gaeff1f61e972d133a04ce3a5f81cf6808)
        - [ ] [randShuffle](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga6a789c8a5cb56c6dd62506179808f763)
        - [ ] [randu](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga1ba1026dca0807b27057ba6a49d258c0)
        - [ ] [reduce](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga4b78072a303f29d9031d56e5638da78e)
        - [ ] [repeat](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga496c3860f3ac44c40b48811333cfda2d)
        - [ ] [scaleAdd](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga9e0845db4135f55dcf20227402f00d98)
        - [ ] [setIdentity](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga388d7575224a4a277ceb98ccaa327c99)
        - [ ] [setRNGSeed](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga757e657c037410d9e19e819569e7de0f)
        - [ ] [solve](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga12b43690dbd31fed96f213eefead2373)
        - [ ] [solveCubic](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga1c3b0b925b085b6e96931ee309e6a1da)
        - [ ] [solvePoly](https://docs.opencv.org/master/d2/de8/group__core__array.html#gac2f5e953016fabcdf793d762f4ec5dce)
        - [ ] [sort](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga45dd56da289494ce874be2324856898f)
        - [ ] [sortIdx](https://docs.opencv.org/master/d2/de8/group__core__array.html#gadf35157cbf97f3cb85a545380e383506)
        - [ ] [SVBackSubst](https://docs.opencv.org/master/d2/de8/group__core__array.html#gab4e620e6fc6c8a27bb2be3d50a840c0b)
        - [ ] [SVDecomp](https://docs.opencv.org/master/d2/de8/group__core__array.html#gab477b5b7b39b370bb03e75b19d2d5109)
        - [ ] [theRNG](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga75843061d150ad6564b5447e38e57722)
        - [ ] [trace](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga3419ac19c7dcd2be4bd552a23e147dd8)
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
        - [ ] [filter2D](https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga27c049795ce870216ddfb366086b5a04)
        - [ ] [getDerivKernels](https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga6d6c23f7bd3f5836c31cfae994fc4aea)
        - [ ] [getGaborKernel](https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#gae84c92d248183bd92fa713ce51cc3599)
        - [ ] [getGaussianKernel](https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#gac05a120c1ae92a6060dd0db190a61afa)
        - [ ] [morphologyDefaultBorderValue](https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga94756fad83d9d24d29c9bf478558c40a)
        - [ ] [morphologyExWithParams](https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga67493776e3ad1a3df63883829375201f)
        - [ ] [pyrMeanShiftFiltering](https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga9fabdce9543bd602445f5db3827e4cc0)
        - [ ] [sepFilter2D](https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga910e29ff7d7b105057d1625a4bf6318d)
        - [ ] [Sobel](https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#gacea54f142e81b6758cb6f375ce782c8d)
        - [ ] [spatialGradient](https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga405d03b20c782b65a4daf54d233239a2)
    
    - [ ] **Geometric Image Transformations - WORK STARTED** The following functions still need implementation:
        - [ ] [convertMaps](https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga9156732fa8f01be9ebd1a194f2728b7f)
        - [ ] [getAffineTransform](https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga8f6d378f9f8eebb5cb55cd3ae295a999)
        - [ ] [getDefaultNewCameraMatrix](https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga744529385e88ef7bc841cbe04b35bfbf)
        - [ ] [getRectSubPix](https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga77576d06075c1a4b6ba1a608850cd614)
        - [ ] [initUndistortRectifyMap](https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga7dfb72c9cf9780a347fbe3d1c47e5d5a)
        - [ ] [initWideAngleProjMap](https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#gaceb049ec48898d1dadd5b50c604429c8)
        - [ ] [invertAffineTransform](https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga57d3505a878a7e1a636645727ca08f51)
        - [ ] [linearPolar](https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#gaa38a6884ac8b6e0b9bed47939b5362f3)
        - [ ] [logPolar](https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#gaec3a0b126a85b5ca2c667b16e0ae022d)
        - [ ] [remap](https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#gab75ef31ce5cdfb5c44b6da5f3b908ea4)
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
    - [ ] Histograms
    - [ ] **Structural Analysis and Shape Descriptors - WORK STARTED** The following functions still need implementation:
        - [ ] [boxPoints](https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#gaf78d467e024b4d7936cf9397185d2f5c)
        - [ ] [connectedComponents](https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#gaedef8c7340499ca391d459122e51bef5)
        - [ ] [connectedComponentsWithStats](https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga107a78bf7cd25dec05fb4dfc5c9e765f)
        - [ ] [fitEllipse](https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#gaf259efaad93098103d6c27b9e4900ffa)
        - [ ] [fitEllipseAMS](https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga69e90cda55c4e192a8caa0b99c3e4550)
        - [ ] [fitEllipseDirect](https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga6421884fd411923a74891998bbe9e813)
        - [ ] [fitLine](https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#gaf849da1fdafa67ee84b1e9a23b93f91f)
        - [ ] [HuMoments](https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#gab001db45c1f1af6cbdbe64df04c4e944)
        - [ ] [intersectConvexConvex](https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga8e840f3f3695613d32c052bec89e782c)
        - [ ] [isContourConvex](https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga8abf8010377b58cbc16db6734d92941b)
        - [ ] [matchShapes](https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#gaadc90cb16e2362c9bd6e7363e6e4c317)
        - [ ] [minAreaRect](https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga3d476a3417130ae5154aea421ca7ead9)
        - [ ] [minEnclosingCircle](https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga8ce13c24081bbc7151e9326f412190f1)
        - [ ] [minEnclosingTriangle](https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga1513e72f6bbdfc370563664f71e0542f)
        - [ ] [pointPolygonTest](https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga1a539e8db2135af2566103705d7a5722)
        - [ ] [rotatedRectangleIntersection](https://docs.opencv.org/master/d3/dc0/group__imgproc__shape.html#ga8740e7645628c59d238b0b22c2abe2d4)

    - [ ] Motion Analysis and Object Tracking
    - [ ] **Feature Detection - WORK STARTED** The following functions still need implementation:
        - [ ] [cornerEigenValsAndVecs](https://docs.opencv.org/master/dd/d1a/group__imgproc__feature.html#ga4055896d9ef77dd3cacf2c5f60e13f1c)
        - [ ] [cornerHarris](https://docs.opencv.org/master/dd/d1a/group__imgproc__feature.html#gac1fc3598018010880e370e2f709b4345)
        - [ ] [cornerMinEigenVal](https://docs.opencv.org/master/dd/d1a/group__imgproc__feature.html#ga3dbce297c1feb859ee36707e1003e0a8)
        - [ ] [createLineSegmentDetector](https://docs.opencv.org/master/dd/d1a/group__imgproc__feature.html#ga6b2ad2353c337c42551b521a73eeae7d)
        - [ ] [HoughLinesPointSet](https://docs.opencv.org/master/dd/d1a/group__imgproc__feature.html#ga2858ef61b4e47d1919facac2152a160e)
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

- [ ] calib3d. Camera Calibration and 3D Reconstruction
- [ ] **features2d. 2D Features Framework - WORK STARTED**
    - [X] **Feature Detection and Description**
    - [ ] **Descriptor Matchers - WORK STARTED** The following functions still need implementation:
        - [ ] [FlannBasedMatcher](https://docs.opencv.org/master/dc/de2/classcv_1_1FlannBasedMatcher.html)
    - [ ] Drawing Function of Keypoints and Matches (https://docs.opencv.org/master/d4/d5d/group__features2d__draw.html)
    - [ ] Object Categorization (https://docs.opencv.org/master/de/d24/group__features2d__category.html)

- [X] **objdetect. Object Detection**
- [ ] **dnn. Deep Neural Network module - WORK STARTED**
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
