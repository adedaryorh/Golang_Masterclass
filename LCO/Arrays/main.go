package main

import "fmt"

func main() {
	fmt.Print("Welcome to array in Go")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[2] = "Peach"

	fmt.Println("Fruit list are: ", fruitList)
	fmt.Println("Fruit list are: ", len(fruitList))

	var vegList = [6]string{"potatoes", "beans", "mushrooms"}
	fmt.Println("Veg list are ", vegList)
	fmt.Println("My Vegetable list is: ", len(vegList))
}
