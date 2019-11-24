package gocv

/*
#include <stdlib.h>
#include "features2d.h"
*/
import "C"

func (c FastFeatureDetectorType) String() string {
	switch c {
	case FastFeatureDetectorType58:
		return "fast-feature-detector-type-58"
	case FastFeatureDetectorType712:
		return "fast-feature-detector-type-712"
	case FastFeatureDetectorType916:
		return "fast-feature-detector-type-916"
	}
	return ""
}

func (c DrawMatchesFlag) String() string {
	switch c {
	case DrawDefault:
		return "draw-default"
	case DrawOverOutImg:
		return "draw-over-out-imt"
	case NotDrawSinglePoints:
		return "draw-single-points"
	case DrawRichKeyPoints:
		return "draw-rich-key-points"
	}
	return ""
}
