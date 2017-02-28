package main

import (
	"fmt"
)

func main() {
	z, data := getSum(2, 3)
	fmt.Println(data, z)
}

func getSum(x int, y int) (z int, data string) {
	return x + y, "result is "
}
