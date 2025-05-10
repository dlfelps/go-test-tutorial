package utils

import (
        "fmt"
        "os"
        "path/filepath"
        "strings"
)

// ReadFile reads the contents of a file
func ReadFile(path string) (string, error) {
        data, err := os.ReadFile(path)
        if err != nil {
                return "", fmt.Errorf("failed to read file %s: %w", path, err)
        }
        return string(data), nil
}

// WriteFile writes content to a file
func WriteFile(path string, content string) error {
        // Create directory if needed
        dir := filepath.Dir(path)
        if err := os.MkdirAll(dir, 0755); err != nil {
                return fmt.Errorf("failed to create directory %s: %w", dir, err)
        }
        
        // Write the file
        err := os.WriteFile(path, []byte(content), 0644)
        if err != nil {
                return fmt.Errorf("failed to write file %s: %w", path, err)
        }
        
        return nil
}

// FileExists checks if a file exists
func FileExists(path string) bool {
        _, err := os.Stat(path)
        return err == nil
}

// CreateTempDir creates a temporary directory for testing
func CreateTempDir(prefix string) (string, error) {
        tempDir, err := os.MkdirTemp("", prefix)
        if err != nil {
                return "", fmt.Errorf("failed to create temp directory: %w", err)
        }
        return tempDir, nil
}

// RemoveTempDir removes a temporary directory
func RemoveTempDir(path string) error {
        return os.RemoveAll(path)
}

// GetTestFiles returns a list of Go test files in a directory
func GetTestFiles(dir string) ([]string, error) {
        var testFiles []string
        
        err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
                if err != nil {
                        return err
                }
                
                if !info.IsDir() && filepath.Ext(path) == ".go" && strings.HasSuffix(path, "_test.go") {
                        testFiles = append(testFiles, path)
                }
                
                return nil
        })
        
        if err != nil {
                return nil, fmt.Errorf("failed to list test files: %w", err)
        }
        
        return testFiles, nil
}
