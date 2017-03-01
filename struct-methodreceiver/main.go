package main

import (
	"fmt"
)

type mydata struct {
	firstname  string
	secondname string
}

func (data mydata) concatenate() string {
	return data.firstname + " " + data.secondname
}

//func receivername functionname inputparams returnparams
//here data mydata is called receivername
// if we use this struct as receiver , then any variable which uses this struct can called this function concatenate
// here variabe x is assigned as type stuct, x can call this method

func main() {
	x := mydata{"shine", "sivadasan"}
	fmt.Println(x.concatenate())
}
