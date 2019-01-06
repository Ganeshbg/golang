package main

import "fmt"

func main() {
	n,_ := fmt.Println("Just say Hello")
	fmt.Println(n)
	foo()
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("I am in foo")
}
