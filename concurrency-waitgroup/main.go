package main

import "fmt"
import "sync"
import "time"

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo: ", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}
func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("bar: ", i)
		time.Sleep(time.Duration(20 * time.Millisecond))
	}
	wg.Done()
}

// Executing concurrent method you need to define wait group and add the number of method to it and
//ask the wait group to wait. Also you need to put waitgroup.Done in those method to take those method from wait group once completed.
//when wg.done is excecuted it will take that method from wait group. wg.wait becomes 0 , the program will execute the main func

// concurrency - idependenty executing lots of things at once
// parallelism - simulatenous execution of (possibly related) computation

//go run -race main.go  ---> to check the race condition
