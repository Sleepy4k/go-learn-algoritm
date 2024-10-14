package main

import "fmt"

func cetakNFibo(n int) {
	var f1, f2, f3 int
	f2 = 0
	f3 = 1

	for i := 1; i <= n; i++ {
		fmt.Println(f3)
		f1 = f2
		f2 = f3
		f3 = f1 + f2
	}
}

func main() {
	cetakNFibo(10)
}
