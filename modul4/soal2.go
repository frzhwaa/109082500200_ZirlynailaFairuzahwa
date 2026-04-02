package main
import "fmt"

func hitungSkor(soal, skor *int) {
	var waktu int
	*soal = 0
	*skor = 0
	for i := 1; i <= 8; i++ {
		fmt.Scan(&waktu)
		if waktu <= 300 {
			*soal += 1
			*skor += waktu
		}
	}
}

func main() {
	var nama, pemenang string
	var soal, skor, soalTerbanyak, minSkor int
	fmt.Scan(&nama)
	minSkor = 999999
	for nama != "Selesai" {
		hitungSkor(&soal, &skor)
		if soal > soalTerbanyak || (soal == soalTerbanyak && skor < minSkor) {
			soalTerbanyak = soal
			minSkor = skor
			pemenang = nama
		}
		fmt.Scan(&nama)
	}
	fmt.Println(pemenang, soalTerbanyak, minSkor)
}