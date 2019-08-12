package main

import (
	"fmt"
)

func main() {

	//Number
	var integer int8
	integer = -10
	var unsignedInteger uint8
	unsignedInteger = 10
	fmt.Println(integer, unsignedInteger) //-10 10

	var floatingPointNumber float32
	floatingPointNumber = 10.12
	fmt.Println(floatingPointNumber) //10.12

	//String
	fmt.Println(len("Hello World")) //11
	fmt.Println("Hello World"[1])   //101
	fmt.Printf("%c\n", "Hello World"[1])
	fmt.Println("Hello " + "World")

	//Boolean
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
