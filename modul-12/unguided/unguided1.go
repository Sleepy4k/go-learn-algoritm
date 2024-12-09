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
	var x []int
	var y int

	for y >= 0 {
		fmt.Scan(&y)

		if y >= 0 {
			x = append(x, y)
		}
	}

	insertionSort(x)

	fmt.Println(x)

	if len(x) < 2 {
		fmt.Println("Data hanya berjumlah 1")
		return
	}

	var curr, gap int

	curr = x[1]
	gap = x[1] - x[0]

	isConstant := true

	for i := 2; i < len(x) - 2; i++ {
		if x[i] - curr != gap {
			isConstant = false
			break
		}

		curr = x[i]
	}

	if isConstant {
		fmt.Printf("Data berjarak %d\n", gap)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}