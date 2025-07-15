/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

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
		}

		model, err = tea.NewProgram(method.NewListView(fpm.SelectedPaths())).Run()
		if err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}

		lvm, ok := model.(method.ListViewModel)
		if !ok {
			fmt.Println("Couldn't extract selected paths.")
		}

		for i, path := range lvm.SelectedPaths() {
			fmt.Printf("%d: %s\n", i, path)
		}
	},
}

func init() {
	rootCmd.AddCommand(methodCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// methodCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// methodCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
