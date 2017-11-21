package main

import "fmt"

var x = 0

func incriment() int {
	x++
	return x
}

func main() {
	fmt.Println(incriment())
	fmt.Println("The value ot the X after running incriment() func for first time ")

	fmt.Println(incriment())
	fmt.Println("The value ot the X after running incriment() func for second time ")
}
