package main
import (
	"fmt"
	"math"
)

func main() {
	var cx1, cx2, cy1, cy2, r1, r2, x, y float64
	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)
	dalam1 := didalam(cx1, cy1, r1, x, y)
	dalam2 := didalam(cx2, cy2, r2, x, y)
	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

func jarak(a, b, c, d float64) float64 {
	var hasil float64
	hasil = math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
	return hasil
}

func didalam(cx, cy, r, x, y float64) bool {
	var hasil bool
	hasil = jarak(x, y, cx, cy) <= r
	return hasil
}