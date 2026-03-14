package main
import "fmt"

func main() {
	var berat int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&berat)
	kg := berat / 1000
	gr := berat % 1000
	biayaKg := kg * 10000
	biayaGr := 0
	if gr >= 500 {
		biayaGr = gr * 5
	} else {
		biayaGr = gr * 15
	}
	total := 0
	if kg >= 10 {
		total = biayaKg
	} else {
		total = biayaKg + biayaGr
	}
	fmt.Println("Detail berat: ", kg, " kg + ", gr, " gr")
	fmt.Println("Detail biaya: Rp. ", biayaKg, " + Rp. ", biayaGr)
	fmt.Println("Total biaya: Rp. ", total)
}