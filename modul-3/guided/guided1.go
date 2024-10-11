package main

import "fmt"

func factorial(n int) int {
	var hasil int = 1

	for i := 1; i <= n; i++ {
		hasil = hasil * i
	}

	return hasil
}

func permutasi(a, b int) int {
	return factorial(a) / factorial(a-b)
}

func kombinasi(a, b int) int {
	return factorial(a) / (factorial(b) * factorial(a-b))
}

func main() {
	var a, b, c, d int

	fmt.Print("Masukan input = ")
	fmt.Scan(&a, &b, &c, &d)

	if a >= c && b >= d {
		fmt.Printf("%d, %d\n", permutasi(a, c), kombinasi(a, c))
		fmt.Printf("%d, %d\n", permutasi(b, d), kombinasi(b, d))
	} else {
		fmt.Println("Syarat tidak terpenuh: a harus >= c dan b harus >= d")
	}
}
