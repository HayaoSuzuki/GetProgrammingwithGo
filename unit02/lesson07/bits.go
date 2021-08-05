package main

import "fmt"

func main() {
	var red uint8 = 255
	red++
	fmt.Println(red)

	var number int8 = 127
	number++
	fmt.Println(number)

	var green uint8 = 3
	fmt.Printf("%08b\n", green)
	green++
	fmt.Printf("%08b\n", green)

	var blue uint8 = 255
	fmt.Printf("%08b\n", blue)
	blue++
	fmt.Printf("%08b\n", blue)

}
