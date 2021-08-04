package main

import "fmt"

func main() {
	year := 2018
	fmt.Printf("%T 型: %v\n", year, year)

	days := 365.2425
	fmt.Printf("%T 型: %[1]v\n", days)

	var red, green, blue uint8 = 0x00, 0x8d, 0xd5
	fmt.Printf("color: #%02x%02x%02x", red, green, blue)
}
