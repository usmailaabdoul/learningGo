package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64
	// byte - alias for uint8
	// rune - alias for uint8
	// float32 float64
	// complex64 complex128

	// Using var
	// var name string = "Usmaila"
	var age int32 = 22
	age = 25

	const isCool = true

	// shorthand
	// name := "Usmaila abdoul"
	size := 2.5

	name, email := "Usmaila", "ismaelabdul77@gmail.com"

	fmt.Println(name, age, isCool, size, email)
	fmt.Printf("%T\n", size)
}
