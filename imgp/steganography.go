package imgp

import (
	"path/filepath"
	"strings"
)

func HideInImage(imgPath string, mnemonic string) (string, error) {
	image, err := LoadImage(imgPath)
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
