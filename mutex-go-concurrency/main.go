package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int
var mg sync.Mutex

func main() {
	wg.Add(2)
	go incrementor("Foo: ")
	go incrementor("Bar: ")
	wg.Wait()
	fmt.Println(counter)
}
func incrementor(ss string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(3 * time.Millisecond))
		mg.Lock()
		counter++
		fmt.Println(ss, i, "Counter: ", counter)
		mg.Unlock()
	}
	wg.Done()
}

//go run -race main.go
// mutex are to avoid race conditions ie two theards working on same resource. Here it is counter global variable
