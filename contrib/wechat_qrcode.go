package contrib

/*
#include <stdlib.h>
#include "wechat_qrcode.h"
*/
import "C"
import (
	"gocv.io/x/gocv"
	"unsafe"
)

type WeChatQRCode struct {
	p C.WeChatQRCode
}

func NewWeChatQRCode(detectProtoTxt, detectCaffe, superProtoTxt, superCaffe string) *WeChatQRCode {
	dp := C.CString(detectProtoTxt)
	dc := C.CString(detectCaffe)
	sp := C.CString(superProtoTxt)
	sc := C.CString(superCaffe)

	defer C.free(unsafe.Pointer(dp))
	defer C.free(unsafe.Pointer(dc))
	defer C.free(unsafe.Pointer(sp))
	defer C.free(unsafe.Pointer(sc))
	return &WeChatQRCode{p: C.NewWeChatQRCode(dp, dc, sp, sc)}
}

func (wq *WeChatQRCode) DetectAndDecode(img gocv.Mat, point *[]gocv.Mat) []string {
	cMats := C.struct_Mats{}
	defer C.WeChatQRCode_Mats_Close(cMats)
	cDecoded := C.CStrings{}
	defer C.WeChatQRCode_CStrings_Close(cDecoded)
	cCodes := C.NewStringsVector()
	defer C.free(unsafe.Pointer(cCodes))

	cDecoded = C.WeChatQRCode_DetectAndDecode((C.WeChatQRCode)(wq.p), (C.Mat)(img.Ptr()), &(cMats), cCodes)
	ps := make([]gocv.Mat, cMats.length)

	for i := C.int(0); i < cMats.length; i++ {
		ps[i] = gocv.NewMat()
		C.WeChatQRCode_Mats_to(cMats, i, (C.Mat)(ps[i].Ptr()))
	}

	*point = ps

	result := make([]string, 0)
	for _, v := range toGoStrings(cDecoded) {
		result = append(result, v)
	}

	return result
}

func toGoStrings(strs C.CStrings) []string {
	length := int(strs.length)
	tmpslice := (*[1 << 20]*C.char)(unsafe.Pointer(strs.strs))[:length:length]
	gostrings := make([]string, length)
	for i, s := range tmpslice {
		gostrings[i] = C.GoString(s)
	}
	return gostrings
}
