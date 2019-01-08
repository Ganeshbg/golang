package main

import (
	"fmt"
)

//Typed constants
const (
	name string = "Ganesh"
	ID int = 1234
	dec float32 = 12.23
)

func main(){

	fmt.Println(name)
	fmt.Println(ID)
	fmt.Println(dec)
	fmt.Printf("%T\n",name)
	fmt.Printf("%T\n",ID)
	fmt.Printf("%T\n",dec)
}