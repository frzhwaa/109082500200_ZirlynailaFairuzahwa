package main
import "fmt"

func factorial(n int, hasil *int){
	*hasil = 1
	var i int
	for i = 1; i <= n; i++ {
		*hasil = *hasil * i
	}
}

func permutation(n, r int, hasil *int) {
	var fn, fnr int
	factorial(n, &fn)
	factorial(n-r, &fnr)
	*hasil = fn / fnr
}

func combination(n, r int, hasil *int) {
	var fn, fr, fnr int
	factorial(n, &fn)
	factorial(r, &fr)
	factorial(n-r, &fnr)
	*hasil = fn / (fr * fnr)
}

func main() {
	var a, b, c, d, hasilP, hasilC int
	fmt.Scan(&a, &b, &c, &d)
	if a >= c && b >= d {
		permutation(a, c, &hasilP)
		combination(a, c, &hasilC)
		fmt.Println(hasilP, hasilC)
		permutation(b, d, &hasilP)
		combination(b, d, &hasilC)
		fmt.Println(hasilP, hasilC)
	}
}