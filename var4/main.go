package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	for {
		var input string
		var result float64

		fmt.Println("Enter expression NumberOperatorNumber:")
		fmt.Scanln(&input)

		if strings.ToLower(input) == "exit" {
			fmt.Println("Exit")
			return
		}

		var operator string
		var operatorIndex int

		for i, ch := range input {
			if ch == '+' || ch == '-' || ch == '*' || ch == '/' {
				operator = string(ch)
				operatorIndex = i
				break
			}
		}

		if operator == "" {
			fmt.Println("Operator not found")
			continue
		}

		// https://www.digitalocean.com/community/tutorials/how-to-convert-data-types-in-go-ru
		firstNumberStr := input[:operatorIndex]
		secondNumberStr := input[operatorIndex+1:]

		firstNumber, err1 := strconv.ParseFloat(firstNumberStr, 64)
		secondNumber, err2 := strconv.ParseFloat(secondNumberStr, 64)

		if err1 != nil || err2 != nil {
			fmt.Println("Invalid format")
			continue
		}

		switch operator {
		case "+":
			result = firstNumber + secondNumber
		case "-":
			result = firstNumber - secondNumber
		case "*":
			result = firstNumber * secondNumber
		case "/":
			if secondNumber != 0 {
				result = firstNumber / secondNumber
			} else {
				fmt.Println("You can't divide by 0")
				continue
			}
		default:
			fmt.Println("Invalid expression")
			continue
		}

		fmt.Printf("%s = %f\n", input, result)
	}
}
