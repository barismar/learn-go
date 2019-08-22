package main

import (
	"fmt"
)

func zero(x int) {
	x = 0
}

func zeroPointer(x *int) {
	*x = 0
}

func one(xPtr *int) {
	*xPtr = 1
}

func swap(x, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func main() {
	a := 5
	zero(a)
	fmt.Println(a) //5

	zeroPointer(&a)
	fmt.Println(a) //0

	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr) // x is 1

	checkMemo := 0
	fmt.Println(&checkMemo)

	var thisIsPointer *int
	thisIsPointer = &a
	fmt.Println(*thisIsPointer)

	x := 1
	y := 2
	swap(&x, &y)
	fmt.Println(x, y)

	var objek int
	objek = 11
	fmt.Println("Objek ", objek, &objek)

	var pointer1 *int
	pointer1 = &objek
	fmt.Println("Pointer1 ", *pointer1, &pointer1)

	var pointer2 *int
	pointer2 = &*pointer1
	fmt.Println("Pointer2 ", *pointer2, &pointer2)

}
