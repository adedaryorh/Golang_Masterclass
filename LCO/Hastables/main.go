package main

import "fmt"

//maps

func main() {
	fmt.Println("Maps in Golang")

	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["RB"] = "ruby"
	languages["PY"] = "python"

	fmt.Println("list of all languages :", languages)
	fmt.Println("list of all languages :", languages["PY"])

	delete(languages, "RB")
	fmt.Println(languages)

	//conditional loops
	for _, value := range languages {
		fmt.Printf("for key v, values is %v\n", value)
	}

}
