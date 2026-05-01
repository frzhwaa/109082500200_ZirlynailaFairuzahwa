package main
import (
	"fmt"
	"math"
)

const max = 100
type bilangan struct {
	info [max]int
	n int
}

func tampilSemua(A bilangan) {
	for i := 0; i < A.n; i++ {
		fmt.Print(A.info[i], " ")
	}
	fmt.Println()
}

func tampilIndeksGanjil(A bilangan) {
	for i := 1; i < A.n; i += 2 {
		fmt.Print(A.info[i], " ")
	}
	fmt.Println()
}

func tampilIndeksGenap(A bilangan) {
	for i := 0; i < A.n; i += 2 {
		fmt.Print(A.info[i], " ")
	}
	fmt.Println()
}

func tampilKelipatanX(A bilangan, x int) {
	for i := 0; i < A.n; i++ {
		if i%x == 0 {
			fmt.Print(A.info[i], " ")
		}
	}
	fmt.Println()
}

func hapusIndex(A *bilangan, idx int) {
	for i := idx; i < A.n-1; i++ {
		A.info[i] = A.info[i+1]
	}
	A.n--
}

func rataRata(A bilangan) float64 {
	total := 0
	for i := 0; i < A.n; i++ {
		total += A.info[i]
	}
	return float64(total) / float64(A.n)
}

func standarDeviasi(A bilangan) float64 {
	avg := rataRata(A)
	var jumlah float64
	for i := 0; i < A.n; i++ {
		jumlah += math.Pow(float64(A.info[i])-avg, 2)
	}
	return math.Sqrt(jumlah / float64(A.n))
}

func frekuensi(A bilangan, x int) int {
	hitung := 0
	for i := 0; i < A.n; i++ {
		if A.info[i] == x {
			hitung++
		}
	}
	return hitung
}

func main() {
	var A bilangan
	var x, idx, cari int
	fmt.Scan(&A.n)
	for i := 0; i < A.n; i++ {
		fmt.Scan(&A.info[i])
	}
	fmt.Println("Seluruh isi array:")
	tampilSemua(A)
	fmt.Println("Indeks ganjil:")
	tampilIndeksGanjil(A)
	fmt.Println("Indeks genap:")
	tampilIndeksGenap(A)
	fmt.Println("Masukkan nilai x untuk indeks kelipatan:")
	fmt.Scan(&x)
	tampilKelipatanX(A, x)
	fmt.Println("Masukkan indeks yang akan dihapus:")
	fmt.Scan(&idx)
	hapusIndex(&A, idx)
	fmt.Println("Array setelah dihapus:")
	tampilSemua(A)
	fmt.Printf("Rata-rata: %.2f\n", rataRata(A))
	fmt.Printf("Standar deviasi: %.2f\n", standarDeviasi(A))
	fmt.Println("Masukkan bilangan yang ingin dicari frekuensinya:")
	fmt.Scan(&cari)
	fmt.Println("Frekuensi:", frekuensi(A, cari))
}