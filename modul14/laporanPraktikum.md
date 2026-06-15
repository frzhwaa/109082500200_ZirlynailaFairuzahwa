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
#### Program tersebut digunakan untuk menyusun nomor rumah kerabat Hercules di setiap daerah secara terurut membesar (ascending) menggunakan algoritma Selection Sort. Program akan membaca jumlah daerah, kemudian untuk setiap daerah membaca sejumlah nomor rumah dan mengurutkannya dari yang terkecil ke terbesar sebelum ditampilkan.
#### Program menggunakan tipe bentukan array bernama arrInt yang bertipe integer dengan kapasitas maksimum NMAX = 1000000. Array ini digunakan untuk menyimpan nomor rumah kerabat pada setiap daerah.
#### Pada fungsi selectionSortAsc, program melakukan pengurutan menggunakan algoritma Selection Sort. Pada setiap iterasi, program mencari nilai terkecil pada bagian array yang belum terurut, kemudian menukarnya dengan elemen paling kiri pada bagian tersebut. Proses ini diulangi hingga seluruh data terurut membesar.
#### Pada fungsi main, program membaca jumlah daerah (**n**). Untuk setiap daerah, program membaca banyak rumah (m) dan nomor-nomor rumah kerabat. Setelah semua nomor rumah pada daerah tersebut dibaca, program memanggil fungsi selectionSortAsc untuk mengurutkan data, kemudian menampilkan hasil pengurutan.
#### Misalkan input yang diberikan adalah:
#### 3
#### 5 2 1 7 9 13
#### 6 189 15 27 39 75 133
#### 3 4 9 1
#### Prosesnya yaitu:
#### Daerah 1:
#### Banyak rumah = 5
#### Nomor rumah = 2 1 7 9 13
#### Setelah diurutkan menjadi:
#### 1 2 7 9 13
#### Daerah 2:
#### Banyak rumah = 6
#### Nomor rumah = 189 15 27 39 75 133
#### Setelah diurutkan menjadi:
#### 15 27 39 75 133 189
#### Daerah 3:
#### Banyak rumah = 3
#### Nomor rumah = 4 9 1
#### Setelah diurutkan menjadi:
#### 1 4 9
#### Jadi, outputnya adalah:
#### 1 2 7 9 13
#### 15 27 39 75 133 189
#### 1 4 9

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
#### Program tersebut digunakan untuk menampilkan nomor rumah kerabat Hercules dengan aturan khusus, yaitu semua nomor rumah ganjil ditampilkan lebih dahulu dalam urutan membesar (ascending), kemudian diikuti nomor rumah genap dalam urutan mengecil (descending). Program memisahkan nomor rumah ganjil dan genap ke dalam dua array yang berbeda, kemudian masing-masing diurutkan menggunakan algoritma Selection Sort.
#### Program menggunakan tipe bentukan array bernama arrint yang bertipe integer dengan kapasitas maksimum MAX = 1000000. Array ganjil digunakan untuk menyimpan nomor rumah ganjil, sedangkan array genap digunakan untuk menyimpan nomor rumah genap.
#### Pada fungsi selectionSort, program melakukan pengurutan menggunakan algoritma Selection Sort secara membesar (ascending). Pada setiap iterasi, program mencari nilai terkecil pada bagian array yang belum terurut, kemudian menukarnya dengan elemen paling kiri pada bagian tersebut hingga seluruh data terurut.
#### Pada fungsi main, program membaca jumlah daerah (n). Untuk setiap daerah, program membaca banyak rumah (m) dan nomor-nomor rumah kerabat. Jika nomor rumah bernilai ganjil maka disimpan ke array ganjil, sedangkan jika bernilai genap disimpan ke array genap. Setelah itu kedua array diurutkan menggunakan selectionSort. Saat menampilkan hasil, array ganjil dicetak dari awal hingga akhir sehingga tetap terurut membesar, sedangkan array genap dicetak dari belakang ke depan sehingga menjadi terurut mengecil.
#### Misalkan input yang diberikan adalah:
#### 3
#### 5 2 1 7 9 13
#### 6 189 15 27 39 75 133
#### 3 4 9 1
#### Prosesnya yaitu:
#### Daerah 1:
#### Banyak rumah = 5
#### Nomor rumah = 2 1 7 9 13
#### Bilangan ganjil:
#### 1 7 9 13
#### Setelah diurutkan:
#### 1 7 9 13
#### Bilangan genap:
#### 2
#### Setelah diurutkan dan dicetak menurun:
#### 2
#### Hasil daerah 1:
#### 1 7 9 13 2
#### Daerah 2:
#### Banyak rumah = 6
#### Nomor rumah = 189 15 27 39 75 133
#### Bilangan ganjil:
#### 189 15 27 39 75 133
#### Setelah diurutkan:
#### 15 27 39 75 133 189
#### Bilangan genap:
#### Tidak ada
#### Hasil daerah 2:
#### 15 27 39 75 133 189
#### Daerah 3:
#### Banyak rumah = 3
#### Nomor rumah = 4 9 1
#### Bilangan ganjil:
#### 9 1
#### Setelah diurutkan:
#### 1 9
#### Bilangan genap:
#### 4
#### Setelah diurutkan dan dicetak menurun:
#### 4
#### Hasil daerah 3:
#### 1 9 4
#### Jadi, outputnya adalah:
#### 1 7 9 13 2
#### 15 27 39 75 133 189
#### 1 9 4
#### Pada daerah pertama, Hercules dapat mengunjungi rumah bernomor ganjil terlebih dahulu tanpa harus menyeberang jalan, kemudian mengunjungi rumah bernomor genap di sisi seberang jalan. Pola yang sama diterapkan pada setiap daerah.

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
#### Program tersebut digunakan untuk membaca sekumpulan bilangan bulat non-negatif, kemudian mengurutkannya menggunakan algoritma Insertion Sort secara membesar (ascending). Setelah data terurut, program akan memeriksa apakah selisih (jarak) antara setiap bilangan yang berurutan selalu sama. Jika sama, program menampilkan "Data berjarak x", sedangkan jika berbeda maka menampilkan "Data berjarak tidak tetap".
#### Program menggunakan tipe bentukan array bernama arr yang bertipe integer dengan kapasitas maksimum 1000. Array ini digunakan untuk menyimpan seluruh bilangan non-negatif yang diinputkan pengguna.
#### Pada fungsi insertionSortAsc, program melakukan pengurutan menggunakan algoritma Insertion Sort secara membesar (ascending). Setiap elemen yang belum terurut disisipkan ke posisi yang sesuai pada bagian array yang sudah terurut sehingga menghasilkan urutan dari nilai terkecil ke terbesar.
#### Pada fungsi main, program membaca bilangan satu per satu. Selama bilangan yang dibaca bernilai non-negatif, bilangan tersebut disimpan ke dalam array. Jika ditemukan bilangan negatif, proses input berhenti. Setelah itu array diurutkan menggunakan Insertion Sort, kemudian program memeriksa apakah selisih antar elemen yang berurutan selalu sama atau tidak.
#### Misalkan input yang diberikan adalah:
#### 31 13 25 43 1 7 19 37 -5
#### Prosesnya yaitu:
#### 31: disimpan ke array
#### 13: disimpan ke array
#### 25: disimpan ke array
#### 43: disimpan ke array
#### 1: disimpan ke array
#### 7: disimpan ke array
#### 19: disimpan ke array
#### 37: disimpan ke array
#### -5: input berhenti karena bilangan negatif
#### Isi array sebelum diurutkan:
#### 31 13 25 43 1 7 19 37
#### Setelah dilakukan Insertion Sort (ascending):
#### 1 7 13 19 25 31 37 43
#### Pemeriksaan jarak:
#### 7 - 1 = 6
#### 13 - 7 = 6
#### 19 - 13 = 6
#### 25 - 19 = 6
#### 31 - 25 = 6
#### 37 - 31 = 6
#### 43 - 37 = 6
#### Semua selisih bernilai 6, sehingga jaraknya tetap.
#### Jadi, outputnya adalah:
#### 1 7 13 19 25 31 37 43
#### Data berjarak 6
#### Program menyimpulkan bahwa data memiliki jarak tetap sebesar 6 karena selisih antara setiap dua bilangan yang berurutan selalu sama setelah data diurutkan.

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
		fmt.Println(i+1, pustaka[i].judul)
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
#### Program tersebut digunakan untuk mengelola data buku pada sebuah perpustakaan. Program dapat membaca data buku, menentukan buku terfavorit berdasarkan rating tertinggi, mengurutkan seluruh buku berdasarkan rating menggunakan algoritma Insertion Sort, menampilkan lima buku dengan rating tertinggi, dan mencari buku berdasarkan rating tertentu menggunakan Binary Search.
#### Program menggunakan tipe bentukan Buku yang memiliki field:
#### id, judul, penulis, penerbit : string
#### eksemplar, tahun, rating : integer
#### Program juga menggunakan tipe bentukan array DaftarBuku yang dapat menyimpan hingga 7919 data buku. Array ini digunakan untuk menyimpan seluruh data buku yang ada di perpustakaan.
#### Pada prosedur DaftarkanBuku, program membaca jumlah buku (n) dan seluruh data buku dari masukan, kemudian menyimpannya ke dalam array pustaka.
#### Pada prosedur CetakTerfavorit, program mencari buku dengan rating tertinggi menggunakan pencarian sekuensial. Setelah ditemukan, program menampilkan informasi buku tersebut berupa judul, penulis, penerbit, dan tahun terbit.
#### Pada prosedur UrutBuku, program mengurutkan seluruh data buku berdasarkan rating secara menurun (descending) menggunakan algoritma Insertion Sort. Buku dengan rating terbesar akan berada di indeks paling depan.
#### Pada prosedur Cetak5Terbaru, program menampilkan lima judul buku dengan rating tertinggi. Jika jumlah buku kurang dari lima, maka semua judul buku akan ditampilkan.
#### Pada prosedur cariBuku, program melakukan pencarian rating menggunakan algoritma Binary Search pada array yang sudah terurut. Jika ditemukan, program menampilkan seluruh informasi buku tersebut. Jika tidak ditemukan, program menampilkan pesan bahwa buku dengan rating tersebut tidak ada.
#### Misalkan input yang diberikan adalah:
#### 6
#### B01 Algoritma Keonho Informatika 10 2020 85
#### B02 StrukturData Sean Erlangga 5 2021 92
#### B03 BasisData Ruto Gramedia 8 2019 78
#### B04 PemrogramanGo Junghwan Informatika 7 2023 95
#### B05 SistemOperasi Jeongwoo Erlangga 6 2022 88
#### B06 JaringanKomputer Haechan Gramedia 4 2021 90
#### 90
#### Prosesnya yaitu:
#### Buku 1:
#### Judul = Algoritma
#### Rating = 85
#### Buku 2:
#### Judul = StrukturData
#### Rating = 92

#### Buku 3:
#### Judul = BasisData
#### Rating = 78
#### Buku 4:
#### Judul = PemrogramanGo
#### Rating = 95
#### Buku 5:
#### Judul = SistemOperasi
#### Rating = 8
#### Buku 6:
#### Judul = JaringanKomputer
#### Rating = 90
#### Menentukan buku terfavorit:
#### Rating tertinggi = 95
#### Buku terfavorit = PemrogramanGo
#### Mengurutkan data dengan Insertion Sort (descending rating):
#### Sebelum diurutkan:
#### Algoritma = 85
#### StrukturData = 92
#### BasisData = 78
#### PemrogramanGo = 95
#### SistemOperasi = 88
#### JaringanKomputer = 90
#### Setelah diurutkan:
#### PemrogramanGo = 95
#### StrukturData = 92
#### JaringanKomputer = 90
#### SistemOperasi = 88
#### Algoritma = 85
#### BasisData = 78
#### Menampilkan 5 buku dengan rating tertinggi:
#### 1. PemrogramanGo
#### 2. StrukturData
#### 3. JaringanKomputer
#### 4. SistemOperasi
#### 5. Algoritma
#### Mencari buku dengan rating 90 menggunakan Binary Search:
#### Data ditemukan pada buku:
#### JaringanKomputer
#### Jadi, outputnya adalah:
#### Buku Terfavorit:
#### Judul : PemrogramanGo
#### Penulis : Junghwan
#### Penerbit : Informatika
#### Tahun : 2023
#### 5 Buku dengan Rating Tertinggi:
#### 1 PemrogramanGo
#### 2 StrukturData
#### 3 JaringanKomputer
#### 4 SistemOperasi
#### 5 Algoritma
#### Data Buku Ditemukan:
#### Judul : JaringanKomputer
#### Penulis : Haechan
#### Penerbit : Gramedia
#### Tahun : 2021
#### Eksemplar : 4
#### Rating : 90
#### Program menyimpulkan bahwa buku PemrogramanGo merupakan buku terfavorit karena memiliki rating tertinggi yaitu 95, sedangkan pencarian rating 90 menghasilkan buku JaringanKomputer.