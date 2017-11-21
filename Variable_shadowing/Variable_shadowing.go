package main

import "fmt"

func max(x int) int {
	return 50 + x
}

func main() {
	max := max(7)
	fmt.Println(max) //max now is result, not the function
}

// DO NOT USE THIS CODING PRACTISCE TO "SHADOW" VARIABLES
