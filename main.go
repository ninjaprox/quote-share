package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/gobold"
)

// loadImageFromFile loads an image from a file path.
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

// loadImageFromURL loads an image from a URL.
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

// overlayTextOnImage overlays text on an image with specified font size and text color.
func overlayTextOnImage(img image.Image, text string, fontSize float64, hexColor string) (image.Image, error) {
	dc := gg.NewContextForImage(img)

	// Load font face with specified size
	font, _ := truetype.Parse(gobold.TTF)
	face := truetype.NewFace(font, &truetype.Options{Size: fontSize})
	dc.SetFontFace(face)

	// Calculate text position
	x := float64(dc.Width()) / 2
	y := float64(dc.Height()) / 2

	// Create a blurred rectangle behind the text
	rectWidth := float64(dc.Width())
	rectHeight := float64(dc.Height())
	rectX := float64(0)
	rectY := float64(0)

	// Draw the blurred rectangle
	blurImg := imaging.Blur(img, 10)
	blurLayer := image.NewRGBA(image.Rect(0, 0, dc.Width(), dc.Height()))
	draw.Draw(blurLayer, blurLayer.Bounds(), blurImg, image.Point{}, draw.Src)
	draw.Draw(dc.Image().(*image.RGBA), image.Rect(int(rectX), int(rectY), int(rectX+rectWidth), int(rectY+rectHeight)), blurLayer, image.Point{int(rectX), int(rectY)}, draw.Over)

	// Set transparency for the rectangle
	dc.SetRGBA(0, 0, 0, 0.5)
	dc.DrawRectangle(rectX, rectY, rectWidth, rectHeight)
	dc.Fill()

	// Draw text centered on the image
	dc.SetHexColor(hexColor)
	dc.DrawStringAnchored(text, x, y, 0.5, 0.5)

	return dc.Image(), nil
}

func main() {
	// Input parameters
	imageSource := "https://images.unsplash.com/photo-1719670046288-f03275608b76?q=80&w=2487&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D" // or "path/to/your/image.jpg"
	text := "Your overlay text here"
	fontSize := 189.0
	hexColor := "#FFFFFF" // Example hex color code
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
	resultImg, err := overlayTextOnImage(img, text, fontSize, hexColor)
	if err != nil {
		fmt.Println("Error overlaying text:", err)
		return
	}

	// Save the resulting image to the local file system
	err = saveImage(resultImg, outputPath)
	if err != nil {
		fmt.Println("Error saving image:", err)
		return
	}

	fmt.Println("Image saved successfully to", outputPath)
}
