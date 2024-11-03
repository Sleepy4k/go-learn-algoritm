package main

import "fmt"

func factorial(n, i int) {
    if i <= n {
        if n % i == 0 {
            fmt.Printf("%d ", i)
        }

        factorial(n, i+1)
    }
}

func main() {
    var n int

    fmt.Print("masukan bilangan: ")
    fmt.Scan(&n)

    factorial(n, 1)
}
