package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"sort"
)

func main() {
	root := "classification_dataset/train"

	classCounts := make(map[string]int)

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && (filepath.Ext(d.Name()) == ".jpg" || filepath.Ext(d.Name()) == ".jpeg" || filepath.Ext(d.Name()) == ".png") {
			class := filepath.Base(filepath.Dir(path))
			classCounts[class]++
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error walking the directory:", err)
		return
	}

	// Sort results
	type classCount struct {
		Class string
		Count int
	}
	var sorted []classCount
	for k, v := range classCounts {
		sorted = append(sorted, classCount{k, v})
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Count < sorted[j].Count
	})

	// Print results
	fmt.Println("Image count per class:")
	for _, cc := range sorted {
		fmt.Printf("%-30s : %d\n", cc.Class, cc.Count)
	}
}
