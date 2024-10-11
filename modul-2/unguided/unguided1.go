package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fixed_order_081 := []string{"merah", "kuning", "hijau", "ungu"}

	reader_081 := bufio.NewReader(os.Stdin)
	is_success_081 := true

	for x_081 := 1; x_081 <= 5; x_081++ {
		fmt.Printf("Percobaan %d: ", x_081)
		
		user_input_081, _ := reader_081.ReadString('\n')
		user_input_081 = strings.TrimSpace(user_input_081)

        user_order_081 := strings.Split(user_input_081, " ")

        for y_081 := 0; y_081 < 4; y_081++ {
            if user_order_081[y_081] != fixed_order_081[y_081] {
                is_success_081 = false
                break
            }
        }

        if !is_success_081 {
            break
        }
    }

    if is_success_081 {
        fmt.Println("BERHASIL : true")
    } else {
        fmt.Println("BERHASIL : false")
    }
}