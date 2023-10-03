package main

import "fmt"

func main() {
	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)
	fmt.Println(c1, c2, c3)

	const (
		c11 = iota
		c22
		c33
	)
	fmt.Println(c11, c22, c33)

	const (
		North = iota * 2
		East
		South
		West
	)
	fmt.Println(West)

	//XX = -2, YY= -4, MM=-5 using iota
	const (
		xx = -(iota + 2) //-2
		_
		yy //-4
		mm //-5
	)
	fmt.Println(xx, yy, mm)
}
