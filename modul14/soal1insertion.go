package main
import "fmt"

const max = 1000

type arr [max]int

func insertionSortAsc(T *arr, n int) {
	var temp, j int
	i := 1
	for i <= n-1 {
		j = i
		temp = T[j]
		for j > 0 && temp < T[j-1] {
			T[j] = T[j-1]
			j = j - 1
		}
		T[j] = temp
		i = i + 1
	}
}

func main() {
	var A arr
	var x, jarak int
	var tetap bool
	n := 0
	fmt.Scan(&x)
	for x >= 0 {
		A[n] = x
		n++
		fmt.Scan(&x)
	}
	insertionSortAsc(&A, n)
	for i := 0; i < n; i++ {
		fmt.Print(A[i])
		if i < n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
	if n <= 1 {
		fmt.Println("Data berjarak 0")
	} else {
		jarak = A[1] - A[0]
		tetap = true
		i := 2
		for i < n && tetap {
			if A[i]-A[i-1] != jarak {
				tetap = false
			}
			i++
		}
		if tetap {
			fmt.Println("Data berjarak", jarak)
		} else {
			fmt.Println("Data berjarak tidak tetap")
		}
	}
}