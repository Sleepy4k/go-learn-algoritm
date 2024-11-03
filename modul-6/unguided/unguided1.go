package main

import "fmt"

func fibo(n int) int {
    if n <= 1 {
        return n
    } else {
        return fibo(n-1) + fibo(n-2)
    }
}

func main() {
    var n int

    fmt.Print("masukan panjang fibonacci: ")
    fmt.Scan(&n)

    fmt.Print("|n")

    for i := 0; i <= n; i++ {
        fmt.Printf("|%d|", i)
    }

    fmt.Print("\n|Sn")

    for i := 0; i <= n; i++ {
        fmt.Printf("|%d|", fibo(i))
    }
}
