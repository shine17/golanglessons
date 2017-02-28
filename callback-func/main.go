package main

func main() {
	ff := []int{2, 2, 3, 8}
	visit(func(n int) {
		println(n)
	}, ff...)

}

func visit(callbk func(int), xs ...int) {
	for _, value := range xs {
		callbk(value)
	}

	//call back function - passing an anonymous func as argument
	// variadic or slice parameter should be at the end of the parameter list
}
