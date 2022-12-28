/* Methods
Go does not have classes. However, you can define methods on types.

A method is a function with a special receiver argument.

The receiver appears in its own argument list between the func keyword and the method name.

In this example, the Abs method has a receiver of type Vertex named v. */


package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}


// another example with fibonacci series==================
// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Vertex struct {
	start, end int
}
// we can remove function x,y parameter and directly use v.x, v.y
func (v Vertex) fibo(x, y int) []int {
	result := []int{}
	result = append(result, 0, 1)
	for i := x; i <= y; i++ {
		val := result[i-1] + result[i-2]
		fmt.Println(val)
		result = append(result, val)

	}
	return result
}

func main() {

	v := Vertex{2, 19}
	f := v.fibo(v.start, v.end)
	fmt.Println(f)

}

/* You can declare a method on non-struct types, too.

In this example we see a numeric type MyFloat with an Abs method.

You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int). */

// method type and receiver argument type must match

package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f,-math.Sqrt2)
	fmt.Println(f.Abs())
}


/* Pointer receivers
You can declare methods with pointer receivers.

// modify fields values from one method to another . its will impact to original data through out struct literal values. so that changed value will go in to next method

This means the receiver type has the literal syntax *T for some type T. (Also, T cannot itself be a pointer such as *int.)

For example, the Scale method here is defined on *Vertex.

Methods with pointer receivers can modify the value to which the receiver points (as Scale does here). Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

Try removing the * from the declaration of the Scale function on line 16 and observe how the program's behavior changes.

With a value receiver, the Scale method operates on a copy of the original Vertex value. (This is the same behavior as for any other function argument.) The Scale method must have a pointer receiver to change the Vertex value declared in the main function */




package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f //30
	v.Y = v.Y * f// 40 
	//30,40 values will go to abs method instead of 3,4 since its modifies the arguments values using pointer as arguments
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}




/* Methods and pointer indirection
Comparing the previous two programs, you might notice that functions with a pointer argument must take a pointer:

var v Vertex
ScaleFunc(v, 5)  // Compile error!
ScaleFunc(&v, 5) // OK

likewise
// pointer method receivers or value method receivers take either a value or pointer as receiver 
// while methods with pointer receivers take either a value or a pointer as the receiver when they are called:

var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
For the statement v.Scale(5), even though v is a value and not a pointer, the method with the pointer receiver is called automatically. That is, as a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver. */

package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	fmt.Println(v.X,v.Y)
	v.X = v.X * f
	v.Y = v.Y * f
	fmt.Println(v.X,v.Y)
}

func ScaleFunc(v *Vertex, f float64) {
fmt.Println(v.X,v.Y)
	v.X = v.X * f
	v.Y = v.Y * f
	fmt.Println(v.X,v.Y)
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}

// likewise
// pointer method receivers or value method receivers take either a value or pointer as receiver 

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}



// -------- own example --------
// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Vertex struct{ a, b int }

func operator(fn func(Vertex) int) int {
	return fn(Vertex{2, 3})
}

func (v *Vertex) show(num int) {
	v.a = v.a * num
	v.b = v.b * num
}

func showFunc(v *Vertex, num int) {
	v.a = v.a * num
	v.b = v.b * num
}

func main() {
	add :=
		func(v Vertex) int { return v.a + v.b }
	v := Vertex{2, 1}
	v.show(3)
	showFunc(&v, 1)

	p :=
		&Vertex{4, 3}
	p.show(2)
	showFunc(p, 1)

	fmt.Println("Hello, 世界", operator(add), v, p)
}
