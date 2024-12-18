package main

import "fmt"

func init() {
	problems["p001"] = problem001
}

func problem001() {
	fmt.Println("Running Problem 001")

	sol := 0
	for i := range 1000 {
		if i%3 == 0 || i%5 == 0 {
			sol += i
		}
	}
	fmt.Println(sol)
}
