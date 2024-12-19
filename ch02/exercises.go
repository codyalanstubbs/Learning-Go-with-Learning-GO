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

	// Exercise 3
	var b byte = 127
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	b = b + 1
	smallI = smallI + 1
	bigI = bigI + 1

	fmt.Println("b")
	fmt.Println(b)
	fmt.Println("smallI")
	fmt.Println(smallI)
	fmt.Println("bigI")
	fmt.Println(bigI)
}
