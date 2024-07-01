package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"strings"

	"github.com/fogleman/gg"
	"golang.org/x/image/font/inconsolata"
)

func loadImageFromFile(filePath string) (image.Image, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func loadImageFromURL(url string) (image.Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, err
	}

	return img, nil
}

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

func overlayTextOnImage(img image.Image, text string) image.Image {
	const W = 800
	const H = 600

	dc := gg.NewContextForImage(img)
	dc.SetFontFace(inconsolata.Bold8x16)
	dc.SetRGB(1, 1, 1)
	dc.DrawStringAnchored(text, W/2, H/2, 0.5, 0.5)

	return dc.Image()
}

func main() {
	// Input parameters
	imageSource := "https://images.unsplash.com/photo-1719670046288-f03275608b76?q=80&w=2487&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D" // or "path/to/your/image.jpg"
	text := "Your overlay text here"
	outputPath := "output.jpg"

	var img image.Image
	var err error

	// Load image from either URL or file system
	if strings.HasPrefix(imageSource, "http://") || strings.HasPrefix(imageSource, "https://") {
		img, err = loadImageFromURL(imageSource)
	} else {
		img, err = loadImageFromFile(imageSource)
	}

	if err != nil {
		fmt.Println("Error loading image:", err)
		return
	}

	// Overlay text on the image
	resultImg := overlayTextOnImage(img, text)

	// Save the resulting image to the local file system
	err = saveImage(resultImg, outputPath)
	if err != nil {
		fmt.Println("Error saving image:", err)
		return
	}

	fmt.Println("Image saved successfully to", outputPath)
}
