package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("45")

	if err != nil {
		panic(err)
	} else {
		fmt.Println(i)
	}

	if i, err = strconv.Atoi("95"); err == nil {
		fmt.Println("The value if i is ", i)
	} else {
		log.Fatal(err)
	}

	if args := os.Args; len(args) != 2 {
		fmt.Println("Please parse an argument ")
	} else if km, err := strconv.Atoi(args[1]); err != nil {
		fmt.Println("please parse an integer args", err)
	} else {
		fmt.Printf("%d km in miles is %vmiles\n ", km, float64(km)*0.621)
	}
}
