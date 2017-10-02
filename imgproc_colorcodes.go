package opencv3

// ColorConversionCode is a color conversion code used on Mat.
//
// For further details, please see:
// http://docs.opencv.org/3.3.0/d7/d1b/group__imgproc__misc.html#ga4e0972be5de079fed4e3a10e24ef5ef0
//
type ColorConversionCode int

const (
	// ColorBGRToBGRA add alpha channel to RGB or BGR image
	ColorBGRToBGRA ColorConversionCode = 0

	// ColorBGRAToBGR removes alpha channel from RGB or BGR image
	ColorBGRAToBGR = 1

	// convert between RGB and BGR color spaces (with or without alpha channel)
	ColorBGRToRGBA = 2

	ColorRGBAToBGR = 3

	ColorBGRToRGB = 4

	ColorBGRAToRGBA = 5

	// convert between RGB/BGR and grayscale
	ColorBGRToGray  = 6
	ColorRGBToGray  = 7
	ColorGrayToBGR  = 8
	ColorGrayToBGRA = 9
	ColorBGRAToGray = 10
	ColorRGBAToGray = 11

	// convert between RGB/BGR and BGR565 (16-bit images)
	ColorBGRToBGR565  = 12
	ColorRGBToBGR565  = 13
	ColorBGR565ToBGR  = 14
	ColorBGR565ToRGB  = 15
	ColorBGRAToBGR565 = 16
	ColorRGBAToBGR565 = 17
	ColorBGR565ToBGRA = 18
	ColorBGR565ToRGBA = 19

	// convert between grayscale to BGR565 (16-bit images)
	ColorGrayToBGR565 = 20
	ColorBGR565ToGray = 21

	// convert between RGB/BGR and BGR555 (16-bit images)
	ColorBGRToBGR555  = 22
	ColorRGBToBGR555  = 23
	ColorBGR555ToBGR  = 24
	ColorBGR555ToRGB  = 25
	ColorBGRAToBGR555 = 26
	ColorRGBAToBGR555 = 27
	ColorBGR555ToBGRA = 28
	ColorBGR555ToRGBA = 29

	// convert between grayscale and BGR555 (16-bit images)
	ColorGrayToBGR555 = 30
	ColorBGR555ToGRAY = 31

	// convert RGB/BGR to CIE XYZ
	ColorBGRToXYZ = 32
	ColorRGBToXYZ = 33
	ColorXYZToBGR = 34
	ColorXYZToRGB = 35

	// convert RGB/BGR to luma-chroma (aka YCC)
	ColorBGRToYCrCb = 36
	ColorRGBToYCrCb = 37
	ColorYCrCbToBGR = 38
	ColorYCrCbToRGB = 39

	// convert RGB/BGR to HSV (hue saturation value)
	ColorBGRToHSV = 40
	ColorRGBToHSV = 41

	// convert RGB/BGR to CIE Lab
	ColorBGRToLab = 44
	ColorRGBToLab = 45

	// convert RGB/BGR to CIE Luv
	ColorBGRToLuv = 50
	ColorRGBToLuv = 51
	// convert RGB/BGR to HLS (hue lightness saturation)
	ColorBGRToHLS = 52
	ColorRGBToHLS = 53

	// backward conversions to RGB/BGR
	ColorHSVToBGR = 54
	ColorHSVToRGB = 55

	ColorLabToBGR = 56
	ColorLabToRGB = 57
	ColorLuvToBGR = 58
	ColorLuvToRGB = 59
	ColorHLSToBGR = 60
	ColorHLSToRGB = 61

	ColorBGRToHSV_FULL = 66
	ColorRGBToHSV_FULL = 67
	ColorBGRToHLS_FULL = 68
	ColorRGBToHLS_FULL = 69

	ColorHSVToBGR_FULL = 70
	ColorHSVToRGB_FULL = 71
	ColorHLSToBGR_FULL = 72
	ColorHLSToRGB_FULL = 73

	ColorLBGRToLab = 74
	ColorLRGBToLab = 75
	ColorLBGRToLuv = 76
	ColorLRGBToLuv = 77

	ColorLabToLBGR = 78
	ColorLabToLRGB = 79
	ColorLuvToLBGR = 80
	ColorLuvToLRGB = 81

	// convert between RGB/BGR and YUV
	ColorBGRToYUV = 82
	ColorRGBToYUV = 83
	ColorYUVToBGR = 84
	ColorYUVToRGB = 85

	// YUV 4:2:0 family to RGB
	ColorYUVToRGBNV12 = 90
	ColorYUVToBGRNV12 = 91
	ColorYUVToRGBNV21 = 92
	ColorYUVToBGRNV21 = 93

	ColorYUVToRGBANV12 = 94
	ColorYUVToBGRANV12 = 95
	ColorYUVToRGBANV21 = 96
	ColorYUVToBGRANV21 = 97

	ColorYUVToRGBYV12 = 98
	ColorYUVToBGRYV12 = 99
	ColorYUVToRGBIYUV = 100
	ColorYUVToBGRIYUV = 101

	ColorYUVToRGBAYV12 = 102
	ColorYUVToBGRAYV12 = 103
	ColorYUVToRGBAIYUV = 104
	ColorYUVToBGRAIYUV = 105

	ColorYUVToGRAY420 = 106

	// YUV 4:2:2 family to RGB
	ColorYUVToRGBUYVY = 107
	ColorYUVToBGRUYVY = 108

	ColorYUVToRGBAUYVY = 111
	ColorYUVToBGRAUYVY = 112

	ColorYUVToRGBYUY2 = 115
	ColorYUVToBGRYUY2 = 116
	ColorYUVToRGBYVYU = 117
	ColorYUVToBGRYVYU = 118

	ColorYUVToRGBAYUY2 = 119
	ColorYUVToBGRAYUY2 = 120
	ColorYUVToRGBAYVYU = 121
	ColorYUVToBGRAYVYU = 122

	ColorYUVToGRAYUYVY = 123
	ColorYUVToGRAYYUY2 = 124

	// alpha premultiplication
	ColorRGBATomRGBA = 125
	ColormRGBAToRGBA = 126

	// RGB to YUV 4:2:0 family
	ColorRGBToYUVI420 = 127
	ColorBGRToYUVI420 = 128

	ColorRGBAToYUVI420 = 129
	ColorBGRAToYUVI420 = 130
	ColorRGBToYUVYV12  = 131
	ColorBGRToYUVYV12  = 132
	ColorRGBAToYUVYV12 = 133
	ColorBGRAToYUVYV12 = 134

	// Demosaicing
	ColorBayerBGToBGR = 46
	ColorBayerGBToBGR = 47
	ColorBayerRGToBGR = 48
	ColorBayerGRToBGR = 49

	ColorBayerBGToGRAY = 86
	ColorBayerGBToGRAY = 87
	ColorBayerRGToGRAY = 88
	ColorBayerGRToGRAY = 89

	// Demosaicing using Variable Number of Gradients
	ColorBayerBGToBGRVNG = 62
	ColorBayerGBToBGRVNG = 63
	ColorBayerRGToBGRVNG = 64
	ColorBayerGRToBGRVNG = 65

	// Edge-Aware Demosaicing
	ColorBayerBGToBGREA = 135
	ColorBayerGBToBGREA = 136
	ColorBayerRGToBGREA = 137
	ColorBayerGRToBGREA = 138

	// Demosaicing with alpha channel
	ColorBayerBGToBGRA = 139
	ColorBayerGBToBGRA = 140
	ColorBayerRGToBGRA = 141
	ColorBayerGRToBGRA = 142

	ColorCOLORCVTMAX = 143
)
