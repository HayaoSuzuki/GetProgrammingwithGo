package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64

func measureTemperature(samples int, sensor func() kelvin) {
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%vK\n", k)
		time.Sleep(time.Second)
	}
}

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

	measureTemperature(3, fakeSensor)
}
