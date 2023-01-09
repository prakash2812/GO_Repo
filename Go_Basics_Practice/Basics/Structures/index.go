/* Structs
A struct is a collection of fields. */


package main

import "fmt"

type Vertex struct {
	Name string
	Age int
	
}

func main() {
	arr:= []int{}
	var data Vertex
	data.Name="prakash"
	fmt.Println(data)
	fmt.Println(Vertex{"arjun", 28})
}

/* Struct fields are accessed using a dot. */

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	id := []Vertex{1, 2}
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
	
	for idx,val:=range id {
		fmt.Println("ket , val", idx,val)
	}
	
}

/* Pointers to structs
Struct fields can be accessed through a struct pointer.

To access the field X of a struct when we have the struct pointer p we could write (*p).X. However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference. */



package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 19
	fmt.Println(v)
}


/* Struct Literals
A struct literal denotes a newly allocated struct value by listing the values of its fields.

You can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.)

The special prefix & returns a pointer to the struct value. */

package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}


// ------------ with array of struct ype
//struct is collection of fields {a;b}
//array/slice is storage of same fields but size will be fixed/dynamic

package main

import (
	"fmt"
)

func main() {
	storage := []struct {
		name   string
		salary float64
	}{{"arjun", 28}, {"prakash", 23}}


	for _, item := range storage {
		fmt.Printf("%T", item)
		updateEmployeeSalary(&item, 23)
	}

	fmt.Printf("update dat is ", storage)

}

func updateEmployeeSalary(data *struct {
	name   string
	salary float64
}, days uint64) {
	data.salary = float64(days) * 2000
}

// -------- same above can be done using struct so we don't need to mention long type for parameters 

package main

import (
	"fmt"
)

type EmpTemplate struct {
	name   string
	salary float64
}

func main() {
	/* var play []struct{a,b int}
	var play1 []EmpTemplate
	play = append(play, struct{a int; b int}{2,3})
	play1 = append(play1, EmpTemplate{"arjun",23}) */

	storage := []EmpTemplate{{"arjun", 28}, {"prakash", 23}}

	for _, item := range storage {
		updateEmployeeSalary(&item, 23)
	}

	fmt.Printf("update dat is ", storage)

}

func updateEmployeeSalary(data *EmpTemplate, days uint64) {
	data.salary = float64(days) * 2000
}
