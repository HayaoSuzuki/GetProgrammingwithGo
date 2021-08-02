package main

import (
	"fmt"
	"math/rand"
)

const distance = 62100000 // km

func getSpaceline() string {
	var line = ""
	if num := rand.Intn(3); num == 0 {
		line = "Space Adventures"
	} else if num == 1 {
		line = "SpaceX"
	} else {
		line = "Virgin Galactic"
	}
	return line
}

func getVelocity() int {
	var velocity = rand.Intn(15) + 16
	return velocity
}

func getFee(velocity int, tripType string) int {
	var fee = velocity + 20
	if tripType == "Round-trip" {
		return fee * 2
	} else {
		return fee
	}
}

func getDaysOfOneWay(velocity int) int {
	return distance / (velocity * 60 * 60 * 24)
}

func getTripType() string {
	var tripType = ""
	if num := rand.Intn(2); num == 0 {
		tripType = "Round-trip"
	} else {
		tripType = "One-way"
	}
	return tripType
}

func main() {
	fmt.Println("Spaceline         Days Trip type Price")
	fmt.Println("=======================================")
	for count := 0; count < 10; count++ {
		var spaceline = getSpaceline()
		var velocity = getVelocity()
		var days = getDaysOfOneWay(velocity)
		var tripType = getTripType()
		var price = getFee(velocity, tripType)
		fmt.Printf("%-16v %4v %-10v $%4v\n", spaceline, days, tripType, price)

	}
}
