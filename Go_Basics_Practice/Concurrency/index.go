// GoRoutines are independently executing functions in same memory/address space
go f()
go g(1,2)
// Channels are typed value that allow goroutine to synchronize and exchange information
// communicate between two goroutines while sending r receiving data synchronously, will wait until it finishes the other channels

c: make(chan int)
go f(){c<-3}
n:= <-c

package main

import (
	"fmt"
	"time"
)

type Ball struct{ num int }

func main() {
	table := make(chan *Ball)
	fmt.Println(table)
	go Player("ping", table)
	go Player("pong", table)

	//fmt.Println("here0", table)
	table <- new(Ball)
	fmt.Println("here1", table)
	time.Sleep(1 * time.Second)
	//w := <-table
	<-table
	//fmt.Println("over", w)

}
func Player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.num++
		fmt.Println("here", name, ball.num)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}

//when we do goroutines they will run independently / separately
// channels will help to make a connection between main and the all goroutines
// also they maintain all go routines each other they share memory by communication, will wait each other and understand each other when to send and recieve.

// its our responsibility to tell or wait until goroutines are done and execute synchronously with the main flow else it will show nothing means we forgot like our brain

//--------------------this can be done using channels or wait WaitGroup

//if channels -->then receiver channel should be mentioned after the go routine. not again inside another thread. should in main thread

package main

import (
	"fmt"
	"time"
)

func main() {
	mych := make(chan int)

	go func(ch chan<- int) {
		ch <- 7

	}(mych)
	go func(ch <-chan int) {
		fmt.Println("receiver", <-ch)

	}(mych)
	go func(ch chan<- int) {
		ch <- 9

	}(mych)
	fmt.Println("value", <-mych, <-mych)

	n :=
		boring("boring")
	for i := 0; i < 5; i++ {
		fmt.Println(<-n)
	}

}

// function return receiver chanel, we can call in main then receive the chanel so that it will wait of main thread not to finish main thread
func boring(msg string) <-chan string {
	ch := make(chan string)

	go func(ch chan string) {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(1 * time.Second)
		}

	}(ch)
	return ch
}


//if wait group --> below examples manually mention wait and all other methods since its third service..

// WaitGroup are like third party service which knows how many r goroutines will add and done or wait until done. so no connection between goroutines by themselves
//wg:= *sync.WaitGroup{} // Pointer and struct(class) it has wg.ADD(int) , Wait() and Done() (class.methods) .. i mean not functions..

// ->>>>>>wg is global declared  since we not needed pointers to main the connection of method modification

func main() {
	mych := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func(ch <-chan int) {
		fmt.Println("value", <-ch)
		wg.Done()
	}(mych)
	go func(ch chan<- int) {
		ch <- 5
		wg.Done()
	}(mych)

	wg.Wait()
}

// wg is global since we passed arguments pointers og wg to main the connection of method modification in each go routines. if we remove pointer and pass a value of wg normally. it will be dead lock at wait method line.

// either you do globally without arguments or continue with pointers 
func main() {
	mych := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		fmt.Println("value", <-ch)
		wg.Done()
	}(mych, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 5
		wg.Done()
	}(mych, wg)

	wg.Wait()
}


// Mutex are block the memory until the go routines finishes writing and block the all other go routines until finished
// because all go routines run in a single memory space and if many threads try to write at once it may cause issues
// mut:= *sync.Mutex{} (class and methods concept)  mut.Block() and mut.UnBlock()

package main

import (
	"fmt"
	"net/http"
)

func main() {

	websites := []string{
		"https://go.dev",
		"https://fb.com",
		"https://google.com",
	}

	for _, val := range websites {
		getStatusCode(val)
	}

}

func getStatusCode(url string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("opps something went wrong", err)
	} else {
		fmt.Printf("%d status code %s", res.StatusCode, url)
	}

}
