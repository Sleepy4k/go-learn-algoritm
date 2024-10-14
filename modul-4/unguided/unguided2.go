package main

import "fmt"    

func hitungSkor(soal *int, skor int) {

}

func main() {
    var nama, pemenang string
    var a, b, c, d, e, f, g, h int
    maxSoal := 0
    minWaktu := 99999

    for {
        fmt.Print("Masukkan nama peserta atau 'selesai' untuk berhenti: ")
        fmt.Scan(&nama)

        if nama == "selesai" || nama == "Selesai" || nama == "SELESAI" {
            break
        }

        hitungSkor(nama, &a, &b)
        fmt.Printf("%s menyelesaikan %d soal dalam %d menit\n", nama, a, b)
        
        if a > maxSoal || (a == maxSoal && b < minWaktu) {
            maxSoal = a
            minWaktu = b
            pemenang = nama
        }
    }

    if maxSoal > 0 {
        fmt.Printf("Pemenangnya adalah %s dengan %d soal diselesaikan dalam %d menit\n", pemenang, maxSoal, minWaktu)
    } else {
        fmt.Println("Tidak ada peserta yang menyelesaikan soal.")
    }
}