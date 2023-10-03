package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the pizzastore!")
	fmt.Println("Please rate our pizza between 1 and 5: ")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating our pizza with: ", input)
	//fmt.Println("Your rating is added plus 1: ", input)

	plusOneRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Your rating is added plus 1: ", plusOneRating+1)
	}
	castNum()

}

// typecasting in numbers
func castNum() {
	var x int = 3
	var y float32 = 3.1

	h := x * int(y)
	fmt.Println(h)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", int(y))

	f := int(float32(x) * y)
	fmt.Println(f)

	s := string(99)
	fmt.Println(s)

	var mystr = fmt.Sprintf("%f", 44.2)
	fmt.Println(mystr)
	fmt.Printf("%T\n", mystr)

	var f1, err = strconv.ParseFloat(mystr, 64)
	_ = err
	fmt.Println(f1 * 3.4)

	var mystr1 = fmt.Sprintf("%d", 34234)
	fmt.Println(mystr1)

	//unicode character
	fmt.Println(string(34234))

	i, err := strconv.Atoi("-50")
	s2 := strconv.Itoa(20)

	fmt.Printf("i type is %T, i value is %v\n", i, i)
	fmt.Printf("s2 type is %T, s2 value is %q\n", s2, s2)

}
