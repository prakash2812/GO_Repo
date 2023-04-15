/* An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

The type []T is a slice with elements of type T.

A slice is formed by specifying two indices, a low and high bound, separated by a colon:

a[low : high]
This selects a half-open range which includes the first element, but excludes the last one.

The following expression creates a slice which includes elements 1 through 3 of a:

a[1:4] */

 /*---------------- Slices are like references to arrays---------------------------- */

/* A slice does not store any data, it just describes a section of an underlying array.

Changing the elements of a slice modifies the corresponding elements of its underlying array.

Other slices that share the same underlying array will see those changes. */

package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

/* ====================================Slice literals================================*/
/* A slice literal is like an array literal without the length.

This is an array literal:

[3]bool{true, true, false}
And this creates the same array as above, then builds a slice that references it:

[]bool{true, true, false}  */
package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		id          int
		isPermanent bool
		name        string
		age         int
		address     string
	}{
		{2, true, "arjun", 28, "Bangalore"},
		{2, true, "prakash", 28, "Hyderabad"},
		{2, true, "rathode", 28, "Bangalore"},
		{2, true, "John", 28, "Bangalore"},
	}
	for idx, val := range s {
		fmt.Println(idx, val)
		fmt.Println(val.id)

		fmt.Println(val.isPermanent)

		fmt.Println(val.name)
		fmt.Println(val.age)

	}
	fmt.Println(s)
}
/* -----------------------------------------Creating a slice with make ---------------------------*/
/* Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.

The make function allocates a zeroed array and returns a slice that refers to that array:

a := make([]int, 5)  // len(a)=5
To specify a capacity, pass a third argument to make:

b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4 */



package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}



/* ------------------------------------------- A slice has both a length and a capacity. ------------------ */

/* The length of a slice is the number of elements it contains.

The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).

You can extend a slice's length by re-slicing it, provided it has sufficient capacity. Try changing one of the slice operations in the example program to extend it beyond its capacity and see what happens. 

Slices can contain any type, including other slices.
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	fmt.Println
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
/*-------------------------------append works on nil slices  -------------     */
/* Appending to a slice
It is common to append new elements to a slice, and so Go provides a built-in append function. The documentation of the built-in package describes append.

func append(s []T, vs ...T) []T
The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.

The resulting value of append is a slice containing all the elements of the original slice plus the provided values.

If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.

(To learn more about slices, read the Slices: usage and internals article.) */

package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// --------------------------------------------------------

package main

import (
	"fmt"
	"sort"
)

func main() {
	data := []string{"hero"} // [hero]
	data[0] = "chang"
	fmt.Println(data) //[chang]

	data1 := append(data, "prakash")//[chang,prakash]
	data1[1] = "change"
	fmt.Println(data1) //[chang change]

	scroes := make([]int, 5) // [0,0,0,0,0]

	big := append(scroes[:4]) // [0,0,0,0]
	big[2] = 10

	fmt.Println(scroes, big) //[0 0 10 0 0] [0 0 10 0]

	data := []string{"sir", "AI", "prakash", "arjun", "david"}
	
	// sorting the data of int and strings from asscending to desending order
	sort.Strings(data)
	fmt.Println(data) //[AI arjun david prakash sir]
	data2 := []int{23,45,6,12,34,576,234,6878,456,12,1}
	sort.Ints(data2)
	fmt.Println(data2)//[1 6 12 12 23 34 45 234 456 576 6878]

	// delete the array using range concept a[low:high]
	courses:= []string{"react","js","go","php","AI","dataScience"}
	index :=3
	courses = append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)

}


// ---------------------
package main

import (
	"fmt"
	"time"
)

func recursive(x int) (int){
	if(x>0){
		return x*recursive(x-1)
	}else{
		return 1
	}	
}
func myFunction (name string,age int) ([]string,int) {	
	return []string{name},age
}
func main(){
	fact:=recursive(4)
	fmt.Println(fact)
	name,age:=myFunction("arjun",28)
	fmt.Println(name,age)
	fmt.Println("time",time.Now())
	// array define using var and :=
	var data = [5]int{1,2,3,4,5} //defined length
	// data1:= [...]int{}//inferred length or undefined length

	// slice 
	// var slice = []int{} // using var or :=
	// slice1:=[]int{}
	// create slice from array
	slice2:= data[2:4] // data[start:end]

	// using make method
	slice3:= make([]int,5,5) // make([]type,length,count)

	//append elements,append slices,copy

	slice4:= append(slice3,4,5,6)

	slice5:=append(slice3,slice4...)
	printSlice
	dest :=make([]int,len(slice5))
	copy(dest,slice5) // copy (dest, source)
	fmt.Println("started")
	for i:=0;i<len(slice5);i++ {
		// fmt.Println(slice5[i])
	}
	fmt.Println(slice2)
	for index,val:= range slice2 {
		fmt.Println(index,val)
	}
}

// ---------------------------------------------------------------------------

package main

import ("fmt")

func main() {
	data := []int{2,3,4}
	data = append(data[:len(data)])
	fmt.Println(data)
}