package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

var blue color.Color = color.RGBA{0, 0, 200, 100}
var red color.Color = color.RGBA{255, 0, 0, 255}
var green color.Color = color.RGBA{0, 255, 0, 255}

func main() {
	file, err := os.Create("rectangle.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()

	img := image.NewRGBA(image.Rect(0, 0, 500, 500))

	draw.Draw(img, img.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	for x := 40; x < 460; x++ {
		y := x/460 + 100
		img.Set(x, y, red)
	}

	for x := 40; x < 460; x++ {
		y := x/460 + 400
		img.Set(x, y, red)
	}

	for x := 40; x < 460; x++ {
		y := x/460 + 100
		img.Set(y, x, green)
	}

	for x := 40; x < 460; x++ {
		y := x/460 + 400
		img.Set(y, x, green)
	}

	png.Encode(file, img)
}
