package main

import "fmt"

func findLowestAndHighest(lowest, highest *float64, n_filled int, kelinci [1000]float64) {
	x := 0

	*lowest = kelinci[0]
	*highest = kelinci[0]

	for x < n_filled {
		if *lowest > kelinci[x] {
			*lowest = kelinci[x]
		}

		if *highest < kelinci[x] {
			*highest = kelinci[x]
		}

		x++
	}
}

func main() {
	x := 0

	fmt.Print("masukan jumlah anak kelinci: ")
	fmt.Scan(&x)

	if x >= 1000 {
		x = 1000
	}

	var kelinci [1000]float64

	for i := 0; i < x; i++ {
		fmt.Printf("masukan berat anak kelinci ke-%d: ", i+1)
		fmt.Scan(&kelinci[i])
	}

	var lowest, highest float64

	findLowestAndHighest(&lowest, &highest, x, kelinci)

	fmt.Printf("Nilai berat kelinci paling kecil yaitu %.2f\n", lowest)
	fmt.Printf("Nilai berat kelinci paling besar yaitu %.2f\n", highest)
}