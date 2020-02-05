package cuda

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

func (c BorderType) String() string {
	switch c {
	case BorderConstant:
		return "border-constant"
	case BorderReplicate:
		return "border-replicate"
	case BorderWrap:
		return "border-wrap"
	case BorderReflect101:
		return "border-reflect101"
	case BorderTransparent:
		return "border-transparent"
	case BorderIsolated:
		return "border-isolated"
	}
	return ""
}
