package main

import "fmt"

const p string = "Death and taxex"
const (
	Pi       = 3.14
	Language = "GO"
	A        = iota //0
	B        = iota //1
	C        = iota //2
)
const (
	D = iota //0
	E = iota //1
	F = iota //2
)
const (
	_  = iota // First iota = 0
	KB = 1 << (iota * 10)
)

func main() {
	const q = 42
	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
	fmt.Println("Number Pi is ", Pi)
	fmt.Println("Language I study right now is ", Language)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println("Next block of constants")
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
	fmt.Println("Bytes to KB, MB, GB, TB")
	fmt.Printf("%b\t\t\t", KB)
	fmt.Printf("%d\n", KB)
}
