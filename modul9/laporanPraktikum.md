# <h1 align="center">Laporan Praktikum Modul 9 - ARRAY </h1>
<p align="center">ZIRLYNAILA FAIRUZAHWA - 109082500200</p>

## Unguided 

### 1. Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila  diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y) berdasarkan dua lingkaran tersebut. Gunakan tipe bentukan titik untuk menyimpan koordinat, dan tipe bentukan lingkaran untuk menyimpan titik pusat lingkaran dan radiusnya.
### Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat.
### Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".
### Fungsi untuk menghitung jarak titik (a, b) dan (c, d) dimana rumus jarak adalah:
### jarak = √(a − c) ^2 + (b − d)^2 
### dan juga fungsi untuk menentukan posisi sebuah titik sembarang berada di dalam suatu  lingkaran atau tidak.
### function jarak(p, q : titik) -> real
### {Mengembalikan jarak antara titik p(x,y) dan titik q(x,y)}
### function didalam(c:lingkaran, p:titik) -> boolean
### {Mengembalikan true apabila titik p(x,y) berada di dalam lingkaran c yang memiliki titik pusat (cx,cy) dan radius r}
### Catatan: Lihat paket math dalam lampiran untuk menggunakan fungsi math.Sqrt() untuk  menghitung akar kuadrat.

#### soal1.go

```go
package main
import (
	"fmt"
	"math"
)

type titik struct {
	x, y int
}

type lingkaran struct {
	pusat titik
	r int
}

func jarak(p, q titik) float64 {
	return math.Sqrt( math.Pow(float64(p.x-q.x), 2) + math.Pow(float64(p.y-q.y), 2))
}

func didalam(c lingkaran, p titik) bool {
	return jarak(c.pusat, p) <= float64(c.r)
}

func main() {
	var l1, l2 lingkaran
	var p titik
	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.r)
	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.r)
	fmt.Scan(&p.x, &p.y)
	dalam1 := didalam(l1, p)
	dalam2 := didalam(l2, p)
	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1](https://github.com/frzhwaa/109082500200_ZirlynailaFairuzahwa/blob/main/modul9/output/soal1.png)
#### Program tersebut meminta pengguna untuk menginputkan data dua buah lingkaran dan satu titik sembarang. Setiap lingkaran memiliki titik pusat berupa koordinat (x, y) serta jari-jari (radius). Setelah semua data dimasukkan, program akan menentukan posisi titik tersebut terhadap kedua lingkaran, apakah berada di dalam lingkaran 1, lingkaran 2, keduanya, atau di luar keduanya.
#### Program menggunakan tipe bentukan titik untuk menyimpan koordinat x dan y, sedangkan tipe bentukan lingkaran digunakan untuk menyimpan titik pusat lingkaran dan nilai radiusnya.
#### Fungsi jarak digunakan untuk menghitung jarak antara dua titik menggunakan rumus
#### jarak = √(a − c) ^2 + (b − d)^2 
#### Fungsi tersebut menerima dua parameter bertipe titik, yaitu titik pusat lingkaran dan titik sembarang. Hasil dari fungsi tersebut berupa nilai jarak dalam bentuk bilangan real (float64).
#### Selanjutnya, fungsi didalam digunakan untuk menentukan apakah sebuah titik berada di dalam lingkaran atau tidak. Caranya adalah dengan membandingkan hasil jarak antara titik pusat lingkaran dan titik sembarang dengan nilai radius lingkaran. Jika jarak lebih kecil atau sama dengan radius, maka titik tersebut berada di dalam lingkaran.
#### Setelah itu, pada fungsi main, program akan membaca input untuk lingkaran 1, lingkaran 2, dan titik sembarang. Program lalu memanggil fungsi didalam untuk masing-masing lingkaran dan menyimpan hasilnya dalam variabel boolean dalam1 dan dalam2.
#### Program menggunakan percabangan if else untuk menentukan output dengan ketentuan:
#### Jika titik berada di dalam kedua lingkaran, maka outputnya "Titik di dalam lingkaran 1 dan 2"
#### Jika hanya di dalam lingkaran 1, maka outputnya "Titik di dalam lingkaran 1"
#### Jika hanya di dalam lingkaran 2, maka outputnya "Titik di dalam lingkaran 2"
#### Jika tidak berada di dalam keduanya, maka outpunya "Titik di luar lingkaran 1 dan 2"
#### Sebagai contoh, jika inputnya adalah:
#### 1 1 5
#### 8 8 4
#### 2 2
#### Maka titik (2,2) akan dihitung jaraknya terhadap pusat lingkaran 1 yaitu (1,1), dan hasilnya lebih kecil dari radius 5 sehingga titik berada di dalam lingkaran 1. Sedangkan terhadap lingkaran 2, jaraknya lebih besar dari radius 4 sehingga titik berada di luar lingkaran 2. Oleh karena itu, output yang dihasilkan adalah:
#### Titik di dalam lingkaran 1

## Unguided 

### 2. Sebuah array digunakan untuk menampung sekumpulan bilangan bulat. Buatlah program yang digunakan untuk mengisi array tersebut sebanyak N elemen nilai. Asumsikan array memiliki kapasitas penyimpanan data sejumlah elemen tertentu. Program dapat menampilkan beberapa informasi berikut: 
### a. Menampilkan keseluruhan isi dari array.
### b. Menampilkan elemen-elemen array dengan indeks ganjil saja.
### c. Menampilkan elemen-elemen array dengan indeks genap saja (asumsi indek ke-0 adalah  genap). 
### d. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x. x bisa diperoleh  dari masukan pengguna. 
### e. Menghapus elemen array pada indeks tertentu, asumsi indeks yang hapus selalu valid.  Tampilkan keseluruhan isi dari arraynya, pastikan data yang dihapus tidak tampil 
### f. Menampilkan rata-rata dari bilangan yang ada di dalam array.
### g. Menampilkan standar deviasi atau simpangan baku dari bilangan yang ada di dalam array  tersebut.
### h. Menampilkan frekuensi dari suatu bilangan tertentu di dalam array yang telah diisi  tersebut.

#### soal2.go

```go
package main
import (
	"fmt"
	"math"
)

const max = 100
type bilangan struct {
	info [max]int
	n int
}

func tampilSemua(A bilangan) {
	for i := 0; i < A.n; i++ {
		fmt.Print(A.info[i], " ")
	}
	fmt.Println()
}

func tampilIndeksGanjil(A bilangan) {
	for i := 1; i < A.n; i += 2 {
		fmt.Print(A.info[i], " ")
	}
	fmt.Println()
}

func tampilIndeksGenap(A bilangan) {
	for i := 0; i < A.n; i += 2 {
		fmt.Print(A.info[i], " ")
	}
	fmt.Println()
}

func tampilKelipatanX(A bilangan, x int) {
	for i := 0; i < A.n; i++ {
		if i%x == 0 {
			fmt.Print(A.info[i], " ")
		}
	}
	fmt.Println()
}

func hapusIndex(A *bilangan, idx int) {
	for i := idx; i < A.n-1; i++ {
		A.info[i] = A.info[i+1]
	}
	A.n--
}

func rataRata(A bilangan) float64 {
	total := 0
	for i := 0; i < A.n; i++ {
		total += A.info[i]
	}
	return float64(total) / float64(A.n)
}

func standarDeviasi(A bilangan) float64 {
	avg := rataRata(A)
	var jumlah float64
	for i := 0; i < A.n; i++ {
		jumlah += math.Pow(float64(A.info[i])-avg, 2)
	}
	return math.Sqrt(jumlah / float64(A.n))
}

func frekuensi(A bilangan, x int) int {
	hitung := 0
	for i := 0; i < A.n; i++ {
		if A.info[i] == x {
			hitung++
		}
	}
	return hitung
}

func main() {
	var A bilangan
	var x, idx, cari int
	fmt.Scan(&A.n)
	for i := 0; i < A.n; i++ {
		fmt.Scan(&A.info[i])
	}
	fmt.Println("Seluruh isi array:")
	tampilSemua(A)
	fmt.Println("Indeks ganjil:")
	tampilIndeksGanjil(A)
	fmt.Println("Indeks genap:")
	tampilIndeksGenap(A)
	fmt.Println("Masukkan nilai x untuk indeks kelipatan:")
	fmt.Scan(&x)
	tampilKelipatanX(A, x)
	fmt.Println("Masukkan indeks yang akan dihapus:")
	fmt.Scan(&idx)
	hapusIndex(&A, idx)
	fmt.Println("Array setelah dihapus:")
	tampilSemua(A)
	fmt.Printf("Rata-rata: %.2f\n", rataRata(A))
	fmt.Printf("Standar deviasi: %.2f\n", standarDeviasi(A))
	fmt.Println("Masukkan bilangan yang ingin dicari frekuensinya:")
	fmt.Scan(&cari)
	fmt.Println("Frekuensi:", frekuensi(A, cari))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 2](https://github.com/frzhwaa/109082500200_ZirlynailaFairuzahwa/blob/main/modul9/output/soal2(1).png)
![Screenshot Output Unguided 2](https://github.com/frzhwaa/109082500200_ZirlynailaFairuzahwa/blob/main/modul9/output/soal2(2).png)
![Screenshot Output Unguided 2](https://github.com/frzhwaa/109082500200_ZirlynailaFairuzahwa/blob/main/modul9/output/soal2(3).png)
#### Program tersebut meminta pengguna untuk menginputkan sejumlah bilangan bulat ke dalam sebuah array. Pertama, pengguna memasukkan jumlah elemen array n, kemudian dilanjutkan dengan mengisi nilai-nilai array sebanyak n elemen. Setelah data dimasukkan, program akan menampilkan berbagai informasi terkait isi array tersebut seperti seluruh isi array, elemen pada indeks tertentu, rata-rata, standar deviasi, hingga frekuensi suatu bilangan.
#### Program menggunakan tipe bentukan bilangan untuk menyimpan array bilangan bulat. Struct ini memiliki dua bagian, yaitu info sebagai array dengan kapasitas maksimum max, dan n sebagai penanda jumlah elemen yang terisi.
#### Fungsi tampilSemua digunakan untuk menampilkan seluruh elemen yang ada di dalam array. Fungsi ini melakukan perulangan dari indeks 0 sampai n-1, lalu mencetak semua nilai yang tersimpan.
#### Fungsi tampilIndeksGanjil digunakan untuk menampilkan elemen-elemen yang berada pada indeks ganjil, yaitu indeks 1, 3, 5, dan seterusnya. Sebaliknya, fungsi tampilIndeksGenap(A) digunakan untuk menampilkan elemen-elemen pada indeks genap, yaitu indeks 0, 2, 4, dan seterusnya.
#### Fungsi tampilKelipatanX digunakan untuk menampilkan elemen-elemen yang berada pada indeks kelipatan bilangan x. Nilai x dimasukkan oleh pengguna. Program akan memeriksa setiap indeks, jika indeks tersebut habis dibagi x, maka elemennya akan ditampilkan.
#### Fungsi hapusIndex digunakan untuk menghapus elemen pada indeks tertentu. Caranya adalah dengan menggeser semua elemen setelah indeks tersebut ke kiri satu posisi, kemudian jumlah elemen n dikurangi satu. Dengan befitu, elemen yang dihapus tidak akan tampil lagi saat array ditampilkan.
#### Fungsi rataRata digunakan untuk menghitung nilai rata-rata seluruh elemen array. Program menjumlahkan semua elemen terlebih dahulu, kemudian hasilnya dibagi dengan jumlah elemen.
#### Fungsi standarDeviasi digunakan untuk menghitung standar deviasi dari data dalam array. Program terlebih dahulu menghitung nilai rata-rata, kemudian mencari selisih setiap elemen terhadap rata-rata, mengkuadratkannya, menjumlahkan semuanya, lalu diakhiri dengan akar kuadrat.
#### Fungsi frekuensi digunakan untuk menghitung berapa kali suatu bilangan tertentu muncul di dalam array. Program akan memeriksa setiap elemen, jika nilainya sama dengan bilangan yang dicari, maka variabel penghitung akan bertambah satu.
#### Pada fungsi main, program akan memanggil semua fungsi tersebut secara berurutan. Program dimulai dengan input jumlah elemen dan isi array, kemudian menampilkan seluruh isi array, elemen indeks ganjil, elemen indeks genap, elemen pada indeks kelipatan x, menghapus elemen pada indeks tertentu, menghitung rata-rata, standar deviasi, dan terakhir menghitung frekuensi suatu bilangan.

## Unguided 

### 3. Sebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang  memenangkan pertandingan bola pada suatu grup pertandingan. Buatlah program yang digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga. 
### Pertama-tama program meminta masukan nama-nama klub yang bertanding, kemudian  program meminta masukan skor hasil pertandingan kedua klub tersebut. Yang disimpan dalam array adalah nama-nama klub yang menang saja. 
### Proses input skor berhenti ketika skor salah satu atau kedua klub tidak valid (negatif). Di akhir  program, tampilkan daftar klub yang memenangkan pertandingan. 
### Perhatikan sesi interaksi pada contoh berikut ini (teks bergaris bawah adalah input/read)
### Klub A : MU
### Klub B : Inter
### Pertandingan 1 : 2 0 // MU = 2 sedangkan Inter = 0
### Pertandingan 2 : 1 2
### Pertandingan 3 : 2 2
### Pertandingan 4 : 0 1
### Pertandingan 5 : 3 2
### Pertandingan 6 : 1 0
### Pertandingan 7 : 5 2
### Pertandingan 8 : 2 3
### Pertandingan 9 : -1 2
### Hasil 1 : MU
### Hasil 2 : Inter
### Hasil 3 : Draw
### Hasil 4 : Inter
### Hasil 5 : MU
### Hasil 6 : MU
### Hasil 7 : MU
### Hasil 8 : Inter
### Pertandingan selesai

#### soal3.go

```go
package main
import "fmt"

const maks = 100
type DaftarKlub struct {
	nama [maks]string
	n int
}

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var hasil DaftarKlub
	var pertandingan int = 1
	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)
	for {
		fmt.Printf("Pertandingan %d : ", pertandingan)
		fmt.Scan(&skorA, &skorB)
		if skorA < 0 || skorB < 0 {
			break
		}
		if skorA > skorB {
			fmt.Println("Hasil", pertandingan, ":", klubA)
			hasil.nama[hasil.n] = klubA
			hasil.n++
		} else if skorB > skorA {
			fmt.Println("Hasil", pertandingan, ":", klubB)
			hasil.nama[hasil.n] = klubB
			hasil.n++
		} else {
			fmt.Println("Hasil", pertandingan, ": Draw")
		}
		pertandingan++
	}
	fmt.Println("Pertandingan selesai")
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 3](https://github.com/frzhwaa/109082500200_ZirlynailaFairuzahwa/blob/main/modul9/output/soal3.png)
#### Program tersebut digunakan untuk merekap hasil pertandingan antara dua klub sepak bola. Pertama, program meminta pengguna untuk memasukkan nama dua klub yang akan bertanding, yaitu Klub A dan Klub B. Setelah itu, program akan meminta input skor dari setiap pertandingan secara berulang sampai salah satu atau kedua skor bernilai negatif. Nilai negatif digunakan sebagai tanda bahwa proses input pertandingan selesai.
#### Program menggunakan tipe bentukan DaftarKlub untuk menyimpan nama-nama klub yang memenangkan pertandingan. Struct tersebut memiliki array nama untuk menampung nama klub pemenang dan variabel n untuk menghitung jumlah data yang sudah tersimpan. Hanya klub yang menang saja yang disimpan, sedangkan jika hasil pertandingan seri (draw), maka tidak disimpan ke dalam array.
#### Pada fungsi main, program pertama-tama membaca input nama klub A dan klub B menggunakan fmt.Scan. Setelah itu, program masuk ke dalam perulangan yang digunakan untuk menerima input skor pertandingan satu per satu.
#### Setiap pertandingan terdiri dari dua skor, yaitu skor untuk Klub A: skorA dan skor untuk Klub B: skorB. Program akan menampilkan nomor pertandingan menggunakan variabel pertandingan yang dimulai dari angka 1 dan terus bertambah setiap kali pertandingan selesai diproses.
#### Di dalam perulangan, program terlebih dahulu memeriksa apakah salah satu skor bernilai negatif. Jika skorA < 0 atau skorB < 0, maka perulangan dihentikan menggunakan break, yang berarti proses pertandingan selesai.
#### Jika skor valid, program akan membandingkan kedua skor tersebut:
#### Jika skorA > skorB, maka Klub A dinyatakan sebagai pemenang
#### Jika skorB > skorA, maka Klub B dinyatakan sebagai pemenang
#### Jika skorA == skorB, maka hasil pertandingan adalah seri atau draw
#### Nama klub yang menang akan langsung ditampilkan sebagai hasil pertandingan, lalu disimpan ke dalam array hasil.nama. Variabel hasil.n akan bertambah satu untuk menandakan jumlah data pemenang yang tersimpan. Jika hasil draw, program hanya menampilkan tulisan "Draw" tanpa menyimpan nama klub ke array.
#### Proses ini akan terus berulang sampai pengguna memasukkan skor negatif. Setelah itu, program akan menampilkan tulisan "Pertandingan selesai" sebagai tanda bahwa seluruh proses telah berakhir.

## Unguided 

### 4. Sebuah array digunakan untuk menampung sekumpulan karakter, Anda diminta untuk membuat sebuah subprogram untuk melakukan membalikkan urutan isi array dan memeriksa apakah membentuk palindrom.
### Lengkapi potongan algoritma berikut ini!
### package main
### import "fmt"
### const NMAX int = 127
### type tabel [NMAX]rune
### tab : tabel
### m : integer
### func isiArray(t *tabel, n *int)
### /*I.S. Data tersedia dalam piranti masukan
### F.S. Array t berisi sejumlah n karakter yang dimasukkan user, Proses input selama karakter bukanlah TITIK dan n <= NMAX */
### func cetakArray(t tabel, n int)
### /*I.S. Terdefinisi array t yang berisi sejumlah n karakter
### F.S. n karakter dalam array muncul di layar */
### func balikanArray(t *tabel, n int)
### /*I.S. Terdefinisi array t yang berisi sejumlah n karakter
### F.S. Urutan isi array t terbalik */
### func main(){
### var tab tabel
### var m int
### // si array tab dengan memanggil prosedur isiArray
### // Balikian isi array tab dengan memanggil balikanArray
### // Cetak is array tab
### }
### Perhatikan sesi interaksi pada contoh berikut ini (teks bergaris bawah adalah input/read)
### Teks : S E N A N G .
### Reverse teks : G N A N E S
### Teks : K A T A K .
### Reverse teks : K A T A K
### Modifikasi program tersebut dengan menambahkan fungsi palindrom. Tambahkan instruksi
### untuk memanggil fungsi tersebut dan menampilkan hasilnya pada program utama.
### *Palindrom adalah teks yang dibaca dari awal atau akhir adalah sama, contoh: KATAK, APA, KASUR_RUSAK.
### func palindrom(t tabel, n int) bool
### /* Mengembalikan true apabila susunan karakter di dalam t membentuk palindrom, dan false apabila sebaliknya. Petunjuk: Manfaatkan prosedur balikanArray */
### Perhatikan sesi interaksi pada contoh berikut ini (teks bergaris bawah adalah input/read)
### Teks : K A T A K
### Palindrom ? true
### Teks : S E N A N G 
### Palindrom ? false

#### soal4.go

```go
package main
import "fmt"
const NMAX int = 127
type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var ch rune
	*n = 0
	fmt.Print("Teks : ")
	for {
		fmt.Scanf("%c", &ch)
		if ch == '.' || *n >= NMAX {
			break
		}
		t[*n] = ch
		*n++
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(string(t[i]), " ")
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	var temp rune
	for i := 0; i < n/2; i++ {
		temp = t[i]
		t[i] = t[n-1-i]
		t[n-1-i] = temp
	}
}

func palindrom(t tabel, n int) bool {
	var balik tabel
	for i := 0; i < n; i++ {
		balik[i] = t[i]
	}
	balikanArray(&balik, n)
	for i := 0; i < n; i++ {
		if t[i] != balik[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int
	isiArray(&tab, &m)
	fmt.Println("Palindrom ?", palindrom(tab, m))
	balikanArray(&tab, m)
	fmt.Print("Reverse teks : ")
	cetakArray(tab, m)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 4](https://github.com/frzhwaa/109082500200_ZirlynailaFairuzahwa/blob/main/modul9/output/soal4(1).png)
![Screenshot Output Unguided 4](https://github.com/frzhwaa/109082500200_ZirlynailaFairuzahwa/blob/main/modul9/output/soal4(2).png)
#### Program tersebut digunakan untuk menyimpan sekumpulan karakter ke dalam sebuah array, kemudian membalik urutan karakter tersebut dan memeriksa apakah susunan karakter tersebut membentuk palindrom. Input karakter dilakukan satu per satu dan akan berhenti ketika pengguna memasukkan tanda titik "." sebagai penanda akhir input.
#### Program menggunakan tipe bentukan tabel, yaitu array dengan kapasitas maksimum NMAX sebanyak 127 elemen bertipe rune. Tipe rune digunakan karena program bekerja dengan karakter, sehingga setiap huruf dapat disimpan sebagai satu elemen array.
#### Fungsi isiArray digunakan untuk mengisi array karakter dari input pengguna. Program akan membaca karakter satu per satu menggunakan fmt.Scanf(), lalu menyimpannya ke dalam array selama karakter yang dimasukkan bukan titik "." dan jumlah elemen belum melebihi kapasitas maksimum. Variabel n digunakan untuk menghitung banyaknya karakter yang berhasil disimpan.
#### Fungsi cetakArray digunakan untuk menampilkan seluruh isi array ke layar. Program melakukan perulangan dari indeks 0 sampai n-1, lalu setiap karakter dicetak satu per satu sehingga pengguna dapat melihat isi array setelah dibalik.
#### Fungsi balikanArray digunakan untuk membalik urutan isi array. Proses pembalikan dilakukan dengan cara menukar elemen pertama dengan elemen terakhir, elemen kedua dengan elemen kedua dari belakang, dan seterusnya sampai mencapai tengah array.
#### Fungsi palindrom digunakan untuk memeriksa apakah susunan karakter membentuk palindrom atau tidak. Palindrom adalah kata atau susunan karakter yang jika dibaca dari depan maupun dari belakang tetap sama, seperti "KATAK", "APA", atau "KASURRUSAK".
#### Fungsi ini bekerja dengan cara menyalin isi array asli ke array baru bernama balik, kemudian array salinan tersebut dibalik menggunakan prosedur balikanArray. Setelah itu, program membandingkan isi array asli dengan array hasil pembalikan. Jika semua elemen sama, maka fungsi mengembalikan nilai true, yang berarti teks tersebut adalah palindrom. Jika ada satu saja yang berbeda, maka hasilnya false.
#### Pada fungsi main, program pertama-tama memanggil prosedur isiArray untuk menerima input karakter dari pengguna. Setelah itu, program memanggil fungsi palindrom untuk mengecek apakah teks tersebut merupakan palindrom, lalu hasilnya ditampilkan. Selanjutnya, program memanggil prosedur balikanArray untuk membalik isi array, kemudian hasil pembalikan ditampilkan menggunakan prosedur cetakArray().