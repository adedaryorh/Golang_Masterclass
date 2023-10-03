package main

import (
	"fmt"
	"math"
)

func main() {
	var x uint8 = 255
	//overflow, x is 0
	x++
	fmt.Println(x)

	a := uint16(255 + 1)
	fmt.Println(a)

	var b int8 = 127
	fmt.Printf("%d\n", b+1)

	b = -128
	b--
	fmt.Println(b)

	f := float32(math.MaxFloat32)
	fmt.Println(f)
	//float number overflows to infinite!!
	f = f * 1.2
	fmt.Println(f)
}
