import "fmt"

func main() {
	hargaBuah := map[string]int {
		"Apel":  5000,
		"Pisang": 3000,
		"Mangga": 7000,
	}

	fmt.Println("Harga Buah:")

	for buah, harga := range hargaBuah {
		fmt.Printf("%s: Rp%d\n", buah, harga)
	}

	fmt.Print("Harga buah Mangga = ", hargaBuah["Mangga"])
}