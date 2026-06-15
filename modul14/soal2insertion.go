package main
import "fmt"

const nMax = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(
			&pustaka[i].id,
			&pustaka[i].judul,
			&pustaka[i].penulis,
			&pustaka[i].penerbit,
			&pustaka[i].eksemplar,
			&pustaka[i].tahun,
			&pustaka[i].rating,
		)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	idxMax := 0
	for i := 1; i < n; i++ {
		if pustaka[i].rating > pustaka[idxMax].rating {
			idxMax = i
		}
	}
	fmt.Println(
		pustaka[idxMax].judul,
		pustaka[idxMax].penulis,
		pustaka[idxMax].penerbit,
		pustaka[idxMax].tahun,
	)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	var temp Buku
	i := 1
	for i <= n-1 {
		j := i
		temp = pustaka[j]
		for j > 0 && temp.rating > pustaka[j-1].rating {
			pustaka[j] = pustaka[j-1]
			j = j - 1
		}
		pustaka[j] = temp
		i = i + 1
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	var batas int
	if n < 5 {
		batas = n
	} else {
		batas = 5
	}
	for i := 0; i < batas; i++ {
		fmt.Println(pustaka[i].judul)
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	var kiri, kanan, tengah int
	var ketemu bool
	kiri = 0
	kanan = n - 1
	ketemu = false
	for kiri <= kanan && !ketemu {
		tengah = (kiri + kanan) / 2
		if pustaka[tengah].rating == r {
			ketemu = true
		} else if r > pustaka[tengah].rating {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	if ketemu {
		fmt.Println(
			pustaka[tengah].judul,
			pustaka[tengah].penulis,
			pustaka[tengah].penerbit,
			pustaka[tengah].tahun,
			pustaka[tengah].eksemplar,
			pustaka[tengah].rating,
		)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}