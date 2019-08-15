package main

import (
	"fmt"
)

//Scope global
var variableOutsideFunc int8 = 50

func main() {
	//Several type
	//1
	var variable int8 = 10
	//2
	var variable2 int8
	variable2 = 20

	//not declaring type
	//1
	var variable3 = 30
	//2
	variable4 := 40

	//Const cannot be changed
	const variableConst int8 = 60

	//multiple assign
	var (
		a = 1
		b = 2
		c = 3
	)

	fmt.Println(variable, variable2, variable3, variable4, variableOutsideFunc, variableConst, a, b, c)
}
