package main
import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    pola(n)
}

func pola(n int) {
    if n == 0 {
        return
    }
    fmt.Print(n, " ")
    pola(n - 1)
    if n != 1 {
        fmt.Print(n, " ")
    }
}