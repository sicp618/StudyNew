package main

import "fmt"

func main() {
	f4()
}

type Student struct {
	Name string
	Age int
}

func f4() {
	mp := make(map[int]Student)
	mp[1] = Student{"a", 1}
	mp[2] = Student{"b", 2}
	k1 := mp[1]
	fmt.Println(k1)
	k1.Name = "c"
	mp[1] = k1

	fmt.Println(k1)
}

func f3() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("f1")
		}
	}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("f2")
		}
	}()

	defer func() {
		func() {
			if err := recover(); err != nil {
				fmt.Println("f3")
			}
		}()
	}()

	panic("panic")
	fmt.Println("end")
	
}

func f2() {
	c := make(chan string, 1)
	close(c)
	fmt.Println(<-c)
	c<-"hello"
}

func f1() {
	var c chan string
	select {
		case c <- "hello":
			fmt.Println("hello")
		default:
			fmt.Println("default")
	}
}