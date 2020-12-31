package main

import "fmt"

func main() {
	// Arrays
	// var fruitArr [2]string

	// Assing values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Declare and assign
	fruitArr := [2]string{"Apple", "Orange"}

	fruitSlice := []string{"Apple", "Orange", "Mango"}

	fmt.Println(fruitArr)
	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:2])
}
