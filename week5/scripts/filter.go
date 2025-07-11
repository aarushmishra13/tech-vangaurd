package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

var (
	out           = "./classification_dataset"
	in            = "./crop"
	aircraftTypes = []string{
		"A10", "A400M", "AG600", "AH64", "An124", "An22", "An225", "An72", "AV8B",
		"B1", "B2", "B21", "B52", "Be200", "C130", "C17", "C2", "C390", "C5",
		"CH47", "CL415", "E2", "E7", "EF2000", "EMB314", "F117", "F14", "F15",
		"F16", "F18", "F22", "F35", "F4", "H6", "Il76", "J10", "J20", "J35", "J36",
		"J50", "JAS39", "JF17", "JH7", "Ka27", "Ka52", "KAAN", "KC135", "KF21",
		"KJ600", "Mi24", "Mi26", "Mi28", "Mi8", "Mig29", "Mig31", "Mirage2000",
		"MQ9", "P3", "Rafale", "RQ4", "SR71", "Su24", "Su25", "Su34", "Su57",
		"TB001", "TB2", "Tornado", "Tu160", "Tu22M", "Tu95", "U2", "UH60", "US2",
		"V22", "V280", "Vulcan", "WZ10", "WZ7", "WZ9", "XB70", "Y20", "YF23", "Z10", "Z19",
	}
)

func main() {
	rand.Seed(time.Now().UnixNano())

	for _, at := range aircraftTypes {
		processFolder(at)
	}
}

func processFolder(class string) {
	srcDir := filepath.Join(in, class)
	trainDir := filepath.Join(out, "train", class)
	valDir := filepath.Join(out, "val", class)

	// Create output folders
	os.MkdirAll(trainDir, os.ModePerm)
	os.MkdirAll(valDir, os.ModePerm)

	files, err := filepath.Glob(filepath.Join(srcDir, "*.jpg"))
	if err != nil {
		fmt.Println("Failed to read images for class:", class)
		return
	}
	if len(files) == 0 {
		fmt.Println("⚠️ No images found for class:", class)
		return
	}

	// Shuffle files
	rand.Shuffle(len(files), func(i, j int) { files[i], files[j] = files[j], files[i] })

	split := int(0.8 * float64(len(files)))

	for i, src := range files {
		dstDir := trainDir
		if i >= split {
			dstDir = valDir
		}
		dst := filepath.Join(dstDir, filepath.Base(src))
		copyFile(src, dst)
	}
	fmt.Printf("✅ %s — %d images (train: %d, val: %d)\n", class, len(files), split, len(files)-split)
}

func copyFile(src, dst string) {
	input, err := ioutil.ReadFile(src)
	if err != nil {
		fmt.Println("❌ Failed to read:", src)
		return
	}
	err = ioutil.WriteFile(dst, input, fs.ModePerm)
	if err != nil {
		fmt.Println("❌ Failed to write:", dst)
	}
}
