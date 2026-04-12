package main
import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    ganjil(1, n)
}

func ganjil(i, n int) {
    if i > n {
        return
    }
    if i % 2 != 0 {
        fmt.Print(i, " ")
    }
    ganjil(i+1, n)
}