package main

import "fmt"

func main() {
    var left_bag_weight_081, right_bag_weight_081 float64

    for {
        fmt.Print("Masukan berat belanjaan di kedua kantong (kg): ")
        fmt.Scan(&left_bag_weight_081, &right_bag_weight_081)

        if left_bag_weight_081 < 0 || right_bag_weight_081 < 0 {
            fmt.Println("Proses selesai")
            break
        }

        total_weight_081 := left_bag_weight_081 + right_bag_weight_081
        if total_weight_081 > 150 {
            fmt.Println("Proses selesai.")
            break
        }

        weight_gap_081 := left_bag_weight_081 - right_bag_weight_081
        if weight_gap_081 < 0 {
            weight_gap_081 = -weight_gap_081
        }

        if weight_gap_081 >= 9 {
            fmt.Println("Sepeda motor Pak Andi akan oleng: true")
        } else {
            fmt.Println("Sepeda motor Pak Andi akan oleng: false") 
        }
    }
}
