package imageGenerator

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

// saveImage saves an image to the local file system.
func saveImage(img image.Image, outputPath string) error {
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	if strings.HasSuffix(outputPath, ".png") {
		err = png.Encode(outFile, img)
	} else {
		err = jpeg.Encode(outFile, img, &jpeg.Options{Quality: 95})
	}

	if err != nil {
		return err
	}

	return nil
}
