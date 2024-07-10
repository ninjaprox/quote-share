package imageGenerator

import (
	"fmt"
	"image"
	"image/draw"
	"math"
	"net/url"
	"os"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
)

func resizeAndCropImage(img image.Image, width, height int) image.Image {
	// Resize the image to cover the target size while maintaining aspect ratio
	resized := imaging.Fill(img, width, height, imaging.Center, imaging.Lanczos)
	return resized
}

func calculateFontSize(text string, width, height int, font *truetype.Font) float64 {
	// Calculate the desired text area (30% of the image area)
	imageArea := float64(width * height)
	desiredTextArea := imageArea * 0.3

	// Calculate initial font size (this is a starting point, we'll adjust it)
	fontSize := math.Sqrt(desiredTextArea / float64(len(text)))

	// Adjust font size to fit 30% of the image area
	for {
		face := truetype.NewFace(font, &truetype.Options{Size: fontSize})
		dc := gg.NewContext(width, height)
		dc.SetFontFace(face)

		// Measure text dimensions
		textWidth, textHeight := dc.MeasureMultilineString(text, 1.5)
		textArea := textWidth * textHeight

		if textArea <= desiredTextArea {
			break
		}

		fontSize *= 0.9 // Reduce font size and try again
	}

	// clamp fontSize within 30 to 100 range
	return math.Max(30, math.Min(100, fontSize))
}

// overlayTextOnImage overlays text on an image with specified font size and text color.
func overlayTextOnImage(img image.Image, text string, hexColor string) (image.Image, error) {
	bounds := img.Bounds()
	dc := gg.NewContext(bounds.Max.X, bounds.Max.Y)

	// Draw the resized image onto the context
	dc.DrawImage(img, 0, 0)

	// Load font face with dynamic size
	fontBytes, _ := os.ReadFile("fonts/Arsenal/Arsenal-Bold.ttf")
	font, _ := truetype.Parse(fontBytes)
	fontSize := calculateFontSize(text, dc.Width(), dc.Height(), font)
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
	dc.DrawStringWrapped(text, x, y, 0.5, 0.5, float64(dc.Width())*0.9, 1.5, gg.AlignCenter)

	return dc.Image(), nil
}

func Generate(text, imageSource string, width, height int) (image.Image, error) {
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

	// Resize and crop the image to the specified dimensions
	img = resizeAndCropImage(img, width, height)

	// Overlay text on the image
	return overlayTextOnImage(img, text, hexColor)
}
