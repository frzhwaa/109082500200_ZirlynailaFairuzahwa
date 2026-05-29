# <h1 align="center">Laporan Praktikum Modul 10 - PENCARIAN NILAI MAX/MIN </h1>
<p align="center">ZIRLYNAILA FAIRUZAHWA - 109082500200</p>

## Unguided 

### 1. Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual.
### Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya N bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual.
### Keluaran terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan terbesar.

#### soal1.go

```go
package main
import "fmt"

const MAX = 1000
type arrBerat [MAX]float64

func main() {
	var berat arrBerat
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Printf("Berat kelinci ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}
	min := berat[0]
	max := berat[0]
	i := 1
	for i < n {
		if min > berat[i] {
			min = berat[i]
		}
		if max < berat[i] {
			max = berat[i]
		}
		i = i + 1
	}
	fmt.Println("Berat terkecil:", min)
	fmt.Println("Berat terbesar:", max)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1](https://github.com/frzhwaa/109082500200_ZirlynailaFairuzahwa/blob/main/modul10/output/soal1.png)
#### Program tersebut digunakan untuk menentukan berat kelinci paling kecil dan paling besar dari sejumlah data berat yang diinputkan pengguna. Program akan membaca banyaknya data kelinci terlebih dahulu, kemudian meminta pengguna memasukkan berat masing-masing kelinci satu per satu.
#### Program menggunakan tipe bentukan array bernama arrBerat yang bertipe float64 dengan kapasitas maksimum MAX = 1000. Array ini digunakan untuk menyimpan seluruh data berat kelinci yang diinputkan.
#### Pada fungsi main, program pertama-tama membaca nilai n yang menunjukkan jumlah kelinci yang akan dimasukkan datanya. Setelah itu, program menggunakan perulangan for untuk menginput berat setiap kelinci dan menyimpannya ke dalam array berat.
#### Setelah seluruh data dimasukkan, program menginisialisasi variabel min dan max dengan nilai berat kelinci pertama, yaitu berat[0]. Variabel min digunakan untuk menyimpan berat terkecil, sedangkan variabel max digunakan untuk menyimpan berat terbesar.
#### Selanjutnya, program menggunakan perulangan while dengan variabel i dimulai dari 1 hingga n-1. Pada setiap perulangan, program akan membandingkan nilai berat[i] dengan nilai min dan max.
#### Jika nilai berat[i] lebih kecil dari min, maka nilai min akan diperbarui dengan berat[i]. Sebaliknya, jika nilai berat[i] lebih besar dari max, maka nilai max akan diperbarui dengan berat[i].
#### Setelah seluruh data diperiksa, program akan menampilkan hasil berupa berat terkecil dan berat terbesar dari semua kelinci yang telah diinputkan.
#### Sebagai contoh, jika inputnya adalah:
#### 5
#### Berat kelinci ke-1: 2
#### Berat kelinci ke-2: 3
#### Berat kelinci ke-3: 4
#### Berat kelinci ke-4: 5
#### Berat kelinci ke-5: 6
#### Maka program akan memeriksa seluruh data berat kelinci. Nilai terkecil yang ditemukan adalah 2, sedangkan nilai terbesar adalah 6. Oleh karena itu, output yang dihasilkan adalah:
#### Berat terkecil: 2
#### Berat terbesar: 6

## Unguided 

### 2. Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual.
### Masukan terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya ikan yang akan dimasukan ke dalam wadah. Baris kedua terdiri dari sejumlah x bilangan riil yang menyatakan banyaknya ikan yang akan dijual.
### Keluaran terdiri dari dua baris. Baris pertama adalah kumpulan bilangan riil yang menyatakan total berat ikan di setiap wadah (jumlah wadah tergantung pada nilai x dan y, urutan ikan yang dimasukan ke dalam wadah sesuai urutan pada masukan baris ke-2). Baris kedua adalah sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah.

#### soal2.go

```go
package main
import "fmt"

const max = 1000
type arrFloat [max]float64

func main() {
	var ikan arrFloat
	var x, y, i, j int
	var total, rata float64
	fmt.Scan(&x, &y)
	for i = 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}
	i = 0
	for i < x {
		total = 0
		j = 0
		for j < y && i < x {
			total = total + ikan[i]
			i = i + 1
			j = j + 1
		}
		fmt.Printf("%.2f ", total)
		rata = rata + total
	}
	fmt.Println()
	fmt.Printf("%.2f\n", rata/float64((x+y-1)/y))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 2](https://github.com/frzhwaa/109082500200_ZirlynailaFairuzahwa/blob/main/modul10/output/soal2.png)
#### Program tersebut digunakan untuk mengelompokkan data berat ikan ke dalam beberapa kelompok, kemudian menghitung total berat setiap kelompok serta menghitung rata-rata total dari seluruh kelompok tersebut.
#### Program menggunakan tipe bentukan array bernama arrFloat yang bertipe float64 dengan kapasitas maksimum 1000. Array ini digunakan untuk menyimpan data berat ikan yang diinputkan pengguna.
#### Pada fungsi main, program terlebih dahulu membaca dua buah bilangan integer yaitu x dan y. Nilai x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya ikan yang akan dimasukan ke dalam wadah.
#### Setelah itu, program menggunakan perulangan for untuk membaca x buah data berat ikan dan menyimpannya ke dalam array ikan.
#### Selanjutnya, program menggunakan perulangan while dengan variabel i untuk memproses data ikan per kelompok. Pada setiap kelompok, variabel total direset menjadi 0 untuk menghitung jumlah berat ikan dalam kelompok tersebut.
#### Di dalamnya terdapat perulangan lain menggunakan variabel j. Perulangan ini akan berjalan selama jumlah data dalam kelompok belum mencapai y dan indeks i masih berada dalam batas jumlah data x.
#### Pada setiap perulangan, nilai ikan[i] akan ditambahkan ke variabel total. Setelah itu, nilai i dan j akan bertambah satu sehingga program berpindah ke data berikutnya.
#### Setelah satu wadah selesai dihitung, program akan menampilkan total berat wadah tersebut dengan format dua angka di belakang koma menggunakan fmt.Printf("%.2f ", total).
#### Nilai total setiap wadah juga akan ditambahkan ke variabel rata. Setelah semua wadah selesai diproses, program menghitung rata-rata total kelompok dengan rumus:
#### rata-rata = jumlah seluruh total kelompok / banyak kelompok
#### Banyak kelompok dihitung menggunakan rumus:
#### (x + y − 1) / y
#### Rumus tersebut digunakan agar pembagian menghasilkan jumlah kelompok yang benar meskipun kelompok terakhir tidak penuh.
#### Sebagai contoh, jika inputnya adalah:
#### 6 2
#### 1.5 2.0 3.5 4.0 5.0 6.0
#### Maka data ikan akan dikelompokkan menjadi:
#### Wadah 1 = 1.5 + 2.0 = 3.5
#### Wadah 2 = 3.5 + 4.0 = 7.5
#### Wadah 3 = 5.0 + 6.0 = 11.0
#### Program kemudian menampilkan total setiap wadah:
#### 3.50 7.50 11.00
#### Jumlah seluruh total wadah adalah:
#### 3.5 + 7.5 + 11.0 = 22.0
#### Banyak wadah ada 3, sehingga rata-ratanya:
#### 22.0 / 3 = 7.33
#### Oleh karena itu, output akhirnya adalah:
#### 3.50 7.50 11.00
#### 7.33

## Unguided 

### 3. Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data  berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.
### Buatlah program dengan spesifikasi subprogram sebagai berikut:
### type arrBalita [100]float64
### func hitungMinMax(arrBerat arrBalita; bMin, bMax *float64) {
### /* I.S. Terdefinisi array dinamis arrBerat
### Proses: Menghitung berat minimum dan maksimum dalam array
### F.S. Menampilkan berat minimum dan maksimum balita */
### ...
### }
### function rerata (arrBerat arrBalita) real {
### /* menghitung dan mengembalikan rerata berat balita dalam array */
### ...
### }
### Perhatikan sesi interaksi pada contoh berikut ini (teks bergaris bawah adalah input/read)
### Masukan banyak data berat balita : 4
### Masukan berat balita ke-1: 5.3
### Masukan berat balita ke-2: 6.2
### Masukan berat balita ke-3: 4.1
### Masukan berat balita ke-4: 9.9
### Berat balita minimum: 4.10 kg
### Berat balita maksimum: 9.90 kg
### Rerata berat balita: 6.38 kg

#### soal3.go

```go
package main
import "fmt"

const MAKS = 100
type arrBalita [MAKS]float64
var n int

func hitungMinMax(arrBerat arrBalita, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]
	i := 1
	for i < n {
		if *bMin > arrBerat[i] {
			*bMin = arrBerat[i]
		}
		if *bMax < arrBerat[i] {
			*bMax = arrBerat[i]
		}
		i = i + 1
	}
}

func rerata(arrBerat arrBalita) float64 {
	var total float64
	total = 0
	i := 0
	for i < n {
		total = total + arrBerat[i]
		i = i + 1
	}
	return total / float64(n)
}

func main() {
	var berat arrBalita
	var min, max, rata float64
	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}
	hitungMinMax(berat, &min, &max)
	rata = rerata(berat)
	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rata)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 3](https://github.com/frzhwaa/109082500200_ZirlynailaFairuzahwa/blob/main/modul10/output/soal3.png)
#### Program tersebut digunakan untuk mengolah data berat balita dengan cara mencari berat minimum, berat maksimum, dan menghitung rata-rata berat balita berdasarkan data yang diinputkan pengguna.
#### Program menggunakan tipe bentukan array bernama arrBalita yang bertipe float64 dengan kapasitas maksimum MAKS = 100. Array ini digunakan untuk menyimpan data berat balita.
#### Program juga menggunakan variabel global n bertipe integer yang berfungsi untuk menyimpan jumlah data balita yang akan diproses.
#### Fungsi hitungMinMax digunakan untuk menghitung berat balita paling kecil dan paling besar dalam array. Fungsi ini menerima parameter berupa array arrBerat serta dua parameter pointer bMin dan bMax yang digunakan untuk menyimpan hasil berat minimum dan maksimum.
#### Pada awal fungsi, nilai bMin dan bMax diisi dengan data pertama dalam array, yaitu arrBerat[0]. Setelah itu, program menggunakan perulangan untuk memeriksa seluruh data mulai dari indeks ke-1 sampai indeks ke n-1.
#### Jika ditemukan data yang lebih kecil dari nilai minimum saat ini, maka nilai minimum akan diperbarui. Sebaliknya, jika ditemukan data yang lebih besar dari nilai maksimum saat ini, maka nilai maksimum akan diperbarui.
#### Fungsi rerata digunakan untuk menghitung rata-rata berat balita. Fungsi ini menerima parameter berupa array arrBerat dan menghasilkan nilai bertipe float64.
#### Di dalam fungsi rerata, program menggunakan variabel total untuk menjumlahkan seluruh berat balita dalam array. Setelah seluruh data dijumlahkan, hasil total dibagi dengan jumlah data n sehingga diperoleh nilai rata-rata.
#### Pada fungsi main, program pertama-tama meminta pengguna memasukkan jumlah data berat balita yang akan diproses. Setelah itu, program menggunakan perulangan for untuk membaca berat setiap balita dan menyimpannya ke dalam array berat.
#### Setelah semua data dimasukkan, program memanggil fungsi hitungMinMax untuk mendapatkan berat minimum dan maksimum, kemudian memanggil fungsi rerata untuk menghitung rata-rata berat balita.
#### Hasil akhirnya ditampilkan menggunakan fmt.Printf dengan format dua angka di belakang koma.
#### Sebagai contoh, jika inputnya adalah:
#### 4
#### 5.3
#### 6.2
#### 4.1
#### 9.9
#### Maka program akan memeriksa seluruh data berat balita. Berat terkecil yang ditemukan adalah 4.10 kg, sedangkan berat terbesar adalah 9.90 kg.
#### Selanjutnya, rata-rata dihitung dengan rumus:
#### (5.3 + 6.2 + 4.1 + 9.9) / 4 = 6.375
#### Karena ditampilkan dengan dua angka di belakang koma, maka hasil rata-ratanya menjadi 6.38 kg.
#### Oleh karena itu, output yang dihasilkan adalah:
#### Berat balita minimum: 4.10 kg
#### Berat balita maksimum: 9.90 kg
#### Rerata berat balita: 6.38 kg