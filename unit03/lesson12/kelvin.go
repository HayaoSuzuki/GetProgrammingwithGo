package main

import "fmt"

func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func celsiusToFahrenheit(c float64) float64 {
	return (c * 9.0 / 5.0) + 32
}

func kelvinToFahrenheit(k float64) float64 {
	return celsiusToFahrenheit(kelvinToCelsius(k))
}

func main() {
	kelvin := 0.0
	celsius := kelvinToCelsius(kelvin)
	fahrenheit := celsiusToFahrenheit(celsius)
	fmt.Println(kelvin, "Kは、", celsius, "Cです。")
	fmt.Println(celsius, "Cは、", fahrenheit, "Fです。")
	fmt.Println(kelvin, "Kは、", kelvinToFahrenheit(kelvin), "Fです。")

}
