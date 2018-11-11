package main

import (
	"fmt"
	"image"
	"os"

	"gocv.io/x/gocv"
)

func main() {
	blockSize := image.Point{X: 10, Y: 10}

	args := os.Args
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "Not enough cmd arguments")
		os.Exit(1)
	}
	orig := gocv.IMRead(args[1], gocv.IMReadColor)
	if orig.Empty() {
		fmt.Fprintf(os.Stderr, "%v is not image file", args[0])
		os.Exit(1)
	}

	// TODO mozaicRectのimage.Rectangle.Maxの値が画像のサイズを上回っている時エラー
	//      Maxは画像の座標の終端よりも1大きい値を指定する
	mozaicRect := image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: orig.Cols(), Y: orig.Rows()},
	}
	mozaicTarget := orig.Region(mozaicRect)

	if err := Mozaic(&mozaicTarget, blockSize); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err.Error())
	}

	origWindow := gocv.NewWindow("orig")
	origWindow.IMShow(orig)
	origWindow.WaitKey(0)
}
