package imageGenerator

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"image"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/patrickmn/go-cache"
)

var imageLoader = newImageLoader(5*time.Minute, 10*time.Minute)

type ImageLoader struct {
	cache *cache.Cache
}

func newImageLoader(defaultExpiration, cleanupInterval time.Duration) *ImageLoader {
	return &ImageLoader{
		cache: cache.New(defaultExpiration, cleanupInterval),
	}
}

// LoadImage loads an image from cache or URL
func (il *ImageLoader) loadImage(url string) ([]byte, error) {
	key := generateKey(url)

	// Try to get the image from cache
	if cachedImage, found := il.cache.Get(key); found {
		return cachedImage.([]byte), nil
	}

	// If not in cache, download the image
	image, err := downloadImage(url)
	if err != nil {
		return nil, err
	}

	// Store the image in cache
	il.cache.Set(key, image, cache.DefaultExpiration)

	return image, nil
}

// generateKey creates a SHA1 hash of the URL
func generateKey(url string) string {
	h := sha1.New()
	h.Write([]byte(url))
	return hex.EncodeToString(h.Sum(nil))
}

// downloadImage downloads an image from a URL
func downloadImage(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

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
	imgBytes, err := imageLoader.loadImage(url)

	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		return nil, err
	}

	return img, nil
}
