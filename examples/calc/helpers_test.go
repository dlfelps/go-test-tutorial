package calc

import (
	"testing"
)

// Helper function to check if results match expectations
func checkResult(t *testing.T, got, want float64, name string) {
	t.Helper() // Marks this as a helper function for better error reporting
	
	if got != want {
		t.Errorf("%s: got %f, want %f", name, got, want)
	}
}

// Helper function to check if errors are as expected
func checkError(t *testing.T, err error, wantErr bool, name string) {
	t.Helper()
	
	if wantErr && err == nil {
		t.Errorf("%s: expected an error but got nil", name)
	}
	
	if !wantErr && err != nil {
		t.Errorf("%s: unexpected error: %v", name, err)
	}
}

// Helper to create a calculator with some operations executed
func setupCalculator(t *testing.T, initialValues ...float64) *Calculator {
	t.Helper()
	
	c := New()
	
	// If initial values are provided, add them to the calculator
	if len(initialValues) > 0 {
		c.Add(initialValues[0], 0) // Store the first value in memory
	}
	
	return c
}

// TestWithHelpers demonstrates using test helpers to simplify tests
func TestWithHelpers(t *testing.T) {
	// Group related tests
	t.Run("BasicOperations", func(t *testing.T) {
		// Setup shared by all tests in this group
		c := setupCalculator(t)
		
		// Test addition with helper
		result := c.Add(5, 10)
		checkResult(t, result, 15, "Add(5, 10)")
		
		// Test multiplication with helper
		result = c.Multiply(4, 5)
		checkResult(t, result, 20, "Multiply(4, 5)")
		
		// Test division with helper
		result, err := c.Divide(20, 4)
		checkError(t, err, false, "Divide(20, 4)")
		checkResult(t, result, 5, "Divide(20, 4)")
		
		// Test division by zero with helper
		result, err = c.Divide(10, 0)
		checkError(t, err, true, "Divide(10, 0)")
	})
	
	// Another group of tests
	t.Run("MemoryOperations", func(t *testing.T) {
		// Parallel test execution
		t.Parallel()
		
		c := setupCalculator(t, 100) // Start with 100 in memory
		
		// Check memory
		checkResult(t, c.GetMemory(), 100, "Initial memory")
		
		// Perform some operations
		c.Add(c.GetMemory(), 50)
		checkResult(t, c.GetMemory(), 150, "After Add")
		
		c.ClearMemory()
		checkResult(t, c.GetMemory(), 0, "After Clear")
	})
}