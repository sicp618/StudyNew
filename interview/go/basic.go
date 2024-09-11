package main

import (
	"fmt"
	"reflect"
)

func init() {
	init2()
	fmt.Println("init")
}

func init2() {
	fmt.Println("init2")
}

func main() {
	fmt.Println("Hello, World!")
	s1 := make([]string, 2)
	fmt.Println("init slice", s1, reflect.TypeOf(s1))

	fns1(s1)
	fmt.Println(s1)

	fns2(&s1)
	fmt.Println(s1)

	s2 := new ([5]string)
	fmt.Println("init array", s2, reflect.TypeOf(s2), reflect.TypeOf(*s2))
	
	fns3(s2)
	fmt.Println(*s2)
}

func fns1(s []string) {
	s[0] = "fn1"
}

func fns2(s *[]string) {
	(*s)[0] = "fn2"
}

func fns3(s *[5]string) {
	s[0] = "fn3"
}

