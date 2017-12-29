package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"io"
	"os"
)

func justforfun22() {
	fileName := flag.String("file", "bc", " Choose an image file (bc|foundry|image|me)")
	flag.Parse()
	// You can register another format here
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	// bc.png, foundry.png, me.png, image.png
	file, err := os.Open("./images/" + *fileName + ".png")

	if err != nil {
		fmt.Println("Error: File could not be opened")
		os.Exit(1)
	}

	defer file.Close()

	pixels, err := getPixels(file)

	if err != nil {
		fmt.Println("Error: Image could not be decoded")
		os.Exit(1)
	}

	// convert pixels to puzzle input
	var input string
	for _, row := range pixels {
		for _, pixel := range row {
			brightness := pixel.R + pixel.G + pixel.B
			if brightness == 0 {
				input += " "
			} else if brightness < 100 {
				input += "░"
			} else if brightness < 100 {
				input += "▓"
			} else {
				input += "█"
			}
		}
		input += "\n"
	}
	advent22B(input)
	// fmt.Println(pixels)
}

// Get the bi-dimensional pixel array
func getPixels(file io.Reader) ([][]Pixel, error) {
	img, _, err := image.Decode(file)

	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels [][]Pixel
	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}

	return pixels, nil
}

// img.At(x, y).RGBA() returns four uint32 values; we want a Pixel
func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}

// Pixel struct example
type Pixel struct {
	R int
	G int
	B int
	A int
}
