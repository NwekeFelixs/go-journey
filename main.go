package main

import (
	"fmt"
)
func main() {
	// Define a string variable
	var str string = "Hello, World!"

	// Print the string variable
	fmt.Println(str)

	// Print the length of the string
	fmt.Println("Length of the string:", len(str))

	// Print the first character of the string
	fmt.Println("First character:", string(str[0]))
}