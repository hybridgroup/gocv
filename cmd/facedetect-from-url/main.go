// What it does:
//
// This example uses the CascadeClassifier class to detect faces from url,
// and draw a rectangle around each of them, before displaying them within a Window.
//
// How to run:
//
// facedetect-from-url [image URL] [classifier XML file] [image file]
//
// 		go run ./cmd/facedetect-from-url/main.go https://raw.githubusercontent.com/hybridgroup/gocv/release/images/face.jpg data/haarcascade_frontalface_default.xml output.jpg
//
// +build example

package main

import (
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"gocv.io/x/gocv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("How to run:\n\tfacedetect-from-url [image URL] [classifier XML file] [image file]")
		return
	}

	// parse args
	imageURL := os.Args[1]
	xmlFile := os.Args[2]
	saveFile := os.Args[3]

	// color for the rect when faces detected
	blue := color.RGBA{0, 0, 255, 0}

	// load classifier to recognize faces
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load(xmlFile) {
		fmt.Printf("Error reading cascade file: %v\n", xmlFile)
		return
	}

	// get image from URL
	res, err := http.Get(imageURL)
	if err != nil {
		log.Fatal(err)
	}

	resByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	img, err := gocv.IMDecode(resByte, 1)
	if err != nil {
		log.Fatal(err)
	}

	rects := classifier.DetectMultiScale(img)
	fmt.Printf("found %d faces\n", len(rects))
	// draw a rectangle around each face on the original image,
	// along with text identifing as "Human"
	for _, r := range rects {
		gocv.Rectangle(&img, r, blue, 3)

		size := gocv.GetTextSize("Human", gocv.FontHersheyPlain, 1.2, 2)
		pt := image.Pt(r.Min.X+(r.Min.X/2)-(size.X/2), r.Min.Y-2)
		gocv.PutText(&img, "Human", pt, gocv.FontHersheyPlain, 1.2, blue, 2)
	}
	gocv.IMWrite(saveFile, img)
	fmt.Printf("saved to %s\n", saveFile)
}
