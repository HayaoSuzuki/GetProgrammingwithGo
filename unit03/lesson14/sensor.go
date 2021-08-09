package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func realSensor() kelvin {
	return 0 // TODO: 本物のセンサを実装する日はくるのか
}

func main() {
	sensor := fakeSensor()
	fmt.Println(sensor)

	sensor = realSensor()
	fmt.Println(sensor)
}
