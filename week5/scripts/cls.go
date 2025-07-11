package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	imageDir     = "images"
	labelDir     = "labels"
	outputImages = "./training_data/images"
	outputLabels = "./training_data/labels"
	trainSplit   = 0.8
)

func main() {
	// seed random for shuffle
	rand.Seed(time.Now().UnixNano())

	// get all .jpg files
	imageFiles, err := filepath.Glob(filepath.Join(imageDir, "*.jpg"))
	if err != nil {
		log.Fatal("Failed to list images:", err)
	}

	// shuffle
	rand.Shuffle(len(imageFiles), func(i, j int) {
		imageFiles[i], imageFiles[j] = imageFiles[j], imageFiles[i]
	})

	// split
	splitIdx := int(float64(len(imageFiles)) * trainSplit)
	train := imageFiles[:splitIdx]
	val := imageFiles[splitIdx:]

	// create folders
	dirs := []string{
		filepath.Join(outputImages, "train"),
		filepath.Join(outputImages, "val"),
		filepath.Join(outputLabels, "train"),
		filepath.Join(outputLabels, "val"),
	}
	for _, d := range dirs {
		if err := os.MkdirAll(d, 0755); err != nil {
			log.Fatal("Failed to create dir:", d)
		}
	}

	// copy function
	copyPair := func(imgPath, splitType string) {
		base := strings.TrimSuffix(filepath.Base(imgPath), filepath.Ext(imgPath))
		labelPath := filepath.Join(labelDir, base+".txt")

		imgOut := filepath.Join(outputImages, splitType, filepath.Base(imgPath))
		lblOut := filepath.Join(outputLabels, splitType, base+".txt")

		if _, err := os.Stat(labelPath); err != nil {
			log.Printf("⚠️  Skipping %s: label not found", imgPath)
			return
		}

		copyFileS(imgPath, imgOut)
		copyFileS(labelPath, lblOut)
	}

	// copy files
	for _, f := range train {
		copyPair(f, "train")
	}
	for _, f := range val {
		copyPair(f, "val")
	}

	fmt.Printf("✅ Done. Train: %d | Val: %d\n", len(train), len(val))
}

func copyFileS(src, dst string) {
	in, err := os.Open(src)
	if err != nil {
		log.Printf("Failed to open %s: %v", src, err)
		return
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		log.Printf("Failed to create %s: %v", dst, err)
		return
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		log.Printf("Failed to copy from %s to %s: %v", src, dst, err)
	}
}
