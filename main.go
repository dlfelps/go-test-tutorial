package main

import (
        "fmt"
        "os"

        "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
        Use:   "gotest-learn",
        Short: "An interactive tutorial for learning Go testing",
        Long: `A command-line educational tool for teaching beginners how to use 
Go's testing framework through interactive tutorials and examples.

Learn concepts of Go testing step-by-step with practical examples.`,
        Run: func(cmd *cobra.Command, args []string) {
                // Display help information if no subcommand is provided
                cmd.Help()
        },
}

// Execute executes the root command.
func Execute() error {
        return rootCmd.Execute()
}

func init() {
        // Add subcommands to the root command
        rootCmd.AddCommand(learnCmd)
        rootCmd.AddCommand(generateCmd)
        rootCmd.AddCommand(runCmd)

        // Custom help template
        rootCmd.SetHelpTemplate(`
Go Testing Tutorial
==================

{{ .Long }}

Available Commands:
{{ range .Commands }}{{ if (or .IsAvailableCommand (eq .Name "help")) }}
  {{ rpad .Name .NamePadding }} {{ .Short }}{{ end }}{{ end }}

Usage:
  {{ .UseLine }}

Flags:
{{ .LocalFlags.FlagUsages | trimTrailingWhitespaces }}

Use "{{ .CommandPath }} [command] --help" for more information about a command.
`)

        // Handle errors
        rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
                // Skip validation for help command
                if cmd.CalledAs() == "help" {
                        return nil
                }
                
                return nil
        }
}

// LEARN COMMAND
var topic string
var interactive bool

var learnCmd = &cobra.Command{
        Use:   "learn [topic number]",
        Short: "Learn about Go testing concepts",
        Long: `Learn about Go testing concepts with explanations and examples.
You can specify a topic number or run in interactive mode to go through all topics.`,
        Run: func(cmd *cobra.Command, args []string) {
                if len(args) > 0 {
                        // Handle topic number
                        fmt.Printf("You selected topic: %s\n", args[0])
                } else if interactive {
                        fmt.Println("Starting interactive tutorial...")
                } else {
                        fmt.Println("Available topics to learn:")
                        fmt.Println("1. Introduction to Go Testing")
                        fmt.Println("2. Writing Basic Tests")
                        fmt.Println("3. Test Assertions and Comparisons")
                        fmt.Println("4. Table-Driven Tests")
                        fmt.Println("5. Benchmarks")
                        fmt.Println("6. Subtests and Test Organization")
                        fmt.Println("7. TestMain Function")
                        fmt.Println("8. Test Helpers")
                        fmt.Println("\nUse 'gotest-learn learn <topic number>' to learn about a specific topic")
                        fmt.Println("Use 'gotest-learn learn --interactive' to start an interactive tutorial")
                }
        },
}

// GENERATE COMMAND
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
                        fmt.Println("Available template types:")
                        fmt.Println("  - basic: Basic test template for simple function testing")
                        fmt.Println("  - table: Table-driven test template for testing multiple cases")
                        fmt.Println("  - benchmark: Benchmark template for measuring performance")
                        fmt.Println("  - subtest: Subtest template for organizing tests into groups")
                        fmt.Println("\nUsage example: gotest-learn generate --type basic --file mycode")
                        return
                }
                
                if fileName == "" {
                        fmt.Println("Error: Please provide a file name using the --file flag")
                        return
                }
                
                fmt.Printf("Generating %s test template for file %s in directory %s\n", templateType, fileName, outputDir)
        },
}

// RUN COMMAND
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
                
                fmt.Println("Running tests...")
                if testFile != "" {
                        fmt.Printf("Testing file: %s\n", testFile)
                } else {
                        fmt.Printf("Testing directory: %s\n", testDir)
                }
                
                if verbose {
                        fmt.Println("Verbose mode enabled")
                }
        },
}

func init() {
        // Learn command flags
        learnCmd.Flags().BoolVarP(&interactive, "interactive", "i", false, "Start an interactive tutorial")

        // Generate command flags
        generateCmd.Flags().StringVarP(&templateType, "type", "t", "", "Type of test template (basic, table, benchmark)")
        generateCmd.Flags().StringVarP(&fileName, "file", "f", "", "Name of the file to test (without _test.go suffix)")
        generateCmd.Flags().StringVarP(&outputDir, "output", "o", ".", "Output directory for the generated test file")

        // Run command flags
        runCmd.Flags().StringVarP(&testFile, "file", "f", "", "Test file to run")
        runCmd.Flags().StringVarP(&testDir, "dir", "d", "", "Directory containing tests to run")
        runCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Run tests in verbose mode")
}

func main() {
        if err := rootCmd.Execute(); err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
}
