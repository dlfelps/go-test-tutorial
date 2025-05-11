package calc

import (
	"testing"
)

// Benchmark Add operation
func BenchmarkAdd(b *testing.B) {
	c := New()

	// Run the Add method b.N times
	for i := 0; i < b.N; i++ {
		c.Add(float64(i), float64(i+1))
	}
}

// Benchmark Operation function with different operations
func BenchmarkOperation(b *testing.B) {
	operations := []string{"+", "-", "*", "/"}

	for _, op := range operations {
		// Create a sub-benchmark for each operation
		b.Run(op, func(b *testing.B) {
			c := New()

			for i := 0; i < b.N; i++ {
				// For division, avoid division by zero
				var val float64 = 1
				if i > 0 {
					val = float64(i)
				}

				c.PerformOperation(op, float64(i+10), val)
			}
		})
	}
}

// Simple benchmark for fixed operations
func BenchmarkFixedOperations(b *testing.B) {
	// Create a calculator to use for benchmarking
	c := New()

	// Reset the timer to exclude setup time
	b.ResetTimer()

	// Report memory allocations
	b.ReportAllocs()

	// Run the benchmark for b.N iterations
	for i := 0; i < b.N; i++ {
		c.Add(10, 20)
		c.Multiply(5, 10)
		c.Subtract(100, 50)
		c.Divide(100, 10)
	}
}
