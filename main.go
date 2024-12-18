package main

import (
	"fmt"
	"os"
)

var problems = make(map[string]func())

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <problem_number>")
		return
	}

	problemNumber := fmt.Sprintf("p%03s", os.Args[1])
	if problem, exists := problems[problemNumber]; exists {
		problem()
	} else {
		fmt.Printf("Problem %s not found\n", problemNumber)
	}
}
