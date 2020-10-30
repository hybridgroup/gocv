package gocv

/*
#include <stdlib.h>
#include "videoio.h"
*/
import "C"
import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"unsafe"
)

// Select preferred API for a capture object.
// Note: Backends are available only if they have been built with your OpenCV binaries
type VideoCaptureAPI int

const (
	// Auto detect == 0
	VideoCaptureAny VideoCaptureAPI = 0

	// Video For Windows (obsolete, removed)
	VideoCaptureVFW VideoCaptureAPI = 200

	// V4L/V4L2 capturing support
	VideoCaptureV4L VideoCaptureAPI = 200

	// Same as VideoCaptureV4L
	VideoCaptureV4L2 VideoCaptureAPI = 200

	// IEEE 1394 drivers
	VideoCaptureFirewire VideoCaptureAPI = 300

	// Same value as VideoCaptureFirewire
	VideoCaptureFireware VideoCaptureAPI = 300

	// Same value as VideoCaptureFirewire
	VideoCaptureIEEE1394 VideoCaptureAPI = 300

	// Same value as VideoCaptureFirewire
	VideoCaptureDC1394 VideoCaptureAPI = 300

	// Same value as VideoCaptureFirewire
	VideoCaptureCMU1394 VideoCaptureAPI = 300

	// QuickTime (obsolete, removed)
	VideoCaptureQT VideoCaptureAPI = 500

	// Unicap drivers (obsolete, removed)
	VideoCaptureUnicap VideoCaptureAPI = 600

	// DirectShow (via videoInput)
	VideoCaptureDshow VideoCaptureAPI = 700

	// PvAPI, Prosilica GigE SDK
	VideoCapturePvAPI VideoCaptureAPI = 800

	// OpenNI (for Kinect)
	VideoCaptureOpenNI VideoCaptureAPI = 900

	// OpenNI (for Asus Xtion)
	VideoCaptureOpenNIAsus VideoCaptureAPI = 910

	// Android - not used
	VideoCaptureAndroid VideoCaptureAPI = 1000

	// XIMEA Camera API
	VideoCaptureXiAPI VideoCaptureAPI = 1100

	// AVFoundation framework for iOS (OS X Lion will have the same API)
	VideoCaptureAVFoundation VideoCaptureAPI = 1200

	// Smartek Giganetix GigEVisionSDK
	VideoCaptureGiganetix VideoCaptureAPI = 1300

	// Microsoft Media Foundation (via videoInput)
	VideoCaptureMSMF VideoCaptureAPI = 1400

	// Microsoft Windows Runtime using Media Foundation
	VideoCaptureWinRT VideoCaptureAPI = 1410

	// RealSense (former Intel Perceptual Computing SDK)
	VideoCaptureIntelPerc VideoCaptureAPI = 1500

	// Synonym for VideoCaptureIntelPerc
	VideoCaptureRealsense VideoCaptureAPI = 1500

	// OpenNI2 (for Kinect)
	VideoCaptureOpenNI2 VideoCaptureAPI = 1600

	// OpenNI2 (for Asus Xtion and Occipital Structure sensors)
	VideoCaptureOpenNI2Asus VideoCaptureAPI = 1610

	// gPhoto2 connection
	VideoCaptureGPhoto2 VideoCaptureAPI = 1700

	// GStreamer
	VideoCaptureGstreamer VideoCaptureAPI = 1800

	// Open and record video file or stream using the FFMPEG library
	VideoCaptureFFmpeg VideoCaptureAPI = 1900

	// OpenCV Image Sequence (e.g. img_%02d.jpg)
	VideoCaptureImages VideoCaptureAPI = 2000

	// Aravis SDK
	VideoCaptureAravis VideoCaptureAPI = 2100

	// Built-in OpenCV MotionJPEG codec
	VideoCaptureOpencvMjpeg VideoCaptureAPI = 2200

	// Intel MediaSDK
	VideoCaptureIntelMFX VideoCaptureAPI = 2300

	// XINE engine (Linux)
	VideoCaptureXINE VideoCaptureAPI = 2400
)

// VideoCaptureProperties are the properties used for VideoCapture operations.
type VideoCaptureProperties int

const (
	// VideoCapturePosMsec contains current position of the
	// video file in milliseconds.
	VideoCapturePosMsec VideoCaptureProperties = 0

	// VideoCapturePosFrames 0-based index of the frame to be
	// decoded/captured next.
	VideoCapturePosFrames VideoCaptureProperties = 1

	// VideoCapturePosAVIRatio relative position of the video file:
	// 0=start of the film, 1=end of the film.
	VideoCapturePosAVIRatio VideoCaptureProperties = 2

	// VideoCaptureFrameWidth is width of the frames in the video stream.
	VideoCaptureFrameWidth VideoCaptureProperties = 3

	// VideoCaptureFrameHeight controls height of frames in the video stream.
	VideoCaptureFrameHeight VideoCaptureProperties = 4

	// VideoCaptureFPS controls capture frame rate.
	VideoCaptureFPS VideoCaptureProperties = 5

	// VideoCaptureFOURCC contains the 4-character code of codec.
	// see VideoWriter::fourcc for details.
	VideoCaptureFOURCC VideoCaptureProperties = 6

	// VideoCaptureFrameCount contains number of frames in the video file.
	VideoCaptureFrameCount VideoCaptureProperties = 7

	// VideoCaptureFormat format of the Mat objects returned by
	// VideoCapture::retrieve().
	VideoCaptureFormat VideoCaptureProperties = 8

	// VideoCaptureMode contains backend-specific value indicating
	// the current capture mode.
	VideoCaptureMode VideoCaptureProperties = 9

	// VideoCaptureBrightness is brightness of the image
	// (only for those cameras that support).
	VideoCaptureBrightness VideoCaptureProperties = 10

	// VideoCaptureContrast is contrast of the image
	// (only for cameras that support it).
	VideoCaptureContrast VideoCaptureProperties = 11

	// VideoCaptureSaturation saturation of the image
	// (only for cameras that support).
	VideoCaptureSaturation VideoCaptureProperties = 12

	// VideoCaptureHue hue of the image (only for cameras that support).
	VideoCaptureHue VideoCaptureProperties = 13

	// VideoCaptureGain is the gain of the capture image.
	// (only for those cameras that support).
	VideoCaptureGain VideoCaptureProperties = 14

	// VideoCaptureExposure is the exposure of the capture image.
	// (only for those cameras that support).
	VideoCaptureExposure VideoCaptureProperties = 15

	// VideoCaptureConvertRGB is a boolean flags indicating whether
	// images should be converted to RGB.
	VideoCaptureConvertRGB VideoCaptureProperties = 16

	// VideoCaptureWhiteBalanceBlueU is currently unsupported.
	VideoCaptureWhiteBalanceBlueU VideoCaptureProperties = 17

	// VideoCaptureRectification is the rectification flag for stereo cameras.
	// Note: only supported by DC1394 v 2.x backend currently.
	VideoCaptureRectification VideoCaptureProperties = 18

	// VideoCaptureMonochrome indicates whether images should be
	// converted to monochrome.
	VideoCaptureMonochrome VideoCaptureProperties = 19

	// VideoCaptureSharpness controls image capture sharpness.
	VideoCaptureSharpness VideoCaptureProperties = 20

	// VideoCaptureAutoExposure controls the DC1394 exposure control
	// done by camera, user can adjust reference level using this feature.
	VideoCaptureAutoExposure VideoCaptureProperties = 21

	// VideoCaptureGamma controls video capture gamma.
	VideoCaptureGamma VideoCaptureProperties = 22

	// VideoCaptureTemperature controls video capture temperature.
	VideoCaptureTemperature VideoCaptureProperties = 23

	// VideoCaptureTrigger controls video capture trigger.
	VideoCaptureTrigger VideoCaptureProperties = 24

	// VideoCaptureTriggerDelay controls video capture trigger delay.
	VideoCaptureTriggerDelay VideoCaptureProperties = 25

	// VideoCaptureWhiteBalanceRedV controls video capture setting for
	// white balance.
	VideoCaptureWhiteBalanceRedV VideoCaptureProperties = 26

	// VideoCaptureZoom controls video capture zoom.
	VideoCaptureZoom VideoCaptureProperties = 27

	// VideoCaptureFocus controls video capture focus.
	VideoCaptureFocus VideoCaptureProperties = 28

	// VideoCaptureGUID controls video capture GUID.
	VideoCaptureGUID VideoCaptureProperties = 29

	// VideoCaptureISOSpeed controls video capture ISO speed.
	VideoCaptureISOSpeed VideoCaptureProperties = 30

	// VideoCaptureBacklight controls video capture backlight.
	VideoCaptureBacklight VideoCaptureProperties = 32

	// VideoCapturePan controls video capture pan.
	VideoCapturePan VideoCaptureProperties = 33

	// VideoCaptureTilt controls video capture tilt.
	VideoCaptureTilt VideoCaptureProperties = 34

	// VideoCaptureRoll controls video capture roll.
	VideoCaptureRoll VideoCaptureProperties = 35

	// VideoCaptureIris controls video capture iris.
	VideoCaptureIris VideoCaptureProperties = 36

	// VideoCaptureSettings is the pop up video/camera filter dialog. Note:
	// only supported by DSHOW backend currently. The property value is ignored.
	VideoCaptureSettings VideoCaptureProperties = 37

	// VideoCaptureBufferSize controls video capture buffer size.
	VideoCaptureBufferSize VideoCaptureProperties = 38

	// VideoCaptureAutoFocus controls video capture auto focus..
	VideoCaptureAutoFocus VideoCaptureProperties = 39

	// VideoCaptureSarNumerator controls the sample aspect ratio: num/den (num)
	VideoCaptureSarNumerator VideoCaptureProperties = 40

	// VideoCaptureSarDenominator controls the sample aspect ratio: num/den (den)
	VideoCaptureSarDenominator VideoCaptureProperties = 41

	// VideoCaptureBackend is the current api backend (VideoCaptureAPI). Read-only property.
	VideoCaptureBackend VideoCaptureProperties = 42

	// VideoCaptureChannel controls the video input or channel number (only for those cameras that support).
	VideoCaptureChannel VideoCaptureProperties = 43

	// VideoCaptureAutoWB controls the auto white-balance.
	VideoCaptureAutoWB VideoCaptureProperties = 44

	// VideoCaptureWBTemperature controls the white-balance color temperature
	VideoCaptureWBTemperature VideoCaptureProperties = 45

	// VideoCaptureCodecPixelFormat shows the the codec's pixel format (4-character code). Read-only property.
	// Subset of AV_PIX_FMT_* or -1 if unknown.
	VideoCaptureCodecPixelFormat VideoCaptureProperties = 46

	// VideoCaptureBitrate displays the video bitrate in kbits/s. Read-only property.
	VideoCaptureBitrate VideoCaptureProperties = 47
)

// VideoCapture is a wrapper around the OpenCV VideoCapture class.
//
// For further details, please see:
// http://docs.opencv.org/master/d8/dfe/classcv_1_1VideoCapture.html
//
type VideoCapture struct {
	p C.VideoCapture
}

// VideoCaptureFile opens a VideoCapture from a file and prepares
// to start capturing. It returns error if it fails to open the file stored in uri path.
func VideoCaptureFile(uri string) (vc *VideoCapture, err error) {
	vc = &VideoCapture{p: C.VideoCapture_New()}

	cURI := C.CString(uri)
	defer C.free(unsafe.Pointer(cURI))

	if !C.VideoCapture_Open(vc.p, cURI) {
		err = fmt.Errorf("Error opening file: %s", uri)
	}

	return
}

// VideoCaptureFile opens a VideoCapture from a file and prepares
// to start capturing. It returns error if it fails to open the file stored in uri path.
func VideoCaptureFileWithAPI(uri string, apiPreference VideoCaptureAPI) (vc *VideoCapture, err error) {
	vc = &VideoCapture{p: C.VideoCapture_New()}

	cURI := C.CString(uri)
	defer C.free(unsafe.Pointer(cURI))

	if !C.VideoCapture_OpenWithAPI(vc.p, cURI, C.int(apiPreference)) {
		err = fmt.Errorf("Error opening file: %s with api backend: %d", uri, apiPreference)
	}

	return
}

// VideoCaptureDevice opens a VideoCapture from a device and prepares
// to start capturing. It returns error if it fails to open the video device.
func VideoCaptureDevice(device int) (vc *VideoCapture, err error) {
	vc = &VideoCapture{p: C.VideoCapture_New()}

	if !C.VideoCapture_OpenDevice(vc.p, C.int(device)) {
		err = fmt.Errorf("Error opening device: %d", device)
	}

	return
}

// VideoCaptureDevice opens a VideoCapture from a device with the api preference.
// It returns error if it fails to open the video device.
func VideoCaptureDeviceWithAPI(device int, apiPreference VideoCaptureAPI) (vc *VideoCapture, err error) {
	vc = &VideoCapture{p: C.VideoCapture_New()}

	if !C.VideoCapture_OpenDeviceWithAPI(vc.p, C.int(device), C.int(apiPreference)) {
		err = fmt.Errorf("Error opening device: %d with api backend: %d", device, apiPreference)
	}

	return
}

// Close VideoCapture object.
func (v *VideoCapture) Close() error {
	C.VideoCapture_Close(v.p)
	v.p = nil
	return nil
}

// Set parameter with property (=key).
func (v *VideoCapture) Set(prop VideoCaptureProperties, param float64) {
	C.VideoCapture_Set(v.p, C.int(prop), C.double(param))
}

// Get parameter with property (=key).
func (v VideoCapture) Get(prop VideoCaptureProperties) float64 {
	return float64(C.VideoCapture_Get(v.p, C.int(prop)))
}

// IsOpened returns if the VideoCapture has been opened to read from
// a file or capture device.
func (v *VideoCapture) IsOpened() bool {
	isOpened := C.VideoCapture_IsOpened(v.p)
	return isOpened != 0
}

// Read reads the next frame from the VideoCapture to the Mat passed in
// as the param. It returns false if the VideoCapture cannot read frame.
func (v *VideoCapture) Read(m *Mat) bool {
	return C.VideoCapture_Read(v.p, m.p) != 0
}

// Grab skips a specific number of frames.
func (v *VideoCapture) Grab(skip int) {
	C.VideoCapture_Grab(v.p, C.int(skip))
}

// CodecString returns a string representation of FourCC bytes, i.e. the name of a codec
func (v *VideoCapture) CodecString() string {
	res := ""
	hexes := []int64{0xff, 0xff00, 0xff0000, 0xff000000}
	for i, h := range hexes {
		res += string(rune(int64(v.Get(VideoCaptureFOURCC)) & h >> (uint(i * 8))))
	}
	return res
}

// ToCodec returns an float64 representation of FourCC bytes
func (v *VideoCapture) ToCodec(codec string) float64 {
	if len(codec) != 4 {
		return -1.0
	}
	c1 := []rune(string(codec[0]))[0]
	c2 := []rune(string(codec[1]))[0]
	c3 := []rune(string(codec[2]))[0]
	c4 := []rune(string(codec[3]))[0]
	return float64((c1 & 255) + ((c2 & 255) << 8) + ((c3 & 255) << 16) + ((c4 & 255) << 24))
}

// VideoWriter is a wrapper around the OpenCV VideoWriter`class.
//
// For further details, please see:
// http://docs.opencv.org/master/dd/d9e/classcv_1_1VideoWriter.html
//
type VideoWriter struct {
	mu *sync.RWMutex
	p  C.VideoWriter
}

// VideoWriterFile opens a VideoWriter with a specific output file.
// The "codec" param should be the four-letter code for the desired output
// codec, for example "MJPG".
//
// For further details, please see:
// http://docs.opencv.org/master/dd/d9e/classcv_1_1VideoWriter.html#a0901c353cd5ea05bba455317dab81130
//
func VideoWriterFile(name string, codec string, fps float64, width int, height int, isColor bool) (vw *VideoWriter, err error) {

	if fps == 0 || width == 0 || height == 0 {
		return nil, fmt.Errorf("one of the numerical parameters "+
			"is equal to zero: FPS: %f, width: %d, height: %d", fps, width, height)
	}

	vw = &VideoWriter{
		p:  C.VideoWriter_New(),
		mu: &sync.RWMutex{},
	}

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cCodec := C.CString(codec)
	defer C.free(unsafe.Pointer(cCodec))

	C.VideoWriter_Open(vw.p, cName, cCodec, C.double(fps), C.int(width), C.int(height), C.bool(isColor))
	return
}

// Close VideoWriter object.
func (vw *VideoWriter) Close() error {
	C.VideoWriter_Close(vw.p)
	vw.p = nil
	return nil
}

// IsOpened checks if the VideoWriter is open and ready to be written to.
//
// For further details, please see:
// http://docs.opencv.org/master/dd/d9e/classcv_1_1VideoWriter.html#a9a40803e5f671968ac9efa877c984d75
//
func (vw *VideoWriter) IsOpened() bool {
	isOpend := C.VideoWriter_IsOpened(vw.p)
	return isOpend != 0
}

// Write the next video frame from the Mat image to the open VideoWriter.
//
// For further details, please see:
// http://docs.opencv.org/master/dd/d9e/classcv_1_1VideoWriter.html#a3115b679d612a6a0b5864a0c88ed4b39
//
func (vw *VideoWriter) Write(img Mat) error {
	vw.mu.Lock()
	defer vw.mu.Unlock()
	C.VideoWriter_Write(vw.p, img.p)
	return nil
}

// OpenVideoCapture return VideoCapture specified by device ID if v is a
// number. Return VideoCapture created from video file, URL, or GStreamer
// pipeline if v is a string.
func OpenVideoCapture(v interface{}) (*VideoCapture, error) {
	switch vv := v.(type) {
	case int:
		return VideoCaptureDevice(vv)
	case string:
		id, err := strconv.Atoi(vv)
		if err == nil {
			return VideoCaptureDevice(id)
		}
		return VideoCaptureFile(vv)
	default:
		return nil, errors.New("argument must be int or string")
	}
}

func OpenVideoCaptureWithAPI(v interface{}, apiPreference VideoCaptureAPI) (*VideoCapture, error) {
	switch vv := v.(type) {
	case int:
		return VideoCaptureDeviceWithAPI(vv, apiPreference)
	case string:
		id, err := strconv.Atoi(vv)
		if err == nil {
			return VideoCaptureDeviceWithAPI(id, apiPreference)
		}
		return VideoCaptureFileWithAPI(vv, apiPreference)
	default:
		return nil, errors.New("argument must be int or string")
	}
}
