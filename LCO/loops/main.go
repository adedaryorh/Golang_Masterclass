package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to loopa in Go")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}
	fmt.Println("**********************")
	//forEach
	for index, day := range days {
		fmt.Printf("index is %v and values is %v\n", index, day)
	}
	rougeValue := 1

	for rougeValue < 10 {
		if rougeValue == 2 {
			goto lco
		}
		if rougeValue == 5 {
			continue
		}
		fmt.Println("Value is ", rougeValue)
		rougeValue++
	}

lco:
	fmt.Println("Jump onto learncodeonline!")
}
