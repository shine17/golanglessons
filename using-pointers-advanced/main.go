package main

import (
	"fmt"
)

func change(z *int) {

	*z = 2
}
func main() {
	a := 5
	change(&a)
	fmt.Println(a) // a changed to 2
}
