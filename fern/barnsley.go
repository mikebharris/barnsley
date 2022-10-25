package fern

import (
	"image"
	"image/color"
	"math/rand"
)

func GenerateBarnsleyFern(width int, height int, iterations int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	xRatio, yRatio := calculateCanvasRatios(width, height)
	fernGreen := color.RGBA{R: 000, G: 128, B: 000, A: 0xff}
	fernRed := color.RGBA{R: 128, G: 000, B: 000, A: 0xff}
	fernBlue := color.RGBA{R: 000, G: 000, B: 128, A: 0xff}

	var x float64 = 0
	var y float64 = 0
	for i := 0; i < iterations; i++ {
		r := rand.Intn(100)
		var colour color.RGBA
		if r == 1 {
			x, y = generateStem(x, y)
			colour = fernGreen
		} else if r < 87 {
			x, y = generateSuccessivelySmallerLeaflets(x, y)
			colour = fernGreen
		} else if r < 94 {
			x, y = generateLargestLeftHandLeaflet(x, y)
			colour = fernRed
		} else {
			x, y = generateLargestRightHandLeaflet(x, y)
			colour = fernBlue
		}
		img.Set(int(x*xRatio)+width/2, int(y*yRatio), colour)
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
