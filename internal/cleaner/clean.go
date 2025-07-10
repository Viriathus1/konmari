package cleaner

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

func CleanUp(directory string, olderThan int, dryRun bool) error {
	// get threshold time
	now := time.Now()
	threshold := now.AddDate(0, 0, -olderThan)

	// file walk through directories
	err := filepath.Walk(directory, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		if info.ModTime().Before(threshold) {
			if dryRun {
				fmt.Printf("[Dry Run] Would delete: %s\n", path)
			} else {
				fmt.Printf("Deleting: %s\n", path)
				err := os.Remove(path)
				if err != nil {
					fmt.Printf("Failed to delete: %s: %v\n", path, err)
				}
			}
		}

		return nil
	})

	return err
}
