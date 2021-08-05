package main

import "fmt"

func main() {
	const lightSpeed = 299792 // km/s
	var distance = 56000000   // km
	fmt.Println(distance/lightSpeed, "秒")

	distance = 401000000
	fmt.Println(distance/lightSpeed, "秒")

	const hoursPerDay = 24
	const spaceX = 100800 // km/h
	distance = 96300000   // km
	fmt.Println(distance/spaceX, "時間")
	fmt.Println(distance/spaceX/hoursPerDay, "日")

}
