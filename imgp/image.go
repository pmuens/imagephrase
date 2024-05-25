package imgp

import (
	"image"
	"image/png"
	"os"
)

type Pixel struct {
	R int
	G int
	B int
	A int
}

type Image struct {
	Width  int
	Height int
	Pixels [][]Pixel
}

// See: https://stackoverflow.com/a/41185404
func LoadImage(path string) (Image, error) {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	file, err := os.Open(path)
	if err != nil {
		return Image{}, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return Image{}, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels [][]Pixel
	for x := 0; x < width; x++ {
		var row []Pixel
		for y := 0; y < height; y++ {
			row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}

	result := Image{
		Width:  width,
		Height: height,
		Pixels: pixels,
	}

	return result, nil
}

func rgbaToPixel(r, g, b, a uint32) Pixel {
	return Pixel{
		R: int(r / 257),
		G: int(g / 257),
		B: int(b / 257),
		A: int(a / 257),
	}
}
