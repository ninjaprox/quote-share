package imageGenerator

import (
	"fmt"
	"image"
	"image/draw"
	"net/url"
	"os"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
)

// overlayTextOnImage overlays text on an image with specified font size and text color.
func overlayTextOnImage(img image.Image, text string, fontSize float64, hexColor string) (image.Image, error) {
	dc := gg.NewContextForImage(img)

	// Load font face with specified size
	fontBytes, _ := os.ReadFile("fonts/Arsenal/Arsenal-Bold.ttf")
	font, _ := truetype.Parse(fontBytes)
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

func Generate(text, imageSource string) (image.Image, error) {
	fontSize := 180.0
	hexColor := "#FFFFFF"

	var img image.Image
	var err error

	// Load image from either URL or file system
	if imageURL, e := url.Parse(imageSource); e == nil && (imageURL.Scheme == "http" || imageURL.Scheme == "https") {
		img, err = loadImageFromURL(imageSource)
	} else {
		err = fmt.Errorf("invalid image URL")
	}

	if err != nil {
		return nil, fmt.Errorf("error loading image: %v", err)
	}

	// Overlay text on the image
	return overlayTextOnImage(img, text, fontSize, hexColor)
}
