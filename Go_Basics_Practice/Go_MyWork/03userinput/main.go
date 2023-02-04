package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the Course rating: ")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thank you for rating: ", input)
	fmt.Printf("rating type is %T", input)
}
