package cmd

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
		
		// Check if Go is installed
		_, err := os.Stat(os.Getenv("GOPATH"))
		if os.IsNotExist(err) {
			return fmt.Errorf("GO environment not properly set up. Make sure Go is installed and GOPATH is configured")
		}
		
		return nil
	}
}
