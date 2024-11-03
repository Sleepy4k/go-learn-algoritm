package main

import "fmt"

func bilangan(n int) {
    if n == 1 {
        fmt.Printf("%d ", n)
        return
    }

    fmt.Printf("%d ", n)
    bilangan(n - 1)
    fmt.Printf("%d ", n)
}

func main() {
    var n int

    fmt.Print("masukan bilangan: ")
    fmt.Scan(&n)

    bilangan(n)
}
