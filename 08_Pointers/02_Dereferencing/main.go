package main

import "fmt"

func main() {
	value := 43
	fmt.Println(value)  //Prints 43
	fmt.Println(&value) // & + variable will print haxadecimal address of variable

	var address = &value  // Creates pointer b to haxadecimal address of variable
	fmt.Println(address)  // Prints pointer
	fmt.Println(*address) // b is pointer to the memory address, * in this case is an operator which allowes to us get actual value stored in this exact memory address.

}
