package gocv

func (c HistCompMethod) String() string {
	switch c {
	case HistCmpCorrel:
		return "hist-cmp-correl"
	case HistCmpChiSqr:
		return "hist-cmp-chi-sqr"
	case HistCmpIntersect:
		return "hist-cmp-intersect"
	case HistCmpBhattacharya:
		return "hist-cmp-bhattacharya"
	case HistCmpChiSqrAlt:
		return "hist-cmp-chi-sqr-alt"
	case HistCmpKlDiv:
		return "hist-cmp-kl-div"
	}
	return ""
}

func (c DistanceTransformLabelTypes) String() string {
	switch c {
	case DistanceLabelCComp:
		return "distance-label-ccomp"
	}
	return ""
}

func (c DistanceTransformMasks) String() string {
	switch c {
	case DistanceMask3:
		return "distance-mask3"
	}
	return ""
}

func (c RetrievalMode) String() string {
	switch c {
	case RetrievalExternal:
		return "retrieval-external"
	case RetrievalList:
		return "retrieval-list"
	case RetrievalCComp:
		return "retrieval-ccomp"
	case RetrievalTree:
		return "retrieval-tree"
	case RetrievalFloodfill:
		return "retrieval-floodfill"
	}
	return ""
}

func (c ContourApproximationMode) String() string {
	switch c {
	case ChainApproxNone:
		return "chain-approx-none"
	case ChainApproxSimple:
		return "chain-approx-simple"
	case ChainApproxTC89L1:
		return "chain-approx-tc89l1"
	case ChainApproxTC89KCOS:
		return "chain-approx-tc89kcos"
	}
	return ""
}

func (c ConnectedComponentsAlgorithmType) String() string {
	switch c {
	case CCL_WU:
		return "ccl-wu"
	case CCL_DEFAULT:
		return "ccl-default"
	case CCL_GRANA:
		return "ccl-grana"
	}
	return ""
}

func (c ConnectedComponentsTypes) String() string {
	switch c {
	case CC_STAT_LEFT:
		return "cc-stat-left"
	case CC_STAT_TOP:
		return "cc-stat-top"
	case CC_STAT_WIDTH:
		return "cc-stat-width"
	case CC_STAT_AREA:
		return "cc-stat-area"
	case CC_STAT_MAX:
		return "cc-stat-max"
	case CC_STAT_HEIGHT:
		return "cc-stat-height"
	}
	return ""
}

func (c TemplateMatchMode) String() string {
	switch c {
	case TmSqdiff:
		return "tm-sq-diff"
	case TmSqdiffNormed:
		return "tm-sq-diff-normed"
	case TmCcorr:
		return "tm-ccorr"
	case TmCcorrNormed:
		return "tm-ccorr-normed"
	case TmCcoeff:
		return "tm-ccoeff"
	case TmCcoeffNormed:
		return "tm-ccoeff-normed"
	}
	return ""
}

func (c MorphShape) String() string {
	switch c {
	case MorphRect:
		return "morph-rect"
	case MorphCross:
		return "morph-cross"
	case MorphEllipse:
		return "morph-ellispe"
	}
	return ""
}

func (c MorphType) String() string {
	switch c {
	case MorphErode:
		return "morph-erode"
	case MorphDilate:
		return "morph-dilate"
	case MorphOpen:
		return "morph-open"
	case MorphClose:
		return "morph-close"
	case MorphGradient:
		return "morph-gradient"
	case MorphTophat:
		return "morph-tophat"
	case MorphBlackhat:
		return "morph-blackhat"
	case MorphHitmiss:
		return "morph-hitmiss"
	}
	return ""
}

func (c BorderType) String() string {
	switch c {
	case BorderConstant:
		return "border-constant"
	case BorderReplicate:
		return "border-replicate"
	case BorderReflect:
		return "border-reflect"
	case BorderWrap:
		return "border-wrap"
	case BorderTransparent:
		return "border-transparent"
	case BorderDefault:
		return "border-default"
	case BorderIsolated:
		return "border-isolated"
	}
	return ""
}

func (c GrabCutMode) String() string {
	switch c {
	case GCInitWithRect:
		return "gc-init-with-rect"
	case GCInitWithMask:
		return "gc-init-with-mask"
	case GCEval:
		return "gc-eval"
	case GCEvalFreezeModel:
		return "gc-eval-freeze-model"
	}
	return ""
}

func (c HoughMode) String() string {
	switch c {
	case HoughStandard:
		return "hough-standard"
	case HoughProbabilistic:
		return "hough-probabilistic"
	case HoughMultiScale:
		return "hough-multi-scale"
	case HoughGradient:
		return "hough-gradient"
	}
	return ""
}

func (c ThresholdType) String() string {
	switch c {
	case ThresholdBinary:
		return "threshold-binary"
	case ThresholdBinaryInv:
		return "threshold-binary-inv"
	case ThresholdTrunc:
		return "threshold-trunc"
	case ThresholdToZero:
		return "threshold-to-zero"
	case ThresholdToZeroInv:
		return "threshold-to-zero-inv"
	case ThresholdMask:
		return "threshold-mask"
	case ThresholdOtsu:
		return "threshold-otsu"
	case ThresholdTriangle:
		return "threshold-triangle"
	}
	return ""
}

func (c AdaptiveThresholdType) String() string {
	switch c {
	case AdaptiveThresholdMean:
		return "adaptative-threshold-mean"
	case AdaptiveThresholdGaussian:
		return "adaptative-threshold-gaussian"
	}
	return ""
}

func (c HersheyFont) String() string {
	switch c {
	case FontHersheySimplex:
		return "font-hershey-simplex"
	case FontHersheyPlain:
		return "font-hershey-plain"
	case FontHersheyDuplex:
		return "font-hershey-duplex"
	case FontHersheyComplex:
		return "font-hershey-complex"
	case FontHersheyTriplex:
		return "font-hershey-triplex"
	case FontHersheyComplexSmall:
		return "font-hershey-complex-small"
	case FontHersheyScriptSimplex:
		return "font-hershey-script-simplex"
	case FontHersheyScriptComplex:
		return "font-hershey-scipt-complex"
	case FontItalic:
		return "font-italic"
	}
	return ""
}

func (c LineType) String() string {
	switch c {
	case Filled:
		return "filled"
	case Line4:
		return "line4"
	case Line8:
		return "line8"
	case LineAA:
		return "line-aa"
	}
	return ""
}

func (c InterpolationFlags) String() string {
	switch c {
	case InterpolationNearestNeighbor:
		return "interpolation-nearest-neighbor"
	case InterpolationLinear:
		return "interpolation-linear"
	case InterpolationCubic:
		return "interpolation-cubic"
	case InterpolationArea:
		return "interpolation-area"
	case InterpolationLanczos4:
		return "interpolation-lanczos4"
	case InterpolationMax:
		return "interpolation-max"
	}
	return ""
}

func (c ColormapTypes) String() string {
	switch c {
	case ColormapAutumn:
		return "colormap-autumn"
	case ColormapBone:
		return "colormap-bone"
	case ColormapJet:
		return "colormap-jet"
	case ColormapWinter:
		return "colormap-winter"
	case ColormapRainbow:
		return "colormap-rainbow"
	case ColormapOcean:
		return "colormap-ocean"
	case ColormapSummer:
		return "colormap-summer"
	case ColormapSpring:
		return "colormap-spring"
	case ColormapCool:
		return "colormap-cool"
	case ColormapHsv:
		return "colormap-hsv"
	case ColormapPink:
		return "colormap-pink"
	case ColormapParula:
		return "colormap-parula"
	}
	return ""
}

func (c DistanceTypes) String() string {
	switch c {
	case DistUser:
		return "dist-user"
	case DistL1:
		return "dist-l1"
	case DistL2:
		return "dist-l2"
	case DistL12:
		return "dist-l12"
	case DistFair:
		return "dist-fair"
	case DistWelsch:
		return "dist-welsch"
	case DistHuber:
		return "dist-huber"
	}
	return ""
}
