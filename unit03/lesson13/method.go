package main

import "fmt"

type celsius float64
type kelvin float64
type fahrenheit float64

func (c celsius) kelvin() kelvin {
	return kelvin(c - 273.15)
}

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32)
}

func (k kelvin) celsius() celsius {
	return celsius(k + 273.15)
}

func (k kelvin) fahrenheit() fahrenheit {
	return fahrenheit(k*9.0/5.0 - 459.67)
}

func (f fahrenheit) celsius() celsius {
	return celsius(5.0 * (f - 32.0) / 9.0)
}

func (f fahrenheit) kelvin() kelvin {
	return kelvin(5.0 * (f + 459.67) / 9.0)
}

func main() {
	var k kelvin = 294.0
	c := k.celsius()
	fmt.Println(k, "Kは", c, "Cです。")

	var moonC celsius = 127.0
	moonK := moonC.kelvin()
	fmt.Println("月の表面温度", moonC, "Cは、", moonK, "Kです。")

}
