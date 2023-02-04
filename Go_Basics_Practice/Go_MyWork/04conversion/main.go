package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to our app")
	fmt.Println("Enter the movie rating---")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thank for giving rating: ", input)
	fmt.Println("vvallue", input, strings.TrimSpace(input))
	increment, ok := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if ok != nil {
		fmt.Println("There is something error")
	} else {

		fmt.Println("Increment 1 more rating", increment+1)
	}

}
