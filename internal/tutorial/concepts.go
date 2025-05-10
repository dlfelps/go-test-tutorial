package tutorial

// Concept represents a Go testing concept with explanation and examples
type Concept struct {
	Title            string
	Description      string
	Example          string
	PracticalExercise string
}

// GetAllConcepts returns all available testing concepts
func GetAllConcepts() []Concept {
	return []Concept{
		getIntroductionConcept(),
		getBasicTestsConcept(),
		getAssertionsConcept(),
		getTableDrivenTestsConcept(),
		getBenchmarkConcept(),
		getSubtestsConcept(),
		getTestMainConcept(),
		getTestHelpersConcept(),
	}
}

func getIntroductionConcept() Concept {
	return Concept{
		Title: "Introduction to Go Testing",
		Description: `Go has a built-in testing framework that makes it easy to write and run tests.
The testing package provides the tools you need to write unit tests, benchmarks,
and examples for your Go code.

Key points about Go testing:

1. Tests are functions in files that end with "_test.go"
2. Test functions start with "Test" followed by a name that starts with a capital letter
3. Test functions take a single parameter: t *testing.T
4. Tests run with the "go test" command
5. The testing framework reports whether tests pass or fail

Go's test files are normal Go files that live alongside your code but are only
compiled and run during testing.`,
		Example: `// math.go
package math

func Add(a, b int) int {
	return a + b
}

// math_test.go
package math

import "testing"

func TestAdd(t *testing.T) {
	got := Add(2, 3)
	want := 5
	
	if got != want {
		t.Errorf("Add(2, 3) = %d; want %d", got, want)
	}
}`,
		PracticalExercise: `Try this:
1. Create a file called "simple.go" with a function that adds two numbers
2. Create a file called "simple_test.go" with a test function
3. Run "go test" in your terminal`,
	}
}

func getBasicTestsConcept() Concept {
	return Concept{
		Title: "Writing Basic Tests",
		Description: `The core of Go testing is the test function. A basic test follows this pattern:

1. Set up any necessary data or state
2. Call the function you're testing with known inputs
3. Check if the result matches what you expect
4. Report success or failure with the t.Errorf method

Test function naming convention:
- Must start with "Test"
- Followed by a name starting with a capital letter
- Should describe what's being tested

The testing.T type provides methods for reporting errors and controlling test execution:
- t.Errorf(): Report an error and continue testing
- t.Fatalf(): Report an error and stop this test function immediately
- t.Logf(): Log information without failing the test
- t.Skip(): Skip this test (useful for tests that only run in certain environments)`,
		Example: `package calculator

import "testing"

func TestAdd(t *testing.T) {
	// Setup test data
	a, b := 2, 3
	expected := 5
	
	// Call the function being tested
	result := Add(a, b)
	
	// Check the result
	if result != expected {
		t.Errorf("Add(%d, %d) = %d; expected %d", 
		         a, b, result, expected)
	}
}

func TestMultiply(t *testing.T) {
	// Multiple test cases
	t.Run("positive numbers", func(t *testing.T) {
		if Multiply(2, 3) != 6 {
			t.Error("Failed to multiply positive numbers")
		}
	})
	
	t.Run("with zero", func(t *testing.T) {
		if Multiply(5, 0) != 0 {
			t.Error("Multiply with zero should return zero")
		}
	})
}`,
		PracticalExercise: `Create a test for a function that checks if a number is even:

1. Create a file called "check.go" with an IsEven function that returns true for even numbers
2. Create "check_test.go" with a TestIsEven function that tests multiple cases
3. Use t.Run() to organize your test cases
4. Run the test with "go test"`,
	}
}

func getAssertionsConcept() Concept {
	return Concept{
		Title: "Test Assertions and Comparisons",
		Description: `Unlike some testing frameworks, Go's standard testing package doesn't provide 
assertion functions. Instead, you write normal Go code to check conditions.

Common assertion patterns:

For Equality:
- Simple values: if got != want { t.Errorf("...") }
- Floating point: if math.Abs(got-want) > 0.0001 { t.Errorf("...") }
- Slices: Use reflect.DeepEqual() or a loop to compare elements
- Structs: Use reflect.DeepEqual() or compare fields individually

For Errors:
- Check for specific errors: if err != expectedErr { t.Errorf("...") }
- Check for any error: if err != nil { t.Errorf("...") }
- Check for no error: if err != nil { t.Fatalf("...") }

If you prefer assertion-style syntax, there are third-party packages like:
- github.com/stretchr/testify/assert
- github.com/matryer/is

But Go's standard approach promotes writing clear, straightforward tests.`,
		Example: `package compare

import (
	"testing"
	"reflect"
)

func TestCompareInts(t *testing.T) {
	// Simple value comparison
	got := 42
	want := 42
	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}

func TestCompareSlices(t *testing.T) {
	// Slice comparison
	got := []int{1, 2, 3}
	want := []int{1, 2, 3}
	
	// Using reflect.DeepEqual
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v, want %v", got, want)
	}
	
	// Alternative: manual comparison
	if len(got) != len(want) {
		t.Errorf("Slices have different lengths: %d vs %d", 
		         len(got), len(want))
		return
	}
	
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("At index %d: got %v, want %v", 
			         i, got[i], want[i])
		}
	}
}

func TestErrorChecking(t *testing.T) {
	// Error checking
	_, err := riskyFunction()
	
	// Check for any error
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	
	// Check for specific error
	_, err = anotherFunction()
	if err != ErrExpected {
		t.Errorf("Got error %v, want %v", err, ErrExpected)
	}
}`,
		PracticalExercise: `Write a test that compares different data structures:

1. Create a function that returns a struct with several fields
2. Write a test that compares the returned struct with an expected value
3. Try comparing slices of structs
4. Experiment with different ways of reporting detailed differences`,
	}
}

func getTableDrivenTestsConcept() Concept {
	return Concept{
		Title: "Table-Driven Tests",
		Description: `Table-driven tests are a pattern for testing multiple cases with the same logic.
This approach makes it easy to add more test cases without duplicating test code.

Key benefits:
- Reduces code duplication
- Makes test cases explicit and easy to read
- Simplifies adding new test cases
- Provides clear documentation of boundary conditions and edge cases

A table-driven test typically consists of:
1. A slice of test case structs, each containing inputs and expected outputs
2. A loop that iterates through each test case
3. Common test logic applied to each case

You can use t.Run() to create a subtest for each test case, which provides better
reporting and the ability to run specific cases.`,
		Example: `package calculator

import "testing"

func TestAdd(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"both positive", 2, 3, 5},
		{"positive and negative", 5, -2, 3},
		{"both negative", -1, -3, -4},
		{"with zero", 0, 8, 8},
	}
	
	// Loop through test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Add(tc.a, tc.b)
			if got != tc.expected {
				t.Errorf("Add(%d, %d) = %d; expected %d", 
				         tc.a, tc.b, got, tc.expected)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	testCases := []struct {
		name        string
		a, b        int
		expected    int
		expectError bool
	}{
		{"normal division", 10, 2, 5, false},
		{"division by zero", 10, 0, 0, true},
		{"negative numbers", -10, -2, 5, false},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Divide(tc.a, tc.b)
			
			// Check error expectation
			if tc.expectError && err == nil {
				t.Error("Expected an error but didn't get one")
			}
			if !tc.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			
			// Only check result if no error expected
			if !tc.expectError && got != tc.expected {
				t.Errorf("Divide(%d, %d) = %d; expected %d", 
				         tc.a, tc.b, got, tc.expected)
			}
		})
	}
}`,
		PracticalExercise: `Write a table-driven test for a string processing function:

1. Create a function that checks if a string is a palindrome
2. Write a table-driven test with various test cases:
   - Normal palindromes ("racecar", "level")
   - Non-palindromes
   - Empty string
   - Single character
   - Case sensitivity
3. Run the tests and see how the reports are organized by case name`,
	}
}

func getBenchmarkConcept() Concept {
	return Concept{
		Title: "Benchmarks",
		Description: `Go testing includes built-in support for benchmarking, which allows you to
measure the performance of your code.

Benchmark functions:
1. Start with "Benchmark" (not "Test")
2. Take a parameter of type *testing.B (not *testing.T)
3. Use b.N for the number of iterations
4. Report allocations with b.ReportAllocs()

Running benchmarks:
- go test -bench=. (run all benchmarks)
- go test -bench=SpecificBenchmark (run specific benchmarks)
- go test -bench=. -benchmem (include memory allocation stats)

The Go testing framework will automatically determine how many iterations (b.N)
to run to get reliable timing information.`,
		Example: `package performance

import "testing"

func BenchmarkFibonacci(b *testing.B) {
	// Reset the timer to exclude setup time
	b.ResetTimer()
	
	// Run the Fibonacci function b.N times
	for i := 0; i < b.N; i++ {
		Fibonacci(20)
	}
}

func BenchmarkSort(b *testing.B) {
	// Testing with different input sizes
	sizes := []int{10, 100, 1000}
	
	for _, size := range sizes {
		// Create a sub-benchmark for each size
		b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
			// Generate input data (once per sub-benchmark)
			data := generateRandomData(size)
			
			// Report memory allocations
			b.ReportAllocs()
			
			// Reset timer before the actual test
			b.ResetTimer()
			
			for i := 0; i < b.N; i++ {
				// Create a copy to avoid measuring slice clear time
				input := make([]int, len(data))
				copy(input, data)
				
				Sort(input)
			}
		})
	}
}`,
		PracticalExercise: `Write a benchmark for comparing two different ways to concatenate strings:

1. Create two functions:
   - One that uses the + operator repeatedly
   - One that uses strings.Builder
2. Write benchmark functions for both
3. Run the benchmarks with "go test -bench=. -benchmem"
4. Compare the performance and memory allocation results`,
	}
}

func getSubtestsConcept() Concept {
	return Concept{
		Title: "Subtests and Test Organization",
		Description: `Go's testing package allows you to organize tests into hierarchical groups
using subtests. This helps with:

1. Grouping related test cases
2. Setting up common test fixtures
3. Running specific groups of tests
4. Parallel test execution

Subtests are created with the t.Run() method, which takes a name and a function.
You can create nested subtests by calling t.Run() inside another subtest.

Parallel tests:
- Call t.Parallel() at the beginning of a test to mark it as parallel
- Parallel tests run concurrently with other parallel tests
- Be careful with shared resources in parallel tests`,
		Example: `package organization

import (
	"testing"
	"time"
)

func TestUserFunctions(t *testing.T) {
	// Setup that's shared across subtests
	db := setupTestDatabase()
	defer db.Close()
	
	// Subtest for user creation
	t.Run("Creation", func(t *testing.T) {
		t.Run("ValidUser", func(t *testing.T) {
			user := User{Name: "Alice", Age: 30}
			id, err := db.CreateUser(user)
			if err != nil {
				t.Fatalf("Failed to create valid user: %v", err)
			}
			if id <= 0 {
				t.Errorf("Expected positive user ID, got %d", id)
			}
		})
		
		t.Run("InvalidUser", func(t *testing.T) {
			user := User{Name: "", Age: -5}
			_, err := db.CreateUser(user)
			if err == nil {
				t.Error("Expected error for invalid user, got nil")
			}
		})
	})
	
	// Subtest for user retrieval
	t.Run("Retrieval", func(t *testing.T) {
		// Setup for this specific subtest
		userId := createTestUser(db, "Bob", 25)
		
		t.Run("ByID", func(t *testing.T) {
			user, err := db.GetUserByID(userId)
			if err != nil {
				t.Fatalf("Failed to get user by ID: %v", err)
			}
			if user.Name != "Bob" {
				t.Errorf("Expected name 'Bob', got '%s'", user.Name)
			}
		})
		
		t.Run("ByName", func(t *testing.T) {
			users, err := db.GetUsersByName("Bob")
			if err != nil {
				t.Fatalf("Failed to get users by name: %v", err)
			}
			if len(users) == 0 {
				t.Error("No users returned for name 'Bob'")
			}
		})
	})
}

func TestParallelExecution(t *testing.T) {
	// These subtests will run in parallel with each other
	t.Run("Parallel1", func(t *testing.T) {
		t.Parallel() // Mark as parallel
		time.Sleep(100 * time.Millisecond)
		// Test logic here
	})
	
	t.Run("Parallel2", func(t *testing.T) {
		t.Parallel() // Mark as parallel
		time.Sleep(100 * time.Millisecond)
		// Test logic here
	})
	
	t.Run("Parallel3", func(t *testing.T) {
		t.Parallel() // Mark as parallel
		time.Sleep(100 * time.Millisecond)
		// Test logic here
	})
}`,
		PracticalExercise: `Organize a complex test suite using subtests:

1. Create a test function for a data structure (like a List or Stack)
2. Use t.Run() to create subtests for different operations (Add, Remove, etc.)
3. Further divide some operations into subcategories with more t.Run() calls
4. Try running specific subtests with "go test -run=TestMyFunction/Add"`,
	}
}

func getTestMainConcept() Concept {
	return Concept{
		Title: "TestMain and Setup/Teardown",
		Description: `The TestMain function allows you to control the testing process and set up
resources that are shared across multiple tests.

When you define a TestMain function in a test file, it will be called once
before any tests in that package are run. You're responsible for calling
m.Run() to execute the tests and os.Exit() with the result.

Use TestMain for:
- Database initialization
- Loading test fixtures
- Configuration setup
- Global resource allocation and cleanup

TestMain runs only once per package, not once per test file.`,
		Example: `package database

import (
	"database/sql"
	"os"
	"testing"
	"log"
)

var testDB *sql.DB

// TestMain controls the testing execution for the entire package
func TestMain(m *testing.M) {
	// Setup
	var err error
	testDB, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("Failed to open test database: %v", err)
	}
	
	// Initialize schema
	if err := initializeTestSchema(testDB); err != nil {
		testDB.Close()
		log.Fatalf("Failed to initialize test schema: %v", err)
	}
	
	// Load test data
	if err := loadTestData(testDB); err != nil {
		testDB.Close()
		log.Fatalf("Failed to load test data: %v", err)
	}
	
	// Run tests
	exitCode := m.Run()
	
	// Teardown
	if err := testDB.Close(); err != nil {
		log.Printf("Warning: failed to close test database: %v", err)
	}
	
	// Return exit code from tests
	os.Exit(exitCode)
}

func TestDatabaseQuery(t *testing.T) {
	// This test can use the testDB that was set up in TestMain
	rows, err := testDB.Query("SELECT * FROM users")
	if err != nil {
		t.Fatalf("Failed to query users: %v", err)
	}
	defer rows.Close()
	
	// Test logic...
}

func TestUserInsertion(t *testing.T) {
	// This test also has access to testDB
	result, err := testDB.Exec(
		"INSERT INTO users (name, email) VALUES (?, ?)",
		"Test User", "test@example.com")
	if err != nil {
		t.Fatalf("Failed to insert user: %v", err)
	}
	
	id, err := result.LastInsertId()
	if err != nil {
		t.Fatalf("Failed to get last insert ID: %v", err)
	}
	
	if id <= 0 {
		t.Errorf("Expected positive ID, got %d", id)
	}
}`,
		PracticalExercise: `Create a TestMain function for a test suite:

1. Define a TestMain function that sets up some shared test resources
2. Write several test functions that use these resources
3. Make sure TestMain properly cleans up after all tests are done
4. Run the tests and observe the order of setup, tests, and teardown`,
	}
}

func getTestHelpersConcept() Concept {
	return Concept{
		Title: "Test Helpers and Utilities",
		Description: `As your test suite grows, you'll want to extract common testing logic into
helper functions to avoid duplication.

Good test helpers:
1. Make tests more readable and concise
2. Hide complex setup and validation logic
3. Provide clear error messages
4. Are marked with t.Helper() to improve error reporting

The t.Helper() method marks a function as a test helper. This improves error
reporting by making test failures report the line number in the test itself,
not in the helper function.

Common types of test helpers:
- Assertion helpers
- Data generation helpers  
- Setup/teardown helpers
- Custom comparison helpers`,
		Example: `package testing_helpers

import (
	"testing"
	"reflect"
)

// Assert that two values are equal
func assertEqual(t *testing.T, got, want interface{}, msg string) {
	t.Helper() // Mark as helper function for better error reporting
	
	if !reflect.DeepEqual(got, want) {
		if msg == "" {
			t.Errorf("Got %v, want %v", got, want)
		} else {
			t.Errorf("%s: Got %v, want %v", msg, got, want)
		}
	}
}

// Assert that an error is not nil
func assertError(t *testing.T, err error, msg string) {
	t.Helper()
	
	if err == nil {
		if msg == "" {
			t.Error("Expected an error but got nil")
		} else {
			t.Errorf("%s: Expected an error but got nil", msg)
		}
	}
}

// Assert that an error is nil
func assertNoError(t *testing.T, err error, msg string) {
	t.Helper()
	
	if err != nil {
		if msg == "" {
			t.Errorf("Unexpected error: %v", err)
		} else {
			t.Errorf("%s: Unexpected error: %v", msg, err)
		}
	}
}

// Example of using helpers in a test
func TestWithHelpers(t *testing.T) {
	result, err := Calculate(10, 5)
	
	assertNoError(t, err, "Calculate function")
	assertEqual(t, result, 15, "Addition result")
	
	_, err = Calculate(10, 0)
	assertError(t, err, "Division by zero")
}

// Create a test user helper
func createTestUser(t *testing.T, db Database) User {
	t.Helper()
	
	user := User{Name: "Test User", Email: "test@example.com"}
	id, err := db.CreateUser(user)
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}
	
	user.ID = id
	return user
}`,
		PracticalExercise: `Create a set of test helpers for a specific domain:

1. Write a small package or application that works with a particular type of data
2. Create test helper functions for common assertions related to that data
3. Use t.Helper() in each helper function
4. Write tests that use these helpers
5. Observe how error messages point to the right location when tests fail`,
	}
}
