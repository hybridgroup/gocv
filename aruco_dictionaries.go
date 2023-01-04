package gocv

/*
#include <stdlib.h>
#include "aruco.h"
#include "core.h"
*/
import "C"

type ArucoDictionaryCode int

const (
	ArucoDict4x4_50         ArucoDictionaryCode = iota
	ArucoDict4x4_100        ArucoDictionaryCode = iota
	ArucoDict4x4_250        ArucoDictionaryCode = iota
	ArucoDict4x4_1000       ArucoDictionaryCode = iota
	ArucoDict5x5_50         ArucoDictionaryCode = iota
	ArucoDict5x5_100        ArucoDictionaryCode = iota
	ArucoDict5x5_250        ArucoDictionaryCode = iota
	ArucoDict5x5_1000       ArucoDictionaryCode = iota
	ArucoDict6x6_50         ArucoDictionaryCode = iota
	ArucoDict6x6_100        ArucoDictionaryCode = iota
	ArucoDict6x6_250        ArucoDictionaryCode = iota
	ArucoDict6x6_1000       ArucoDictionaryCode = iota
	ArucoDict7x7_50         ArucoDictionaryCode = iota
	ArucoDict7x7_100        ArucoDictionaryCode = iota
	ArucoDict7x7_250        ArucoDictionaryCode = iota
	ArucoDict7x7_1000       ArucoDictionaryCode = iota
	ArucoDictArucoOriginal  ArucoDictionaryCode = iota
	ArucoDictAprilTag_16h5  ArucoDictionaryCode = iota ///< 4x4 bits, minimum hamming distance between any two codes = 5, 30 codes
	ArucoDictAprilTag_25h9  ArucoDictionaryCode = iota ///< 5x5 bits, minimum hamming distance between any two codes = 9, 35 codes
	ArucoDictAprilTag_36h10 ArucoDictionaryCode = iota ///< 6x6 bits, minimum hamming distance between any two codes = 10, 2320 codes
	ArucoDictAprilTag_36h11 ArucoDictionaryCode = iota ///< 6x6 bits, minimum hamming distance between any two codes = 11, 587 codes
)

type ArucoDictionary struct {
	p C.ArucoDictionary
}

func GetPredefinedDictionary(dictionaryId ArucoDictionaryCode) ArucoDictionary {
	var p C.ArucoDictionary = C.getPredefinedDictionary(C.int(dictionaryId))
	return ArucoDictionary{p: p}
}
