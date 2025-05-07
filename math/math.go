package math

// Calculator interface defines basic math operations
type Calculator interface {
	Add(a, b int) int
	Subtract(a, b int) int
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
