package main

import "testing"

func TestCalculator(t *testing.T) {
        // TODO: Replace these values with actual test data
        input := "test"
        expected := "expected"
        
        // TODO: Call your function with the input
        actual := Calculator(input)
        
        // TODO: Check if the result matches what you expect
        if actual != expected {
                t.Errorf("Calculator(%v) = %v; expected %v", input, actual, expected)
        }
}
