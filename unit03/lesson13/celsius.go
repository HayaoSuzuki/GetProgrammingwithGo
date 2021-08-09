package main

import "fmt"

func main() {
	type celsius float64

	var temperature celsius = 20
	const degree = 20
	temperature += degree
	//var warmUp float64 = 10
	//temperature += warmUp
	fmt.Println(temperature)
}
