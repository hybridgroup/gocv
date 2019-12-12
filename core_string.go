package gocv

func (c MatType) String() string {
	switch c {
	case MatTypeCV8U:
		return "CV8U"
	case MatTypeCV8UC2:
		return "CV8UC2"
	case MatTypeCV8UC3:
		return "CV8UC3"
	case MatTypeCV8UC4:
		return "CV8UC4"
	case MatTypeCV16U:
		return "CV16U"
	case MatTypeCV16UC2:
		return "CV16UC2"
	case MatTypeCV16UC3:
		return "CV16UC3"
	case MatTypeCV16UC4:
		return "CV16UC4"
	case MatTypeCV16S:
		return "CV16S"
	case MatTypeCV16SC2:
		return "CV16SC2"
	case MatTypeCV16SC3:
		return "CV16SC3"
	case MatTypeCV16SC4:
		return "CV16SC4"
	case MatTypeCV32S:
		return "CV32S"
	case MatTypeCV32SC2:
		return "CV32SC2"
	case MatTypeCV32SC3:
		return "CV32SC3"
	case MatTypeCV32SC4:
		return "CV32SC4"
	case MatTypeCV32F:
		return "CV32F"
	case MatTypeCV32FC2:
		return "CV32FC2"
	case MatTypeCV32FC3:
		return "CV32FC3"
	case MatTypeCV32FC4:
		return "CV32FC4"
	case MatTypeCV64F:
		return "CV64F"
	case MatTypeCV64FC2:
		return "CV64FC2"
	case MatTypeCV64FC3:
		return "CV64FC3"
	case MatTypeCV64FC4:
		return "CV64FC4"
	}
	return ""
}

func (c CompareType) String() string {
	switch c {
	case CompareEQ:
		return "eq"
	case CompareGT:
		return "gt"
	case CompareGE:
		return "ge"
	case CompareLT:
		return "lt"
	case CompareLE:
		return "le"
	case CompareNE:
		return "ne"
	}
	return ""
}

func (c CovarFlags) String() string {
	switch c {
	case CovarScrambled:
		return "covar-scrambled"
	case CovarNormal:
		return "covar-normal"
	case CovarUseAvg:
		return "covar-use-avg"
	case CovarScale:
		return "covar-scale"
	case CovarRows:
		return "covar-rows"
	case CovarCols:
		return "covar-cols"
	}
	return ""
}

func (c DftFlags) String() string {
	switch c {
	case DftForward:
		return "dft-forward"
	case DftInverse:
		return "dft-inverse"
	case DftScale:
		return "dft-scale"
	case DftRows:
		return "dft-rows"
	case DftComplexOutput:
		return "dft-complex-output"
	case DftRealOutput:
		return "dft-real-output"
	case DftComplexInput:
		return "dft-complex-input"
	}
	return ""
}

func (c RotateFlag) String() string {
	switch c {
	case Rotate90Clockwise:
		return "rotate-90-clockwise"
	case Rotate180Clockwise:
		return "rotate-180-clockwise"
	case Rotate90CounterClockwise:
		return "rotate-90-counter-clockwise"
	}
	return ""
}

func (c KMeansFlags) String() string {
	switch c {
	case KMeansRandomCenters:
		return "kmeans-random-centers"
	case KMeansPPCenters:
		return "kmeans-pp-centers"
	case KMeansUseInitialLabels:
		return "kmeans-use-initial-labels"
	}
	return ""
}

func (c NormType) String() string {
	switch c {
	case NormInf:
		return "norm-inf"
	case NormL1:
		return "norm-l1"
	case NormL2:
		return "norm-l2"
	case NormL2Sqr:
		return "norm-l2-sqr"
	case NormHamming:
		return "norm-hamming"
	case NormHamming2:
		return "norm-hamming2"
	case NormRelative:
		return "norm-relative"
	case NormMinMax:
		return "norm-minmax"
	}
	return ""
}

func (c TermCriteriaType) String() string {
	switch c {
	case Count:
		return "count"
	case EPS:
		return "eps"
	}
	return ""
}

func (c SolveDecompositionFlags) String() string {
	switch c {
	case SolveDecompositionLu:
		return "solve-decomposition-lu"
	case SolveDecompositionSvd:
		return "solve-decomposition-svd"
	case SolveDecompositionEing:
		return "solve-decomposition-eing"
	case SolveDecompositionCholesky:
		return "solve-decomposition-cholesky"
	case SolveDecompositionQr:
		return "solve-decomposition-qr"
	case SolveDecompositionNormal:
		return "solve-decomposition-normal"
	}
	return ""
}

func (c ReduceTypes) String() string {
	switch c {
	case ReduceSum:
		return "reduce-sum"
	case ReduceAvg:
		return "reduce-avg"
	case ReduceMax:
		return "reduce-max"
	case ReduceMin:
		return "reduce-min"
	}
	return ""
}

func (c SortFlags) String() string {
	switch c {
	case SortEveryRow:
		return "sort-every-row"
	case SortEveryColumn:
		return "sort-every-column"
	case SortDescending:
		return "sort-descending"
	}
	return ""
}
