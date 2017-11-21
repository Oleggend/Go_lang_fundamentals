package main

import "fmt"

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	incriment := wrapper()
	fmt.Println(incriment())
	fmt.Println("The value ot the X after running incriment() func for first time ")

	fmt.Println(incriment())
	fmt.Println("The value ot the X after running incriment() func for second time ")
}
