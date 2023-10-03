package main

import "fmt"

func main() {
	fmt.Println("welcome to functions in golang")
	greeter()
	result := adder(4, 6)
	fmt.Println("Result is: ", result)

	proResult, myMsg := proAdder(2, 5, 6, 7, 3)
	fmt.Println("proresult is ", proResult)
	fmt.Println("My Message  is ", myMsg)
}
func greeter() {
	fmt.Println("Hello world from Golang")
}
func adder(valOne int, valTwo int) int {
	return valTwo + valOne
}
func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "This is the total value returned "
}
