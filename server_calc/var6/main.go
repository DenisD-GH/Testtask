package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./history.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTable := `CREATE TABLE IF NOT EXISTS calculations (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        expression TEXT,
        result TEXT
    );`

	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, db)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
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

	saveExpression := `INSERT INTO calculations(expression, result) VALUES (?, ?)`
	_, err := db.Exec(saveExpression, task, result)
	if err != nil {
		http.Error(w, "Failed to save calculation", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, result)
}
