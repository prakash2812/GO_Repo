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