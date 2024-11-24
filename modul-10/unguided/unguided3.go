package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for _, berat := range arrBerat {
		if berat < *bMin && berat != 0 {
			*bMin = berat
		}

		if berat > *bMax {
			*bMax = berat
		}
	}
}

func rerata(arrBerat arrBalita) float64 {
	var total float64
	var count int

	for _, berat := range arrBerat {
		if berat != 0 {
			total += berat
			count++
		}
	}

	return total / float64(count)
}

func main() {
	n := 0

	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(&n)

	var balita arrBalita

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&balita[i])
	}

	var bMin, bMax float64

	hitungMinMax(balita, &bMin, &bMax)

	fmt.Printf("Berat balita minimum: %.2f\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f\n", bMax)

	fmt.Printf("Rerata berat balita: %.2f\n", rerata(balita))
}