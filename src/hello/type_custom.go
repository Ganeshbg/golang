package main 

import "fmt"

type gan int
func main() {

	var x gan
	x = 34
	fmt.Printf("%T\n",x)
	fmt.Println(x)
	var y int 
	//conversion not casting
	y = int(x)
	fmt.Println(y)
}