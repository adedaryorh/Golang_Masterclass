package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the time class in golang")
	currentTime := time.Now()
	fmt.Println(currentTime.Format("01-02-2006 15:04:05 Monday"))
	currentDate := time.Date(2020, time.May, 21, 23, 23, 0, 0, time.UTC)
	fmt.Println(currentDate.Format("01-02-2006 15:04:05 Monday"))
	fmt.Println(currentDate)
}
