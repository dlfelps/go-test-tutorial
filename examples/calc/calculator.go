package calc

import (
	"errors"
	"fmt"
)

// Calculator provides basic arithmetic operations
type Calculator struct {
	// Memory stores the last calculated result
	Memory float64
}

// Add performs addition of two numbers
func (c *Calculator) Add(a, b float64) float64 {
	c.Memory = a + b
	return c.Memory
}

// Subtract performs subtraction of two numbers
func (c *Calculator) Subtract(a, b float64) float64 {
	c.Memory = a - b
	return c.Memory
}

// Multiply performs multiplication of two numbers
func (c *Calculator) Multiply(a, b float64) float64 {
	c.Memory = a * b
	return c.Memory
}

// Divide performs division of two numbers
func (c *Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	c.Memory = a / b
	return c.Memory, nil
}

// ClearMemory resets the calculator's memory
func (c *Calculator) ClearMemory() {
	c.Memory = 0
}

// GetMemory returns the current value in memory
func (c *Calculator) GetMemory() float64 {
	return c.Memory
}

// New creates a new Calculator instance
func New() *Calculator {
	return &Calculator{Memory: 0}
}

// PerformOperation handles different operations with a common interface
func (c *Calculator) PerformOperation(op string, a, b float64) (float64, error) {
	switch op {
	case "+":
		return c.Add(a, b), nil
	case "-":
		return c.Subtract(a, b), nil
	case "*":
		return c.Multiply(a, b), nil
	case "/":
		return c.Divide(a, b)
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}