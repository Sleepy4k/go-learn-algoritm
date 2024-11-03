package main

import "fmt"

func ganjil(n int) {
    if n == 1 {
        fmt.Print("1 ")
        return
    }

    ganjil(n-1)

    if n % 2 == 1 {
        fmt.Printf("%d ", n)
    }
}

func main() {
    var n int

    fmt.Print("masukan bilangan: ")
    fmt.Scan(&n)

    ganjil(n)
}
