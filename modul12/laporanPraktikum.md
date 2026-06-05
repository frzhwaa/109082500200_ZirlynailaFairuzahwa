# <h1 align="center">Laporan Praktikum Modul 12 - SEARCHING </h1>
<p align="center">ZIRLYNAILA FAIRUZAHWA - 109082500200</p>

## Unguided 

### 1. Pada pemilihan ketua RT yang baru saja berlangsung, terdapat 20 calon ketua yang bertanding  memperebutkan suara warga. Perhitungan suara dapat segera dilakukan karena warga cukup mengisi formulir dengan nomor dari calon ketua RT yang dipilihnya. Seperti biasa, selalu ada pengisian yang tidak tepat atau dengan nomor pilihan di luar yang tersedia, sehingga data juga harus divalidasi. Tugas Anda untuk membuat program mencari siapa yang memenangkan pemilihan ketua RT.
### Buatlah program pilkart yang akan membaca, memvalidasi, dan menghitung suara yang diberikan dalam pemilihan ketua RT tersebut. 
### Masukan hanya satu baris data saja, berisi bilangan bulat valid yang kadang tersisipi dengan  data tidak valid. Data valid adalah integer dengan nilai di antara 1 s.d. 20 (inklusif). Data berakhir jika ditemukan sebuah bilangan dengan nilai 0. 
### Keluaran dimulai dengan baris berisi jumlah data suara yang terbaca, diikuti baris yang berisi  berapa banyak suara yang valid. Kemudian sejumlah baris yang mencetak data para calon apa saja yang mendapatkan suara.

#### soal1.go

```go
package main
import "fmt"

const JUMLAH_CALON = 20
type ArrSuara [JUMLAH_CALON + 1]int

func bacaSuara(suara *ArrSuara, masuk *int, sah *int) {
	var x int
	fmt.Scan(&x)
	for x != 0 {
		*masuk++
		if x >= 1 && x <= JUMLAH_CALON {
			suara[x]++
			*sah++
		}
		fmt.Scan(&x)
	}
}

func tampilkanHasil(suara ArrSuara) {
	for i := 1; i <= JUMLAH_CALON; i++ {
		if suara[i] > 0 {
			fmt.Printf("%d: %d\n", i, suara[i])
		}
	}
}

func main() {
	var suara ArrSuara
	var totalMasuk, totalSah int
	bacaSuara(&suara, &totalMasuk, &totalSah)
	fmt.Println("Suara masuk:", totalMasuk)
	fmt.Println("Suara sah:", totalSah)
	tampilkanHasil(suara)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1](https://github.com/frzhwaa/109082500200_ZirlynailaFairuzahwa/blob/main/modul12/output/soal1.png)
#### Program tersebut digunakan untuk menghitung jumlah suara sah dan tidak sah dari sejumlah data suara yang diinputkan pengguna. Program akan membaca suara satu per satu, kemudian menghitung total suara masuk, total suara sah, dan jumlah suara untuk setiap calon.
#### Program menggunakan tipe bentukan array bernama ArrSuara yang bertipe int dengan kapasitas maksimum JUMLAH_CALON + 1. Array ini digunakan untuk menyimpan jumlah suara masing-masing calon (indeks 1 sampai 20).
#### Pada fungsi bacaSuara, program membaca input suara berupa angka. Jika angka yang dimasukkan adalah 0, maka proses input berhenti. Jika angka berada dalam rentang 1 sampai 20, maka suara dianggap sah dan ditambahkan ke calon yang sesuai. Setiap input (selain 0) akan menambah jumlah total suara masuk.
#### Pada fungsi tampilkanHasil, program menampilkan jumlah suara untuk setiap calon yang memperoleh suara (hanya calon dengan suara > 0).
#### Pada fungsi main, program menginisialisasi array suara dan variabel penghitung. Kemudian memanggil bacaSuara untuk membaca input, menampilkan total suara masuk, total suara sah, dan hasil perolehan suara tiap calon.
#### Misalkan input yang diberikan adalah:
#### 7 19 3 2 78 3 1 -3 18 19 0
#### Prosesnya yaitu:
#### 7: suara sah untuk calon 7  
#### 19: suara sah untuk calon 19  
#### 3: suara sah untuk calon 3  
#### 2: suara sah untuk calon 2  
#### 78: tidak sah (di luar 1–20)  
#### 3: suara sah untuk calon 3  
#### 1: suara sah untuk calon 1  
#### -3: tidak sah (di luar 1–20)  
#### 18: suara sah untuk calon 18  
#### 19: suara sah untuk calon 19  
#### 0: input berhenti  
#### Jadi, outputnya adalah:
#### Suara masuk: 10
#### Suara sah: 8
#### 1: 1
#### 2: 1
#### 3: 2
#### 7: 1
#### 18: 1
#### 19: 2

## Unguided 

### 2. Berdasarkan program sebelumnya, buat program pilkart yang mencari siapa pemenang  pemilihan ketua RT. Sekaligus juga ditentukan bahwa wakil ketua RT adalah calon yang mendapatkan suara terbanyak kedua. Jika beberapa calon mendapatkan suara terbanyak yang sama, ketua terpilih adalah dengan nomor peserta yang paling kecil dan wakilnya dengan  nomor peserta terkecil berikutnya.
### Masukan hanya satu baris data saja, berisi bilangan bulat valid yang kadang tersisipi dengan  data tidak valid. Data valid adalah bilangan bulat dengan nilai di antara 1 s.d. 20 (inklusif). Data berakhir jika ditemukan sebuah bilangan dengan nilai 0. 
### Keluaran dimulai dengan baris berisi jumlah data suara yang terbaca, diikuti baris yang berisi berapa banyak suara yang valid. Kemudian tercetak calon nomor berapa saja yang menjadi pasangan ketua RT dan wakil ketua RT yang baru.

#### soal2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 2](https://github.com/frzhwaa/109082500200_ZirlynailaFairuzahwa/blob/main/modul12/output/soal2.png)
#### Program tersebut digunakan untuk menentukan calon ketua dan wakil ketua RT berdasarkan jumlah suara terbanyak dari data suara yang diinputkan pengguna. Program akan membaca suara satu per satu, menghitung total suara masuk, total suara sah, lalu mencari dua calon dengan suara tertinggi.
#### Program menggunakan tipe bentukan array bernama arrSuara yang bertipe int dengan kapasitas maksimum MAX = 20. Array ini digunakan untuk menyimpan jumlah suara masing-masing calon (indeks 1 sampai 20).
#### Pada fungsi ReadSuara, program membaca input suara berupa angka. Jika angka yang dimasukkan adalah 0, maka proses input berhenti. Jika angka berada dalam rentang 1 sampai 20, maka suara dianggap sah dan ditambahkan ke calon yang sesuai. Setiap input (selain 0) akan menambah jumlah total suara masuk.  
#### Pada fungsi cariPemenang, program mencari dua calon dengan suara terbanyak. Variabel maks1 menyimpan suara terbanyak (ketua). Variabel maks2 menyimpan suara terbanyak kedua (wakil). Program melakukan perulangan dari calon 1 sampai 20, membandingkan jumlah suara, lalu memperbarui nilai ketua dan wakil sesuai hasil.
#### Pada fungsi main, program menginisialisasi array suara dan variabel penghitung. Kemudian memanggil ReadSuara untuk membaca input, lalu cariPemenang untuk menentukan ketua dan wakil.
#### Misalkan input yang diberikan adalah:
#### 7 19 3 2 78 3 1 -3 18 19 0
#### Prosesnya yaitu:
#### 7: suara sah untuk calon 7  
#### 19: suara sah untuk calon 19  
#### 3: suara sah untuk calon 3  
#### 2: suara sah untuk calon 2  
#### 78: tidak sah (di luar 1–20)  
#### 3: suara sah untuk calon 3  
#### 1: suara sah untuk calon 1  
#### -3: tidak sah (di luar 1–20)  
#### 18: suara sah untuk calon 18  
#### 19: suara sah untuk calon 19  
#### 0: input berhenti
#### Jadi, outputnya adalah:
#### Suara masuk: 10  
#### Suara sah: 8  
#### Ketua RT: 19
#### Wakil ketua: 3

## Unguided 

### 3. Diberikan n data integer positif dalam keadaan terurut membesar dan sebuah integer lain k,  apakah bilangan k tersebut ada dalam daftar bilangan yang diberikan? Jika ya, berikan indeksnya, jika tidak sebutkan "TIDAK ADA". 
### Masukan terdiri dari dua baris. Baris pertama berisi dua buah integer positif, yaitu n dan k. n  menyatakan banyaknya data, dimana 1 < n <= 1000000. k adalah bilangan yang ingin dicari. Baris kedua berisi n buah data integer positif yang sudah terurut membesar. 
### Keluaran terdiri dari satu baris saja, yaitu sebuah bilangan yang menyatakan posisi data yang  dicari (k) dalam kumpulan data yang diberikan. Posisi data dihitung dimulai dari angka 0. Atau memberikan keluaran "TIDAK ADA" jika data k tersebut tidak ditemukan dalam kumpulan. 
### Program yang dibangun harus menggunakan subprogram dengan mengikuti kerangka yang  sudah diberikan berikut ini.
### package main
### import "fmt"
### const NMAX = 1000000
### var data [NMAX]int
### func main(){
### /* buatlah kode utama yang membaca baris pertama (n dan k). kemudian data diisi oleh prosedur isiArray(n), dan pencarian oleh fungsi posisi(n,k), dan setelah itu output dicetak. */
### }
### func isiArray(n int){
### /* I.S. terdefinisi integer n, dan sejumlah n data sudah siap pada piranti masukan.
### F.S. Array data berisi n (<=NMAX) bilangan */
### }
### func posisi(n, k int) int {
### /* mengembalikan posisi k dalam array data dengan n elemen. Posisi dimulai dari posisi 0. Jika tidak ada kembalikan -1 */
### }

#### soal3.go

```go
package main
import "fmt"

const NMAX = 1000000
var data [NMAX]int

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	isiArray(n)
	idx := posisi(n, k)
	if idx == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(idx)
	}
}

func isiArray(n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	var kiri, kanan, tengah int
	kiri = 0
	kanan = n - 1
	for kiri <= kanan {
		tengah = (kiri + kanan) / 2
		if data[tengah] == k {
			return tengah
		} else if k < data[tengah] {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}
	return -1
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 3](https://github.com/frzhwaa/109082500200_ZirlynailaFairuzahwa/blob/main/modul12/output/soal3.png)
#### Program tersebut digunakan untuk mencari posisi suatu bilangan dalam array menggunakan algoritma binary search. Program akan membaca jumlah data n dan nilai yang dicari k, kemudian membaca isi array sebanyak n elemen, lalu mencari apakah k ada di dalam array.
#### Program menggunakan tipe bentukan array global bernama data yang bertipe int dengan kapasitas maksimum NMAX = 1000000. Array ini digunakan untuk menyimpan seluruh data bilangan yang diinputkan.
#### Pada fungsi main, program pertama-tama membaca nilai n (jumlah data) dan k (nilai yang dicari). Setelah itu, program memanggil fungsi isiArray untuk mengisi array dengan n bilangan yang diinputkan pengguna. Kemudian program memanggil fungsi posisi untuk mencari nilai k dalam array menggunakan binary search. Jika nilai ditemukan, program menampilkan indeks posisi nilai tersebut. Jika tidak ditemukan, program menampilkan "TIDAK ADA".  
#### Pada fungsi isiArray, program membaca n bilangan dari input dan menyimpannya ke dalam array data.
#### Pada fungsi posisi, program melakukan pencarian dengan binary search, yaitu:  
#### Variabel kiri menyimpan indeks awal (0).  
#### Variabel kanan menyimpan indeks akhir (n-1).  
#### Variabel tengah menyimpan indeks tengah dari array.  
#### Jika data[tengah] == k, maka nilai ditemukan dan indeks dikembalikan.  
#### Jika k < data[tengah], pencarian dilanjutkan ke bagian kiri array.  
#### Jika k > data[tengah], pencarian dilanjutkan ke bagian kanan array.  
#### Jika tidak ditemukan, fungsi mengembalikan -1.
#### Misalkan input yang diberikan adalah:
#### 12 534
#### 1 3 8 16 32 123 323 323 534 543 823 999
#### Proses:
#### n = 12, k = 534
#### Array berisi 12 bilangan terurut.  
#### Binary search menemukan bahwa 534 ada di indeks ke-8 (indeks mulai dari 0).  
#### Jadi, outputnya adalah:
#### 8
#### Misalkan input:
#### 12 535
#### 1 3 8 16 32 123 323 323 534 543 823 999
#### Proses:
#### n = 12, k = 535 
#### Array berisi 12 bilangan terurut.  
#### Binary search tidak menemukan 535 dalam array.  
#### Jadi, outputnya adalah:
#### TIDAK ADA