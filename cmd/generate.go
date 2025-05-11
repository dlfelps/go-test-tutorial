package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"go-test-tutorial/internal/tutorial"
	"go-test-tutorial/internal/utils"

	"github.com/spf13/cobra"
)

var templateType string
var fileName string
var outputDir string

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate Go test files",
	Long: `Generate Go test template files based on different testing patterns.
You can choose from various templates like basic test, table-driven test, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		if templateType == "" {
			listTemplates()
			return
		}

		if fileName == "" {
			fmt.Println("Error: Please provide a file name using the --file flag")
			return
		}

		generateTestFile()
	},
}

func init() {
	generateCmd.Flags().StringVarP(&templateType, "type", "t", "", "Type of test template (basic, table, benchmark)")
	generateCmd.Flags().StringVarP(&fileName, "file", "f", "", "Name of the file to test (without _test.go suffix)")
	generateCmd.Flags().StringVarP(&outputDir, "output", "o", ".", "Output directory for the generated test file")
}

func listTemplates() {
	fmt.Println("Available template types:")
	for _, tmpl := range tutorial.GetAllTemplates() {
		fmt.Printf("  - %s: %s\n", tmpl.Name, tmpl.Description)
	}
	fmt.Println("\nUsage example: gotest-learn generate --type basic --file mycode")
}

func generateTestFile() {
	// Validate template type
	template, found := tutorial.GetTemplateByName(templateType)
	if !found {
		fmt.Printf("Error: Template type '%s' not found\n", templateType)
		listTemplates()
		return
	}

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		return
	}

	// Clean up the filename
	fileName = strings.TrimSuffix(fileName, ".go")
	fileName = strings.TrimSuffix(fileName, "_test")

	// Prepare output file path
	outputFile := filepath.Join(outputDir, fileName+"_test.go")

	// Check if file already exists
	if _, err := os.Stat(outputFile); err == nil {
		fmt.Printf("File '%s' already exists. Overwrite? (y/N): ", outputFile)
		var response string
		fmt.Scanln(&response)
		if strings.ToLower(response) != "y" {
			fmt.Println("Operation cancelled")
			return
		}
	}

	// Try to read the source file to get package name and function names
	sourceFile := filepath.Join(outputDir, fileName+".go")
	packageName := "main" // default

	sourceExists := false
	if _, err := os.Stat(sourceFile); err == nil {
		sourceExists = true
		content, err := utils.ReadFile(sourceFile)
		if err == nil {
			// Extract package name from source
			pkgLines := utils.ExtractPackageName(content)
			if pkgLines != "" {
				packageName = pkgLines
			}
		}
	}

	// Generate the content
	content := template.GetContent(packageName, fileName)

	// Write the file
	if err := utils.WriteFile(outputFile, content); err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}

	fmt.Printf("Successfully generated test file: %s\n", outputFile)

	if !sourceExists {
		fmt.Printf("\nNotice: Source file '%s' not found. You might want to create it first.\n", sourceFile)
	}

	fmt.Println("\nNext steps:")
	fmt.Println("1. Edit the generated test file to match your actual code")
	fmt.Println("2. Run the test with 'gotest-learn run --file " + fileName + "_test.go'")
	fmt.Println("   or with Go's standard test command: 'go test " + outputDir + "'")
}
