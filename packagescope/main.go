package main

import (
	"fmt"
)

var x = 23

func main() {
	println(x)

	y := 34
	println(y)
	fmt.Printf("%d is in block level scope and cannot be accessed in foo function \n", y)

	fmt.Print("y cannot be accessed before declaring or intializing line since order matters")

	foo()
}

func foo() {
	println(x)
}
