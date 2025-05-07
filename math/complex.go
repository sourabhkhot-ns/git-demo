package math

import "fmt"

// ComplexCalculator defines basic math operations for complex numbers
type ComplexCalculator interface {
	Add(a, b complex128) complex128
	Subtract(a, b complex128) complex128
}

// BasicComplexCalculator implements the ComplexCalculator interface
type BasicComplexCalculator struct {
	Name string
}

// NewBasicComplexCalculator creates a new instance of BasicComplexCalculator
func NewBasicComplexCalculator(name string) *BasicComplexCalculator {
	return &BasicComplexCalculator{
		Name: name,
	}
}

// Add implements addition for complex numbers
func (c *BasicComplexCalculator) Add(a, b complex128) complex128 {
	return a + b
}

// Subtract implements subtraction for complex numbers
func (c *BasicComplexCalculator) Subtract(a, b complex128) complex128 {
	return a - b
}

// Optional: To print complex numbers nicely
func FormatComplex(z complex128) string {
	return fmt.Sprintf("%.2f + %.2fi", real(z), imag(z))
}
