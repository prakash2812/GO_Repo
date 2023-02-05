package main

import "fmt"

func main() {
	fmt.Println("For loops concept")

	courses := []string{"react", "javascipt", "ai", "node"}

	// normal loop
	for i := 0; i < len(courses); i++ {
		fmt.Println(i)
	}

	// range loop with two values
	for _, value := range courses {
		fmt.Println(value)
	}

	// range loop with one value
	for idx := range courses {
		fmt.Println(courses[idx])
	}

	// while loop
	rougueValue := 1
	for rougueValue < 10 {

		if rougueValue == 2 {
			goto godev
		}

		if rougueValue == 5 {
			rougueValue++
			continue //  skip the 5 and it continues iteration
			// break it breaks iteration
		}
		fmt.Println("Rougur value", rougueValue)
		rougueValue++
	}

	// label with go

godev:
	fmt.Println("welcome to go.dev learning site")
}
