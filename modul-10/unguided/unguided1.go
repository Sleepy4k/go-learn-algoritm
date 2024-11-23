package main

import "fmt"

type data struct {
	berat float64
}

func findLowestAndHighest(lowest, highest *float64, n_filled int, kelinci [1000]data) {
	x := 0

	*lowest = kelinci[0].berat
	*highest = kelinci[0].berat

	for x < n_filled {
		if *lowest > kelinci[x].berat {
			*lowest = kelinci[x].berat
		}

		if *highest < kelinci[x].berat {
			*highest = kelinci[x].berat
		}

		x++
	}
}

func main() {
	var kelinci [1000]data
	x := 0

	fmt.Print("masukan jumlah anak kelinci: ")
	fmt.Scan(&x)

	if x >= 1000 {
		x = 1000
	}

	for i := 0; i < x; i++ {
		y := 0.0

		fmt.Printf("masukan berat anak kelinci ke-%d: ", i+1)
		fmt.Scan(&y)

		kelinci[i] = data{y}
	}

	var lowest, highest float64

	findLowestAndHighest(&lowest, &highest, x, kelinci)

	fmt.Printf("Nilai berat kelinci paling kecil yaitu %.2f\n", lowest)
	fmt.Printf("Nilai berat kelinci paling besar yaitu %.2f\n", highest)
}