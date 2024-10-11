package main

import (
    "fmt"
    "math"
)

func jarak(a, b, c, d float64) float64 {
    return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}

func diDalam(cx, cy, r, x, y float64) bool {
    return jarak(cx, cy, x, y) <= r
}

func main() {
	var (
		cx1, cy1, r1 float64
		cx2, cy2, r2 float64
		x, y float64
	)

    fmt.Println("Masukkan koordinat titik pusat dan radius lingkaran 1:")
    fmt.Scan(&cx1, &cy1, &r1)
    fmt.Println("Masukkan koordinat titik pusat dan radius lingkaran 2:")
    fmt.Scan(&cx2, &cy2, &r2)
    fmt.Println("Masukkan koordinat titik sembarang:")
    fmt.Scan(&x, &y)

    diDalam1 := diDalam(cx1, cy1, r1, x, y)
    diDalam2 := diDalam(cx2, cy2, r2, x, y)

    if diDalam1 && diDalam2 {
        fmt.Println("Titik di dalam lingkaran 1 dan 2")
    } else if diDalam1 {
        fmt.Println("Titik di dalam lingkaran 1")
    } else if diDalam2 {
        fmt.Println("Titik di dalam lingkaran 2")
    } else {
        fmt.Println("Titik di luar lingkaran 1 dan 2")
    }
}
