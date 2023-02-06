/* Defer ---------------------- */
//  wheever you see defer in  parent function , imagine it's places at bootom lines of that function
// if you have nested function and in that another defer, it do same as above again it keep defer at bootom of line and executes
// now parent defers will be called

//	defer invokes function after surrounding functions executed
//
// it places this at the last line of closing braces {}
// It follows last in fist out data structure (LIFO)
package main

import "fmt"

// in main function defer keep at end of lines so it calls mydefer related function
// again in that mydefer it keep defer at last lines and calls out LIFO
// after my defer executed it calls main defer funciton in LIFO way
func main() {
	fmt.Println("Welcome to defer concept")

	defer fmt.Println("world")
	defer fmt.Println("one")
	defer fmt.Println("two")

	fmt.Println("Hello")

	mydefer()
}

func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("my defer ends")
}

/*
Welcome to defer concept
Hello
my defer ends
4
3
2
1
0
two
one
world

*/
