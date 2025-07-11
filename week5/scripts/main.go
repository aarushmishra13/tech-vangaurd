package main

import (
	"encoding/csv"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

var classMapping = map[string]int{
	"A10": 0, "A400M": 1, "AG600": 2, "AH64": 3, "An124": 4, "An22": 5, "An225": 6, "An72": 7,
	"AV8B": 8, "B1": 9, "B2": 10, "B21": 11, "B52": 12, "Be200": 13, "C130": 14, "C17": 15,
	"C2": 16, "C390": 17, "C5": 18, "CH47": 19, "CL415": 20, "E2": 21, "E7": 22, "EF2000": 23,
	"EMB314": 24, "F117": 25, "F14": 26, "F15": 27, "F16": 28, "F18": 29, "F22": 30, "F35": 31,
	"F4": 32, "H6": 33, "Il76": 34, "J10": 35, "J20": 36, "J35": 37, "J36": 38, "J50": 39,
	"JAS39": 40, "JF17": 41, "JH7": 42, "Ka27": 43, "Ka52": 44, "KAAN": 45, "KC135": 46,
	"KF21": 47, "KJ600": 48, "Mi24": 49, "Mi26": 50, "Mi28": 51, "Mi8": 52, "Mig29": 53,
	"Mig31": 54, "Mirage2000": 55, "MQ9": 56, "P3": 57, "Rafale": 58, "RQ4": 59, "SR71": 60,
	"Su24": 61, "Su25": 62, "Su34": 63, "Su57": 64, "TB001": 65, "TB2": 66, "Tornado": 67,
	"Tu160": 68, "Tu22M": 69, "Tu95": 70, "U2": 71, "UH60": 72, "US2": 73, "V22": 74, "V280": 75,
	"Vulcan": 76, "WZ10": 77, "WZ7": 78, "WZ9": 79, "XB70": 80, "Y20": 81, "YF23": 82,
	"Z10": 83, "Z19": 84,
}

func convertCSVToYOLO(csvPath, outputDir string, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(csvPath)
	if err != nil {
		log.Printf("Failed to open %s: %v\n", csvPath, err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil || len(rows) < 2 {
		log.Printf("Failed to read/parse CSV: %s\n", csvPath)
		return
	}

	var lines []string
	for _, row := range rows[1:] {
		width, _ := strconv.ParseFloat(row[1], 64)
		height, _ := strconv.ParseFloat(row[2], 64)
		className := row[3]
		xmin, _ := strconv.ParseFloat(row[4], 64)
		ymin, _ := strconv.ParseFloat(row[5], 64)
		xmax, _ := strconv.ParseFloat(row[6], 64)
		ymax, _ := strconv.ParseFloat(row[7], 64)

		classID, ok := classMapping[className]
		if !ok {
			log.Printf("Unknown class: %s in %s\n", className, csvPath)
			continue
		}

		xCenter := (xmin + xmax) / 2 / width
		yCenter := (ymin + ymax) / 2 / height
		bbWidth := (xmax - xmin) / width
		bbHeight := (ymax - ymin) / height

		lines = append(lines, fmt.Sprintf("%d %.6f %.6f %.6f %.6f", classID, xCenter, yCenter, bbWidth, bbHeight))
	}

	txtName := strings.TrimSuffix(filepath.Base(csvPath), filepath.Ext(csvPath)) + ".txt"
	outPath := filepath.Join(outputDir, txtName)
	os.WriteFile(outPath, []byte(strings.Join(lines, "\n")), 0644)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run convert.go <csv_dir> <output_dir>")
		return
	}

	csvDir := os.Args[1]
	outputDir := os.Args[2]
	os.MkdirAll(outputDir, 0755)

	files, err := os.ReadDir(csvDir)
	if err != nil {
		log.Fatal("Failed to read input directory:", err)
	}

	var wg sync.WaitGroup
	sem := make(chan struct{}, 16) // concurrency limit

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".csv") {
			continue
		}
		wg.Add(1)
		sem <- struct{}{}
		go func(f fs.DirEntry) {
			defer func() { <-sem }()
			csvPath := filepath.Join(csvDir, f.Name())
			convertCSVToYOLO(csvPath, outputDir, &wg)
		}(file)
	}

	wg.Wait()
	fmt.Println("✅ Conversion complete.")
}
