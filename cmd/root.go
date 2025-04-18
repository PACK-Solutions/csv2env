package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "csv2env",
	Short: "A tool to generate .env files from templates and CSV files",
	Long: `.env Generator

A command-line tool that generates .env files from a .properties template and CSV files containing property values.
This tool takes a .properties template file with placeholders in the format #PLACEHOLDER# and replaces them with values from a CSV file.
The CSV file should have property keys in the header row and values in the subsequent row.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
}
