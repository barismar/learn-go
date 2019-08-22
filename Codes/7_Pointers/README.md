The implementation on the main.go file.

* Pointer is variable that point to memory location of another variable.
* new(type) is function that take type as an argument allocates enough memory to fit a value of that type and returns a pointer to it.

Problem
* How do you get the memory address of a variable?
    &variable

* How do you assign a value to a pointer?
    pointer *int
    pointer = &variable

* How do you create a new pointer?
    pointer *type

* What is the value of x after running this program:
```
func square(x *float64) {
  *x = *x * *x
}
func main() {
  x := 1.5
  square(&x)//2.25
}
```
* Write a program that can swap two integers (x := 1; y := 2; swap(&x, &y) should give you x=2 and y=1).

```
func swap(x, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
func main() {
  x := 1
	y := 2
	swap(&x, &y)
	fmt.Println(x, y)
}
```