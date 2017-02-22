package main

import (
	"fmt"
)

const ss float64 = 1.14

func main() {
	var a float64

	fmt.Scan(&a)
	c := a * ss
	fmt.Println(c)
}
