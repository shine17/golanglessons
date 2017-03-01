package main

import (
	"fmt"
)

func main() {

	x := []string{"a", "b", "c", "d", "e", "f", "g"}
	fmt.Println(x)
	fmt.Println(x[2])   // get data from position 2 which is c
	fmt.Println(x[2:4]) // get data from postion 2 to 4 -1 which is 3 and result is c,d
}
