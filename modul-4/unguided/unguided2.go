package main

import "fmt"

func hitungSkor(soal, skor *int) {
	var waktu int

	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)

		if waktu <= 300 {
			*soal++
			*skor += waktu
		}
	}
}

func main() {
	var nama, pemenang string
	var soal, skor int

	soalTertinggi := 0
	skorTerendah := 8 * 301

	for {
		soal = 0
		skor = 0

		fmt.Print("Masukan data nama dan waktu atau 'selesai' untuk berhenti: ")
		fmt.Scan(&nama)

		if nama == "selesai" || nama == "Selesai" || nama == "SELESAI" {
			break
		}

		hitungSkor(&soal, &skor)

		if soal > soalTertinggi || (soal == soalTertinggi && skor < skorTerendah) {
			soalTertinggi = soal
			skorTerendah = skor
			pemenang = nama
		}
	}

	if soalTertinggi > 0 {
		fmt.Printf("%s %d %d\n", pemenang, soalTertinggi, skorTerendah)
	} else {
		fmt.Println("Tidak ada peserta yang menyelesaikan soal.")
	}
}
