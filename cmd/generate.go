package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	propertiesFile string
	csvFile        string
	outputFile     string
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a .env file from a template and CSV file",
	Long: `Generate a .env file by replacing placeholders in a .properties template with values from a CSV file.

Example:
  csv2env generate -t example/template.properties -c example/client1.csv -o .env`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Validate required flags
		if propertiesFile == "" || csvFile == "" {
			return fmt.Errorf("both template and csv flags are required")
		}

		// Read the properties template file
		templateContent, err := readPropertiesTemplate(propertiesFile)
		if err != nil {
			return fmt.Errorf("error reading properties template: %v", err)
		}

		// Read the CSV file
		propertyValues, err := readCSVFile(csvFile)
		if err != nil {
			return fmt.Errorf("error reading CSV file: %v", err)
		}

		// Generate the .env file
		err = generateEnvFile(templateContent, propertyValues, outputFile)
		if err != nil {
			return fmt.Errorf("error generating .env file: %v", err)
		}

		fmt.Printf("Successfully generated %s\n", outputFile)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Define flags for the generate command
	generateCmd.Flags().StringVarP(&propertiesFile, "template", "t", "", "Path to the .properties template file (required)")
	generateCmd.Flags().StringVarP(&csvFile, "csv", "c", "", "Path to the CSV file containing property values (required)")
	generateCmd.Flags().StringVarP(&outputFile, "output", "o", ".env", "Path to the output .env file")

	// Mark required flags
	generateCmd.MarkFlagRequired("template")
	generateCmd.MarkFlagRequired("csv")
}

// readPropertiesTemplate reads the .properties template file and returns its content
func readPropertiesTemplate(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read properties template file: %w", err)
	}
	return string(data), nil
}

// readCSVFile reads the CSV file and returns a map of property keys to values
func readCSVFile(filePath string) (map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV file: %w", err)
	}

	if len(records) < 2 {
		return nil, fmt.Errorf("CSV file must contain at least a header row and a data row")
	}

	// First row is assumed to be headers (property keys)
	headers := records[0]
	// Second row contains the values
	values := records[1]

	if len(headers) != len(values) {
		return nil, fmt.Errorf("number of headers and values in CSV do not match")
	}

	propertyValues := make(map[string]string)
	for i, header := range headers {
		propertyValues[header] = values[i]
	}

	return propertyValues, nil
}

// generateEnvFile generates a .env file by replacing placeholders in the template with values from the CSV
func generateEnvFile(templateContent string, propertyValues map[string]string, outputPath string) error {
	// Create output directory if it doesn't exist
	outputDir := filepath.Dir(outputPath)
	if outputDir != "." {
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}
	}

	// Replace placeholders in the template
	result := templateContent
	for key, value := range propertyValues {
		placeholder := fmt.Sprintf("#%s#", key)
		result = strings.ReplaceAll(result, placeholder, value)
	}

	// Write the result to the output file
	err := os.WriteFile(outputPath, []byte(result), 0644)
	if err != nil {
		return fmt.Errorf("failed to write .env file: %w", err)
	}

	return nil
}
