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

func drawColorBackground(ctx *gg.Context, color string) {
	ctx.SetHexColor(color)
	ctx.DrawRectangle(0, 0, float64(ctx.Width()), float64(ctx.Height()))
	ctx.Fill()
}

func drawImageBackground(ctx *gg.Context, img image.Image) {
	// Draw the blurred image
	blurImg := imaging.Blur(img, 10)
	blurLayer := image.NewRGBA(image.Rect(0, 0, ctx.Width(), ctx.Height()))
	draw.Draw(blurLayer, blurLayer.Bounds(), blurImg, image.Point{}, draw.Src)
	draw.Draw(ctx.Image().(*image.RGBA), image.Rect(0, 0, ctx.Width(), ctx.Height()), blurLayer, image.Point{0, 0}, draw.Over)

	// Draw the dark overlay for readability
	ctx.SetRGBA(0, 0, 0, 0.5)
	ctx.DrawRectangle(0, 0, float64(ctx.Width()), float64(ctx.Height()))
	ctx.Fill()
}

func drawText(ctx *gg.Context, text string, color string) {
	// Load font face with dynamic size
	fontBytes, _ := os.ReadFile("fonts/Arsenal/Arsenal-Bold.ttf")
	font, _ := truetype.Parse(fontBytes)
	fontSize := calculateFontSize(text, ctx.Width(), ctx.Height(), font)
	face := truetype.NewFace(font, &truetype.Options{Size: fontSize})
	ctx.SetFontFace(face)

	// Calculate text position
	x := float64(ctx.Width()) / 2
	y := float64(ctx.Height()) / 2

	// Draw text centered on the image
	ctx.SetHexColor(color)
	ctx.DrawStringWrapped(text, x, y, 0.5, 0.5, float64(ctx.Width())*0.9, 1.5, gg.AlignCenter)
}

func Generate(text, background string, width, height int) (image.Image, error) {
	textColor := "#FFFFFF"
	backgroundColor := ""

	var img image.Image

	if background == "red" {
		backgroundColor = "#CC0000"
	} else if background == "white" {
		textColor = "#000000"
		backgroundColor = "#FFFFFF"
	} else if imageURL, err := url.Parse(background); err == nil && (imageURL.Scheme == "http" || imageURL.Scheme == "https") {
		img, err = loadImageFromURL(background)
		if err != nil {
			return nil, fmt.Errorf("failed to load image")
		}
	} else {
		return nil, fmt.Errorf("invalid image URL")
	}

	ctx := gg.NewContext(width, height)
	if backgroundColor != "" {
		drawColorBackground(ctx, backgroundColor)
	}
	if img != nil {
		img = resizeAndCropImage(img, width, height)
		drawImageBackground(ctx, img)
	}
	drawText(ctx, text, textColor)

	return ctx.Image(), nil
}
