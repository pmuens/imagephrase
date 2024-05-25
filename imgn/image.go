package imgn

import (
	"image"
	"image/color"
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

type RGBA struct {
	R uint8
	G uint8
	B uint8
	A uint8
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

// See: https://yourbasic.org/golang/create-image
func SaveImage(path string, img Image) error {
	topLeft := image.Point{0, 0}
	bottomRight := image.Point{img.Width, img.Height}

	result := image.NewRGBA(image.Rectangle{topLeft, bottomRight})

	for x := 0; x < img.Width; x++ {
		for y := 0; y < img.Height; y++ {
			rgba := pixelToRgba(img.Pixels[x][y])
			result.Set(x, y, color.RGBA{
				R: rgba.R,
				G: rgba.G,
				B: rgba.B,
				A: rgba.A,
			})
		}
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	png.Encode(file, result)

	return nil
}

func rgbaToPixel(r, g, b, a uint32) Pixel {
	return Pixel{
		R: int(r / 257),
		G: int(g / 257),
		B: int(b / 257),
		A: int(a / 257),
	}
}

func pixelToRgba(pixel Pixel) RGBA {
	return RGBA{
		R: uint8(pixel.R),
		G: uint8(pixel.G),
		B: uint8(pixel.B),
		A: uint8(pixel.A),
	}
}
