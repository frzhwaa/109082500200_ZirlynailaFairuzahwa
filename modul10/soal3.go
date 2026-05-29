package main
import "fmt"

const MAKS = 100
type arrBalita [MAKS]float64
var n int

func hitungMinMax(arrBerat arrBalita, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]
	i := 1
	for i < n {
		if *bMin > arrBerat[i] {
			*bMin = arrBerat[i]
		}
		if *bMax < arrBerat[i] {
			*bMax = arrBerat[i]
		}
		i = i + 1
	}
}

func rerata(arrBerat arrBalita) float64 {
	var total float64
	total = 0
	i := 0
	for i < n {
		total = total + arrBerat[i]
		i = i + 1
	}
	return total / float64(n)
}

func main() {
	var berat arrBalita
	var min, max, rata float64
	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}
	hitungMinMax(berat, &min, &max)
	rata = rerata(berat)
	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rata)
}