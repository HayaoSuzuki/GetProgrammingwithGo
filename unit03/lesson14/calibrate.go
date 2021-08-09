package main

import "fmt"

type kelvin float64
type sensor func() kelvin

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func realSensor() kelvin {
	return 0 // TODO: 本物のセンサを実装する日はくるのか
}

func main() {

	sensor := calibrate(realSensor, 5)
	fmt.Println(sensor())
}
