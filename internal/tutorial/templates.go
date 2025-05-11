package tutorial

import (
	"strings"
)

// Template represents a test template that can be generated
type Template struct {
	Name        string
	Description string
	Content     string
}

// GetContent returns the template content with the package name
// and function name substituted
func (t *Template) GetContent(packageName, fileName string) string {
	content := t.Content

	// Replace package name
	content = strings.Replace(content, "{{PACKAGE}}", packageName, -1)

	// Replace function name based on the file name
	funcName := strings.Title(strings.TrimSuffix(fileName, ".go"))
	content = strings.Replace(content, "{{FUNCTION}}", funcName, -1)

	return content
}

// GetAllTemplates returns all available test templates
func GetAllTemplates() []Template {
	return []Template{
		getBasicTemplate(),
		getTableDrivenTemplate(),
		getBenchmarkTemplate(),
		getSubtestTemplate(),
	}
}

// GetTemplateByName returns a template by its name
func GetTemplateByName(name string) (Template, bool) {
	for _, tmpl := range GetAllTemplates() {
		if tmpl.Name == name {
			return tmpl, true
		}
	}
	return Template{}, false
}

func getBasicTemplate() Template {
	return Template{
		Name:        "basic",
		Description: "Basic test template for simple function testing",
		Content: `package {{PACKAGE}}

import "testing"

func Test{{FUNCTION}}(t *testing.T) {
        // TODO: Replace these values with actual test data
        input := "test"
        expected := "expected"
        
        // TODO: Call your function with the input
        actual := {{FUNCTION}}(input)
        
        // TODO: Check if the result matches what you expect
        if actual != expected {
                t.Errorf("{{FUNCTION}}(%v) = %v; expected %v", input, actual, expected)
        }
}
`,
	}
}

func getTableDrivenTemplate() Template {
	return Template{
		Name:        "table",
		Description: "Table-driven test template for testing multiple cases",
		Content: `package {{PACKAGE}}

import "testing"

func Test{{FUNCTION}}(t *testing.T) {
        // Define test cases
        testCases := []struct {
                name     string
                input    string // TODO: Change to actual input type
                expected string // TODO: Change to actual output type
        }{
                {
                        name:     "case 1",
                        input:    "test input 1",
                        expected: "expected output 1",
                },
                {
                        name:     "case 2",
                        input:    "test input 2",
                        expected: "expected output 2",
                },
                // TODO: Add more test cases here
        }
        
        // Run all test cases
        for _, tc := range testCases {
                t.Run(tc.name, func(t *testing.T) {
                        // TODO: Call your function with the test case input
                        actual := {{FUNCTION}}(tc.input)
                        
                        // TODO: Check if the result matches the expected output
                        if actual != tc.expected {
                                t.Errorf("{{FUNCTION}}(%v) = %v; expected %v", 
                                         tc.input, actual, tc.expected)
                        }
                })
        }
}
`,
	}
}

func getBenchmarkTemplate() Template {
	return Template{
		Name:        "benchmark",
		Description: "Benchmark template for measuring performance",
		Content: `package {{PACKAGE}}

import "testing"

// First, a basic test to verify functionality
func Test{{FUNCTION}}(t *testing.T) {
        // TODO: Replace with actual test data
        input := "test"
        expected := "expected"
        
        actual := {{FUNCTION}}(input)
        
        if actual != expected {
                t.Errorf("{{FUNCTION}}(%v) = %v; expected %v", input, actual, expected)
        }
}

// Benchmark to measure performance
func Benchmark{{FUNCTION}}(b *testing.B) {
        // TODO: Replace with actual benchmark data
        input := "test"
        
        // Reset the timer to exclude setup time
        b.ResetTimer()
        
        // Run the function b.N times
        for i := 0; i < b.N; i++ {
                {{FUNCTION}}(input)
        }
}

// Benchmark with different input sizes
func Benchmark{{FUNCTION}}Sizes(b *testing.B) {
        // TODO: Define your benchmark cases with different input sizes
        benchCases := []struct {
                name  string
                input string // TODO: Change to actual input type
        }{
                {"small", "small input"},
                {"medium", "medium input that is a bit longer"},
                {"large", "large input that is much longer than the others and tests performance with bigger data"},
        }
        
        for _, bc := range benchCases {
                b.Run(bc.name, func(b *testing.B) {
                        // Report memory allocations
                        b.ReportAllocs()
                        
                        for i := 0; i < b.N; i++ {
                                {{FUNCTION}}(bc.input)
                        }
                })
        }
}
`,
	}
}

func getSubtestTemplate() Template {
	return Template{
		Name:        "subtest",
		Description: "Subtest template for organizing tests into groups",
		Content: `package {{PACKAGE}}

import "testing"

func Test{{FUNCTION}}(t *testing.T) {
        // Group tests by category
        t.Run("Category1", func(t *testing.T) {
                // Subtest 1
                t.Run("TestCase1", func(t *testing.T) {
                        // TODO: Replace with actual test data
                        input := "test 1"
                        expected := "expected 1"
                        
                        actual := {{FUNCTION}}(input)
                        
                        if actual != expected {
                                t.Errorf("{{FUNCTION}}(%v) = %v; expected %v", 
                                         input, actual, expected)
                        }
                })
                
                // Subtest 2
                t.Run("TestCase2", func(t *testing.T) {
                        // TODO: Replace with actual test data
                        input := "test 2"
                        expected := "expected 2"
                        
                        actual := {{FUNCTION}}(input)
                        
                        if actual != expected {
                                t.Errorf("{{FUNCTION}}(%v) = %v; expected %v", 
                                         input, actual, expected)
                        }
                })
        })
        
        // Another group of tests
        t.Run("Category2", func(t *testing.T) {
                // Setup for this category
                // TODO: Add any setup code here
                
                // You can define helper functions inside the test
                runTest := func(t *testing.T, input, expected string) {
                        t.Helper() // Mark as helper function for better error reporting
                        
                        actual := {{FUNCTION}}(input)
                        
                        if actual != expected {
                                t.Errorf("{{FUNCTION}}(%v) = %v; expected %v", 
                                         input, actual, expected)
                        }
                }
                
                // Subtest 3
                t.Run("TestCase3", func(t *testing.T) {
                        runTest(t, "test 3", "expected 3")
                })
                
                // Subtest 4
                t.Run("TestCase4", func(t *testing.T) {
                        runTest(t, "test 4", "expected 4")
                })
                
                // More subtests can be added here
        })
}
`,
	}
}
