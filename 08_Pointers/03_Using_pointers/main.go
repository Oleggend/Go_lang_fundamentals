package main

import "fmt"

func main() {
	value := 43
	fmt.Println(value)  //Prints 43
	fmt.Println(&value) // & + variable will print haxadecimal address of variable

	var address = &value  // Creates pointer b to haxadecimal address of variable
	fmt.Println(address)  // Prints pointer
	fmt.Println(*address) // b is pointer to the memory address, * in this case is an operator which allowes to us get actual value stored in this exact memory address.

	*address = 42      // In this line ve using hex address to change valu stored at this address
	fmt.Println(value) // At firs variable "value" has value of 43 (line 6), but code in line 14 will change value from 43 to 42

}
