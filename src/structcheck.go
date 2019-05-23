package main

import (
	"fmt"
)

type Address struct {
	Type string
	Ipadd string
}
func main() {
	add := []Address{}
	ad :=Address{Type:"e",Ipadd:"12.14"}
	ad1 := Address{Type:"i",Ipadd:"12.10"}
	
	add = append(add,ad1)
	add = append(add,ad)

	fmt.Println(add)

	//y := add[.Type == 'e']
	fmt.Println(y)

}