package main

import "fmt"

func f1(x, y int) float64 {
	var hasil float64
	hasil = float64(2*x) - 0.5 * float64(y) + 3.0
	return hasil
}

// x dan y adalah pass by value
// hasil adala pass bu reference
func f2(x, y int, hasil *float64) {
	*hasil = float64(2*x) - 0.5 * float64(y) + 3.0
}

func output(c, d float64) {
	fmt.Printf("Hasil dari f2: %.2f\n", c)
	fmt.Printf("Hasil dari f1: %.2f\n", d)
}

func main() {
	var a, b int
	var c float64

	fmt.Scan(&a, &b)

	f2(a, b, &c)
	output(c, f1(b, a))
}