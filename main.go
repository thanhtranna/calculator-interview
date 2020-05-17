package main

import (
	"fmt"

	calculator "calculator/calculator"
)

func main() {
	calculator := new(calculator.Calculator)
	data, _ := calculator.Evaluate("A = 1 + 2")
	fmt.Println("result", data)

	data, _ = calculator.Evaluate("A + 2")
	fmt.Println("result", data)

	data, _ = calculator.Evaluate("B = 2 + A")
	fmt.Println("result", data)
}
