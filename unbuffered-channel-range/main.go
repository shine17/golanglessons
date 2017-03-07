package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for val := range c {
		fmt.Println(val)
	}
	// here we don't require time.sleep since the execution stop while ranging over the channel
}
