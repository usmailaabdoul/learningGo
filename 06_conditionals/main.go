package main

import "fmt"

func main() {
	x := 5
	y := 10

	// if else
	if x <= y {
		fmt.Printf("%d is less or equal than %d\n", x, y)
	} else {
		fmt.Printf("%d is greater than %d\n", y, x)
	}

	// else if
	color := "pink"

	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("color is NOT blue or red")
	}

	// Switch

}
