package tutorial

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Example represents a complete example with source and test code
type Example struct {
	Name        string
	Description string
	SourceCode  string
	TestCode    string
}

// GetAllExamples returns all available examples
func GetAllExamples() []Example {
	return []Example{
		getBasicExample(),
		getTableDrivenExample(),
		getBenchmarkExample(),
		getSubtestExample(),
		getTestMainExample(),
	}
}

// CreateExampleFiles creates the source and test files for an example
func CreateExampleFiles(example Example, directory string) error {
	// Create directory if it doesn't exist
	if err := os.MkdirAll(directory, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	// Extract package name from source code
	packageName := extractPackageName(example.SourceCode)
	if packageName == "" {
		packageName = "main"
	}

	// Create source file
	sourceFileName := strings.ToLower(example.Name) + ".go"
	sourcePath := filepath.Join(directory, sourceFileName)
	if err := os.WriteFile(sourcePath, []byte(example.SourceCode), 0644); err != nil {
		return fmt.Errorf("failed to write source file: %v", err)
	}

	// Create test file
	testFileName := strings.ToLower(example.Name) + "_test.go"
	testPath := filepath.Join(directory, testFileName)
	if err := os.WriteFile(testPath, []byte(example.TestCode), 0644); err != nil {
		return fmt.Errorf("failed to write test file: %v", err)
	}

	return nil
}

// Helper function to extract package name from source code
func extractPackageName(source string) string {
	lines := strings.Split(source, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "package ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "package "))
		}
	}
	return ""
}

func getBasicExample() Example {
	return Example{
		Name:        "Calculator",
		Description: "A simple calculator with basic tests",
		SourceCode: `package calculator

// Add returns the sum of two integers
func Add(a, b int) int {
	return a + b
}

// Subtract returns the difference between two integers
func Subtract(a, b int) int {
	return a - b
}

// Multiply returns the product of two integers
func Multiply(a, b int) int {
	return a * b
}

// Divide returns the quotient of two integers
// Returns an error if attempting to divide by zero
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}
`,
		TestCode: `package calculator

import (
	"testing"
)

func TestAdd(t *testing.T) {
	// Test case
	result := Add(2, 3)
	expected := 5
	
	if result != expected {
		t.Errorf("Add(2, 3) = %d; expected %d", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	// Test case
	result := Subtract(5, 3)
	expected := 2
	
	if result != expected {
		t.Errorf("Subtract(5, 3) = %d; expected %d", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	// Test case
	result := Multiply(4, 3)
	expected := 12
	
	if result != expected {
		t.Errorf("Multiply(4, 3) = %d; expected %d", result, expected)
	}
}

func TestDivide(t *testing.T) {
	// Test successful division
	result, err := Divide(10, 2)
	expected := 5
	
	if err != nil {
		t.Errorf("Divide(10, 2) returned error: %v", err)
	}
	
	if result != expected {
		t.Errorf("Divide(10, 2) = %d; expected %d", result, expected)
	}
	
	// Test division by zero
	_, err = Divide(10, 0)
	if err == nil {
		t.Error("Divide(10, 0) did not return an error")
	}
}
`,
	}
}

func getTableDrivenExample() Example {
	return Example{
		Name:        "StringUtils",
		Description: "String utilities with table-driven tests",
		SourceCode: `package stringutils

import (
	"strings"
)

// Reverse returns the reverse of a string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsPalindrome checks if a string reads the same backward as forward
func IsPalindrome(s string) bool {
	// Convert to lowercase and remove non-alphanumeric characters
	s = strings.ToLower(s)
	var result strings.Builder
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			result.WriteRune(r)
		}
	}
	
	cleaned := result.String()
	return cleaned == Reverse(cleaned)
}

// CountWords counts the number of words in a string
func CountWords(s string) int {
	if len(s) == 0 {
		return 0
	}
	
	return len(strings.Fields(s))
}
`,
		TestCode: `package stringutils

import (
	"testing"
)

func TestReverse(t *testing.T) {
	// Define test cases
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"single character", "a", "a"},
		{"palindrome", "racecar", "racecar"},
		{"normal string", "hello", "olleh"},
		{"with spaces", "hello world", "dlrow olleh"},
		{"with unicode", "こんにちは", "はちにんこ"},
	}
	
	// Run test cases
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Reverse(tc.input)
			if got != tc.expected {
				t.Errorf("Reverse(%q) = %q; expected %q", 
				         tc.input, got, tc.expected)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty string", "", true},
		{"single character", "a", true},
		{"simple palindrome", "racecar", true},
		{"palindrome with spaces", "never odd or even", true},
		{"palindrome with punctuation", "A man, a plan, a canal: Panama", true},
		{"palindrome with mixed case", "Able was I ere I saw Elba", true},
		{"non-palindrome", "hello", false},
		{"almost palindrome", "almost tsomla", false},
	}
	
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsPalindrome(tc.input)
			if got != tc.expected {
				t.Errorf("IsPalindrome(%q) = %v; expected %v", 
				         tc.input, got, tc.expected)
			}
		})
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"empty string", "", 0},
		{"single word", "hello", 1},
		{"two words", "hello world", 2},
		{"multiple words", "the quick brown fox", 4},
		{"extra spaces", "  spaced   out   ", 2},
	}
	
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := CountWords(tc.input)
			if got != tc.expected {
				t.Errorf("CountWords(%q) = %d; expected %d", 
				         tc.input, got, tc.expected)
			}
		})
	}
}
`,
	}
}

func getBenchmarkExample() Example {
	return Example{
		Name:        "Fibonacci",
		Description: "Fibonacci implementations with benchmarks",
		SourceCode: `package fibonacci

// FibRecursive calculates the nth Fibonacci number recursively
func FibRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return FibRecursive(n-1) + FibRecursive(n-2)
}

// FibIterative calculates the nth Fibonacci number iteratively
func FibIterative(n int) int {
	if n <= 1 {
		return n
	}
	
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// FibMemoized calculates the nth Fibonacci number using memoization
func FibMemoized(n int) int {
	cache := make(map[int]int)
	
	var fib func(int) int
	fib = func(n int) int {
		if n <= 1 {
			return n
		}
		
		// Check if we've already calculated this value
		if val, ok := cache[n]; ok {
			return val
		}
		
		// Calculate and cache the value
		cache[n] = fib(n-1) + fib(n-2)
		return cache[n]
	}
	
	return fib(n)
}
`,
		TestCode: `package fibonacci

import (
	"testing"
)

// Test implementations for correctness
func TestFibonacci(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{10, 55},
		{20, 6765},
	}
	
	implementations := map[string]func(int) int{
		"Recursive": FibRecursive,
		"Iterative": FibIterative,
		"Memoized":  FibMemoized,
	}
	
	for name, impl := range implementations {
		t.Run(name, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(string(rune(tc.n+'0')), func(t *testing.T) {
					result := impl(tc.n)
					if result != tc.expected {
						t.Errorf("Fib(%d) = %d; expected %d", 
							tc.n, result, tc.expected)
					}
				})
			}
		})
	}
}

// Benchmark recursive implementation
func BenchmarkFibRecursive(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
	}{
		{"n=5", 5},
		{"n=10", 10},
		{"n=20", 20},
	}
	
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FibRecursive(bm.n)
			}
		})
	}
}

// Benchmark iterative implementation
func BenchmarkFibIterative(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
	}{
		{"n=5", 5},
		{"n=10", 10},
		{"n=20", 20},
		{"n=100", 100}, // Can handle larger values
	}
	
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FibIterative(bm.n)
			}
		})
	}
}

// Benchmark memoized implementation
func BenchmarkFibMemoized(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
	}{
		{"n=5", 5},
		{"n=10", 10},
		{"n=20", 20},
		{"n=100", 100}, // Can handle larger values
	}
	
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// Note: b.N iterations means the cache is rebuilt each time
				// This is intentional to measure the full algorithm performance
				FibMemoized(bm.n)
			}
		})
	}
}
`,
	}
}

func getSubtestExample() Example {
	return Example{
		Name:        "Stack",
		Description: "Stack implementation with subtests",
		SourceCode: `package stack

import "errors"

var (
	ErrEmptyStack = errors.New("stack is empty")
)

// Stack represents a last-in-first-out (LIFO) stack of elements
type Stack struct {
	elements []interface{}
}

// New creates a new empty stack
func New() *Stack {
	return &Stack{
		elements: make([]interface{}, 0),
	}
}

// Push adds an element to the top of the stack
func (s *Stack) Push(value interface{}) {
	s.elements = append(s.elements, value)
}

// Pop removes and returns the top element from the stack
// Returns an error if the stack is empty
func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, ErrEmptyStack
	}
	
	index := len(s.elements) - 1
	element := s.elements[index]
	s.elements = s.elements[:index]
	return element, nil
}

// Peek returns the top element without removing it
// Returns an error if the stack is empty
func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, ErrEmptyStack
	}
	
	return s.elements[len(s.elements)-1], nil
}

// IsEmpty returns true if the stack contains no elements
func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

// Size returns the number of elements in the stack
func (s *Stack) Size() int {
	return len(s.elements)
}

// Clear removes all elements from the stack
func (s *Stack) Clear() {
	s.elements = make([]interface{}, 0)
}
`,
		TestCode: `package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	// Group tests for stack operations
	t.Run("Basic Operations", func(t *testing.T) {
		// Test creating a new stack
		t.Run("New", func(t *testing.T) {
			s := New()
			if s == nil {
				t.Fatal("New() returned nil")
			}
			
			if !s.IsEmpty() {
				t.Error("New stack should be empty")
			}
			
			if s.Size() != 0 {
				t.Errorf("New stack size = %d; expected 0", s.Size())
			}
		})
		
		// Test pushing elements
		t.Run("Push", func(t *testing.T) {
			s := New()
			
			s.Push(1)
			if s.Size() != 1 {
				t.Errorf("Stack size after one Push = %d; expected 1", s.Size())
			}
			
			s.Push("two")
			if s.Size() != 2 {
				t.Errorf("Stack size after two Pushes = %d; expected 2", s.Size())
			}
			
			if s.IsEmpty() {
				t.Error("Stack should not be empty after Push")
			}
		})
		
		// Test peeking at elements
		t.Run("Peek", func(t *testing.T) {
			s := New()
			
			// Peek on empty stack
			_, err := s.Peek()
			if err != ErrEmptyStack {
				t.Errorf("Peek on empty stack = %v; expected %v", err, ErrEmptyStack)
			}
			
			// Peek after pushing
			s.Push(42)
			value, err := s.Peek()
			if err != nil {
				t.Errorf("Peek error = %v; expected nil", err)
			}
			
			if value != 42 {
				t.Errorf("Peek value = %v; expected 42", value)
			}
			
			// Size should not change after Peek
			if s.Size() != 1 {
				t.Errorf("Stack size after Peek = %d; expected 1", s.Size())
			}
		})
		
		// Test popping elements
		t.Run("Pop", func(t *testing.T) {
			s := New()
			
			// Pop on empty stack
			_, err := s.Pop()
			if err != ErrEmptyStack {
				t.Errorf("Pop on empty stack = %v; expected %v", err, ErrEmptyStack)
			}
			
			// Pop after pushing
			s.Push("first")
			s.Push("second")
			
			value, err := s.Pop()
			if err != nil {
				t.Errorf("Pop error = %v; expected nil", err)
			}
			
			if value != "second" {
				t.Errorf("Pop value = %v; expected \"second\"", value)
			}
			
			if s.Size() != 1 {
				t.Errorf("Stack size after Pop = %d; expected 1", s.Size())
			}
			
			// Pop again
			value, err = s.Pop()
			if err != nil {
				t.Errorf("Pop error = %v; expected nil", err)
			}
			
			if value != "first" {
				t.Errorf("Pop value = %v; expected \"first\"", value)
			}
			
			if s.Size() != 0 {
				t.Errorf("Stack size after second Pop = %d; expected 0", s.Size())
			}
			
			if !s.IsEmpty() {
				t.Error("Stack should be empty after popping all elements")
			}
		})
		
		// Test clearing the stack
		t.Run("Clear", func(t *testing.T) {
			s := New()
			
			// Push elements
			s.Push(1)
			s.Push(2)
			s.Push(3)
			
			// Clear
			s.Clear()
			
			if !s.IsEmpty() {
				t.Error("Stack should be empty after Clear")
			}
			
			if s.Size() != 0 {
				t.Errorf("Stack size after Clear = %d; expected 0", s.Size())
			}
		})
	})
	
	// Test complex sequence of operations
	t.Run("Operation Sequence", func(t *testing.T) {
		s := New()
		
		// Sequence of operations
		s.Push(1)
		s.Push(2)
		s.Push(3)
		
		val, _ := s.Pop()
		if val != 3 {
			t.Errorf("First Pop = %v; expected 3", val)
		}
		
		s.Push(4)
		
		val, _ = s.Peek()
		if val != 4 {
			t.Errorf("Peek after sequence = %v; expected 4", val)
		}
		
		val, _ = s.Pop()
		if val != 4 {
			t.Errorf("Second Pop = %v; expected 4", val)
		}
		
		val, _ = s.Pop()
		if val != 2 {
			t.Errorf("Third Pop = %v; expected 2", val)
		}
		
		if s.Size() != 1 {
			t.Errorf("Size after sequence = %d; expected 1", s.Size())
		}
	})
}
`,
	}
}

func getTestMainExample() Example {
	return Example{
		Name:        "Database",
		Description: "Database operations with TestMain setup",
		SourceCode: `package database

import (
	"errors"
	"sync"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUserExists   = errors.New("user already exists")
)

// User represents a user in the database
type User struct {
	ID    int
	Name  string
	Email string
}

// Database provides operations for storing and retrieving users
type Database struct {
	mu    sync.RWMutex
	users map[int]User
	nextID int
}

// NewDatabase creates a new in-memory database
func NewDatabase() *Database {
	return &Database{
		users: make(map[int]User),
		nextID: 1,
	}
}

// AddUser adds a new user to the database
func (db *Database) AddUser(user User) (int, error) {
	db.mu.Lock()
	defer db.mu.Unlock()
	
	// Check if user with same email exists
	for _, u := range db.users {
		if u.Email == user.Email {
			return 0, ErrUserExists
		}
	}
	
	// Assign ID and store user
	user.ID = db.nextID
	db.users[user.ID] = user
	db.nextID++
	
	return user.ID, nil
}

// GetUser retrieves a user by ID
func (db *Database) GetUser(id int) (User, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	
	user, ok := db.users[id]
	if !ok {
		return User{}, ErrUserNotFound
	}
	
	return user, nil
}

// UpdateUser updates an existing user
func (db *Database) UpdateUser(user User) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	
	if _, ok := db.users[user.ID]; !ok {
		return ErrUserNotFound
	}
	
	// Check if email conflicts with another user
	for id, u := range db.users {
		if u.Email == user.Email && id != user.ID {
			return ErrUserExists
		}
	}
	
	db.users[user.ID] = user
	return nil
}

// DeleteUser removes a user from the database
func (db *Database) DeleteUser(id int) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	
	if _, ok := db.users[id]; !ok {
		return ErrUserNotFound
	}
	
	delete(db.users, id)
	return nil
}

// GetAllUsers returns all users in the database
func (db *Database) GetAllUsers() []User {
	db.mu.RLock()
	defer db.mu.RUnlock()
	
	users := make([]User, 0, len(db.users))
	for _, user := range db.users {
		users = append(users, user)
	}
	
	return users
}
`,
		TestCode: `package database

import (
	"os"
	"testing"
)

var testDB *Database

// TestMain sets up the test database before running tests
func TestMain(m *testing.M) {
	// Setup
	testDB = NewDatabase()
	
	// Add some test data
	setupTestData(testDB)
	
	// Run tests
	exitCode := m.Run()
	
	// Cleanup (not strictly necessary for in-memory DB, but good practice)
	testDB = nil
	
	os.Exit(exitCode)
}

func setupTestData(db *Database) {
	users := []User{
		{Name: "Alice Smith", Email: "alice@example.com"},
		{Name: "Bob Johnson", Email: "bob@example.com"},
		{Name: "Charlie Brown", Email: "charlie@example.com"},
	}
	
	for _, user := range users {
		db.AddUser(user)
	}
}

func TestAddUser(t *testing.T) {
	// Test adding a new user
	t.Run("New User", func(t *testing.T) {
		user := User{
			Name:  "Dave Davis",
			Email: "dave@example.com",
		}
		
		id, err := testDB.AddUser(user)
		if err != nil {
			t.Fatalf("Failed to add new user: %v", err)
		}
		
		if id <= 0 {
			t.Errorf("Got invalid user ID: %d", id)
		}
		
		// Verify user was added
		retrieved, err := testDB.GetUser(id)
		if err != nil {
			t.Fatalf("Failed to get added user: %v", err)
		}
		
		if retrieved.Name != user.Name {
			t.Errorf("Got name %q, expected %q", retrieved.Name, user.Name)
		}
		
		if retrieved.Email != user.Email {
			t.Errorf("Got email %q, expected %q", retrieved.Email, user.Email)
		}
	})
	
	// Test adding a user with duplicate email
	t.Run("Duplicate Email", func(t *testing.T) {
		user := User{
			Name:  "Another Alice",
			Email: "alice@example.com", // Email already exists
		}
		
		_, err := testDB.AddUser(user)
		if err != ErrUserExists {
			t.Errorf("Got error %v, expected %v", err, ErrUserExists)
		}
	})
}

func TestGetUser(t *testing.T) {
	// Test getting an existing user
	t.Run("Existing User", func(t *testing.T) {
		// Get ID of an existing user (assuming ID 1 exists from setup)
		user, err := testDB.GetUser(1)
		if err != nil {
			t.Errorf("Failed to get existing user: %v", err)
		}
		
		if user.ID != 1 {
			t.Errorf("Got user ID %d, expected 1", user.ID)
		}
		
		if user.Name == "" {
			t.Error("Got empty name for existing user")
		}
	})
	
	// Test getting a non-existent user
	t.Run("Non-existent User", func(t *testing.T) {
		// Try to get a user with an ID that doesn't exist (assuming 9999 doesn't exist)
		_, err := testDB.GetUser(9999)
		if err != ErrUserNotFound {
			t.Errorf("Got error %v, expected %v", err, ErrUserNotFound)
		}
	})
}

func TestUpdateUser(t *testing.T) {
	// Test updating an existing user
	t.Run("Valid Update", func(t *testing.T) {
		// First, get an existing user
		user, err := testDB.GetUser(2)
		if err != nil {
			t.Fatalf("Failed to get user for update test: %v", err)
		}
		
		// Modify the user
		user.Name = "Updated Name"
		
		// Update
		err = testDB.UpdateUser(user)
		if err != nil {
			t.Errorf("Failed to update user: %v", err)
		}
		
		// Verify the update
		updated, err := testDB.GetUser(2)
		if err != nil {
			t.Fatalf("Failed to get user after update: %v", err)
		}
		
		if updated.Name != "Updated Name" {
			t.Errorf("Got name %q after update, expected %q", 
			         updated.Name, "Updated Name")
		}
	})
	
	// Test updating with a duplicate email
	t.Run("Duplicate Email Update", func(t *testing.T) {
		// Get two different users
		user1, _ := testDB.GetUser(1)
		user2, _ := testDB.GetUser(3)
		
		// Try to update user2 with user1's email
		user2.Email = user1.Email
		
		err := testDB.UpdateUser(user2)
		if err != ErrUserExists {
			t.Errorf("Got error %v, expected %v", err, ErrUserExists)
		}
	})
}

func TestDeleteUser(t *testing.T) {
	// Add a user specifically for deletion
	user := User{
		Name:  "Temporary User",
		Email: "temp@example.com",
	}
	
	id, err := testDB.AddUser(user)
	if err != nil {
		t.Fatalf("Failed to add user for deletion test: %v", err)
	}
	
	// Test deleting the user
	t.Run("Delete Existing", func(t *testing.T) {
		err := testDB.DeleteUser(id)
		if err != nil {
			t.Errorf("Failed to delete user: %v", err)
		}
		
		// Verify user was deleted
		_, err = testDB.GetUser(id)
		if err != ErrUserNotFound {
			t.Errorf("Got error %v after deletion, expected %v", 
			         err, ErrUserNotFound)
		}
	})
	
	// Test deleting a non-existent user
	t.Run("Delete Non-existent", func(t *testing.T) {
		err := testDB.DeleteUser(9999)
		if err != ErrUserNotFound {
			t.Errorf("Got error %v, expected %v", err, ErrUserNotFound)
		}
	})
}

func TestGetAllUsers(t *testing.T) {
	users := testDB.GetAllUsers()
	
	// We should have at least the users from setup
	if len(users) < 3 {
		t.Errorf("Got %d users, expected at least 3", len(users))
	}
	
	// Verify all users have valid IDs
	for _, user := range users {
		if user.ID <= 0 {
			t.Errorf("Found user with invalid ID: %d", user.ID)
		}
		
		if user.Name == "" {
			t.Errorf("Found user with empty name: %+v", user)
		}
		
		if user.Email == "" {
			t.Errorf("Found user with empty email: %+v", user)
		}
	}
}
`,
	}
}
