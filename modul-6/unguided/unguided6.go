package main

import "fmt"

func pangkat(x, y int) int {
    if y == 0 {
        return 1
    } else {
        return x * pangkat(x, y-1)
    }
}

func main() {
    var x, y int

    fmt.Print("masukan bilangan x dan y: ")
    fmt.Scan(&x, &y)

    fmt.Println(pangkat(x, y))
}
