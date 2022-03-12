package gocv

/*
#include <stdlib.h>
#include "wechat_qrcode.h"
*/
import "C"
import (
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

func (wq *WeChatQRCode) DetectAndDecode(img Mat, point *[]Mat) []string {
	cMats := C.struct_Mats{}
	defer C.Mats_Close(cMats)
	cDecoded := C.CStrings{}
	defer C.CStrings_Close(cDecoded)
	cCodes := C.NewStringsVector()
	defer C.free(unsafe.Pointer(cCodes))

	cDecoded = C.WeChatQRCode_DetectAndDecode((C.WeChatQRCode)(wq.p), (C.Mat)(img.Ptr()), &(cMats), cCodes)
	ps := make([]Mat, cMats.length)

	for i := C.int(0); i < cMats.length; i++ {
		ps[i].p = C.Mats_get(cMats, i)
	}

	*point = ps

	result := make([]string, 0)
	for _, v := range toGoStrings(cDecoded) {
		result = append(result, v)
	}

	return result
}
