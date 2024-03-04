package main

import "fmt"

var outX int = 0

func outerFunc() func(int) int {
	x := 10
	return func(y int) int {
		x++
		outX++
		return x + y + outX
	}
}

func main() {
	innerFunc := outerFunc()
	fmt.Println(innerFunc(20))
	fmt.Println(innerFunc(20))

	f1(innerFunc)
	in2 := outerFunc()
	fmt.Println(in2(20))
}

func f1(in func(int) int) {
	fmt.Println(in(20))
}
