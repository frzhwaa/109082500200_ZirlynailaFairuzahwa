package main
import "fmt"

const JUMLAH_CALON = 20
type ArrSuara [JUMLAH_CALON + 1]int

func bacaSuara(suara *ArrSuara, masuk *int, sah *int) {
	var x int
	fmt.Scan(&x)
	for x != 0 {
		*masuk++
		if x >= 1 && x <= JUMLAH_CALON {
			suara[x]++
			*sah++
		}
		fmt.Scan(&x)
	}
}

func tampilkanHasil(suara ArrSuara) {
	for i := 1; i <= JUMLAH_CALON; i++ {
		if suara[i] > 0 {
			fmt.Printf("%d: %d\n", i, suara[i])
		}
	}
}

func main() {
	var suara ArrSuara
	var totalMasuk, totalSah int
	bacaSuara(&suara, &totalMasuk, &totalSah)
	fmt.Println("Suara masuk:", totalMasuk)
	fmt.Println("Suara sah:", totalSah)
	tampilkanHasil(suara)
}