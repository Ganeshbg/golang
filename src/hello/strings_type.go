package main

import "fmt"

func main() {
	s := "Hello World!"
	fmt.Println(s)
	bs := []byte(s)
	for i := 0; i<len(s);i++ {
		fmt.Printf("%#U\n",bs[i])
	}
}