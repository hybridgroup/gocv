package gocv

func (c CalibFlag) String() string {
	switch c {
	case CalibUseIntrinsicGuess:
		return "calib-use-intrinsec-guess"
	case CalibRecomputeExtrinsic:
		return "calib-recompute-extrinsic"
	case CalibCheckCond:
		return "calib-check-cond"
	case CalibFixSkew:
		return "calib-fix-skew"
	case CalibFixK1:
		return "calib-fix-k1"
	case CalibFixK2:
		return "calib-fix-k2"
	case CalibFixK3:
		return "calib-fix-k3"
	case CalibFixK4:
		return "calib-fix-k4"
	case CalibFixIntrinsic:
		return "calib-fix-intrinsic"
	case CalibFixPrincipalPoint:
		return "calib-fix-principal-point"
	}
	return ""
}
