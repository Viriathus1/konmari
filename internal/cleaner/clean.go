package cleaner

import (
	"fmt"
	"io/fs"
	"math"
	"os"
	"path/filepath"
	"time"
)

func CleanUp(directory string, olderThan int, dryRun bool) error {
	// get threshold time
	now := time.Now()
	threshold := now.AddDate(0, 0, -int(math.Abs(float64(olderThan))))
	paths := []string{}

	// file walk through directories
	err := filepath.Walk(directory, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		if info.ModTime().Before(threshold) {
			paths = append(paths, path)
		}

		return nil
	})

	if err == nil {
		DeleteFiles(paths, dryRun)
	}

	return err
}

func DeleteFiles(paths []string, dryRun bool) {
	for _, path := range paths {
		if dryRun {
			fmt.Printf("[Dry Run] Would delete: %s\n", path)
		} else {
			fmt.Printf("Deleting: %s\n", path)
			err := os.Remove(path)
			if err != nil {
				fmt.Printf("Failed to delete: %s: %v", path, err)
			}
		}
	}
}
