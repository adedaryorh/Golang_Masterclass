package main

import "fmt"

func main() {
	fmt.Println("Structs in Go")

	//no inheritance in GO; no super class or parents
	Adedayo := User{"Adedayo", "adedaryorh@gmail.com", true, 25}
	fmt.Printf("Adedayo details are: %+v\n", Adedayo)
	fmt.Printf("Name is  %v and email is %v.", Adedayo.Name, Adedayo.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
