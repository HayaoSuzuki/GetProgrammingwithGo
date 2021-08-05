package main

import "fmt"

func main() {
	const lightSpeed = 299792 // km/s
	const secondsPerDay = 86400

	var distance int64 = 41.3e12 // km
	fmt.Println("アルファ・ケンタウリまでの距離は、", distance, "km.")

	var days = distance / lightSpeed / secondsPerDay
	fmt.Println("光の速度で、", days, "日かかる。")

	// 以下、普通の整数型では扱えない
	//distance = 24e18 // km
	//fmt.Println("アンドロメダ銀河までの距離は、", distance, "km.")
	//days = distance / lightSpeed / secondsPerDay
	//fmt.Println("光の速度で、", days, "日かかる。")
}
