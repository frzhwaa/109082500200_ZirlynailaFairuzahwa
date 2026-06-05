package main
import "fmt"

const MAX = 20
type arrSuara [MAX + 1]int

func ReadSuara(suara *arrSuara, masuk *int, sah *int) {
	var x int
	fmt.Scan(&x)
	for x != 0 {
		*masuk++
		if x >= 1 && x <= MAX {
			suara[x]++
			*sah++
		}
		fmt.Scan(&x)
	}
}

func cariPemenang(suara arrSuara, ketua, wakil *int) {
	var maks1, maks2 int
	for i := 1; i <= MAX; i++ {
		if suara[i] > maks1 {
			maks2 = maks1
			*wakil = *ketua
			maks1 = suara[i]
			*ketua = i
		} else if suara[i] > maks2 {
			maks2 = suara[i]
			*wakil = i
		}
	}
}

func main() {
	var suara arrSuara
	var masuk, sah, ketua, wakil int
	ReadSuara(&suara, &masuk, &sah)
	cariPemenang(suara, &ketua, &wakil)
	fmt.Println("Suara masuk:", masuk)
	fmt.Println("Suara sah:", sah)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}