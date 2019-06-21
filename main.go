package main

import (
	"image"
	"io"
	"log"
	"os"
	"strings"

	"github.com/fogleman/gg"
)

func main() {
	arg := os.Args[1]

	writeToImage(arg)
}

func writeToImage(file_name string) {
	im, err := gg.LoadImage(file_name)
	const S = 1024
	if err != nil {
		log.Fatal(err)
	}

	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace("impact.ttf", 96); err != nil {
		panic(err)
	}

	dc.DrawRoundedRectangle(0, 0, 512, 512, 0)
	dc.DrawImage(im, 0, 0)
	dc.DrawStringAnchored("Hello, world!", S/2, S-96, 0.5, 0.5)
	dc.Clip()
	dc.SavePNG("out.png")

	pixel, err := getPixels(strings.NewReader(file_name))
	for _, element := range pixel {
		print("%s\n", element)
	}
}

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

func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}

type Pixel struct {
	R int
	G int
	B int
	A int
}
