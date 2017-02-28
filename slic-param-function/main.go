package main

import (
	"fmt"
)

func main() {
	c := []float64{1, 3, 36, 89}
	av := avg(c...)
	fmt.Println(av)
}

func avg(num ...float64) float64 {
	fmt.Println(num)
	fmt.Printf("%T \n", num)
	var total = 0.0
	for _, value := range num {
		total += value
	}

	return total / float64(len(num))
}

// avg function is called a variadic function since it take a slice of float64 values
