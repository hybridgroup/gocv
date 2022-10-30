// What it does:
//
// This example demonstrates a couple of uses of the XPhoto module.
// It can use the GrayworldWB class with BalanceWhite image
// to save an image file on disk.
//
// This example can also use the Inpaint functions with inpaint algorithms type
// to save an image file on disk.
//
// How to run:
//
// 		go run ./cmd/xphoto/main.go -i -g
//

package main

import (
	"flag"
	"fmt"
	"os"

	"gocv.io/x/gocv"
	"gocv.io/x/gocv/contrib"
)

func Inpaint() {

	jpgImage := gocv.IMRead("./images/space_shuttle.jpg", gocv.IMReadColor)

	if jpgImage.Empty() {
		fmt.Printf("Invalid read of Source Mat in TestInpaint test\n")
		return
	}

	src := gocv.NewMat()
	defer src.Close()
	sizeImage := jpgImage.Size()
	jpgImage.ConvertTo(&src, gocv.MatTypeCV8UC3)

	maskFsrFast := gocv.NewMatWithSizes(sizeImage, gocv.MatTypeCV8UC1)
	defer maskFsrFast.Close()

	dstShitMap := gocv.NewMat()
	defer dstShitMap.Close()
	contrib.Inpaint(&src, &maskFsrFast, &dstShitMap, contrib.FsrFast)

	dstFsrFast := gocv.NewMat()
	defer dstFsrFast.Close()
	contrib.Inpaint(&src, &maskFsrFast, &dstFsrFast, contrib.FsrFast)

	dstFsrBest := gocv.NewMat()
	defer dstFsrBest.Close()
	contrib.Inpaint(&src, &maskFsrFast, &dstFsrBest, contrib.FsrFast)

	if dstShitMap.Empty() || dstShitMap.Rows() != src.Rows() || dstShitMap.Cols() != src.Cols() || dstShitMap.Type() != src.Type() {
		fmt.Printf("Invlalid TestInpaint ShitMap test\n")
		return
	}
	fmt.Printf("ShitMap : MAT %d <> %d : %d\n", dstShitMap.Rows(), src.Rows(), dstShitMap.Type())
	gocv.IMWrite("ShitMap_inpaint.png", dstShitMap)

	if dstFsrFast.Empty() || dstFsrFast.Rows() != src.Rows() || dstFsrFast.Cols() != src.Cols() || dstFsrFast.Type() != src.Type() {
		fmt.Printf("Invlalid TestInpaint FsrFast test\n")
		return
	}
	fmt.Printf("FsrFast : MAT %d <> %d : %d\n", dstFsrFast.Rows(), src.Rows(), dstFsrFast.Type())
	gocv.IMWrite("FsrFast_inpaint.png", dstFsrFast)

	if dstFsrBest.Empty() || dstFsrBest.Rows() != src.Rows() || dstFsrBest.Cols() != src.Cols() || dstFsrBest.Type() != src.Type() {
		fmt.Printf("Invlalid TestInpaint FsrBest test\n")
		return
	}
	fmt.Printf("FsrBest : MAT %d <> %d : %d\n", dstFsrBest.Rows(), src.Rows(), dstFsrBest.Type())
	gocv.IMWrite("FsrBest_inpaint.png", dstFsrBest)

}

func BalanceWhite() {

	fileGrayWorld := "grayworld_space_shuttle.png"
	src := gocv.IMRead("./images/space_shuttle.jpg", gocv.IMReadColor)

	if src.Empty() {
		fmt.Printf("Invalid read of Source Mat in TestInpaint test\n")
		return
	}
	defer src.Close()

	fmt.Println("using GrayworldWB with white balance function")
	dst := gocv.NewMat()
	defer dst.Close()

	grayworldwb := contrib.NewGrayworldWB()
	defer grayworldwb.Close()

	grayworldwb.SetSaturationThreshold(0.7)
	grayworldwb.BalanceWhite(src, &dst)

	gocv.IMWrite(fileGrayWorld, dst)
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("How to run:\n\tmain [-i] [-b]")
		return
	}
	balanceWhitePtr := flag.Bool("b", false, "GrayworldWB functions")
	inpaintPtr := flag.Bool("i", false, "Inpaint functions")
	flag.Parse()

	if *balanceWhitePtr {
		BalanceWhite()
	}

	if *inpaintPtr {
		Inpaint()
	}

}
