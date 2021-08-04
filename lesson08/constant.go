package main

import "fmt"

func main() {
	const distance = 24000000000000000000
	const lightSpeed = 299792
	const secondPerDay = 86400
	const days = distance / lightSpeed / secondPerDay
	fmt.Println("アンドロメダ銀河まで光速で、", days, "日の距離。")
}
