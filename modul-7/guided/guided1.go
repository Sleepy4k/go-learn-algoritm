//Guided 2 - Slice
package main

import "fmt"

func sudahAda(daftarTeman []string, nama string) bool {
	for _, teman := range daftarTeman {
		if teman == nama {
			return true
		}
	}
	return false
}

func main() {
	daftarTeman := []string{"Andi", "Budi", "Cici"}

	namaBaru := []string{"Dewi", "Budi", "Eka"}

	for _, nama := range namaBaru {
		if !sudahAda(daftarTeman, nama) {
			daftarTeman = append(daftarTeman, nama)
		} else {
			fmt.Println("Nama", nama, "sudah ada dalam daftar.")
		}
	}

	fmt.Println("Daftar Teman:", daftarTeman)
}