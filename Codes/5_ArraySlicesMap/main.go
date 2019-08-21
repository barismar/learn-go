package main

import (
	"fmt"
)

func main() {
	//Array
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))

	//Slices

	arr := [5]float64{1, 2, 3, 4, 5}
	y := arr[0:5]
	fmt.Println(y)

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
	//slice1 has [1,2,3] and slice2 has [1,2,3,4,5]

	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice3, slice4)
	fmt.Println(slice3, slice4)

	//Map
	z := make(map[string]int)
	z["key"] = 10
	fmt.Println(z["key"])
}
