package main

import (
	"fmt"
)

const MyJwtToken string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"

func main() {
	var name string = "Tunde"
	fmt.Println(name)
	fmt.Printf("I am variable %v new guy here of type: %T \n ", name, " ")

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n ", isLoggedIn)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type: %T \n ", smallValue)

	var alliasValue string
	fmt.Println(alliasValue)
	fmt.Printf("Variable is of type: %T \n ", alliasValue)

	myNewValue := "Tunday"
	fmt.Println(myNewValue)

	fmt.Println(MyJwtToken)

	car, cost := "bmw", 5000
	fmt.Println(car, cost)

	car, year := "bmw", 2018
	fmt.Println(car, year)

	var opened bool = false
	opened, file := true, "a.txt"
	_, _ = opened, file

	//multiple declaration
	var (
		salary    float64
		firstName string
		gender    bool
	)
	fmt.Println(salary, firstName, gender)

	var (
		a, b, ol int
	)
	fmt.Println(a, b, ol)

	var i, j int
	i, j = 5, 6
	//j, i = i, j
	_, _ = i, j
	fmt.Println(i, j)

	sum := 5 + 2.3
	fmt.Println(sum)

	const friend string = "tunde"
	const l, s, m int = 6, 8, 9
	fmt.Println(l, s, m)
	const (
		k int = 7
		o
		t
	)
	fmt.Println(k, o, t)

	//typed const
	const g float64 = 5.1
	//untyped or typeless const
	const p = 6.7
	//constant expression are always evaluated at compile time
	const c float64 = g * p
	const str = "hello" + "Go!"
	const y = 5 > 10
	fmt.Println(y)

	const r int = 5
	//This will not work cuy go doesn't allow it
	//const qa float64 = 2.4 * r
	const re = 5.6 * 5
	fmt.Println(re)
	fmt.Printf("The type of re: %T\n", re)

}

//Rules of const
/*
1. u cannot change a const
const temp int = 100
temp = 50
2. u cannot initiate a const at runtime
const power = math.power(2,3)
3. u cannot use a var to initialize const
t := 5
const tc = t
4. u can len function to initialize in a const
cuz len is a built in func
that's available at compile time but math is a function available at runtime
const l1 = len("hello")
*/
