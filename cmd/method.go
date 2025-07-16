/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/Viriathus1/konmari/internal/cleaner"
	"github.com/Viriathus1/konmari/internal/method"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// methodCmd represents the method command
var methodCmd = &cobra.Command{
	Use:   "method",
	Short: "Launch the interactive Konmari Method",
	Run: func(cmd *cobra.Command, args []string) {
		model, err := tea.NewProgram(method.NewFilePicker()).Run()
		if err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}

		fpm, ok := model.(method.FilePickerModel)
		if !ok {
			fmt.Println("Couldn't extract selected paths.")
			os.Exit(1)
		}

		potentialPaths := fpm.SelectedPaths()
		if len(potentialPaths) <= 0 {
			fmt.Println("No files selected.")
			return
		}

		model, err = tea.NewProgram(method.NewListView(potentialPaths)).Run()
		if err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}

		lvm, ok := model.(method.ListViewModel)
		if !ok {
			fmt.Println("Couldn't extract selected paths.")
			os.Exit(1)
		}

		cleaner.DeleteFiles(lvm.SelectedPaths(), false)
	},
}

func init() {
	rootCmd.AddCommand(methodCmd)
}
