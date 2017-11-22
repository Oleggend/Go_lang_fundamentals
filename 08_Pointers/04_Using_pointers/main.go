package main

import "fmt"

func zero(x int) {
	x = 0
	fmt.Println("Variable x in function zero has hardcoded value: ", x)

	fmt.Println("The memory address of variable x in zero func is: ", &x)
}

func main() {
	x := 5
	fmt.Println("Variable x in func main has hardcoded value - ", x)
	fmt.Println("The memory address of variable x in main func BEFORE CALLING ZERO FUNC is: ", &x)
	fmt.Println()
	fmt.Println("Calling zero function")
	zero(x)
	fmt.Println("The memory address of variable x in main func AFTER CALLING ZERO FUNC is:", &x)
	fmt.Println("Variable x value is: ", x)

}
