package main

import "fmt"

func main() {
	// Exercise 1
	var i int32 = 20
	var f float32 = float32(i)
	fmt.Println(i)
	fmt.Println(f)

	// Exercise 2
	const value = 20
	var i2 int32 = value
	var f2 float32 = value
	fmt.Println(i2)
	fmt.Println(f2)
}
