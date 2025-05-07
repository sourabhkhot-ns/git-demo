package main

import (
	"fmt"
	"git-demo/math"
)

func main() {
	// Using constructor for integer calculator
	intCalc := math.NewBasicCalculator("MyCalculator")

	// Perform integer calculations
	sum := intCalc.Add(10, 5)
	diff := intCalc.Subtract(10, 5)

	// Print results
	fmt.Printf("Calculator: %s\n", intCalc.Name)
	fmt.Printf("10 + 5 = %d\n", sum)
	fmt.Printf("10 - 5 = %d\n", diff)

	// Demonstrate interface usage
	var calcInterface math.Calculator = intCalc
	fmt.Printf("Using interface: 7 + 3 = %d\n", calcInterface.Add(7, 3))

	// Using constructor for complex calculator
	complexCalc := math.NewBasicComplexCalculator("MyComplexCalc")

	a := complex(2, 3) // 2 + 3i
	b := complex(1, 4) // 1 + 4i

	complexSum := complexCalc.Add(a, b)
	complexDiff := complexCalc.Subtract(a, b)

	fmt.Println("Complex Calculator:", complexCalc.Name)
	fmt.Println("Sum:", math.FormatComplex(complexSum))
	fmt.Println("Difference:", math.FormatComplex(complexDiff))
}
