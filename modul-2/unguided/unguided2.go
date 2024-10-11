package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader_081 := bufio.NewReader(os.Stdin)

	fmt.Print("N: ")
	var total_n_081 int

	for {
		user_input_081, err := reader_081.ReadString('\n')
		if err != nil {
			fmt.Printf("Error membaca input: %v\n", err)
			return
		}

		total_n_081, err =
			strconv.Atoi(strings.TrimSpace(user_input_081))
		if err != nil || total_n_081 <= 0 {
			fmt.Println("Harap masukkan bilangan bulat positif.")
		} else {
			break
		}
	}

	var total_pita_081 string
	var total_bunga_081 int

	for x_081 := 1; x_081 <= total_n_081; x_081++ {
		fmt.Printf("Bunga %d: ", x_081)

		user_input_081, err := reader_081.ReadString('\n')

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		user_input_081 = strings.TrimSpace(user_input_081)

		if strings.ToUpper(user_input_081) == "SELESAI" {
			break
		}

		if total_pita_081 == "" {
			total_pita_081 = user_input_081
		} else {
			total_pita_081 = total_pita_081 + " – " + user_input_081
		}

		total_bunga_081++
	}

	fmt.Println("Pita:", total_pita_081)
	fmt.Printf("Bunga: %d\n", total_bunga_081)
}
