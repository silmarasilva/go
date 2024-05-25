package main

import "fmt"

func main() {

	var numero int = 10000
	fmt.Println(numero)

	var numero1 int16 = 100
	fmt.Println(numero1)

	var numero2 int64 = 100000000000000
	fmt.Println(numero2)

	var numero3 uint32 = 10000
	fmt.Println(numero3)

	// alias
	// INT32 = RUNE
	var numero4 rune = 123456
	fmt.Println(numero4)

	//alias
	// INT8 = BYTE
	var numero5 byte = 123
	fmt.Println(numero5)
}