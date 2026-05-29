package main
import "fmt"

const MAX = 1000
type arrBerat [MAX]float64

func main() {
	var berat arrBerat
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Printf("Berat kelinci ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}
	min := berat[0]
	max := berat[0]
	i := 1
	for i < n {
		if min > berat[i] {
			min = berat[i]
		}
		if max < berat[i] {
			max = berat[i]
		}
		i = i + 1
	}
	fmt.Println("Berat terkecil:", min)
	fmt.Println("Berat terbesar:", max)
}