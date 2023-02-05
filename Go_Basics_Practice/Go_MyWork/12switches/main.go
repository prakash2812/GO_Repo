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
