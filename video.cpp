#include "video.h"

BackgroundSubtractorMOG2 BackgroundSubtractorMOG2_Create() {
    try {
		return new cv::Ptr<cv::BackgroundSubtractorMOG2>(cv::createBackgroundSubtractorMOG2());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return NULL;
    }
}

BackgroundSubtractorMOG2 BackgroundSubtractorMOG2_CreateWithParams(int history, double varThreshold, bool detectShadows) {
    try {
		return new cv::Ptr<cv::BackgroundSubtractorMOG2>(cv::createBackgroundSubtractorMOG2(history,varThreshold,detectShadows));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return NULL;
    }
}

BackgroundSubtractorKNN BackgroundSubtractorKNN_Create() {
    try {
		return new cv::Ptr<cv::BackgroundSubtractorKNN>(cv::createBackgroundSubtractorKNN());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return NULL;
    }
}

BackgroundSubtractorKNN BackgroundSubtractorKNN_CreateWithParams(int history, double dist2Threshold, bool detectShadows) {
    try {
		return new cv::Ptr<cv::BackgroundSubtractorKNN>(cv::createBackgroundSubtractorKNN(history,dist2Threshold,detectShadows));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return NULL;
    }
}

void BackgroundSubtractorMOG2_Close(BackgroundSubtractorMOG2 b) {
    delete b;
}

OpenCVResult BackgroundSubtractorMOG2_Apply(BackgroundSubtractorMOG2 b, Mat src, Mat dst) {
  	try {
		(*b)->apply(*src, *dst);
      	OpenCVResult result = {0, NULL};
      	return result;
  	} catch(const cv::Exception& e) {
      	OpenCVResult result = {e.code, e.what()};
      	return result;
    }
}

OpenCVResult BackgroundSubtractorMOG2_ApplyWithParams(BackgroundSubtractorMOG2 b, Mat src, Mat dst, double learningRate) {
    try {
		(*b)->apply(*src, *dst, learningRate);
		OpenCVResult result = {0, NULL};
		return result;
	} catch(const cv::Exception& e) {
		OpenCVResult result = {e.code, e.what()};
		return result;
  	}
}

void BackgroundSubtractorKNN_Close(BackgroundSubtractorKNN k) {
    delete k;
}

OpenCVResult BackgroundSubtractorKNN_Apply(BackgroundSubtractorKNN k, Mat src, Mat dst) {
    try {
		(*k)->apply(*src, *dst);
		OpenCVResult result = {0, NULL};
		return result;
	} catch(const cv::Exception& e) {
		OpenCVResult result = {e.code, e.what()};
		return result;
  	}
}

OpenCVResult CalcOpticalFlowFarneback(Mat prevImg, Mat nextImg, Mat flow, double scale, int levels,
                              int winsize, int iterations, int polyN, double polySigma, int flags) {
	try {
		cv::calcOpticalFlowFarneback(*prevImg, *nextImg, *flow, scale, levels, winsize, iterations, polyN, polySigma, flags);
		OpenCVResult result = {0, NULL};
		return result;
	} catch(const cv::Exception& e) {
		OpenCVResult result = {e.code, e.what()};
		return result;
  	}
}

OpenCVResult CalcOpticalFlowPyrLK(Mat prevImg, Mat nextImg, Mat prevPts, Mat nextPts, Mat status, Mat err) {
    try {
		cv::calcOpticalFlowPyrLK(*prevImg, *nextImg, *prevPts, *nextPts, *status, *err);
		OpenCVResult result = {0, NULL};
		return result;
	} catch(const cv::Exception& e) {
		OpenCVResult result = {e.code, e.what()};
		return result;
	}
}

OpenCVResult CalcOpticalFlowPyrLKWithParams(Mat prevImg, Mat nextImg, Mat prevPts, Mat nextPts, Mat status, Mat err, Size winSize, int maxLevel, TermCriteria criteria, int flags, double minEigThreshold){
    try {
		cv::Size sz(winSize.width, winSize.height);
		cv::calcOpticalFlowPyrLK(*prevImg, *nextImg, *prevPts, *nextPts, *status, *err, sz, maxLevel, *criteria, flags, minEigThreshold);
		OpenCVResult result = {0, NULL};
		return result;
	} catch(const cv::Exception& e) {
		OpenCVResult result = {e.code, e.what()};
		return result;
  	}
}

double FindTransformECC(Mat templateImage, Mat inputImage, Mat warpMatrix, int motionType, TermCriteria criteria, Mat inputMask, int gaussFiltSize){
    try {
		return cv::findTransformECC(*templateImage, *inputImage, *warpMatrix, motionType, *criteria, *inputMask, gaussFiltSize);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return 0.0;
	}
}

bool Tracker_Init(Tracker self, Mat image, Rect boundingBox) {
    try {
		cv::Rect bb(boundingBox.x, boundingBox.y, boundingBox.width, boundingBox.height);

		(*self)->init(*image, bb);
		return true;
	} catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return false;
	}
}

bool Tracker_Update(Tracker self, Mat image, Rect* boundingBox) {
    try {
		cv::Rect bb;
		bool ret = (*self)->update(*image, bb);
		boundingBox->x = int(bb.x);
		boundingBox->y = int(bb.y);
		boundingBox->width = int(bb.width);
		boundingBox->height = int(bb.height);
		return ret;
	} catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return false;
	}
}

TrackerMIL TrackerMIL_Create() {
    try {
		return new cv::Ptr<cv::TrackerMIL>(cv::TrackerMIL::create());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return NULL;
    }
}

void TrackerMIL_Close(TrackerMIL self) {
    delete self;
}

TrackerGOTURN TrackerGOTURN_Create(void){
    try {
		return new cv::Ptr<cv::TrackerGOTURN>(cv::TrackerGOTURN::create());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return NULL;
    }
}

TrackerGOTURN TrackerGOTURN_CreateWithParams(const char* modelBin, const char* modelTxt){
    try {
		cv::TrackerGOTURN::Params params;
		params.modelBin = modelBin;
		params.modelTxt = modelTxt;
	  
		return new cv::Ptr<cv::TrackerGOTURN>(cv::TrackerGOTURN::create(params));
	} catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return NULL;
    }
}


void TrackerGOTURN_Close(TrackerGOTURN tr) {
    delete tr;
}

KalmanFilter KalmanFilter_New(int dynamParams, int measureParams) {
    try {
		return new cv::KalmanFilter(dynamParams, measureParams, 0, CV_32F);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return NULL;
    }
}

KalmanFilter KalmanFilter_NewWithParams(int dynamParams, int measureParams, int controlParams, int type) {
    try {
		return new cv::KalmanFilter(dynamParams, measureParams, controlParams, type);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return NULL;
    }
}

OpenCVResult KalmanFilter_Init(KalmanFilter kf, int dynamParams, int measureParams) {
    try {
		kf->init(dynamParams, measureParams, 0, CV_32F);
		OpenCVResult result = {0, NULL};
		return result;
	} catch(const cv::Exception& e) {
		OpenCVResult result = {e.code, e.what()};
		return result;
  	}
}

OpenCVResult KalmanFilter_InitWithParams(KalmanFilter kf, int dynamParams, int measureParams, int controlParams, int type) {
    try {
		kf->init(dynamParams, measureParams, controlParams, type);
		OpenCVResult result = {0, NULL};
		return result;
	} catch(const cv::Exception& e) {
		OpenCVResult result = {e.code, e.what()};
		return result;
  	}
}

void KalmanFilter_Close(KalmanFilter kf) {
    delete kf;
}

Mat KalmanFilter_Predict(KalmanFilter kf) {
    try {
		return new cv::Mat(kf->predict());
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return new cv::Mat();
    }
}

Mat KalmanFilter_PredictWithParams(KalmanFilter kf, Mat control) {
    try {
		return new cv::Mat(kf->predict(*control));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return new cv::Mat();
    }
}

Mat KalmanFilter_Correct(KalmanFilter kf, Mat measurement) {
    try {
		return new cv::Mat(kf->correct(*measurement));
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return new cv::Mat();
    }
}

Mat KalmanFilter_GetStatePre(KalmanFilter kf) {
    try {
		return new cv::Mat(kf->statePre);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return new cv::Mat();
    }
}

Mat KalmanFilter_GetStatePost(KalmanFilter kf) {
    try {
		return new cv::Mat(kf->statePost);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return new cv::Mat();
    }
}

Mat KalmanFilter_GetTransitionMatrix(KalmanFilter kf) {
    try {
		return new cv::Mat(kf->transitionMatrix);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return new cv::Mat();
    }
}

Mat KalmanFilter_GetControlMatrix(KalmanFilter kf) {
    try {
		return new cv::Mat(kf->controlMatrix);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return new cv::Mat();
    }
}

Mat KalmanFilter_GetMeasurementMatrix(KalmanFilter kf) {
    try {
		return new cv::Mat(kf->measurementMatrix);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return new cv::Mat();
    }
}

Mat KalmanFilter_GetProcessNoiseCov(KalmanFilter kf) {
    try {
		return new cv::Mat(kf->processNoiseCov);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return new cv::Mat();
    }
}

Mat KalmanFilter_GetMeasurementNoiseCov(KalmanFilter kf) {
    try {
		return new cv::Mat(kf->measurementNoiseCov);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return new cv::Mat();
    }
}

Mat KalmanFilter_GetErrorCovPre(KalmanFilter kf) {
    try {
		return new cv::Mat(kf->errorCovPre);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return new cv::Mat();
    }
}

Mat KalmanFilter_GetGain(KalmanFilter kf) {
    try {
		return new cv::Mat(kf->gain);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return new cv::Mat();
    }
}

Mat KalmanFilter_GetErrorCovPost(KalmanFilter kf) {
    try {
		return new cv::Mat(kf->errorCovPost);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return new cv::Mat();
    }
}

Mat KalmanFilter_GetTemp1(KalmanFilter kf) {
    try {
		return new cv::Mat(kf->temp1);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return new cv::Mat();
    }
}

Mat KalmanFilter_GetTemp2(KalmanFilter kf) {
    try {
		return new cv::Mat(kf->temp2);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return new cv::Mat();
    }
}

Mat KalmanFilter_GetTemp3(KalmanFilter kf) {
    try {
		return new cv::Mat(kf->temp3);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return new cv::Mat();
    }
}

Mat KalmanFilter_GetTemp4(KalmanFilter kf) {
    try {
		return new cv::Mat(kf->temp4);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return new cv::Mat();
    }
}

Mat KalmanFilter_GetTemp5(KalmanFilter kf) {
    try {
		return new cv::Mat(kf->temp5);
    } catch(const cv::Exception& e){
        setExceptionInfo(e.code, e.what());
		return new cv::Mat();
    }
}

void KalmanFilter_SetStatePre(KalmanFilter kf, Mat statePre) {
  	kf->statePre = *statePre;
}

void KalmanFilter_SetStatePost(KalmanFilter kf, Mat statePost) {
  	kf->statePost = *statePost;
}

void KalmanFilter_SetTransitionMatrix(KalmanFilter kf, Mat transitionMatrix) {
  	kf->transitionMatrix = *transitionMatrix;
}

void KalmanFilter_SetControlMatrix(KalmanFilter kf, Mat controlMatrix) {
  	kf->controlMatrix = *controlMatrix;
}

void KalmanFilter_SetMeasurementMatrix(KalmanFilter kf, Mat measurementMatrix) {
  	kf->measurementMatrix = *measurementMatrix;
}

void KalmanFilter_SetProcessNoiseCov(KalmanFilter kf, Mat processNoiseCov) {
  	kf->processNoiseCov = *processNoiseCov;
}

void KalmanFilter_SetMeasurementNoiseCov(KalmanFilter kf, Mat measurementNoiseCov) {
  	kf->measurementNoiseCov = *measurementNoiseCov;
}

void KalmanFilter_SetErrorCovPre(KalmanFilter kf, Mat errorCovPre) {
  	kf->errorCovPre = *errorCovPre;
}

void KalmanFilter_SetGain(KalmanFilter kf, Mat gain) {
  	kf->gain = *gain;
}

void KalmanFilter_SetErrorCovPost(KalmanFilter kf, Mat errorCovPost) {
  	kf->errorCovPost = *errorCovPost;
}
