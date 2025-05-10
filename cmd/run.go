package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"go-test-tutorial/internal/utils"

	"github.com/spf13/cobra"
)

var testFile string
var testDir string
var verbose bool

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run Go tests and show results",
	Long: `Run Go tests and show the results with helpful explanations.
You can run tests for a specific file or directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		if testFile == "" && testDir == "" {
			testDir = "."
		}
		
		runTests()
	},
}

func init() {
	runCmd.Flags().StringVarP(&testFile, "file", "f", "", "Test file to run")
	runCmd.Flags().StringVarP(&testDir, "dir", "d", "", "Directory containing tests to run")
	runCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Run tests in verbose mode")
}

func runTests() {
	var args []string
	args = append(args, "test")
	
	if verbose {
		args = append(args, "-v")
	}
	
	// If a specific file is provided
	if testFile != "" {
		// Clean up the file name
		if !filepath.IsAbs(testFile) {
			testFile = filepath.Join(".", testFile)
		}
		
		dir := filepath.Dir(testFile)
		
		// Extract the test name pattern if possible
		fileBase := filepath.Base(testFile)
		
		// Check if the file exists
		if _, err := os.Stat(testFile); os.IsNotExist(err) {
			fmt.Printf("Error: Test file '%s' not found\n", testFile)
			return
		}
		
		// Extract test function names
		content, err := utils.ReadFile(testFile)
		if err != nil {
			fmt.Printf("Error reading test file: %v\n", err)
			return
		}
		
		testFuncs := utils.ExtractTestFunctionNames(content)
		if len(testFuncs) > 0 {
			pattern := "^("
			for i, f := range testFuncs {
				if i > 0 {
					pattern += "|"
				}
				pattern += f
			}
			pattern += ")$"
			
			args = append(args, "-run", pattern)
		}
		
		args = append(args, dir)
	} else if testDir != "" {
		// Run tests in the specified directory
		args = append(args, testDir)
	}
	
	fmt.Println("Running tests...")
	fmt.Printf("Command: go %s\n\n", utils.JoinArgs(args))
	
	output, err := utils.ExecuteCommand("go", args...)
	fmt.Println(output)
	
	if err != nil {
		fmt.Println("\nTest execution failed. This might be because:")
		fmt.Println("1. There are actual test failures")
		fmt.Println("2. The test file has syntax errors")
		fmt.Println("3. The test depends on code that doesn't exist yet")
		fmt.Println("\nCheck the output above for specific error messages.")
	} else {
		fmt.Println("\nAll tests passed successfully!")
		fmt.Println("\nNext steps you might want to try:")
		fmt.Println("1. Add more test cases to cover edge cases")
		fmt.Println("2. Learn about benchmarks with 'gotest-learn learn 5'")
		fmt.Println("3. Explore table-driven tests with 'gotest-learn learn 4'")
	}
}
