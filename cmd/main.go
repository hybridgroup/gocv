package main

import (
	"fmt"
	"image"
	"os"

	"github.com/hybridgroup/gocv/cmd/captest"
	"github.com/hybridgroup/gocv/cmd/capwindow"
	"github.com/hybridgroup/gocv/cmd/faceblur"
	"github.com/hybridgroup/gocv/cmd/facedetect"
	"github.com/hybridgroup/gocv/cmd/mjpegstreamer"
	pvlblur "github.com/hybridgroup/gocv/cmd/pvl/faceblur"
	pvlsmile "github.com/hybridgroup/gocv/cmd/pvl/smiledetect"
	"github.com/hybridgroup/gocv/cmd/saveimage"
	"github.com/hybridgroup/gocv/cmd/savevideo"
	"github.com/hybridgroup/gocv/cmd/showimage"
	"github.com/hybridgroup/gocv/cmd/version"

	"github.com/urfave/cli"
)

func main() {
	var (
		cameraID                                   int
		classifierFile, host, imageFile, videoFile string
	)
	app := cli.NewApp()
	app.Name = "GoCV Command line utilities"
	app.Usage = ""
	app.Flags = []cli.Flag{}
	app.Commands = []cli.Command{
		{
			Name: "captest",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:        "camera-id",
					Value:       0,
					Usage:       "Camera ID",
					Destination: &cameraID,
				},
			},
			Usage: "Capture image from webcam by trying to read 100 frames",
			Action: func(c *cli.Context) {
				captest.Run(cameraID)
			},
		},
		{
			Name: "capwindow",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:        "camera-id",
					Value:       0,
					Usage:       "Camera ID",
					Destination: &cameraID,
				},
			},
			Usage: "Capture frames from a connected webcam and displays the video in Window class",
			Action: func(c *cli.Context) {
				capwindow.Run(cameraID)
			},
		},
		{
			Name: "faceblur",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:        "camera-id",
					Value:       0,
					Usage:       "Camera ID",
					Destination: &cameraID,
				},
				cli.StringFlag{
					Name:        "clf-file",
					Value:       "demo.xml",
					Usage:       "Classifier XML file",
					Destination: &classifierFile,
				},
			},
			Usage: "Capture frames from a connected webcam, detects a face and diplays the blurred face",
			Action: func(c *cli.Context) {
				faceblur.Run(cameraID, classifierFile)
			},
		},
		{
			Name: "facedetect",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:        "camera-id",
					Value:       0,
					Usage:       "Camera ID",
					Destination: &cameraID,
				},
				cli.StringFlag{
					Name:        "clf-file",
					Value:       "demo.xml",
					Usage:       "Classifier XML file",
					Destination: &classifierFile,
				},
			},
			Usage: "Capture frames from a connected webcam, detects a face and diplay a rectangle around the face",
			Action: func(c *cli.Context) {
				facedetect.Run(cameraID, classifierFile)
			},
		},
		{
			Name: "mstreamer",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:        "camera-id",
					Value:       0,
					Usage:       "Camera ID",
					Destination: &cameraID,
				},
				cli.StringFlag{
					Name:        "host",
					Value:       "127.0.0.1:8080",
					Usage:       "Host to stream MPJEG frames",
					Destination: &host,
				},
			},
			Usage: "Capture frames from a connected webcam and stream them",
			Action: func(c *cli.Context) {
				mjpegstreamer.Run(cameraID, host)
			},
		},
		{
			Name: "pvlfaceblur",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:        "camera-id",
					Value:       0,
					Usage:       "Camera ID",
					Destination: &cameraID,
				},
			},
			Usage: "Use Intel CV SDK PVL for face detection, then display blurred face",
			Action: func(c *cli.Context) {
				pvlblur.Run(cameraID)
			},
		},
		{
			Name: "pvlsmiledetect",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:        "camera-id",
					Value:       0,
					Usage:       "Camera ID",
					Destination: &cameraID,
				},
			},
			Usage: "Use Intel CV SDK PVL for face detection, to detect smiles",
			Action: func(c *cli.Context) {
				pvlsmile.Run(cameraID)
			},
		},
		{
			Name: "saveimage",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:        "camera-id",
					Value:       0,
					Usage:       "Camera ID",
					Destination: &cameraID,
				},
				cli.StringFlag{
					Name:        "image-file",
					Value:       "image.jpg",
					Usage:       "File to save image to",
					Destination: &imageFile,
				},
			},
			Usage: "Save image from connected webcam to disk",
			Action: func(c *cli.Context) {
				saveimage.Run(cameraID, imageFile)
			},
		},
		{
			Name: "savevideo",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:        "camera-id",
					Value:       0,
					Usage:       "Camera ID",
					Destination: &cameraID,
				},
				cli.StringFlag{
					Name:        "video-file",
					Value:       "video.mjpeg",
					Usage:       "File to save video to",
					Destination: &videoFile,
				},
			},
			Usage: "Save video from connected webcam to disk",
			Action: func(c *cli.Context) {
				savevideo.Run(cameraID, videoFile)
			},
		},
		{
			Name: "showimage",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "image-file",
					Value:       "image.jpg",
					Usage:       "File to save video to",
					Destination: &videoFile,
				},
			},
			Usage: "Show the passed in image file in an OpenCV window",
			Action: func(c *cli.Context) {
				showimage.Run(videoFile)
			},
		},
		{
			Name:  "version",
			Usage: "Display the current OpenCV library version",
			Action: func(c *cli.Context) {
				version.Run()
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
