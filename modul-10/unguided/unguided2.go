package main

import "fmt"

func main() {
	var x, y, wadah int

	fmt.Print("Masukan jumlah ikan dan jumlah ikan per wadah: ")
	fmt.Scan(&x, &y)

	var ikan, totalBerat, rerata [1000]float64

	for i := 0; i < x; i++ {
		fmt.Printf("masukan berat ikan ke-%d: ", i+1)
		fmt.Scan(&ikan[i])
	}

	if x % y == 0 {
		wadah = x / y
	} else {
		wadah = (x / y) + 1
	}

	for i := 0; i < x; i++ {
		totalBerat[i/y] += ikan[i]
	}

	for i := 0; i < wadah; i++ {
		rerata[i] = totalBerat[i] / float64(y)
	}

	for i := 0; i < wadah; i++ {
		fmt.Printf("wadah %d: berat %.2f kg, rata-rata %.2f\n", i+1, totalBerat[i], rerata[i])
	}
}