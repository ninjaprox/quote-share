package handlers

import (
	"fmt"
	"image/jpeg"
	"net/http"

	"vinhis.me/quote-share/lib/imageGenerator"
)

// handleQuote handles the /v1/quote endpoint.
func HandleQuote(w http.ResponseWriter, r *http.Request) {
	// Get the text and image parameters from the request
	text := r.URL.Query().Get("text")
	imageSource := r.URL.Query().Get("image")
	download := r.URL.Query().Get("download")

	if text == "" || imageSource == "" {
		http.Error(w, "Missing text or image parameter", http.StatusBadRequest)
		return
	}

	width, height := 1200, 1200
	resultImg, err := imageGenerator.Generate(text, imageSource, width, height)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error overlaying text: %v", err), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to image/jpeg
	w.Header().Set("Content-Type", "image/jpeg")
	if download == "true" {
		w.Header().Set("Content-Disposition", "attachment; filename=\"quote-share.jpg\"")
	}

	// Encode and write the image to the response
	if err := jpeg.Encode(w, resultImg, &jpeg.Options{Quality: 95}); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding image: %v", err), http.StatusInternalServerError)
		return
	}
}
