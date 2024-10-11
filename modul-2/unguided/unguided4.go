package main 

import "fmt"

func hitungFungsiFK(K_081 int) float64 { 
    return float64((4*K_081+2)*(4*K_081+2)) / float64((4*K_081+1)*(4*K_081+3)) 
} 

func hitungAkar2Hampiran(K_081 int) float64 { 
    akar2_081 := 1.0 
    for x_081 := 0; x_081 <= K_081; x_081++ { 
        akar2_081 *= float64((4*x_081+2)*(4*x_081+2)) / float64((4*x_081+1)*(4*x_081+3)) 
    }

    return akar2_081 
} 

func main() { 
    var choice_081 int 
    var K_081 int 

    fmt.Println("Pilih opsi:") 
    fmt.Println("1. Hitung f(K)") 
    fmt.Println("2. Hitung hampiran akar 2") 
    fmt.Print("Masukkan choice_081 (1/2): ") 
    fmt.Scan(&choice_081) 
    fmt.Print("Masukkan nilai K: ") 
    fmt.Scan(&K_081)

    if choice_081 == 1 { 
        fK_081 := hitungFungsiFK(K_081) 
        fmt.Printf("Nilai f(K) = %.10f\n", fK_081) 
    } else if choice_081 == 2 { 
        akar2_081 := hitungAkar2Hampiran(K_081) 
        fmt.Printf("Nilai hampiran akar 2 = %.10f\n", akar2_081) 
    } else { 
        fmt.Println("choice_081 tidak valid!") 
    } 
} 
