package main
import "fmt"

const MAX = 1000000

type arrint [MAX]int

func selectionSort(T *arrint, n int) {
	var t, j, idx_min int
	i := 1
	for i <= n-1 {
		idx_min = i - 1
		j = i
		for j < n {
			if T[idx_min] > T[j] {
				idx_min = j
			}
			j = j + 1
		}
		t = T[idx_min]
		T[idx_min] = T[i-1]
		T[i-1] = t
		i = i + 1
	}
}

func main() {
	var n, m, x, nGanjil, nGenap int
	var ganjil, genap arrint
	var pertama bool
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		nGanjil = 0
		nGenap = 0
		fmt.Scan(&m)
		for j := 0; j < m; j++ {
			fmt.Scan(&x)
			if x%2 == 1 {
				ganjil[nGanjil] = x
				nGanjil++
			} else {
				genap[nGenap] = x
				nGenap++
			}
		}
		selectionSort(&ganjil, nGanjil)
		selectionSort(&genap, nGenap)
		pertama = true
		for j := 0; j < nGanjil; j++ {
			if !pertama {
				fmt.Print(" ")
			}
			fmt.Print(ganjil[j])
			pertama = false
		}
		for j := nGenap - 1; j >= 0; j-- {
			if !pertama {
				fmt.Print(" ")
			}
			fmt.Print(genap[j])
			pertama = false
		}
		fmt.Println()
	}
}