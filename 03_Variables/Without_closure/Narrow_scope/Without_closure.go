package main

import "fmt"

func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println("The value ot the X after running incriment() func for first time ")

	fmt.Println(increment())
	fmt.Println("The value ot the X after running incriment() func for second time ")
}
