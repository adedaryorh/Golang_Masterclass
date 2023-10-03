package main

import "fmt"

func main() {
	fmt.Println("welcome to a class on pointer")
	var one *int
	fmt.Println("value of pointer is", one)

	myFirstValue := 56
	var newValue = &myFirstValue

	fmt.Println("Value of actual pointer is", newValue)
	fmt.Println("Value of actual pointer is", *newValue)

	*newValue = *newValue + 2
	fmt.Println("New Value is: ", myFirstValue)

}
