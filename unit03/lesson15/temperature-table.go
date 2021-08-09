package main

import "fmt"

type celsius float64

type getRowFunc func(row int) (string, string)

const valueFormat = "%.1f"
const rowFormat = "%8s %8s\n"

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32)
}

type fahrenheit float64

func (f fahrenheit) celsius() celsius {
	return celsius(5.0 * (f - 32.0) / 9.0)
}

func celsiusToFahrenheit(row int) (string, string) {
	c := celsius(row*5 - 40)
	f := c.fahrenheit()
	cell1 := fmt.Sprintf(valueFormat, c)
	cell2 := fmt.Sprintf(valueFormat, f)

	return cell1, cell2
}

func fahrenheitToCelsius(row int) (string, string) {
	f := fahrenheit(row*5 - 40)
	c := f.celsius()
	cell1 := fmt.Sprintf(valueFormat, f)
	cell2 := fmt.Sprintf(valueFormat, c)

	return cell1, cell2
}

func drawTable(rows int, getRow getRowFunc) {
	for row := 0; row < rows; row++ {
		c1, c2 := getRow(row)
		fmt.Printf(rowFormat, c1, c2)

	}
	fmt.Println()
}

func main() {
	drawTable(29, celsiusToFahrenheit)
	fmt.Println()
	drawTable(29, fahrenheitToCelsius)
}
