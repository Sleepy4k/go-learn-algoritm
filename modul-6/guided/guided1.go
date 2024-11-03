package main

import "fmt"

func main() {
    cetak(5)
}

func cetak(x int) {
    fmt.Println(x)
    cetak(x + 1)
}
