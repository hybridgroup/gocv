package gocv

// ColorConversionCode is a color conversion code used on Mat.
//
// For further details, please see:
// http://docs.opencv.org/master/d7/d1b/group__imgproc__misc.html#ga4e0972be5de079fed4e3a10e24ef5ef0
//
type ColorConversionCode int

const (
	// ColorBGRToBGRA adds alpha channel to BGR image.
	ColorBGRToBGRA ColorConversionCode = 0

	// ColorBGRAToBGR removes alpha channel from BGR image.
	ColorBGRAToBGR ColorConversionCode = 1

	// ColorBGRToRGBA converts from BGR to RGB with alpha channel.
	ColorBGRToRGBA ColorConversionCode = 2

	// ColorRGBAToBGR converts from RGB with alpha to BGR color space.
	ColorRGBAToBGR ColorConversionCode = 3

	// ColorBGRToRGB converts from BGR to RGB without alpha channel.
	ColorBGRToRGB ColorConversionCode = 4

	// ColorBGRAToRGBA converts from BGR with alpha channel
	// to RGB with alpha channel.
	ColorBGRAToRGBA ColorConversionCode = 5

	// ColorBGRToGray converts from BGR to grayscale.
	ColorBGRToGray ColorConversionCode = 6

	// ColorRGBToGray converts from RGB to grayscale.
	ColorRGBToGray ColorConversionCode = 7

	// ColorGrayToBGR converts from grayscale to BGR.
	ColorGrayToBGR ColorConversionCode = 8

	// ColorGrayToBGRA converts from grayscale to BGR with alpha channel.
	ColorGrayToBGRA ColorConversionCode = 9

	// ColorBGRAToGray converts from BGR with alpha channel to grayscale.
	ColorBGRAToGray ColorConversionCode = 10

	// ColorRGBAToGray converts from RGB with alpha channel to grayscale.
	ColorRGBAToGray ColorConversionCode = 11

	// ColorBGRToBGR565 converts from BGR to BGR565 (16-bit images).
	ColorBGRToBGR565 ColorConversionCode = 12

	// ColorRGBToBGR565 converts from RGB to BGR565 (16-bit images).
	ColorRGBToBGR565 ColorConversionCode = 13

	// ColorBGR565ToBGR converts from BGR565 (16-bit images) to BGR.
	ColorBGR565ToBGR ColorConversionCode = 14

	// ColorBGR565ToRGB converts from BGR565 (16-bit images) to RGB.
	ColorBGR565ToRGB ColorConversionCode = 15

	// ColorBGRAToBGR565 converts from BGRA (with alpha channel)
	// to BGR565 (16-bit images).
	ColorBGRAToBGR565 ColorConversionCode = 16

	// ColorRGBAToBGR565 converts from RGBA (with alpha channel)
	// to BGR565 (16-bit images).
	ColorRGBAToBGR565 ColorConversionCode = 17

	// ColorBGR565ToBGRA converts from BGR565 (16-bit images)
	// to BGRA (with alpha channel).
	ColorBGR565ToBGRA ColorConversionCode = 18

	// ColorBGR565ToRGBA converts from BGR565 (16-bit images)
	// to RGBA (with alpha channel).
	ColorBGR565ToRGBA ColorConversionCode = 19

	// ColorGrayToBGR565 converts from grayscale
	// to BGR565 (16-bit images).
	ColorGrayToBGR565 ColorConversionCode = 20

	// ColorBGR565ToGray converts from BGR565 (16-bit images)
	// to grayscale.
	ColorBGR565ToGray ColorConversionCode = 21

	// ColorBGRToBGR555 converts from BGR to BGR555 (16-bit images).
	ColorBGRToBGR555 ColorConversionCode = 22

	// ColorRGBToBGR555 converts from RGB to BGR555 (16-bit images).
	ColorRGBToBGR555 ColorConversionCode = 23

	// ColorBGR555ToBGR converts from BGR555 (16-bit images) to BGR.
	ColorBGR555ToBGR ColorConversionCode = 24

	// ColorBGR555ToRGB converts from BGR555 (16-bit images) to RGB.
	ColorBGR555ToRGB ColorConversionCode = 25

	// ColorBGRAToBGR555 converts from BGRA (with alpha channel)
	// to BGR555 (16-bit images).
	ColorBGRAToBGR555 ColorConversionCode = 26

	// ColorRGBAToBGR555 converts from RGBA (with alpha channel)
	// to BGR555 (16-bit images).
	ColorRGBAToBGR555 ColorConversionCode = 27

	// ColorBGR555ToBGRA converts from BGR555 (16-bit images)
	// to BGRA (with alpha channel).
	ColorBGR555ToBGRA ColorConversionCode = 28

	// ColorBGR555ToRGBA converts from BGR555 (16-bit images)
	// to RGBA (with alpha channel).
	ColorBGR555ToRGBA ColorConversionCode = 29

	// ColorGrayToBGR555 converts from grayscale to BGR555 (16-bit images).
	ColorGrayToBGR555 ColorConversionCode = 30

	// ColorBGR555ToGRAY converts from BGR555 (16-bit images) to grayscale.
	ColorBGR555ToGRAY ColorConversionCode = 31

	// ColorBGRToXYZ converts from BGR to CIE XYZ.
	ColorBGRToXYZ ColorConversionCode = 32

	// ColorRGBToXYZ converts from RGB to CIE XYZ.
	ColorRGBToXYZ ColorConversionCode = 33

	// ColorXYZToBGR converts from CIE XYZ to BGR.
	ColorXYZToBGR ColorConversionCode = 34

	// ColorXYZToRGB converts from CIE XYZ to RGB.
	ColorXYZToRGB ColorConversionCode = 35

	// ColorBGRToYCrCb converts from BGR to luma-chroma (aka YCC).
	ColorBGRToYCrCb ColorConversionCode = 36

	// ColorRGBToYCrCb converts from RGB to luma-chroma (aka YCC).
	ColorRGBToYCrCb ColorConversionCode = 37

	// ColorYCrCbToBGR converts from luma-chroma (aka YCC) to BGR.
	ColorYCrCbToBGR ColorConversionCode = 38

	// ColorYCrCbToRGB converts from luma-chroma (aka YCC) to RGB.
	ColorYCrCbToRGB ColorConversionCode = 39

	// ColorBGRToHSV converts from BGR to HSV (hue saturation value).
	ColorBGRToHSV ColorConversionCode = 40

	// ColorRGBToHSV converts from RGB to HSV (hue saturation value).
	ColorRGBToHSV ColorConversionCode = 41

	// ColorBGRToLab converts from BGR to CIE Lab.
	ColorBGRToLab ColorConversionCode = 44

	// ColorRGBToLab converts from RGB to CIE Lab.
	ColorRGBToLab ColorConversionCode = 45

	// ColorBGRToLuv converts from BGR to CIE Luv.
	ColorBGRToLuv ColorConversionCode = 50

	// ColorRGBToLuv converts from RGB to CIE Luv.
	ColorRGBToLuv ColorConversionCode = 51

	// ColorBGRToHLS converts from BGR to HLS (hue lightness saturation).
	ColorBGRToHLS ColorConversionCode = 52

	// ColorRGBToHLS converts from RGB to HLS (hue lightness saturation).
	ColorRGBToHLS ColorConversionCode = 53

	// ColorHSVToBGR converts from HSV (hue saturation value) to BGR.
	ColorHSVToBGR ColorConversionCode = 54

	// ColorHSVToRGB converts from HSV (hue saturation value) to RGB.
	ColorHSVToRGB ColorConversionCode = 55

	// ColorLabToBGR converts from CIE Lab to BGR.
	ColorLabToBGR ColorConversionCode = 56

	// ColorLabToRGB converts from CIE Lab to RGB.
	ColorLabToRGB ColorConversionCode = 57

	// ColorLuvToBGR converts from CIE Luv to BGR.
	ColorLuvToBGR ColorConversionCode = 58

	// ColorLuvToRGB converts from CIE Luv to RGB.
	ColorLuvToRGB ColorConversionCode = 59

	// ColorHLSToBGR converts from HLS (hue lightness saturation) to BGR.
	ColorHLSToBGR ColorConversionCode = 60

	// ColorHLSToRGB converts from HLS (hue lightness saturation) to RGB.
	ColorHLSToRGB ColorConversionCode = 61

	// ColorBGRToHSVFull converts from BGR to HSV (hue saturation value) full.
	ColorBGRToHSVFull ColorConversionCode = 66

	// ColorRGBToHSVFull converts from RGB to HSV (hue saturation value) full.
	ColorRGBToHSVFull ColorConversionCode = 67

	// ColorBGRToHLSFull converts from BGR to HLS (hue lightness saturation) full.
	ColorBGRToHLSFull ColorConversionCode = 68

	// ColorRGBToHLSFull converts from RGB to HLS (hue lightness saturation) full.
	ColorRGBToHLSFull ColorConversionCode = 69

	// ColorHSVToBGRFull converts from HSV (hue saturation value) to BGR full.
	ColorHSVToBGRFull ColorConversionCode = 70

	// ColorHSVToRGBFull converts from HSV (hue saturation value) to RGB full.
	ColorHSVToRGBFull ColorConversionCode = 71

	// ColorHLSToBGRFull converts from HLS (hue lightness saturation) to BGR full.
	ColorHLSToBGRFull ColorConversionCode = 72

	// ColorHLSToRGBFull converts from HLS (hue lightness saturation) to RGB full.
	ColorHLSToRGBFull ColorConversionCode = 73

	// ColorLBGRToLab converts from LBGR to CIE Lab.
	ColorLBGRToLab ColorConversionCode = 74

	// ColorLRGBToLab converts from LRGB to CIE Lab.
	ColorLRGBToLab ColorConversionCode = 75

	// ColorLBGRToLuv converts from LBGR to CIE Luv.
	ColorLBGRToLuv ColorConversionCode = 76

	// ColorLRGBToLuv converts from LRGB to CIE Luv.
	ColorLRGBToLuv ColorConversionCode = 77

	// ColorLabToLBGR converts from CIE Lab to LBGR.
	ColorLabToLBGR ColorConversionCode = 78

	// ColorLabToLRGB converts from CIE Lab to LRGB.
	ColorLabToLRGB ColorConversionCode = 79

	// ColorLuvToLBGR converts from CIE Luv to LBGR.
	ColorLuvToLBGR ColorConversionCode = 80

	// ColorLuvToLRGB converts from CIE Luv to LRGB.
	ColorLuvToLRGB ColorConversionCode = 81

	// ColorBGRToYUV converts from BGR to YUV.
	ColorBGRToYUV ColorConversionCode = 82

	// ColorRGBToYUV converts from RGB to YUV.
	ColorRGBToYUV ColorConversionCode = 83

	// ColorYUVToBGR converts from YUV to BGR.
	ColorYUVToBGR ColorConversionCode = 84

	// ColorYUVToRGB converts from YUV to RGB.
	ColorYUVToRGB ColorConversionCode = 85

	// ColorYUVToRGBNV12 converts from YUV 4:2:0 to RGB NV12.
	ColorYUVToRGBNV12 ColorConversionCode = 90

	// ColorYUVToBGRNV12 converts from YUV 4:2:0 to BGR NV12.
	ColorYUVToBGRNV12 ColorConversionCode = 91

	// ColorYUVToRGBNV21 converts from YUV 4:2:0 to RGB NV21.
	ColorYUVToRGBNV21 ColorConversionCode = 92

	// ColorYUVToBGRNV21 converts from YUV 4:2:0 to BGR NV21.
	ColorYUVToBGRNV21 ColorConversionCode = 93

	// ColorYUVToRGBANV12 converts from YUV 4:2:0 to RGBA NV12.
	ColorYUVToRGBANV12 ColorConversionCode = 94

	// ColorYUVToBGRANV12 converts from YUV 4:2:0 to BGRA NV12.
	ColorYUVToBGRANV12 ColorConversionCode = 95

	// ColorYUVToRGBANV21 converts from YUV 4:2:0 to RGBA NV21.
	ColorYUVToRGBANV21 ColorConversionCode = 96

	// ColorYUVToBGRANV21 converts from YUV 4:2:0 to BGRA NV21.
	ColorYUVToBGRANV21 ColorConversionCode = 97

	ColorYUVToRGBYV12 ColorConversionCode = 98
	ColorYUVToBGRYV12 ColorConversionCode = 99
	ColorYUVToRGBIYUV ColorConversionCode = 100
	ColorYUVToBGRIYUV ColorConversionCode = 101

	ColorYUVToRGBAYV12 ColorConversionCode = 102
	ColorYUVToBGRAYV12 ColorConversionCode = 103
	ColorYUVToRGBAIYUV ColorConversionCode = 104
	ColorYUVToBGRAIYUV ColorConversionCode = 105

	ColorYUVToGRAY420 ColorConversionCode = 106

	// YUV 4:2:2 family to RGB
	ColorYUVToRGBUYVY ColorConversionCode = 107
	ColorYUVToBGRUYVY ColorConversionCode = 108

	ColorYUVToRGBAUYVY ColorConversionCode = 111
	ColorYUVToBGRAUYVY ColorConversionCode = 112

	ColorYUVToRGBYUY2 ColorConversionCode = 115
	ColorYUVToBGRYUY2 ColorConversionCode = 116
	ColorYUVToRGBYVYU ColorConversionCode = 117
	ColorYUVToBGRYVYU ColorConversionCode = 118

	ColorYUVToRGBAYUY2 ColorConversionCode = 119
	ColorYUVToBGRAYUY2 ColorConversionCode = 120
	ColorYUVToRGBAYVYU ColorConversionCode = 121
	ColorYUVToBGRAYVYU ColorConversionCode = 122

	ColorYUVToGRAYUYVY ColorConversionCode = 123
	ColorYUVToGRAYYUY2 ColorConversionCode = 124

	// alpha premultiplication
	ColorRGBATomRGBA ColorConversionCode = 125
	ColormRGBAToRGBA ColorConversionCode = 126

	// RGB to YUV 4:2:0 family
	ColorRGBToYUVI420 ColorConversionCode = 127
	ColorBGRToYUVI420 ColorConversionCode = 128

	ColorRGBAToYUVI420 ColorConversionCode = 129
	ColorBGRAToYUVI420 ColorConversionCode = 130
	ColorRGBToYUVYV12  ColorConversionCode = 131
	ColorBGRToYUVYV12  ColorConversionCode = 132
	ColorRGBAToYUVYV12 ColorConversionCode = 133
	ColorBGRAToYUVYV12 ColorConversionCode = 134

	// Demosaicing
	ColorBayerBGToBGR ColorConversionCode = 46
	ColorBayerGBToBGR ColorConversionCode = 47
	ColorBayerRGToBGR ColorConversionCode = 48
	ColorBayerGRToBGR ColorConversionCode = 49

	ColorBayerBGToGRAY ColorConversionCode = 86
	ColorBayerGBToGRAY ColorConversionCode = 87
	ColorBayerRGToGRAY ColorConversionCode = 88
	ColorBayerGRToGRAY ColorConversionCode = 89

	// Demosaicing using Variable Number of Gradients
	ColorBayerBGToBGRVNG ColorConversionCode = 62
	ColorBayerGBToBGRVNG ColorConversionCode = 63
	ColorBayerRGToBGRVNG ColorConversionCode = 64
	ColorBayerGRToBGRVNG ColorConversionCode = 65

	// Edge-Aware Demosaicing
	ColorBayerBGToBGREA ColorConversionCode = 135
	ColorBayerGBToBGREA ColorConversionCode = 136
	ColorBayerRGToBGREA ColorConversionCode = 137
	ColorBayerGRToBGREA ColorConversionCode = 138

	// Demosaicing with alpha channel
	ColorBayerBGToBGRA ColorConversionCode = 139
	ColorBayerGBToBGRA ColorConversionCode = 140
	ColorBayerRGToBGRA ColorConversionCode = 141
	ColorBayerGRToBGRA ColorConversionCode = 142

	ColorCOLORCVTMAX ColorConversionCode = 143
)
