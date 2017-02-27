package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func myhandler(rw http.ResponseWriter, r *http.Request) {
	response, err := GetResponse()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(rw, string(response))

}

func main() {
	http.HandleFunc("/", myhandler)
	http.ListenAndServe("localhost:3456", nil)
}

type mydata struct {
	Stuff Data
}

type Data struct {
	Fruit       Fruits
	Vegetatable Vegetatables
}

type Fruits map[string]int
type Vegetatables map[string]int

func GetResponse() ([]byte, error) {
	myfruits := make(map[string]int)
	myfruits["apple"] = 12
	myfruits["banana"] = 55

	myvegetable := make(map[string]int)
	myvegetable["carrot"] = 15
	myvegetable["cabbage"] = 74

	d := Data{myfruits, myvegetable}
	x := mydata{d}
	return json.MarshalIndent(x, "", "  ")
}
