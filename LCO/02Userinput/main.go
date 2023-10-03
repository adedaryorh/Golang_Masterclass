package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Println("******************")

	fmt.Println("os.Args", os.Args)
	fmt.Println("1st argument path: ", os.Args[0])
	fmt.Println("2nd argument: ", os.Args[1])
	fmt.Println("3rd argument: ", os.Args[2])
	fmt.Println("Total number of args inside os.Args ", len(os.Args))

	//convert string to float
	var result, err = strconv.ParseFloat(os.Args[1], 64)
	fmt.Printf("%T\n", os.Args[1])
	fmt.Printf("%T\n", result)
	_ = err

	userName := "Accepting a User Input"
	fmt.Println(userName)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza: ")

	// comma ok !! err.
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	//fmt.Println("Type of this rating is: ", reflect.TypeOf(input)
	fmt.Printf("Type of this rating is:  %T", input)

}
