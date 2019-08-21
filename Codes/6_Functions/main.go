package main

import "fmt"

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

//with no return
func printSomething() {
	fmt.Println("HAI")
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

//Recursion
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	//basic with returning values
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))

	//basic no returning values
	printSomething()

	//Variadic Function
	fmt.Println(add(1, 2, 3))

	//Closure (add function on function)
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())

	factorialAns := factorial(5)
	fmt.Println(factorialAns)

	//Defer
	defer func() {
		str := recover()
		fmt.Println(str)
		fmt.Println("cek")
	}()
	panic("PANIC")
}
