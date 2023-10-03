package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in golang")

	rand.Seed(time.Now().UnixNano())
	diceNum := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNum)

	switch diceNum {
	case 1:
		fmt.Println("Dice value is 1 and u can open")
	case 2:
		fmt.Println("You can move to 2 spot ")
	case 3:
		fmt.Println("You can move to 3 spot ")
		fallthrough
	case 4:
		fmt.Println("You can move to 4 spot ")
		fallthrough
	case 5:
		fmt.Println("You can move to 5 spot ")
	case 6:
		fmt.Println("You can move to 6 spot and play again ")
	default:
		fmt.Println("What is this ?")
	}
}
