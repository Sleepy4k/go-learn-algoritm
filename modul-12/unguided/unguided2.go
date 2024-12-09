package main

import "fmt"

const nMax int = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating int
}

type DaftarBuku [nMax]Buku

var Pustaka DaftarBuku = DaftarBuku{}
var nPustaka int

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
    for i := 0; i < n; i++ {
        var buku Buku
        fmt.Print("\nMasukkan ID Buku: ")
        fmt.Scan(&buku.id)
        fmt.Print("Masukkan Judul Buku: ")
        fmt.Scan(&buku.judul)
        fmt.Print("Masukkan Penulis Buku: ")
        fmt.Scan(&buku.penulis)
        fmt.Print("Masukkan Penerbit Buku: ")
        fmt.Scan(&buku.penerbit)
        fmt.Print("Masukkan Tahun Terbit: ")
        fmt.Scan(&buku.tahun)
        fmt.Print("Masukkan Eksemplar Buku: ")
        fmt.Scan(&buku.eksemplar)
        fmt.Print("Masukkan Rating Buku: ")
        fmt.Scan(&buku.rating)

        pustaka[nPustaka] = buku
        nPustaka++
    }

    fmt.Println("\n")
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
    if n == 0 {
        fmt.Println("Tidak ada buku yang terdaftar")
        return
    }

    terfavorit := pustaka[0]
    for i := 1; i < n; i++ {
        if pustaka[i].rating > terfavorit.rating {
            terfavorit = pustaka[i]
        }
    }

    fmt.Printf("Buku dengan rating tertinggi:\n")
    fmt.Printf("Judul: %s\nPenulis: %s\nPenerbit: %s\nTahun: %d\nEksemplar: %d\nRating: %d\n", terfavorit.judul, terfavorit.penulis, terfavorit.penerbit, terfavorit.tahun, terfavorit.eksemplar, terfavorit.rating)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
    if n == 0 {
        fmt.Println("Tidak ada buku yang terdaftar")
        return
    }

    for i := 1; i < n; i++ {
        temp := pustaka[i]
        j := i - 1

        for j >= 0 && pustaka[j].rating < temp.rating {
            pustaka[j + 1] = pustaka[j]
            j--
        }

        pustaka[j + 1] = temp
    }
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
    if n == 0 {
        fmt.Println("Tidak ada buku yang terdaftar")
        return
    }

    fmt.Println("5 Buku dengan terbaru:")

	for i, buku := range pustaka {
        if i == 5 || i == n {
            break
        }
        fmt.Printf("Judul: %s\nPenulis: %s\nPenerbit: %s\nTahun: %d\nEksemplar: %d\nRating: %d\n\n", buku.judul, buku.penulis, buku.penerbit, buku.tahun, buku.eksemplar, buku.rating)
    }
}

func CariBuku(pustaka DaftarBuku, n, r int) {
    for i := 1; i < n; i++ {
        temp := pustaka[i]
        j := i - 1

        for j >= 0 && pustaka[j].rating < temp.rating {
            pustaka[j + 1] = pustaka[j]
            j--
        }

        pustaka[j + 1] = temp
    }

    fmt.Printf("Buku dengan rating %d:\n", r)
    for _, buku := range pustaka {
        if buku.rating == r {
            fmt.Printf("Judul: %s\nPenulis: %s\nPenerbit: %s\nTahun: %d\nEksemplar: %d\nRating: %d\n\n", buku.judul, buku.penulis, buku.penerbit, buku.tahun, buku.eksemplar, buku.rating)
        }
    }
}

func main() {
    var n int
    fmt.Print("Masukkan jumlah buku yang ingin didaftarkan: ")
    fmt.Scan(&n)

    if n <= 0 {
        fmt.Println("Pastikan jumlah buku yang didaftarkan lebih dari 0")
        return
    }

    DaftarkanBuku(&Pustaka, n)

    CetakTerfavorit(Pustaka, nPustaka)

    UrutBuku(&Pustaka, nPustaka)

    Cetak5Terbaru(Pustaka, nPustaka)

    var r int
    fmt.Print("Masukkan rating buku yang ingin dicari: ")
    fmt.Scan(&r)

    if r <= 0 {
        fmt.Println("Pastikan rating buku yang dicari lebih dari 0")
        return
    }

    CariBuku(Pustaka, nPustaka, r)
}