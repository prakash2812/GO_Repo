package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) []int {
	result:= []int{}
	result=append(result,0,1)
	return func(size int) []int{
		val:= result[size-1]+result[size-2]
		result=append(result,val)
		return result
	}
}

func main() {
	f := fibonacci()
	fmt.Println(f)
	for i := 2; i < 10; i++ {
		fmt.Println(f(i))
	}
}
