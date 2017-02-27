package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type mydata struct {
	Stuff Data
}

type Data struct {
	Fruit       Fruits
	Vegetatable Vegetatables
}

type Fruits map[string]int
type Vegetatables map[string]int

func main() {
	payload, err := http.Get("http://localhost:3456")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(payload.Body)
	if err != nil {
		panic(err)
	}
	var p mydata
	err = json.Unmarshal(body, &p)
	if err != nil {
		panic(err)
	}
	fmt.Println("veggie=", p.Stuff.Vegetatable, "\n", "fruit=", p.Stuff.Fruit)

}
