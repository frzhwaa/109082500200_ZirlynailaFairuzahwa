package main
import "fmt"

const maks = 100
type DaftarKlub struct {
	nama [maks]string
	n int
}

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var hasil DaftarKlub
	var pertandingan int = 1
	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)
	for {
		fmt.Printf("Pertandingan %d : ", pertandingan)
		fmt.Scan(&skorA, &skorB)
		if skorA < 0 || skorB < 0 {
			break
		}
		if skorA > skorB {
			fmt.Println("Hasil", pertandingan, ":", klubA)
			hasil.nama[hasil.n] = klubA
			hasil.n++
		} else if skorB > skorA {
			fmt.Println("Hasil", pertandingan, ":", klubB)
			hasil.nama[hasil.n] = klubB
			hasil.n++
		} else {
			fmt.Println("Hasil", pertandingan, ": Draw")
		}
		pertandingan++
	}
	fmt.Println("Pertandingan selesai")
}