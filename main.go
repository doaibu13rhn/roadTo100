package main

import (
	"fmt"
	"strconv"
)

func main() {
	// multiple konstanta
	const (
		a = 183
		b = " adalah tinggi ku"
	)

	fmt.Println(a, b)

	// tipe data boolean
	isPass := false
	fmt.Printf("passing grade %t\n ", isPass)

	//tipe data integer
	rewardPoint := 79
	fmt.Println("your point is ", rewardPoint, "point")

	//tipe data float
	discount := 4.5
	fmt.Println("Congrats, you got ", discount, "discount")

	//konversi tipe data numerik
	var aNumber int32 = 100
	var numberInFloat float32 = float32(aNumber)
	fmt.Printf("Nilainya adalah %.2f\n", numberInFloat)

	//konversi tipe data string
	var numberInString string = strconv.Itoa(int(aNumber))
	fmt.Printf("Nilainya adalah %s\n", numberInString)

	//operator aritmatika
	total := 6
	fmt.Println(total == 8)
}
