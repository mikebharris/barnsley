//go:build wasm || linux || ignore
// +build wasm linux ignore

package main

import (
	"math/rand"
	"syscall/js"
)

func main() {
	const iterations = 1000

	doc := js.Global().Get("document")
	canvas := doc.Call("getElementById", "fernImg")
	width := doc.Get("body").Get("clientWidth").Float()
	height := doc.Get("body").Get("clientHeight").Float()
	canvas.Set("width", width)
	canvas.Set("height", height)
	ctx := canvas.Call("getContext", "2d")

	done := make(chan struct{}, 0)

	var barnsleyFunc js.Func
	barnsleyFunc = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		xRatio, yRatio := calculateCanvasRatios(width, height)
		fernGreen := "#008800"
		fernRed := "#880000"
		fernBlue := "#000088"

		var x float64 = 0
		var y float64 = 0
		for i := 0; i < iterations; i++ {
			r := rand.Intn(100)
			var colour string
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
			ctx.Set("fillStyle", colour)
			ctx.Call("fillRect", int(x*xRatio)+int(width)/2, int(y*yRatio), 1, 1)
		}
		js.Global().Call("requestAnimationFrame", barnsleyFunc)
		return nil
	})
	defer barnsleyFunc.Release()

	js.Global().Call("requestAnimationFrame", barnsleyFunc)
	<-done
}

func calculateCanvasRatios(width float64, height float64) (float64, float64) {
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
