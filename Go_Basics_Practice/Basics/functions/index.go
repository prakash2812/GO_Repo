/* A function can take zero or more arguments.

In this example, add takes two parameters of type int.

Notice that the type comes after the variable name.

(For more about why types look the way they do, see the article on Go's declaration syntax.) */


package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}

/* When two or more consecutive named function parameters share a type, you can omit the type from all but the last.

In this example, we shortened

x int, y int
to

x, y int */


package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}


/* Multiple results
A function can return any number of results.

The swap function returns two strings. */

package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}



/* Named return values
Go's return values may be named. If so, they are treated as variables defined at the top of the function.

These names should be used to document the meaning of the return values.

A return statement without arguments returns the named return values. This is known as a "naked" return.

Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in long er functions.*/


package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}

/* Function values
Functions are values too. They can be passed around just like other values.

Function values may be used as function arguments and return values. */

package main

import (
	"fmt"
	"math"
)

func operators (fn func(int,int) int) int{
	return fn(3,4)
} 

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	add := func (x,y int) int{
		return x+y
	}
	fmt.Println(operators(add))
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}


/* Function closures
Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

For example, the adder function returns a closure. Each closure is bound to its own sum variable.
 */


 package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func isEvenNumber () func (int) bool {
	return func (y int) bool{
	return y%2 == 0
	}
}

func main() {
	pos, neg, isEven := adder(), adder(), isEvenNumber()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
			isEven(i),
		)
	}
}



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
