package main

import (
	"fmt"
	"sync"
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
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	storage := []EmpTemplate{{"arjun", 28}, {"prakash", 23}}

	wg.Add(3)
	go updateEmployeeSalary(storage, wg, mut)
	go updateEmployeeSalary(storage, wg, mut)
	go updateEmployeeSalary(storage, wg, mut)
	// }
	wg.Wait()
	mut.Lock()
	fmt.Println("update date is ", storage)
	mut.Unlock()

}

func updateEmployeeSalary(storage []EmpTemplate, wg *sync.WaitGroup, mut *sync.Mutex) {
	defer wg.Done()
	for index, item := range storage {
		mut.Lock()
		storage[index].salary *= item.salary
		mut.Unlock()
	}
	mut.Lock()
	fmt.Println("show storage after each iteration of goroutine", storage)
	mut.Unlock()
}
