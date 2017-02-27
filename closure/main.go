package main

import "fmt"

func main() {
	//var x int = 42
	x := 42
	fmt.Println(x)
	{
		y := "something"
		fmt.Println(y)
	}
	// A closure is a function value that references variables from outside its body.
	//The function may access and assign to the referenced variables;
	// in this sense the function is "bound" to the variables.
}
