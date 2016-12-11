package main

import (
	"fmt"
)

func main() {

	for i := 3000; i < 8000; i++ {
		fmt.Printf("%d \t %b \t %X \t %q \n", i, i, i, i)
	}
}
