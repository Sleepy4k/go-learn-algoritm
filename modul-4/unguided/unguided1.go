package main

import "fmt"

func factorial(n int, hasil *int) {
    *hasil = 1

    for i := 1; i <= n; i++ {
        *hasil *= i
    }
}

func permutasi(a, b int, hasil *int) {
    var x, y int

    factorial(a, &x)
    factorial(a-b, &y)

    *hasil = x / y
}

func kombinasi(a, b int, hasil *int) {
    var x, y, z int

    factorial(a, &x)
    factorial(b, &y)
    factorial(a-b, &z)

    *hasil = x / (y * z)
}

func main() {
    var a, b, c, d int

    fmt.Print("Masukan input = ")
    fmt.Scan(&a, &b, &c, &d)

    if a >= c && b >= d {
        var x, y int

        permutasi(a, c, &x)
        kombinasi(a, c, &y)
        fmt.Printf("%d %d\n", x, y)

        permutasi(b, d, &x)
        kombinasi(b, d, &y)
        fmt.Printf("%d %d\n", x, y)
    } else {
        fmt.Println("Syarat tidak terpenuh: a harus >= c dan b harus >= d")
    }
}
