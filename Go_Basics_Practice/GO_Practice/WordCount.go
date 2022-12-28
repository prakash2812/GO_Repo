// Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.

package main

import (
	"fmt"
	"strings"
)

func wordCount(s string) map[string]int {
	arr := strings.Fields(s)
	fmt.Println("The array:", arr)
	// The array: [I ate a donut. Then I ate another donut.]
	m := make(map[string]int)
	for _, val := range arr {
		_, ok := m[val]
		if ok {
			m[val] += 1
		} else {
			m[val] = 1
		}

	}
	return m
}

func main() {
	fmt.Println("The map is ",wordCount("I ate a donut. Then I ate another donut.")) 
	// map[string]int{"I":2, "Then":1, "a":1, "another":1, "ate":2, "donut.":2}
}
