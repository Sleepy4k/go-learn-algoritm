package main

import "fmt"

func cetakDeret(n int) {
    for n > 1 {
        if n%2 == 0 {
            n /= 2
        } else {
            n = 3*n + 1
        }

        fmt.Printf("%d ", n)
    }
}

func main() {
    var n int

    fmt.Print("Masukan data n: ")
    fmt.Scan(&n)

    fmt.Printf("%d ", n)
    cetakDeret(n)
}
