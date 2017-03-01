package main

import (
	"encoding/json"
	"fmt"
)

type mydata struct {
	FirstName  string
	SecondName string
}

func main() {
	p := mydata{FirstName: "shine", SecondName: "sivadasan"}
	x, _ := json.Marshal(p)
	fmt.Println(x)
	fmt.Printf("%T \n", x)
	fmt.Println(string(x))
}
