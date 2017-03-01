package main

import (
	"fmt"
)

func main() {

	dd := make(map[string]int)
	dd["x"] = 1
	dd["y"] = 2
	dd["z"] = 3

	fmt.Println(dd)
	fmt.Println(dd["x"])

	delete(dd, "y")
	fmt.Println(dd)

	_, re := dd["y"]
	fmt.Println(re)

}
