package main

import "fmt"

func main() {
	var nums [5]int
	var chars [3]rune

	fmt.Println("Masukkan 5 angka integer (32-127):")

	for i := 0; i < 5; i++ {
		fmt.Scanf("%d", &nums[i])
	}

	fmt.Println("Masukkan 3 karakter:")
	for i := 0; i < 3; i++ {
		fmt.Scanf("%c", &chars[i])
		if chars[i] == '\n' {
			i--
		}
	}

	fmt.Println("Keluaran:")

	for i := 0; i < 5; i++ {
		fmt.Printf("%c", nums[i])
	}

	fmt.Println()

	for i := 0; i < 3; i++ {
		fmt.Printf("%c", chars[i])
	}

	fmt.Println()
}