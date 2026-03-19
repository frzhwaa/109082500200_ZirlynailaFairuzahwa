# <h1 align="center">Laporan Praktikum Modul 3 - FUNGSI </h1>
<p align="center">ZIRLYNAILA FAIRUZAHWA - 109082500200</p>

## Unguided 

### 1. Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p)
### Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi, dengan syarat a ≥ c dan b ≥ d.
### Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d.
### Catatan: permutasi (P) dan kombinasi (C) dari n terhadap r (n ≥ r) dapat dihitung dengan menggunakan persamaan berikut!
### P(n, r) = n! / (n−r)!, sedangkan C(n, r) = n! / r!(n−r)!
### Selesaikan program tersebut dengan memanfaatkan subprogram yang diberikan berikut ini!
### function factorial(n: integer) → integer
### {mengembalikan nilai faktorial dari n}
### function permutation(n,r : integer) → integer
### {Mengembalikan hasil n permutasi r, dan n >= r}
### function combination(n,r : integer) → integer
### {Mengembalikan hasil n kombinasi r, dan n >= r}

#### soal1.go

```go
package main
import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	if a >= c && b >= d {
		fmt.Println(permutation(a, c), combination(a, c))
		fmt.Println(permutation(b, d), combination(b, d))
	}
}

func factorial(n int) int{
	var hasil int = 1
	var i int
	for i = 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}

func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1](https://github.com/frzhwaa/109082500200_ZirlynailaFairuzahwa/blob/main/modul3/output/soal1.png)
#### Program tersebut meminta pengguna untuk menginputkan empat buah bilangan bulat yaitu a, b, c, dan d. Setelah data dimasukkan, program akan mengecek apakah memenuhi kondisi a >= c dan b >= d. Jika kondisi terpenuhi, maka program akan menghitung nilai permutasi dan kombinasi menggunakan fungsi permutation dan combination yang telah dibuat. Fungsi factorial digunakan untuk menghitung nilai faktorial dari suatu bilangan dengan cara mengalikan semua bilangan dari 1 sampai n. Nilai tersebut kemudian digunakan dalam fungsi permutation dengan rumus n!/(n−r)! dan fungsi combination dengan rumus n!/r!(n−r)!.
#### Setelah perhitungan selesai, program akan menampilkan hasil dalam dua baris. Baris pertama berisi hasil permutasi(a, c) dan kombinasi(a, c), sedangkan baris kedua berisi permutasi(b, d) dan kombinasi(b, d). Sebagai contoh, jika diinputkan 5 10 3 10, maka akan menghasilkan 60 10 pada baris pertama dan 3628800 1 pada baris kedua. Contoh lain, jika diinputkan 8 0 2 0, maka hasilnya adalah 56 28 pada baris pertama dan 1 1 pada baris kedua.

## Unguided 

### 2. Diberikan tiga buah fungsi matematika yaitu f(x) = x^2, g(x) = x − 2 dan h(x) = x + 1. Fungsi komposisi (fogoh)(x) artinya adalah f(g(h(x))). Tuliskan f(x), g(x) dan h(x) dalam bentuk function.
### Masukan terdiri dari sebuah bilangan bulat a, b dan c yang dipisahkan oleh spasi.
### Keluaran terdiri dari tiga baris. Baris pertama adalah (fogoh)(a), baris kedua (gohof)(b), dan baris ketiga adalah (hofog)(c)!

#### soal2.go

```go
package main
import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(fogoh(a))
	fmt.Println(gohof(b))
	fmt.Println(hofog(c))
}

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

func fogoh(x int) int{
	return f(g(h(x)))
}

func gohof(x int) int{
	return g(h(f(x)))
}

func hofog(x int) int{
	return h(f(g(x)))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 2](https://github.com/frzhwaa/109082500200_ZirlynailaFairuzahwa/blob/main/modul3/output/soal2.png)
#### Program tersebut meminta pengguna untuk menginputkan tiga buah bilangan bulat, yaitu a, b, dan c. Setelah data dimasukkan, program menggunakan tiga fungsi matematika yaitu f(x)=x^2, g(x)=x-2, dan h(x)=x+1, lalu menghitung komposisi fungsi. Komposisi yang digunakan adalah (fogoh)(a)=f(g(h(a))), (gohof)(b)=g(h(f(b))), dan (hofog)(c)=h(f(g(c))). Di dalam fungsi main, program memanggil tiga fungsi komposisi yaitu fogoh(a), gohof(b), dan hofog(c). Masing-masing fungsi komposisi akan memanggil fungsi lain secara bertingkat, dimulai dari fungsi paling dalam ke fungsi paling luar sesuai urutan komposisinya.
#### Sebagai contoh, untuk input 7 2 10, pada fogoh(7) terjadi pemanggilan h(7)=8, lalu g(8)=6, dan f(6)=36. Pada gohof(2) dipanggil f(2)=4, lalu h(4)=5, kemudian g(5)=3. Sedangkan hofog(10) memanggil g(10)=8, lalu f(8)=64, dan terakhir h(64)=65, sehingga outputnya adalah 36, 3, dan 65. Proses yang sama berlaku untuk input 5 5 5 yang menghasilkan 16, 24, 10 serta input 3 8 4 yang menghasilkan 4, 63, 5.

## Unguided 

### 3. Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y) berdasarkan dua lingkaran tersebut.
### Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat.
### Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".
### Fungsi untuk menghitung jarak titik (a, b) dan (c, d) dimana rumus jarak adalah:
### jarak = √(a − c)^2 + (b − d)^2
### dan juga fungsi untuk menentukan posisi sebuah titik sembarang berada di dalam suatu lingkaran atau tidak.
### function jarak(a,b,c,d : real) -> real
### {Mengembalikan jarak antara titik (a,b) dan titik (c,d)}
### function didalam(cx,cy,r,x,y : real) -> boolean
### {Mengembalikan true apabila titik (x,y) berada di dalam lingkaran yang memiliki titik pusat (cx,cy) dan radius r}
### Catatan: Lihat paket math dalam lampiran untuk menggunakan fungsi math.Sqrt() untuk menghitung akar kuadrat.

#### soal3.go

```go
package main
import (
	"fmt"
	"math"
)

func main() {
	var cx1, cx2, cy1, cy2, r1, r2, x, y float64
	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)
	dalam1 := didalam(cx1, cy1, r1, x, y)
	dalam2 := didalam(cx2, cy2, r2, x, y)
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

func jarak(a, b, c, d float64) float64 {
	var hasil float64
	hasil = math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
	return hasil
}

func didalam(cx, cy, r, x, y float64) bool {
	var hasil bool
	hasil = jarak(x, y, cx, cy) <= r
	return hasil
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 3](https://github.com/frzhwaa/109082500200_ZirlynailaFairuzahwa/blob/main/modul3/output/soal3.png)
#### Program tersebut meminta pengguna untuk menginputkan tiga baris data, yaitu baris pertama dan kedua berisi koordinat titik pusat serta radius dari dua lingkaran (cx1, cy1, r1) dan (cx2, cy2, r2), serta baris ketiga berisi titik sembarang (x, y). Program kemudian menggunakan fungsi jarak(a,b,c,d) untuk menghitung jarak antara titik (x,y) dengan pusat lingkaran menggunakan rumus √(a-c)^2 + (b-d)^2. Selanjutnya fungsi didalam(cx,cy,r,x,y) akan dipanggil untuk mengecek apakah titik berada di dalam lingkaran (jarak <= radius). Di dalam fungsi main, hasil dari dua pemanggilan fungsi didalam disimpan ke variabel dalam1 dan dalam2, lalu digunakan dalam percabangan if-else untuk menentukan apakah titik berada di dalam lingkaran 1, lingkaran 2, keduanya, atau di luar keduanya.
#### Sebagai contoh, pada input 1 1 5, 8 8 4, dan 2 2, titik lebih dekat ke pusat lingkaran 1 sehingga outputnya "Titik di dalam lingkaran 1". Pada input 1 2 3, 4 5 6, dan 7 8, titik berada dalam lingkaran 2 sehingga hasilnya "Titik di dalam lingkaran 2". Pada input 5 10 15, -15 4 20, dan 0 0, titik berada di dalam kedua lingkaran sehingga outputnya "Titik di dalam lingkaran 1 dan 2". Sedangkan pada input 1 1 5, 8 8 4, dan 15 20, titik berada di luar kedua lingkaran sehingga menghasilkan "Titik di luar lingkaran 1 dan 2".