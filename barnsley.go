package main

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
)

func main() {
	width := 1024
	height := 1256
	var iters = 10000000
	png.Encode(os.Stdout, generateBarnsleyFern(width, height, iters))
}

func generateBarnsleyFern(width int, height int, iters int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	xratio, yratio := calculateCanvasRatios(width, height)
	fernGreen := color.RGBA{R: 000, G: 128, B: 000, A: 0xff}

	var x float64 = 0
	var y float64 = 0
	for i := 0; i < iters; i++ {
		r := rand.Intn(100)
		if r == 1 {
			x, y = generateStem(x, y)
		} else if r < 87 {
			x, y = generateSuccessivelySmallerLeaflets(x, y)
		} else if r < 94 {
			x, y = generateLargestLeftHandLeaflet(x, y)
		} else {
			x, y = generateLargestRightHandLeaflet(x, y)
		}
		img.Set(int(x*xratio)+width/2, int(y*yratio), fernGreen)
	}
	return img
}

func calculateCanvasRatios(width int, height int) (float64, float64) {
	return float64(width) / 5.26, float64(height) / 10.02
}

func generateStem(x float64, y float64) (float64, float64) {
	x = 0
	y = y * 0.16
	return x, y
}

func generateSuccessivelySmallerLeaflets(x float64, y float64) (float64, float64) {
	t := 0.85*x + 0.04*y
	y = -0.04*x + 0.85*y + 1.60
	x = t
	return x, y
}

func generateLargestLeftHandLeaflet(x float64, y float64) (float64, float64) {
	t := 0.20*x + -0.26*y
	y = 0.23*x + 0.22*y + 1.60
	x = t
	return x, y
}

func generateLargestRightHandLeaflet(x float64, y float64) (float64, float64) {
	t := -0.15*x + 0.28*y
	y = 0.26*x + 0.24*y + 0.44
	x = t
	return x, y
}
