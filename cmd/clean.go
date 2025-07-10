/*
Copyright © 2025 Viriathus1 viriathus1_dev@proton.me
*/
package cmd

import (
	"fmt"

	"github.com/Viriathus1/konmari/internal/cleaner"
	"github.com/spf13/cobra"
)

var (
	dir       string
	olderThan int
	dryRun    bool
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean up old files",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Cleaning %s files older than %d days (dry-run: %v)\n", dir, olderThan, dryRun)
		err := cleaner.CleanUp(dir, olderThan, dryRun)
		if err != nil {
			fmt.Printf("Error during cleanup: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)

	cleanCmd.Flags().StringVarP(&dir, "dir", "d", ".", "Target Directory")
	cleanCmd.Flags().IntVar(&olderThan, "days", 30, "Delete files older than this amount of days")
	cleanCmd.Flags().BoolVar(&dryRun, "dry-run", true, "Perform a dry run")
}
