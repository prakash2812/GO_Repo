/* The range form of the for loop iterates over a slice or map.

When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index. */


package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

/*-------------------------------------- Basic Types */

/* bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
The example shows variables of several types, and also that variable declarations may be "factored" into blocks, as with import statements.

The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type. */


package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

// --------------- switch cases ---------------------------------------------
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("switch case dice example")

	rand.Seed(int64(time.Now().Nanosecond()))
	random := rand.Intn(6) + 1

	switch random {
	case 1:
		fmt.Println("you can begin and move to 1 step")
	case 2:
		fmt.Println("you can move to 2 step")
	case 3:
		fmt.Println("you can move to 3 step")
		// fall though it continue to below case until break
		fallthrough
	case 4:
		fmt.Println("you can move to 4 step")
		fallthrough
	case 5:
		fmt.Println("you can move to 5 step")
	case 6:
		fmt.Println("you can move to 6 step")
	default:
		fmt.Println("This is not expected number [0>=n<=6]")
	}
}

// -------------------- loops concepts -------------------------------------------

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


