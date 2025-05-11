package calc

import (
	"testing"
)

func TestBasicOperations(t *testing.T) {
	// Create a new calculator for each test
	c := New()
	
	// Test addition
	if got := c.Add(2, 3); got != 5 {
		t.Errorf("Add(2, 3) = %f; want 5", got)
	}
	
	// Test subtraction
	if got := c.Subtract(10, 4); got != 6 {
		t.Errorf("Subtract(10, 4) = %f; want 6", got)
	}
	
	// Test multiplication
	if got := c.Multiply(3, 4); got != 12 {
		t.Errorf("Multiply(3, 4) = %f; want 12", got)
	}
	
	// Test division
	if got, err := c.Divide(12, 4); err != nil || got != 3 {
		t.Errorf("Divide(12, 4) = %f, %v; want 3, nil", got, err)
	}
	
	// Test division by zero
	if _, err := c.Divide(5, 0); err == nil {
		t.Error("Divide(5, 0) did not return an error")
	}
}

func TestTableDriven(t *testing.T) {
	// Define a table of test cases
	testCases := []struct {
		name     string
		op       string
		a, b     float64
		expected float64
		wantErr  bool
	}{
		{"addition", "+", 2, 3, 5, false},
		{"subtraction", "-", 5, 2, 3, false},
		{"multiplication", "*", 4, 5, 20, false},
		{"division", "/", 10, 2, 5, false},
		{"division by zero", "/", 10, 0, 0, true},
		{"unknown operation", "?", 10, 5, 0, true},
	}
	
	// Run each test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := New() // Fresh calculator for each subtest
			
			result, err := c.PerformOperation(tc.op, tc.a, tc.b)
			
			// Check error expectation
			if tc.wantErr && err == nil {
				t.Errorf("Expected an error for %s(%f, %f)", tc.op, tc.a, tc.b)
			}
			if !tc.wantErr && err != nil {
				t.Errorf("Unexpected error for %s(%f, %f): %v", tc.op, tc.a, tc.b, err)
			}
			
			// If we don't expect an error, check the result
			if !tc.wantErr && result != tc.expected {
				t.Errorf("%s(%f, %f) = %f; want %f", tc.op, tc.a, tc.b, result, tc.expected)
			}
			
			// Test that memory is updated
			if !tc.wantErr && c.GetMemory() != result {
				t.Errorf("Memory = %f; want %f", c.GetMemory(), result)
			}
		})
	}
}

func TestMemory(t *testing.T) {
	c := New()
	
	// Memory should start at 0
	if mem := c.GetMemory(); mem != 0 {
		t.Errorf("Initial memory = %f; want 0", mem)
	}
	
	// Memory should be updated after operations
	c.Add(5, 10)
	if mem := c.GetMemory(); mem != 15 {
		t.Errorf("Memory after Add = %f; want 15", mem)
	}
	
	// Test Clear
	c.ClearMemory()
	if mem := c.GetMemory(); mem != 0 {
		t.Errorf("Memory after Clear = %f; want 0", mem)
	}
}