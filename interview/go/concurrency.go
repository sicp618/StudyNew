package main

import (
	"fmt"
	"sort"
)

const BB = 1

func main() {
	mp := make(map[string][]int)
	mp["a"] = []int{1, 2, 3}
	if v, ok := mp["a"]; ok {
		v[0] = 100
		fmt.Println(v)
	} else {
		fmt.Println("not found")
	}
	j1 := make([]int, 5)
	fmt.Println(j1)
	is := []int32{3, 2, 1}
	f4(is)
	fmt.Println(is)
	a := 3
	switch a {
		case 1:
		case 2:
			
	}
}

func f4(is []int32) {
	sort.Slice(is, func(i, j int) bool {
		return is[i] < is[j]
	})
}

func f3() {
	c := boring3(boring1(), boring1())
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
}

func boring3(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1: c <- s
			case s := <-input2: c <- s
			}
		}
	}()

	return c
}

type Message struct {
	str string
	wait chan bool
}

func f2() {
	c := make(chan bool)
	msg := Message{"ping", c}
	fmt.Println(msg.str)
	c <- true

}

func boring2(c chan Message) {
	for i := 0; i < 5; i++ {
		msg1 := <-c
		msg2 := <-c
		msg1.wait <- true
		msg2.wait <- true
	}
}

func f1() {
	c := boring1()
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
}

func boring1() <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			c <- fmt.Sprintf("boring %d", i)
		}
	}()
	return c
}
