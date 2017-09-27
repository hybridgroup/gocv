package opencv3

// imgproc color conversion codes
const (
	// ColorBGR2BGRA add alpha channel to RGB or BGR image
	ColorBGR2BGRA = 0
	ColorRGB2RGBA = ColorBGR2BGRA

	ColorBGRA2BGR = 1 //!< remove alpha channel from RGB or BGR image
	ColorRGBA2RGB = ColorBGRA2BGR

	ColorBGR2RGBA = 2 //!< convert between RGB and BGR color spaces (with or without alpha channel)
	ColorRGB2BGRA = ColorBGR2RGBA

	ColorRGBA2BGR = 3
	ColorBGRA2RGB = ColorRGBA2BGR

	ColorBGR2RGB = 4
	ColorRGB2BGR = ColorBGR2RGB

	ColorBGRA2RGBA = 5
	ColorRGBA2BGRA = ColorBGRA2RGBA

	ColorBGR2GRAY  = 6 //!< convert between RGB/BGR and grayscale, @ref Colorconvert_rgb_gray "color conversions"
	ColorRGB2GRAY  = 7
	ColorGRAY2BGR  = 8
	ColorGRAY2RGB  = ColorGRAY2BGR
	ColorGRAY2BGRA = 9
	ColorGRAY2RGBA = ColorGRAY2BGRA
	ColorBGRA2GRAY = 10
	ColorRGBA2GRAY = 11

	ColorBGR2BGR565  = 12 //!< convert between RGB/BGR and BGR565 (16-bit images)
	ColorRGB2BGR565  = 13
	ColorBGR5652BGR  = 14
	ColorBGR5652RGB  = 15
	ColorBGRA2BGR565 = 16
	ColorRGBA2BGR565 = 17
	ColorBGR5652BGRA = 18
	ColorBGR5652RGBA = 19

	ColorGRAY2BGR565 = 20 //!< convert between grayscale to BGR565 (16-bit images)
	ColorBGR5652GRAY = 21

	ColorBGR2BGR555  = 22 //!< convert between RGB/BGR and BGR555 (16-bit images)
	ColorRGB2BGR555  = 23
	ColorBGR5552BGR  = 24
	ColorBGR5552RGB  = 25
	ColorBGRA2BGR555 = 26
	ColorRGBA2BGR555 = 27
	ColorBGR5552BGRA = 28
	ColorBGR5552RGBA = 29

	ColorGRAY2BGR555 = 30 //!< convert between grayscale and BGR555 (16-bit images)
	ColorBGR5552GRAY = 31

	ColorBGR2XYZ = 32 //!< convert RGB/BGR to CIE XYZ, @ref Colorconvert_rgb_xyz "color conversions"
	ColorRGB2XYZ = 33
	ColorXYZ2BGR = 34
	ColorXYZ2RGB = 35

	ColorBGR2YCrCb = 36 //!< convert RGB/BGR to luma-chroma (aka YCC), @ref Colorconvert_rgb_ycrcb "color conversions"
	ColorRGB2YCrCb = 37
	ColorYCrCb2BGR = 38
	ColorYCrCb2RGB = 39

	ColorBGR2HSV = 40 //!< convert RGB/BGR to HSV (hue saturation value), @ref Colorconvert_rgb_hsv "color conversions"
	ColorRGB2HSV = 41

	ColorBGR2Lab = 44 //!< convert RGB/BGR to CIE Lab, @ref Colorconvert_rgb_lab "color conversions"
	ColorRGB2Lab = 45

	ColorBGR2Luv = 50 //!< convert RGB/BGR to CIE Luv, @ref Colorconvert_rgb_luv "color conversions"
	ColorRGB2Luv = 51
	ColorBGR2HLS = 52 //!< convert RGB/BGR to HLS (hue lightness saturation), @ref Colorconvert_rgb_hls "color conversions"
	ColorRGB2HLS = 53

	ColorHSV2BGR = 54 //!< backward conversions to RGB/BGR
	ColorHSV2RGB = 55

	ColorLab2BGR = 56
	ColorLab2RGB = 57
	ColorLuv2BGR = 58
	ColorLuv2RGB = 59
	ColorHLS2BGR = 60
	ColorHLS2RGB = 61

	ColorBGR2HSV_FULL = 66 //!<
	ColorRGB2HSV_FULL = 67
	ColorBGR2HLS_FULL = 68
	ColorRGB2HLS_FULL = 69

	ColorHSV2BGR_FULL = 70
	ColorHSV2RGB_FULL = 71
	ColorHLS2BGR_FULL = 72
	ColorHLS2RGB_FULL = 73

	ColorLBGR2Lab = 74
	ColorLRGB2Lab = 75
	ColorLBGR2Luv = 76
	ColorLRGB2Luv = 77

	ColorLab2LBGR = 78
	ColorLab2LRGB = 79
	ColorLuv2LBGR = 80
	ColorLuv2LRGB = 81

	ColorBGR2YUV = 82 //!< convert between RGB/BGR and YUV
	ColorRGB2YUV = 83
	ColorYUV2BGR = 84
	ColorYUV2RGB = 85

	//! YUV 4:2:0 family to RGB
	ColorYUV2RGB_NV12 = 90
	ColorYUV2BGR_NV12 = 91
	ColorYUV2RGB_NV21 = 92
	ColorYUV2BGR_NV21 = 93
	ColorYUV420sp2RGB = ColorYUV2RGB_NV21
	ColorYUV420sp2BGR = ColorYUV2BGR_NV21

	ColorYUV2RGBA_NV12 = 94
	ColorYUV2BGRA_NV12 = 95
	ColorYUV2RGBA_NV21 = 96
	ColorYUV2BGRA_NV21 = 97
	ColorYUV420sp2RGBA = ColorYUV2RGBA_NV21
	ColorYUV420sp2BGRA = ColorYUV2BGRA_NV21

	ColorYUV2RGB_YV12 = 98
	ColorYUV2BGR_YV12 = 99
	ColorYUV2RGB_IYUV = 100
	ColorYUV2BGR_IYUV = 101
	ColorYUV2RGB_I420 = ColorYUV2RGB_IYUV
	ColorYUV2BGR_I420 = ColorYUV2BGR_IYUV
	ColorYUV420p2RGB  = ColorYUV2RGB_YV12
	ColorYUV420p2BGR  = ColorYUV2BGR_YV12

	ColorYUV2RGBA_YV12 = 102
	ColorYUV2BGRA_YV12 = 103
	ColorYUV2RGBA_IYUV = 104
	ColorYUV2BGRA_IYUV = 105
	ColorYUV2RGBA_I420 = ColorYUV2RGBA_IYUV
	ColorYUV2BGRA_I420 = ColorYUV2BGRA_IYUV
	ColorYUV420p2RGBA  = ColorYUV2RGBA_YV12
	ColorYUV420p2BGRA  = ColorYUV2BGRA_YV12

	ColorYUV2GRAY_420  = 106
	ColorYUV2GRAY_NV21 = ColorYUV2GRAY_420
	ColorYUV2GRAY_NV12 = ColorYUV2GRAY_420
	ColorYUV2GRAY_YV12 = ColorYUV2GRAY_420
	ColorYUV2GRAY_IYUV = ColorYUV2GRAY_420
	ColorYUV2GRAY_I420 = ColorYUV2GRAY_420
	ColorYUV420sp2GRAY = ColorYUV2GRAY_420
	ColorYUV420p2GRAY  = ColorYUV2GRAY_420

	//! YUV 4:2:2 family to RGB
	ColorYUV2RGB_UYVY = 107
	ColorYUV2BGR_UYVY = 108
	//ColorYUV2RGB_VYUY = 109,
	//ColorYUV2BGR_VYUY = 110,
	ColorYUV2RGB_Y422 = ColorYUV2RGB_UYVY
	ColorYUV2BGR_Y422 = ColorYUV2BGR_UYVY
	ColorYUV2RGB_UYNV = ColorYUV2RGB_UYVY
	ColorYUV2BGR_UYNV = ColorYUV2BGR_UYVY

	ColorYUV2RGBA_UYVY = 111
	ColorYUV2BGRA_UYVY = 112
	//ColorYUV2RGBA_VYUY = 113,
	//ColorYUV2BGRA_VYUY = 114,
	ColorYUV2RGBA_Y422 = ColorYUV2RGBA_UYVY
	ColorYUV2BGRA_Y422 = ColorYUV2BGRA_UYVY
	ColorYUV2RGBA_UYNV = ColorYUV2RGBA_UYVY
	ColorYUV2BGRA_UYNV = ColorYUV2BGRA_UYVY

	ColorYUV2RGB_YUY2 = 115
	ColorYUV2BGR_YUY2 = 116
	ColorYUV2RGB_YVYU = 117
	ColorYUV2BGR_YVYU = 118
	ColorYUV2RGB_YUYV = ColorYUV2RGB_YUY2
	ColorYUV2BGR_YUYV = ColorYUV2BGR_YUY2
	ColorYUV2RGB_YUNV = ColorYUV2RGB_YUY2
	ColorYUV2BGR_YUNV = ColorYUV2BGR_YUY2

	ColorYUV2RGBA_YUY2 = 119
	ColorYUV2BGRA_YUY2 = 120
	ColorYUV2RGBA_YVYU = 121
	ColorYUV2BGRA_YVYU = 122
	ColorYUV2RGBA_YUYV = ColorYUV2RGBA_YUY2
	ColorYUV2BGRA_YUYV = ColorYUV2BGRA_YUY2
	ColorYUV2RGBA_YUNV = ColorYUV2RGBA_YUY2
	ColorYUV2BGRA_YUNV = ColorYUV2BGRA_YUY2

	ColorYUV2GRAY_UYVY = 123
	ColorYUV2GRAY_YUY2 = 124
	//CV_YUV2GRAY_VYUY    = CV_YUV2GRAY_UYVY,
	ColorYUV2GRAY_Y422 = ColorYUV2GRAY_UYVY
	ColorYUV2GRAY_UYNV = ColorYUV2GRAY_UYVY
	ColorYUV2GRAY_YVYU = ColorYUV2GRAY_YUY2
	ColorYUV2GRAY_YUYV = ColorYUV2GRAY_YUY2
	ColorYUV2GRAY_YUNV = ColorYUV2GRAY_YUY2

	//! alpha premultiplication
	ColorRGBA2mRGBA = 125
	ColormRGBA2RGBA = 126

	//! RGB to YUV 4:2:0 family
	ColorRGB2YUV_I420 = 127
	ColorBGR2YUV_I420 = 128
	ColorRGB2YUV_IYUV = ColorRGB2YUV_I420
	ColorBGR2YUV_IYUV = ColorBGR2YUV_I420

	ColorRGBA2YUV_I420 = 129
	ColorBGRA2YUV_I420 = 130
	ColorRGBA2YUV_IYUV = ColorRGBA2YUV_I420
	ColorBGRA2YUV_IYUV = ColorBGRA2YUV_I420
	ColorRGB2YUV_YV12  = 131
	ColorBGR2YUV_YV12  = 132
	ColorRGBA2YUV_YV12 = 133
	ColorBGRA2YUV_YV12 = 134

	//! Demosaicing
	ColorBayerBG2BGR = 46
	ColorBayerGB2BGR = 47
	ColorBayerRG2BGR = 48
	ColorBayerGR2BGR = 49

	ColorBayerBG2RGB = ColorBayerRG2BGR
	ColorBayerGB2RGB = ColorBayerGR2BGR
	ColorBayerRG2RGB = ColorBayerBG2BGR
	ColorBayerGR2RGB = ColorBayerGB2BGR

	ColorBayerBG2GRAY = 86
	ColorBayerGB2GRAY = 87
	ColorBayerRG2GRAY = 88
	ColorBayerGR2GRAY = 89

	//! Demosaicing using Variable Number of Gradients
	ColorBayerBG2BGR_VNG = 62
	ColorBayerGB2BGR_VNG = 63
	ColorBayerRG2BGR_VNG = 64
	ColorBayerGR2BGR_VNG = 65

	ColorBayerBG2RGB_VNG = ColorBayerRG2BGR_VNG
	ColorBayerGB2RGB_VNG = ColorBayerGR2BGR_VNG
	ColorBayerRG2RGB_VNG = ColorBayerBG2BGR_VNG
	ColorBayerGR2RGB_VNG = ColorBayerGB2BGR_VNG

	//! Edge-Aware Demosaicing
	ColorBayerBG2BGR_EA = 135
	ColorBayerGB2BGR_EA = 136
	ColorBayerRG2BGR_EA = 137
	ColorBayerGR2BGR_EA = 138

	ColorBayerBG2RGB_EA = ColorBayerRG2BGR_EA
	ColorBayerGB2RGB_EA = ColorBayerGR2BGR_EA
	ColorBayerRG2RGB_EA = ColorBayerBG2BGR_EA
	ColorBayerGR2RGB_EA = ColorBayerGB2BGR_EA

	//! Demosaicing with alpha channel
	ColorBayerBG2BGRA = 139
	ColorBayerGB2BGRA = 140
	ColorBayerRG2BGRA = 141
	ColorBayerGR2BGRA = 142

	ColorBayerBG2RGBA = ColorBayerRG2BGRA
	ColorBayerGB2RGBA = ColorBayerGR2BGRA
	ColorBayerRG2RGBA = ColorBayerBG2BGRA
	ColorBayerGR2RGBA = ColorBayerGB2BGRA

	ColorCOLORCVT_MAX = 143
)
