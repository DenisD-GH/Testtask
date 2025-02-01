package main

import (
	"fmt"
)

func main() {
	var firstNumber float64
	var secondNumber float64
	var operator string
	var result float64

	fmt.Println("Enter expression \"Number Operation Number\":")
	fmt.Scanf("%f%s%f", &firstNumber, &operator, &secondNumber)

	switch operator {
	case "+":
		result = firstNumber + secondNumber
		fmt.Println("Result:", result)
	case "-":
		result = firstNumber - secondNumber
		fmt.Println("Result:", result)
	case "*":
		result = firstNumber * secondNumber
		fmt.Println("Result:", result)
	case "/":
		if secondNumber != 0 {
			result = firstNumber / secondNumber
			fmt.Println("Result:", result)
		} else {
			fmt.Println("You can't divide by 0")
		}
	default:
		fmt.Println("Invalid input")
	}
}
