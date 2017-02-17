package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
//	"log"
)

func main() {
	resp, _ := http.Get("http://shop.nordstrom.com")
    /*resp, err := http.Get("http://shop.nordstrom.com")
    if err != nil{
        log.Fatal(err)
    }*/
	page, _ := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
	fmt.Println(string(page))
}
// Here _ is called blank indentifier.
// Basically we specify this if we don't want to assign it to a varible and use it.
// here the return value is error which we don't want to use