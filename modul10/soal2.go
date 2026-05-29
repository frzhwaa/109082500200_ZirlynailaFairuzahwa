package main
import "fmt"

const max = 1000
type arrFloat [max]float64

func main() {
	var ikan arrFloat
	var x, y, i, j int
	var total, rata float64
	fmt.Scan(&x, &y)
	for i = 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}
	i = 0
	for i < x {
		total = 0
		j = 0
		for j < y && i < x {
			total = total + ikan[i]
			i = i + 1
			j = j + 1
		}
		fmt.Printf("%.2f ", total)
		rata = rata + total
	}
	fmt.Println()
	fmt.Printf("%.2f\n", rata/float64((x+y-1)/y))
}