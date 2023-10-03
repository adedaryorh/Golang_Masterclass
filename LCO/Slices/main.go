package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in go ")

	var fruitList = []string{"Apples", "Tomatoes", "Mango"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana", "Cashew")
	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	//memory management
	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 344
	highScores[2] = 465
	highScores[3] = 2865
	//highScores[4] = 286

	highScores = append(highScores, 453, 747, 878)

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)

	//removing values form slices based on index

	var courses = []string{"python", "javascript", "swift", "java", "ruby", "kotlin"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
