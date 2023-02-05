package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("WElcome time related things")

	presentTime := time.Now()

	fmt.Println("Present time", presentTime)

	fmt.Println("Format time is ", presentTime.Format("01-02-2006 15:04:05 Monday"))

	createDate := time.Date(2004, time.December, 28, 12, 28, 28, 0, time.Local)
	fmt.Println("create  date", createDate)
	fmt.Println("format create date", createDate.Format("01-03-2005 Wednesday"))

}
