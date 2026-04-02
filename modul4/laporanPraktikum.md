# <h1 align="center">Laporan Praktikum Modul 4 - PROSEDUR </h1>
<p align="center">ZIRLYNAILA FAIRUZAHWA - 109082500200</p>

## Unguided 

### 1. Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p)
### Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi, dengan syarat a ≥ c dan b ≥ d.
### Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d.
### Catatan: permutasi (P) dan kombinasi (C) dari n terhadap r (n ≥ r) dapat dihitung dengan menggunakan persamaan berikut!
### P(n, r) = n! / (n−r)!, sedangkan C(n, r) = n! / r!(n−r)!
### Selesaikan program tersebut dengan memanfaatkan prosedur yang diberikan berikut ini!
### procedure factorial(in n: integer, in/out hasil:integer)
### {I.S. terdefinisi bilangan bulat positif n
### F.S. hasil berisi nilai faktorial dari n}
### procedure permutation(in n,r : integer, in/out hasil:integer)
### {I.S. terdefinisi bilangan bulat positif n dan r, dan n >= r
### F.S. hasil berisi nilai dari n permutasi r}
### procedure combination(in n,r : integer, in/out hasil:integer)
### {I.S. terdefinisi bilangan bulat positif n dan r, dan n >= r}

#### soal1.go

```go
package main
import "fmt"

func factorial(n int, hasil *int){
	*hasil = 1
	var i int
	for i = 1; i <= n; i++ {
		*hasil = *hasil * i
	}
}

func permutation(n, r int, hasil *int) {
	var fn, fnr int
	factorial(n, &fn)
	factorial(n-r, &fnr)
	*hasil = fn / fnr
}

func combination(n, r int, hasil *int) {
	var fn, fr, fnr int
	factorial(n, &fn)
	factorial(r, &fr)
	factorial(n-r, &fnr)
	*hasil = fn / (fr * fnr)
}

func main() {
	var a, b, c, d, hasilP, hasilC int
	fmt.Scan(&a, &b, &c, &d)
	if a >= c && b >= d {
		permutation(a, c, &hasilP)
		combination(a, c, &hasilC)
		fmt.Println(hasilP, hasilC)
		permutation(b, d, &hasilP)
		combination(b, d, &hasilC)
		fmt.Println(hasilP, hasilC)
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1](https://github.com/frzhwaa/109082500200_ZirlynailaFairuzahwa/blob/main/modul4/output/soal1.png)
#### Program tersebut meminta pengguna untuk menginputkan empat buah bilangan bulat yaitu a, b, c, dan d. Setelah data dimasukkan, program akan mengecek apakah memenuhi kondisi a >= c dan b >= d. Jika kondisi terpenuhi, maka program akan menghitung nilai permutasi dan kombinasi menggunakan prosedur permutation dan combination yang telah dibuat.
#### Prosedur factorial digunakan untuk menghitung nilai faktorial dari suatu bilangan dengan cara mengalikan semua bilangan dari 1 sampai n. Prosedur ini menerima input berupa n dan menghasilkan nilai faktorial melalui parameter hasil (pointer). Awalnya nilai hasil diisi 1, kemudian dilakukan perulangan dari 1 sampai n, dan setiap langkah nilai hasil dikalikan dengan i sehingga diperoleh n!.
#### Prosedur permutation digunakan untuk menghitung permutasi dengan rumus n!/(n−r)!. Prosedur ini menerima input n dan r, kemudian memanggil prosedur factorial untuk menghitung n! dan (n−r)!. Hasil dari kedua faktorial tersebut disimpan dalam variabel sementara, lalu dilakukan pembagian n! dengan (n−r)! dan hasilnya disimpan ke parameter hasil.
#### Prosedur combination digunakan untuk menghitung kombinasi dengan rumus n!/(r!(n−r)!). Prosedur ini menerima input n dan r, kemudian memanggil prosedur factorial untuk menghitung n!, r!, dan (n−r)!. Ketiga nilai tersebut disimpan dalam variabel sementara, lalu dihitung dengan membagi n! dengan hasil perkalian r! dan (n−r)!, kemudian hasilnya disimpan ke parameter hasil.
#### Setelah perhitungan selesai, program akan menampilkan hasil dalam dua baris. Baris pertama berisi hasil permutation(a, c) dan combination(a, c), sedangkan baris kedua berisi permutation(b, d) dan combination(b, d). Jika kondisi a >= c dan b >= d tidak terpenuhi, maka program tidak akan menampilkan output. Sebagai contoh, jika diinputkan 5 10 3 10, maka akan menghasilkan 60 10 pada baris pertama dan 3628800 1 pada baris kedua. Contoh lain, jika diinputkan 8 0 2 0, maka hasilnya adalah 56 28 pada baris pertama dan 1 1 pada baris kedua.

## Unguided 

### 2. Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan soal paling banyak dalam waktu paling singkat adalah pemenangnya.
### Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal. Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca di dalam prosedur.
### prosedure hitungSkor(in/out soal, skor : integer)
### Setiap baris masukan dimulai dengan satu string nama peserta tersebut diikuti dengan adalah 8 integer yang menyatakan berapa lama (dalam menit) peserta tersebut menyelesaikan soal. Jika tidak berhasil atau tidak mengirimkan jawaban maka otomatis dianggap menyelesaikan dalam waktu 5 jam 1 menit (301 menit).
### Satu baris keluaran berisi nama pemenang, jumlah soal yang diselesaikan, dan nilai yang diperoleh. Nilai adalah total waktu yang dibutuhkan untuk menyelesaikan soal yang berhasil diselesaikan.

#### soal2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 2](https://github.com/frzhwaa/109082500200_ZirlynailaFairuzahwa/blob/main/modul4/output/soal2.png)
#### Program tersebut digunakan untuk menentukan pemenang dalam kompetisi pemrograman tingkat nasional. Setiap peserta diberikan 8 soal yang harus diselesaikan dalam waktu maksimal 5 jam (300 menit). Peserta yang menyelesaikan soal paling banyak dengan total waktu paling singkat akan menjadi pemenang.
#### Program meminta pengguna untuk menginputkan nama peserta satu per satu. Untuk setiap peserta, program akan membaca 8 buah waktu pengerjaan soal. Proses input akan terus dilakukan sampai pengguna memasukkan kata "Selesai".
#### Program dibuat secara modular dengan menggunakan prosedur hitungSkor untuk menghitung jumlah soal yang berhasil diselesaikan dan total nilai (skor) dari setiap peserta. Prosedur ini menerima parameter soal dan skor sebagai input/output. Pada awalnya nilai soal dan skor diinisialisasi menjadi 0. Kemudian dilakukan perulangan sebanyak 8 kali untuk membaca waktu pengerjaan setiap soal. Jika waktu yang diinputkan kurang dari atau sama dengan 300 menit, maka soal dianggap berhasil diselesaikan sehingga nilai soal bertambah 1 dan nilai skor ditambahkan dengan waktu tersebut. Jika waktu lebih dari 300 menit (misalnya 301), maka soal dianggap tidak berhasil diselesaikan sehingga tidak menambah jumlah soal maupun skor.
#### Di dalam program utama, setiap nama peserta yang dimasukkan akan diproses menggunakan prosedur hitungSkor untuk mendapatkan jumlah soal yang berhasil diselesaikan dan total nilainya. Program kemudian membandingkan hasil setiap peserta untuk menentukan pemenang. Pemenang adalah peserta dengan jumlah soal terbanyak. Jika terdapat jumlah soal yang sama, maka dipilih peserta dengan nilai (total waktu) yang lebih kecil. Nilai awal skor minimum diinisialisasi dengan nilai besar (saya menggunakan 999999) agar memudahkan proses perbandingan.
#### Setelah semua data peserta dimasukkan dan input "Selesai" diberikan, program akan menampilkan nama pemenang, jumlah soal yang diselesaikan, dan total nilai yang diperoleh. Sebagai contoh, jika diinputkan Astuti 20 50 301 301 61 71 75 10 Bertha 25 47 301 26 50 60 65 21 Selesai, maka hasilnya adalah Bertha 7 294 Hal ini karena Astuti menyelesaikan 6 soal dengan total waktu 287 menit, sedangkan Bertha menyelesaikan 7 soal dengan total waktu 294 menit, sehingga Bertha menjadi pemenang karena menyelesaikan soal lebih banyak.

## Unguided 

### 3. Skiena dan Revilla dalam Programming Challenges mendefinisikan sebuah deret bilangan. Deret dimulai dengan sebuah bilangan bulat n. Jika bilangan n saat itu genap, maka suku berikutnya adalah 1⁄2n, tetapi jika ganjil maka suku berikutnya bernilai 3n+1. Rumus yang sama digunakan terus menerus untuk mencari suku berikutnya. Deret berakhir ketika suku terakhir bernilai 1. Sebagai contoh jika dimulai dengan n=22, maka deret bilangan yang diperoleh adalah:
### 22 11 34 17 52 26 13 40 20 10 5 16 8 4 2 1
### Untuk suku awal sampai dengan 1000000, diketahui deret selalu mencapai suku dengan nilai 1.
### Buat program skiena yang akan mencetak setiap suku dari deret yang dijelaskan di atas untuk nilai suku awal yang diberikan. Pencetakan deret harus dibuat dalam prosedur cetakDeret yang mempunyai 1 parameter formal, yaitu nilai dari suku awal.
### prosedure cetakDeret(in n : integer )
### Masukan berupa satu bilangan integer positif yang lebih kecil dari 1000000.
### Keluaran terdiri dari satu baris saja. Setiap suku dari deret tersebut dicetak dalam baris yang dan dipisahkan oleh sebuah spasi.

#### soal3.go

```go
package main
import "fmt"

func cetakDeret(n int) {
	for n != 1 {
		fmt.Print(n, " ")
		if n % 2 == 0 {
			n = n / 2
		} else {
			n = 3 * n + 1
		}
	}
	fmt.Print(1)
}

func main() {
	var n int
	fmt.Scan(&n)
	cetakDeret(n)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 3](https://github.com/frzhwaa/109082500200_ZirlynailaFairuzahwa/blob/main/modul4/output/soal3.png)
#### Program tersebut digunakan untuk mencetak deret bilangan berdasarkan aturan Skiena dan Revilla. Deret dimulai dari sebuah bilangan bulat positif n yang lebih kecil dari 1000000, kemudian deret akan berlanjut hingga mencapai nilai 1.
#### Program meminta pengguna untuk menginputkan sebuah bilangan bulat n sebagai suku awal deret. Selanjutnya program akan memanggil prosedur cetakDeret untuk menampilkan deret bilangan tersebut. Prosedur cetakDeret menerima satu parameter yaitu n sebagai nilai awal. Di dalam prosedur, dilakukan perulangan selama nilai n tidak sama dengan 1. Pada setiap langkah, nilai n akan dicetak terlebih dahulu. Kemudian program akan mengecek apakah n merupakan bilangan genap atau ganjil. Jika n genap, maka nilai n selanjutnya menjadi n dibagi 2. Jika n ganjil, maka nilai n selanjutnya menjadi 3n + 1. Proses ini dilakukan terus menerus hingga n bernilai 1. Setelah perulangan selesai, angka 1 dicetak sebagai suku terakhir dari deret.
#### Hasil keluaran berupa satu baris yang berisi seluruh suku deret, dipisahkan oleh spasi. Sebagai contoh, jika diinputkan 22, maka hasilnya adalah 22 11 34 17 52 26 13 40 20 10 5 16 8 4 2 1. Deret tersebut diperoleh dengan mengikuti aturan, yaitu setiap bilangan genap dibagi 2 dan setiap bilangan ganjil dikalikan 3 lalu ditambah 1, hingga akhirnya mencapai nilai 1.