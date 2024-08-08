package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// task := r.URL.Query().Get("task")
	task := strings.ReplaceAll(r.URL.Query().Get("task"), " ", "+")
	if task == "" {
		http.Error(w, "Task parameter is missing", http.StatusBadRequest)
		return
	}

	var result float64
	var operator string
	var operatorIndex int

	for i, ch := range task {
		if ch == '+' || ch == '-' || ch == '*' || ch == '/' {
			operator = string(ch)
			operatorIndex = i
			break
		}
	}

	if operator == "" {
		http.Error(w, "Operator not found", http.StatusBadRequest)
		return
	}

	firstNumberStr := task[:operatorIndex]
	secondNumberStr := task[operatorIndex+1:]

	firstNumber, err1 := strconv.ParseFloat(firstNumberStr, 64)
	secondNumber, err2 := strconv.ParseFloat(secondNumberStr, 64)

	if err1 != nil || err2 != nil {
		http.Error(w, "Invalid format", http.StatusBadRequest)
		return
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
			http.Error(w, "You can't divide by 0", http.StatusBadRequest)
			return
		}
	default:
		http.Error(w, "Invalid expression", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "%f\n", result)
}
