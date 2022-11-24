package gocv

/*
#include <stdlib.h>
#include "photo.h"
*/
import "C"

import (
	"image"
	"unsafe"
)

//SeamlessCloneFlags seamlessClone algorithm flags
type SeamlessCloneFlags int

// MergeMertens is a wrapper around the cv::MergeMertens.
type MergeMertens struct {
	p unsafe.Pointer // This unsafe pointer will in fact be a C.MergeMertens
}

// AlignMTB is a wrapper around the cv::AlignMTB.
type AlignMTB struct {
	p unsafe.Pointer // This unsafe pointer will in fact be a C.AlignMTB
}

const (
	// NormalClone The power of the method is fully expressed when inserting objects with complex outlines into a new background.
	NormalClone SeamlessCloneFlags = iota

	// MixedClone The classic method, color-based selection and alpha masking might be time consuming and often leaves an undesirable halo. Seamless cloning, even averaged with the original image, is not effective. Mixed seamless cloning based on a loose selection proves effective.
	MixedClone

	// MonochromeTransfer Monochrome transfer allows the user to easily replace certain features of one object by alternative features.
	MonochromeTransfer
)

// ColorChange mix two differently colored versions of an image seamlessly.
//
// For further details, please see:
// https://docs.opencv.org/master/df/da0/group__photo__clone.html#ga6684f35dc669ff6196a7c340dc73b98e
//
func ColorChange(src, mask Mat, dst *Mat, red_mul, green_mul, blue_mul float32) {
	C.ColorChange(src.p, mask.p, dst.p, C.float(red_mul), C.float(green_mul), C.float(blue_mul))
}

// SeamlessClone blend two image by Poisson Blending.
//
// For further details, please see:
// https://docs.opencv.org/master/df/da0/group__photo__clone.html#ga2bf426e4c93a6b1f21705513dfeca49d
//
func SeamlessClone(src, dst, mask Mat, p image.Point, blend *Mat, flags SeamlessCloneFlags) {
	cp := C.struct_Point{
		x: C.int(p.X),
		y: C.int(p.Y),
	}

	C.SeamlessClone(src.p, dst.p, mask.p, cp, blend.p, C.int(flags))
}

// IlluminationChange modifies locally the apparent illumination of an image.
//
// For further details, please see:
// https://docs.opencv.org/master/df/da0/group__photo__clone.html#gac5025767cf2febd8029d474278e886c7
//
func IlluminationChange(src, mask Mat, dst *Mat, alpha, beta float32) {
	C.IlluminationChange(src.p, mask.p, dst.p, C.float(alpha), C.float(beta))
}

// TextureFlattening washes out the texture of the selected region, giving its contents a flat aspect.
//
// For further details, please see:
// https://docs.opencv.org/master/df/da0/group__photo__clone.html#gad55df6aa53797365fa7cc23959a54004
//
func TextureFlattening(src, mask Mat, dst *Mat, lowThreshold, highThreshold float32, kernelSize int) {
	C.TextureFlattening(src.p, mask.p, dst.p, C.float(lowThreshold), C.float(highThreshold), C.int(kernelSize))
}

// FastNlMeansDenoisingColoredMulti denoises the selected images.
//
// For further details, please see:
// https://docs.opencv.org/master/d1/d79/group__photo__denoise.html#gaa501e71f52fb2dc17ff8ca5e7d2d3619
//
func FastNlMeansDenoisingColoredMulti(src []Mat, dst *Mat, imgToDenoiseIndex int, temporalWindowSize int) {
	cMatArray := make([]C.Mat, len(src))
	for i, r := range src {
		cMatArray[i] = (C.Mat)(r.p)
	}
	matsVector := C.struct_Mats{
		mats:   (*C.Mat)(&cMatArray[0]),
		length: C.int(len(src)),
	}
	C.FastNlMeansDenoisingColoredMulti(matsVector, dst.p, C.int(imgToDenoiseIndex), C.int(temporalWindowSize))
}

// FastNlMeansDenoisingColoredMulti denoises the selected images.
//
// For further details, please see:
// https://docs.opencv.org/master/d1/d79/group__photo__denoise.html#gaa501e71f52fb2dc17ff8ca5e7d2d3619
//
func FastNlMeansDenoisingColoredMultiWithParams(src []Mat, dst *Mat, imgToDenoiseIndex int, temporalWindowSize int, h float32, hColor float32, templateWindowSize int, searchWindowSize int) {
	cMatArray := make([]C.Mat, len(src))
	for i, r := range src {
		cMatArray[i] = (C.Mat)(r.p)
	}
	matsVector := C.struct_Mats{
		mats:   (*C.Mat)(&cMatArray[0]),
		length: C.int(len(src)),
	}
	C.FastNlMeansDenoisingColoredMultiWithParams(matsVector, dst.p, C.int(imgToDenoiseIndex), C.int(temporalWindowSize), C.float(h), C.float(hColor), C.int(templateWindowSize), C.int(searchWindowSize))
}

// NewMergeMertens returns returns a new MergeMertens white LDR merge algorithm.
// of type MergeMertens with default parameters.
// MergeMertens algorithm merge the ldr image should result in a HDR image.
//
// For further details, please see:
// https://docs.opencv.org/master/d6/df5/group__photo__hdr.html
// https://docs.opencv.org/master/d7/dd6/classcv_1_1MergeMertens.html
// https://docs.opencv.org/master/d6/df5/group__photo__hdr.html#ga79d59aa3cb3a7c664e59a4b5acc1ccb6
//
func NewMergeMertens() MergeMertens {
	return MergeMertens{p: unsafe.Pointer(C.MergeMertens_Create())}
}

// NewMergeMertensWithParams returns a new MergeMertens white LDR merge algorithm
// of type MergeMertens with customized parameters.
// MergeMertens algorithm merge the ldr image should result in a HDR image.
//
// For further details, please see:
// https://docs.opencv.org/master/d6/df5/group__photo__hdr.html
// https://docs.opencv.org/master/d7/dd6/classcv_1_1MergeMertens.html
// https://docs.opencv.org/master/d6/df5/group__photo__hdr.html#ga79d59aa3cb3a7c664e59a4b5acc1ccb6
//
func NewMergeMertensWithParams(contrast_weight float32, saturation_weight float32, exposure_weight float32) MergeMertens {
	return MergeMertens{p: unsafe.Pointer(C.MergeMertens_CreateWithParams(C.float(contrast_weight), C.float(saturation_weight), C.float(exposure_weight)))}
}

// Close MergeMertens.
func (b *MergeMertens) Close() error {
	C.MergeMertens_Close((C.MergeMertens)(b.p)) // Here the unsafe pointer is cast into the right type
	b.p = nil
	return nil
}

// BalanceWhite computes merge LDR images using the current MergeMertens.
// Return a image MAT : 8bits 3 channel image ( RGB 8 bits )
// For further details, please see:
// https://docs.opencv.org/master/d7/dd6/classcv_1_1MergeMertens.html#a2d2254b2aab722c16954de13a663644d
//
func (b *MergeMertens) Process(src []Mat, dst *Mat) {
	cMatArray := make([]C.Mat, len(src))
	for i, r := range src {
		cMatArray[i] = (C.Mat)(r.p)
	}
	// Conversion function from a Golang slice into an array of matrices that are understood by OpenCV
	matsVector := C.struct_Mats{
		mats:   (*C.Mat)(&cMatArray[0]),
		length: C.int(len(src)),
	}
	C.MergeMertens_Process((C.MergeMertens)(b.p), matsVector, dst.p)
	// Convert a series of double [0.0,1.0] to [0,255] with Golang
	dst.ConvertToWithParams(dst, MatTypeCV8UC3, 255.0, 0.0)
}

// NewAlignMTB returns an AlignMTB for converts images to median threshold bitmaps.
// of type AlignMTB converts images to median threshold bitmaps (1 for pixels
// brighter than median luminance and 0 otherwise) and than aligns the resulting
// bitmaps using bit operations.

// For further details, please see:
// https://docs.opencv.org/master/d6/df5/group__photo__hdr.html
// https://docs.opencv.org/master/d7/db6/classcv_1_1AlignMTB.html
// https://docs.opencv.org/master/d6/df5/group__photo__hdr.html#ga2f1fafc885a5d79dbfb3542e08db0244
//
func NewAlignMTB() AlignMTB {
	return AlignMTB{p: unsafe.Pointer(C.AlignMTB_Create())}
}

// NewAlignMTBWithParams returns an AlignMTB for converts images to median threshold bitmaps.
// of type AlignMTB converts images to median threshold bitmaps (1 for pixels
// brighter than median luminance and 0 otherwise) and than aligns the resulting
// bitmaps using bit operations.

// For further details, please see:
// https://docs.opencv.org/master/d6/df5/group__photo__hdr.html
// https://docs.opencv.org/master/d7/db6/classcv_1_1AlignMTB.html
// https://docs.opencv.org/master/d6/df5/group__photo__hdr.html#ga2f1fafc885a5d79dbfb3542e08db0244
//
func NewAlignMTBWithParams(max_bits int, exclude_range int, cut bool) AlignMTB {
	return AlignMTB{p: unsafe.Pointer(C.AlignMTB_CreateWithParams(C.int(max_bits), C.int(exclude_range), C.bool(cut)))}
}

// Close AlignMTB.
func (b *AlignMTB) Close() error {
	C.AlignMTB_Close((C.AlignMTB)(b.p))
	b.p = nil
	return nil
}

// Process computes an alignment using the current AlignMTB.
//
// For further details, please see:
// https://docs.opencv.org/master/d7/db6/classcv_1_1AlignMTB.html#a37b3417d844f362d781f34155cbcb201
//
func (b *AlignMTB) Process(src []Mat, dst *[]Mat) {

	cSrcArray := make([]C.Mat, len(src))
	for i, r := range src {
		cSrcArray[i] = r.p
	}
	cSrcMats := C.struct_Mats{
		mats:   (*C.Mat)(&cSrcArray[0]),
		length: C.int(len(src)),
	}

	cDstMats := C.struct_Mats{}

	C.AlignMTB_Process((C.AlignMTB)(b.p), cSrcMats, &cDstMats)

	// Pass the matrices by reference from an OpenCV/C++ to a GoCV::Mat object
	for i := C.int(0); i < cDstMats.length; i++ {
		var tempdst Mat
		tempdst.p = C.Mats_get(cDstMats, i)
		*dst = append(*dst, tempdst)
	}
	return
}

// FastNlMeansDenoising performs image denoising using Non-local Means Denoising algorithm
// http://www.ipol.im/pub/algo/bcm_non_local_means_denoising/
//
// For further details, please see:
// https://docs.opencv.org/4.x/d1/d79/group__photo__denoise.html#ga4c6b0031f56ea3f98f768881279ffe93
//
func FastNlMeansDenoising(src Mat, dst *Mat) {
	C.FastNlMeansDenoising(src.p, dst.p)
}

// FastNlMeansDenoisingWithParams performs image denoising using Non-local Means Denoising algorithm
// http://www.ipol.im/pub/algo/bcm_non_local_means_denoising/
//
// For further details, please see:
// https://docs.opencv.org/4.x/d1/d79/group__photo__denoise.html#ga4c6b0031f56ea3f98f768881279ffe93
//
func FastNlMeansDenoisingWithParams(src Mat, dst *Mat, h float32, templateWindowSize int, searchWindowSize int) {
	C.FastNlMeansDenoisingWithParams(src.p, dst.p, C.float(h), C.int(templateWindowSize), C.int(searchWindowSize))
}

// FastNlMeansDenoisingColored is a modification of fastNlMeansDenoising function for colored images.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d1/d79/group__photo__denoise.html#ga21abc1c8b0e15f78cd3eff672cb6c476
//
func FastNlMeansDenoisingColored(src Mat, dst *Mat) {
	C.FastNlMeansDenoisingColored(src.p, dst.p)
}

// FastNlMeansDenoisingColoredWithParams is a modification of fastNlMeansDenoising function for colored images.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d1/d79/group__photo__denoise.html#ga21abc1c8b0e15f78cd3eff672cb6c476
//
func FastNlMeansDenoisingColoredWithParams(src Mat, dst *Mat, h float32, hColor float32, templateWindowSize int, searchWindowSize int) {
	C.FastNlMeansDenoisingColoredWithParams(src.p, dst.p, C.float(h), C.float(hColor), C.int(templateWindowSize), C.int(searchWindowSize))
}

// DetailEnhance filter enhances the details of a particular image
//
// For further details, please see:
// https://docs.opencv.org/4.x/df/dac/group__photo__render.html#gae5930dd822c713b36f8529b21ddebd0c
//
func DetailEnhance(src Mat, dst *Mat, sigma_s, sigma_r float32) {
	C.DetailEnhance(src.p, dst.p, C.float(sigma_s), C.float(sigma_r))
}

type EdgeFilter int

const (
	// RecursFilter Recursive Filtering.
	RecursFilter EdgeFilter = 1

	// NormconvFilter Normalized Convolution Filtering.
	NormconvFilter = 2
)

// EdgePreservingFilter filtering is the fundamental operation in image and video processing.
// Edge-preserving smoothing filters are used in many different applications.
//
// For further details, please see:
// https://docs.opencv.org/4.x/df/dac/group__photo__render.html#gafaee2977597029bc8e35da6e67bd31f7
//
func EdgePreservingFilter(src Mat, dst *Mat, filter EdgeFilter, sigma_s, sigma_r float32) {
	C.EdgePreservingFilter(src.p, dst.p, C.int(filter), C.float(sigma_s), C.float(sigma_r))
}

// PencilSketch pencil-like non-photorealistic line drawing.
//
// For further details, please see:
// https://docs.opencv.org/4.x/df/dac/group__photo__render.html#gae5930dd822c713b36f8529b21ddebd0c
//
func PencilSketch(src Mat, dst1, dst2 *Mat, sigma_s, sigma_r, shade_factor float32) {
	C.PencilSketch(src.p, dst1.p, dst2.p, C.float(sigma_s), C.float(sigma_r), C.float(shade_factor))
}

// Stylization aims to produce digital imagery with a wide variety of effects
// not focused on photorealism. Edge-aware filters are ideal for stylization,
// as they can abstract regions of low contrast while preserving, or enhancing,
// high-contrast features.
//
// For further details, please see:
// https://docs.opencv.org/4.x/df/dac/group__photo__render.html#gacb0f7324017df153d7b5d095aed53206
//
func Stylization(src Mat, dst *Mat, sigma_s, sigma_r float32) {
	C.Stylization(src.p, dst.p, C.float(sigma_s), C.float(sigma_r))
}

// InpaintMethods is the methods for inpainting process.
type InpaintMethods int

const (
	// NS inpaints using Navier-Stokes based method, created by Bertalmio, Marcelo,
	// Andrea L. Bertozzi, and Guillermo Sapiro in 2001
	NS InpaintMethods = 0

	// Telea inpaints using Fast Marching Method proposed by Alexandru Telea in 2004.
	Telea InpaintMethods = 1
)

// Inpaint reconstructs the selected image area from the pixel near the area boundary.
// The function may be used to remove dust and scratches from a scanned photo, or to
// remove undesirable objects from still images or video.
//
// For further details, please see:
// https://docs.opencv.org/4.x/d7/d8b/group__photo__inpaint.html#gaedd30dfa0214fec4c88138b51d678085
//
func Inpaint(src Mat, mask Mat, dst *Mat, inpaintRadius float32, algorithmType InpaintMethods) {
	C.PhotoInpaint(C.Mat(src.Ptr()), C.Mat(mask.Ptr()), C.Mat(dst.Ptr()), C.float(inpaintRadius), C.int(algorithmType))
}
