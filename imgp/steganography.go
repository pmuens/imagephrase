package imgp

import (
	"path/filepath"
	"strings"
)

const (
	LsbsPerPixel  = 3                                  // Red, Green, Blue.
	WordBitLength = 11                                 // log_2(2048).
	PixelsPerWord = (WordBitLength + 1) / LsbsPerPixel // Adding 1 so that division is without remainder.
)

func HideInImage(imgPath string, mnemonic string) (string, error) {
	numbers, err := WordsToInts(mnemonic)
	if err != nil {
		return "", err
	}

	image, err := LoadImage(imgPath)
	if err != nil {
		return "", err
	}

	err = hideNumbersInPixels(numbers, image)
	if err != nil {
		return "", err
	}

	ext := filepath.Ext(imgPath)
	dir, file := filepath.Split(imgPath)
	newFileName := strings.TrimSuffix(file, ext) + ".modified" + ext
	newImgPath := filepath.Join(dir, newFileName)

	err = SaveImage(newImgPath, image)
	if err != nil {
		return "", err
	}

	return newImgPath, nil
}

// See: https://stackoverflow.com/a/6059487
func hideNumbersInPixels(numbers []int, image Image) error {
	// TODO: Add support for hiding information in more than one row of pixels.
	row := image.Pixels[0]

	for i, number := range numbers {
		// Grab 4 pixels.
		low := i * PixelsPerWord
		high := low + PixelsPerWord
		chunk := row[low:high]

		index := 0

		// Iterate over number (which represents a word), 3 bits at a
		//	time (starting from LSb).
		for j := 0; j < WordBitLength; j += LsbsPerPixel {
			pixel := chunk[index]

			// Red.
			chunk[index].R = (pixel.R & (^1)) | (number & 1)
			number >>= 1

			// Green.
			chunk[index].G = (pixel.G & (^1)) | (number & 1)
			number >>= 1

			// Blue.
			chunk[index].B = (pixel.B & (^1)) | (number & 1)
			number >>= 1

			index++
		}
	}

	return nil
}
