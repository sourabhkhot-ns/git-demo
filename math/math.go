package math

import "errors"

// Calculator interface defines basic math operations
type Calculator interface {
	Add(a, b int) int
	Subtract(a, b int) int
	Multiply(a, b int) int
	Divide(a, b int) (float64, error)
}

// BasicCalculator implements the Calculator interface
type BasicCalculator struct {
	Name string
}

// NewBasicCalculator creates a new instance of BasicCalculator
func NewBasicCalculator(name string) *BasicCalculator {
	return &BasicCalculator{
		Name: name,
	}
}

// Add implements addition operation
func (c *BasicCalculator) Add(a, b int) int {
	return a + b
}

// Subtract implements subtraction operation
func (c *BasicCalculator) Subtract(a, b int) int {
	return a - b
}

// Multiply implements multiplication operation
func (c *BasicCalculator) Multiply(a, b int) int {
	return a * b
}

// Divide implements division operation
func (c *BasicCalculator) Divide(a, b int) (float64, error) {
	if b == 0 {
		return 0.0, errors.New("cannot divide by zero")
	}
	return float64(a) / float64(b), nil
}
