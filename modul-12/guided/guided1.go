package main

import "fmt"

func insertionSort(arr []int) {
	n := len(arr)

	for i := 1; i < n; i++ {
		temp := arr[i]
		j := i

		for j > 0 && arr[j-1] > temp {
			arr[j] = arr[j-1]
			j--
		}

		arr[j] = temp
	}
}

func main() {
	data := []int{5, 2, 9, 1, 4, 6}
	fmt.Println("Data sebelum diurutkan:", data)
	insertionSort(data)
	fmt.Println("Data setelah diurutkan (descending):", data)
}