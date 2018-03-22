# Roadmap

This is a list of all of the functionality areas within OpenCV, and OpenCV Contrib.

Any section listed with an "X" means that all of the relevant OpenCV functionality has been wrapped for use within GoCV.

Any section listed with **WORK STARTED** indicates that some work has been done, but not all functionality in that module has been completed. If there are any functions listed under a section marked **WORK STARTED**, it indicates that that function still requires a wrapper implemented.

And any section that is simply listed, indicates that so far, no work has been done on that module.

Your pull requests will be greatly appreciated!

## Modules list

- [ ] core. Core functionality
    - [ ] **Basic structures - WORK STARTED**
    - [ ] **Operations on arrays - WORK STARTED**. The following functions still need implementation:
        - [ ] [checkRange](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga2bd19d89cae59361416736f87e3c7a64)
        - [ ] [determinant](https://docs.opencv.org/master/d2/de8/group__core__array.html#gaf802bd9ca3e07b8b6170645ef0611d0c)
        - [ ] [eigenNonSymmetric](https://docs.opencv.org/master/d2/de8/group__core__array.html#gaf51987e03cac8d171fbd2b327cf966f6)
        - [ ] [Mahalanobis](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga4493aee129179459cbfc6064f051aa7d)
        - [ ] [min](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga9af368f182ee76d0463d0d8d5330b764)
        - [ ] [minMaxIdx](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga7622c466c628a75d9ed008b42250a73f)
        - [ ] [mixChannels](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga51d768c270a1cdd3497255017c4504be)
        - [ ] [mulSpectrums](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga3ab38646463c59bf0ce962a9d51db64f)
        - [ ] [mulTransposed](https://docs.opencv.org/master/d2/de8/group__core__array.html#gadc4e49f8f7a155044e3be1b9e3b270ab)
        - [ ] [patchNaNs](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga62286befb7cde3568ff8c7d14d5079da)
        - [ ] [PCABackProject](https://docs.opencv.org/master/d2/de8/group__core__array.html#gab26049f30ee8e94f7d69d82c124faafc)
        - [ ] [PCACompute](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga4e2073c7311f292a0648f04c37b73781)
        - [ ] [PCAProject](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga6b9fbc7b3a99ebfd441bbec0a6bc4f88)
        - [ ] [perspectiveTransform](https://docs.opencv.org/master/d2/de8/group__core__array.html#gad327659ac03e5fd6894b90025e6900a7)
        - [ ] [phase](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga9db9ca9b4d81c3bde5677b8f64dc0137)
        - [ ] [polarToCart](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga581ff9d44201de2dd1b40a50db93d665)
        - [ ] [PSNR](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga07aaf34ae31d226b1b847d8bcff3698f)
        - [ ] [randn](https://docs.opencv.org/master/d2/de8/group__core__array.html#gaeff1f61e972d133a04ce3a5f81cf6808)
        - [ ] [randShuffle](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga6a789c8a5cb56c6dd62506179808f763)
        - [ ] [randu](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga1ba1026dca0807b27057ba6a49d258c0)
        - [ ] [reduce](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga4b78072a303f29d9031d56e5638da78e)
        - [ ] [repeat](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga496c3860f3ac44c40b48811333cfda2d)
        - [ ] [rotate](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga4ad01c0978b0ce64baa246811deeac24)
        - [ ] [scaleAdd](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga9e0845db4135f55dcf20227402f00d98)
        - [ ] [setIdentity](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga388d7575224a4a277ceb98ccaa327c99)
        - [ ] [setRNGSeed](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga757e657c037410d9e19e819569e7de0f)
        - [ ] [solve](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga12b43690dbd31fed96f213eefead2373)
        - [ ] [solveCubic](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga1c3b0b925b085b6e96931ee309e6a1da)
        - [ ] [solvePoly](https://docs.opencv.org/master/d2/de8/group__core__array.html#gac2f5e953016fabcdf793d762f4ec5dce)
        - [ ] [sort](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga45dd56da289494ce874be2324856898f)
        - [ ] [sortIdx](https://docs.opencv.org/master/d2/de8/group__core__array.html#gadf35157cbf97f3cb85a545380e383506)
        - [ ] [sqrt](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga186222c3919657890f88df5a1f64a7d7)
        - [ ] [SVBackSubst](https://docs.opencv.org/master/d2/de8/group__core__array.html#gab4e620e6fc6c8a27bb2be3d50a840c0b)
        - [ ] [SVDecomp](https://docs.opencv.org/master/d2/de8/group__core__array.html#gab477b5b7b39b370bb03e75b19d2d5109)
        - [ ] [theRNG](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga75843061d150ad6564b5447e38e57722)
        - [ ] [trace](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga3419ac19c7dcd2be4bd552a23e147dd8)
        - [ ] [transform](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga393164aa54bb9169ce0a8cc44e08ff22)
        - [ ] [transpose](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga46630ed6c0ea6254a35f447289bd7404)
        - [ ] [vconcat](https://docs.opencv.org/master/d2/de8/group__core__array.html#ga744f53b69f6e4f12156cdde4e76aed27)
    - [ ] XML/YAML Persistence
    - [ ] Clustering
    - [ ] Utility and system functions and macros
    - [ ] OpenGL interoperability
    - [ ] Intel IPP Asynchronous C/C++ Converters
    - [ ] Optimization Algorithms
    - [ ] OpenCL support 

- [ ] imgproc. Image processing
    - [ ] **Image Filtering - WORK STARTED**
    - [ ] **Geometric Image Transformations - WORK STARTED**
    - [ ] **Miscellaneous Image Transformations - WORK STARTED**
    - [ ] **Drawing Functions - WORK STARTED**
    - [ ] ColorMaps in OpenCV
    - [ ] Planar Subdivision
    - [ ] Histograms
    - [ ] Structural Analysis and Shape Descriptors
    - [ ] **Motion Analysis and Object Tracking - WORK STARTED**
    - [ ] **Feature Detection - WORK STARTED**
    - [ ] **Object Detection - WORK STARTED**

- [X] **imgcodecs. Image file reading and writing.**
- [X] **videoio. Video I/O**
- [X] **highgui. High-level GUI**
- [ ] **video. Video Analysis - WORK STARTED**
    - [X] **Motion Analysis**
    - [ ] **Object Tracking - WORK STARTED** (https://docs.opencv.org/master/dc/d6b/group__video__track.html)
        - [ ] [buildOpticalFlowPyramid] https://docs.opencv.org/master/dc/d6b/group__video__track.html#ga86640c1c470f87b2660c096d2b22b2ce
        - [ ] [estimateRigidTransform] https://docs.opencv.org/master/dc/d6b/group__video__track.html#ga762cbe5efd52cf078950196f3c616d48
        - [ ] [findTransformECC] https://docs.opencv.org/master/dc/d6b/group__video__track.html#ga7ded46f9a55c0364c92ccd2019d43e3a
        - [ ] [meanShift] https://docs.opencv.org/master/dc/d6b/group__video__track.html#ga7ded46f9a55c0364c92ccd2019d43e3a
        - [ ] [CamShift] https://docs.opencv.org/master/dc/d6b/group__video__track.html#gaef2bd39c8356f423124f1fe7c44d54a1
        - [ ] [DualTVL1OpticalFlow] https://docs.opencv.org/master/dc/d47/classcv_1_1DualTVL1OpticalFlow.html
        - [ ] [FarnebackOpticalFlow] https://docs.opencv.org/master/de/d9e/classcv_1_1FarnebackOpticalFlow.html
        - [ ] [KalmanFilter] https://docs.opencv.org/master/dd/d6a/classcv_1_1KalmanFilter.html
        - [ ] [SparsePyrLKOpticalFlow] https://docs.opencv.org/master/d7/d08/classcv_1_1SparsePyrLKOpticalFlow.html

- [ ] calib3d. Camera Calibration and 3D Reconstruction
- [ ] **features2d. 2D Features Framework - WORK STARTED**
    - [X] **Feature Detection and Description**
    - [ ] Descriptor Matchers (https://docs.opencv.org/master/d8/d9b/group__features2d__match.html)
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
- [ ] img_hash. The module brings implementations of different image hashing algorithms.
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
