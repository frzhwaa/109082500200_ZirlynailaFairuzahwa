package main
import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    baris(1, n)
}

func baris(i, n int) {
    if i > n {
        return
    }
    bintang(i)
    fmt.Println()
    baris(i+1, n)
}

func bintang(j int) {
    if j == 0 {
        return
    }
    bintang(j-1)
    fmt.Print("*")
}