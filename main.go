package main

import (
	"github.com/mikebharris/barnsley/fern"
	"image/png"
	"os"
)

func main() {
	width := 1024
	height := 1256
	var iterations = 10000000

	png.Encode(os.Stdout, fern.GenerateBarnsleyFern(width, height, iterations))
}
