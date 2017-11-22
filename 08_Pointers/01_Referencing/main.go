package main

import "fmt"

func main() {
	q := 43
	fmt.Println(q)
	fmt.Println(&q)

	var b = &q
	fmt.Println(b)
}
