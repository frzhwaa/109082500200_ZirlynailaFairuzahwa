# <h1 align="center">Laporan Praktikum Modul 14 - SORTING </h1>
<p align="center">ZIRLYNAILA FAIRUZAHWA - 109082500200</p>

## Unguided 

### 1 Selection. Hercules, preman terkenal seantero ibukota, memiliki kerabat di banyak daerah. Tentunya Hercules sangat suka mengunjungi semua kerabatnya itu. 
### Diberikan masukan nomor rumah dari semua kerabatnya di suatu daerah, buatlah program  rumahkerabat yang akan menyusun nomor-nomor rumah kerabatnya secara terurut membesar menggunakan algoritma selection sort. 
### Masukan dimulai dengan sebuah integer n (0 < n < 1000), banyaknya daerah kerabat  Hercules tinggal. Isi n baris berikutnya selalu dimulai dengan sebuah integer m (0 < m < 1000000) yang menyatakan banyaknya rumah kerabat di daerah tersebut, diikuti dengan rangkaian bilangan bulat positif, nomor rumah para kerabat. 
### Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar di masing- masing daerah.

#### soal1selection.go

```go
package main
import "fmt"

const NMAX = 1000000

type arrInt [NMAX]int

func selectionSortAsc(T *arrInt, n int) {
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
	var n, m int
	var rumah arrInt
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&m)
		for j := 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}
		selectionSortAsc(&rumah, m)
		for j := 0; j < m; j++ {
			fmt.Print(rumah[j])
			if j < m-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1](https://github.com/frzhwaa/109082500200_ZirlynailaFairuzahwa/blob/main/modul14/output/soal1selection.png)
#### ...

## Unguided

### 2 Selection. Belakangan diketahui ternyata Hercules itu tidak berani menyeberang jalan, maka selalu diusahakan agar hanya menyeberang jalan sesedikit mungkin, hanya diujung jalan. Karena nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah program kerabat dekat yang akan menampilkan nomor rumah mulai dari nomor yang ganjil lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor genap terurut mengecil.
### Format Masukan masih persis sama seperti sebelumnya.
### Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar untuk nomor ganjil, diikuti dengan terurut mengecil untuk nomor genap, di masing-masing daerah.

#### soal2selection.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 2](https://github.com/frzhwaa/109082500200_ZirlynailaFairuzahwa/blob/main/modul14/output/soal2selection.png)
#### ...

## Unguided 

### 1 Insertion. Buatlah sebuah program yang digunakan untuk membaca data integer seperti contoh yang  diberikan di bawah ini, kemudian diurutkan (menggunakan metoda insertion sort), dan memeriksa apakah data yang terurut berjarak sama terhadap data sebelumnya. 
### Masukan terdiri dari sekumpulan bilangan bulat yang diakhiri oleh bilangan negatif. Hanya  bilangan non negatif saja yang disimpan ke dalam array. 
### Keluaran terdiri dari dua baris. Baris pertama adalah isi dari array setelah dilakukan pengurutan, sedangkan baris kedua adalah status jarak setiap bilangan yang ada di dalam  array. "Data berjarak x" atau "data berjarak tidak tetap".

#### soal1insertion.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 3](https://github.com/frzhwaa/109082500200_ZirlynailaFairuzahwa/blob/main/modul14/output/soal1insertion.png)
#### ....

## Unguided 

### 2 Insertion. Sebuah program perpustakaan digunakan untuk mengelola data buku di dalam suatu  perpustakaan. Misalnya terdefinisi struct dan array seperti berikut ini:
### const nMax : integer = 7919
### type Buku = <
### id, judul, penulis, penerbit : string
### eksemplar, tahun, rating : integer >
### type DaftarBuku = array [ 1..nMax] of Buku
### Pustaka : DaftarBuku
### nPustaka: integer
### Masukan terdiri dari beberapa baris. Baris pertama adalah bilangan bulat N yang menyatakan banyaknya data buku yang ada di dalam perpustakaan. N baris berikutnya,  masing-masingnya adalah data buku sesuai dengan atribut atau field pada struct. Baris terakhir adalah bilangan bulat yang menyatakan rating buku yang akan dicari.
### Keluaran terdiri dari beberapa baris. Baris pertama adalah data buku terfavorit, baris kedua  adalah lima judul buku dengan rating tertinggi, selanjutnya baris terakhir adalah data buku yang dicari sesuai rating yang diberikan pada masukan baris terakhir. 
### Lengkapi subprogram-subprogram dibawah ini, sesuai dengan I.S. dan F.S yang diberikan.
### procedure DaftarkanBuku(in/out pustaka : DaftarBuku, n : integer)
### {I.S. sejumlah n data buku telah siap para piranti masukan
### F.S. n berisi sebuah nilai, dan pustaka berisi sejumlah n data buku}
### procedure CetakTerfavorit(in pustaka : DaftarBuku, in n : integer)
### {I.S. array pustaka berisi n buah data buku dan belum terurut
### F.S. Tampilan data buku (judul, penulis, penerbit, tahun)
### terfavorit, yaitu memiliki rating tertinggi}
### procedure UrutBuku( in/out pustaka : DaftarBuku, n : integer )
### {I.S. Array pustaka berisi n data buku
### F.S. Array pustaka terurut menurun/mengecil terhadap rating.
### Catatan: Gunakan metoda Insertion sort}
### procedure Cetak5Terbaru( in pustaka : DaftarBuku, n integer )
### {I.S. pustaka berisi n data buku yang sudah terurut menurut rating
### F.S. Laporan 5 judul buku dengan rating tertinggi
### Catatan: Isi pustaka mungkin saja kurang dari 5}
### procedure CariBuku(in pustaka : DaftarBuku, n : integer, r : integer )
### {I.S. pustaka berisi n data buku yang sudah terurut menurut rating
### F.S. Laporan salah satu buku (judul, penulis, penerbit, tahun,
### eksemplar, rating) dengan rating yang diberikan. Jika tidak ada buku
### dengan rating yang ditanyakan, cukup tuliskan “Tidak ada buku dengan
### rating seperti itu”. Catatan: Gunakan pencarian biner/belah dua.}

#### soal2insertion.go

```go
package main
import "fmt"

const nMax = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
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
	fmt.Println("Buku Terfavorit:")
	fmt.Println("Judul    :", pustaka[idxMax].judul)
	fmt.Println("Penulis  :", pustaka[idxMax].penulis)
	fmt.Println("Penerbit :", pustaka[idxMax].penerbit)
	fmt.Println("Tahun    :", pustaka[idxMax].tahun)
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
	fmt.Println("5 Buku dengan Rating Tertinggi:")
	for i := 0; i < batas; i++ {
		fmt.Println(i+1, ".", pustaka[i].judul)
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
		fmt.Println("Data Buku Ditemukan:")
		fmt.Println("Judul     :", pustaka[tengah].judul)
		fmt.Println("Penulis   :", pustaka[tengah].penulis)
		fmt.Println("Penerbit  :", pustaka[tengah].penerbit)
		fmt.Println("Tahun     :", pustaka[tengah].tahun)
		fmt.Println("Eksemplar :", pustaka[tengah].eksemplar)
		fmt.Println("Rating    :", pustaka[tengah].rating)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var pustaka DaftarBuku
	var n, ratingCari int
	DaftarkanBuku(&pustaka, &n)
	fmt.Scan(&ratingCari)
	CetakTerfavorit(pustaka, n)
	UrutBuku(&pustaka, n)
	Cetak5Terbaru(pustaka, n)
	CariBuku(pustaka, n, ratingCari)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 4](https://github.com/frzhwaa/109082500200_ZirlynailaFairuzahwa/blob/main/modul14/output/soal2insertion.png)
#### ...