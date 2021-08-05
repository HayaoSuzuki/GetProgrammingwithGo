package main

import "fmt"

func main() {
	age := 41
	marsAge := float64(age)
	marsDays := 687.0
	earthDays := 365.2425
	marsAge = marsAge * earthDays / marsDays
	fmt.Println("私は火星では、", marsAge, "歳です。")
	fmt.Println("私は火星では、", int(marsAge), "歳です。")
}
