package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome to files concepts")

	content := "this is data for store in files"
	filepath := "./test.text"
	file, err := createFile(filepath)
	checkNillError(err)

	defer file.Close()

	length, err := writeFile(file, content)
	checkNillError(err)
	fmt.Println("The write Length of file ", length)

	databyte, err := ReadFile(filepath)
	checkNillError(err)
	fmt.Println("The read file data is ", string(databyte))

}

func createFile(path string) (*os.File, error) {
	file, err := os.Create(path)
	return file, err
}

func writeFile(file *os.File, content string) (int, error) {
	length, err := io.WriteString(file, content)
	return length, err
}

func ReadFile(filepath string) ([]byte, error) {
	file, error := ioutil.ReadFile(filepath)
	return file, error
}

func checkNillError(err error) {
	if err != nil {
		panic(err)
	}
}
