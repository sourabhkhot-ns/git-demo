package main

import (
	"fmt"
	"git-demo/math"
)

func main() {
	// Using constructor
	calc := math.NewBasicCalculator("MyCalculator")
	
	// Perform calculations
	sum := calc.Add(10, 5)
	diff := calc.Subtract(10, 5)
	
	// Print results
	fmt.Printf("Calculator: %s\n", calc.Name)
	fmt.Printf("10 + 5 = %d\n", sum)
	fmt.Printf("10 - 5 = %d\n", diff)
	
	// Demonstrate interface usage
	var calcInterface math.Calculator = calc
	fmt.Printf("Using interface: 7 + 3 = %d\n", calcInterface.Add(7, 3))
}
