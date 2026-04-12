# <h1 align="center">Laporan Praktikum Modul 5 - REKURSIF </h1>
<p align="center">ZIRLYNAILA FAIRUZAHWA - 109082500200</p>

## Unguided 

### 1. Deret fibonacci adalah sebuah deret dengan nilai suku ke-0 dan ke-1 adalah 0 dan 1, dan nilai suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat diformulasikan Sn = Sn−1 + Sn−2 . Berikut ini adalah contoh nilai deret fibonacci hingga suku ke-10. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonacci tersebut.

#### soal1.go

```go
package main
import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    fmt.Println(fibonacci(n))
}

func fibonacci(n int) int {
    if n == 0 {
        return 0
    } else if n == 1 {
        return 1
    } else {
        return fibonacci(n-1) + fibonacci(n-2)
    }
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1](https://github.com/frzhwaa/109082500200_ZirlynailaFairuzahwa/blob/main/modul5/output/soal1.png)
#### Program tersebut meminta pengguna untuk menginputkan sebuah bilangan bulat yaitu n. Setelah nilai n dimasukkan, program akan memanggil fungsi fibonacci(n) untuk menghitung nilai suku ke-n dari deret Fibonacci, kemudian hasilnya akan ditampilkan ke layar.
#### Fungsi fibonacci digunakan untuk menghitung nilai deret Fibonacci secara rekursif. Fungsi ini memiliki dua kondisi dasar (base case), yaitu jika n == 0 maka hasilnya 0, dan jika n == 1 maka hasilnya 1. Kedua kondisi ini digunakan agar rekursi dapat berhenti.
#### Jika nilai n lebih dari 1, maka fungsi akan memanggil dirinya sendiri dengan rumus fibonacci(n-1) + fibonacci(n-2). Artinya, nilai Fibonacci ke-n diperoleh dari penjumlahan dua suku sebelumnya, yaitu suku ke(n-1) dan suku ke(n-2).
#### Proses rekursi ini akan terus berulang hingga mencapai kondisi dasar. Setiap pemanggilan fungsi akan menunggu hasil dari pemanggilan berikutnya sampai semua perhitungan selesai, kemudian hasil akhirnya dikembalikan ke fungsi utama (main).
#### Setelah perhitungan selesai, program akan menampilkan hasil nilai Fibonacci ke-n. Sebagai contoh, jika diinputkan n = 5, maka perhitungannya adalah:
#### fibonacci(5) = fibonacci(4) + fibonacci(3)
#### fibonacci(4) = fibonacci(3) + fibonacci(2)
#### fibonacci(3) = fibonacci(2) + fibonacci(1)
#### fibonacci(2) = fibonacci(1) + fibonacci(0)
#### Sehingga hasil akhirnya adalah 5.

## Unguided 

### 2. Buatlah sebuah program yang digunakan untuk menampilkan pola bintang berikut ini dengan menggunakan fungsi rekursif. N adalah masukan dari user.

#### soal2.go

```go
package main
import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    baris(1, n)
}

func baris(i, n int) {
    if i > n {
        return
    }
    bintang(i)
    fmt.Println()
    baris(i+1, n)
}

func bintang(j int) {
    if j == 0 {
        return
    }
    bintang(j-1)
    fmt.Print("*")
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 2](https://github.com/frzhwaa/109082500200_ZirlynailaFairuzahwa/blob/main/modul5/output/soal2.png)
#### Program tersebut meminta pengguna untuk menginputkan sebuah bilangan bulat yaitu n. Setelah nilai n dimasukkan, program akan memanggil prosedur baris(1, n) untuk menampilkan pola bintang secara bertahap dari 1 hingga n baris.
#### Prosedur baris digunakan untuk mengatur jumlah baris yang akan ditampilkan. Prosedur ini menerima parameter i sebagai nomor baris saat ini dan n sebagai batas maksimal baris. Jika i > n, maka proses dihentikan (base case). Jika belum, maka program akan mencetak bintang sebanyak i buah dengan memanggil prosedur bintang(i), lalu pindah ke baris baru, dan memanggil kembali baris(i+1, n).
#### Prosedur bintang digunakan untuk mencetak jumlah bintang dalam satu baris menggunakan rekursi. Jika j == 0, maka proses dihentikan (base case). Jika tidak, maka fungsi akan memanggil dirinya sendiri dengan bintang(j-1) terlebih dahulu, kemudian mencetak satu karakter "*". Hal ini menyebabkan bintang dicetak dari jumlah kecil ke besar secara berurutan.
#### Kombinasi antara prosedur baris dan bintang menghasilkan pola segitiga bintang yang meningkat setiap barisnya. Setiap baris akan memiliki jumlah bintang sesuai dengan nomor baris tersebut.
#### Setelah semua proses selesai, program akan menampilkan pola bintang ke layar. Sebagai contoh, jika diinputkan n = 5, maka hasilnya adalah:
#### *
#### **
#### ***
#### ****
#### *****

## Unguided 

### 3. Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja yang habis membagi N.
### Masukan terdiri dari sebuah bilangan bulat positif N.
### Keluaran terdiri dari barisan bilangan yang menjadi faktor dari N (terurut dari 1 hingga N ya).

#### soal3.go

```go
package main
import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    faktor(n, 1)
}

func faktor(n, i int) {
    if i > n {
        return
    }
    if n % i == 0 {
        fmt.Print(i, " ")
    }
    faktor(n, i + 1)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 3](https://github.com/frzhwaa/109082500200_ZirlynailaFairuzahwa/blob/main/modul5/output/soal3.png)
#### Program tersebut meminta pengguna untuk menginputkan sebuah bilangan bulat yaitu n. Setelah nilai n dimasukkan, program akan memanggil prosedur faktor(n, 1) untuk mencari dan menampilkan semua faktor dari bilangan tersebut.
#### Prosedur faktor digunakan untuk mengecek setiap bilangan dari i = 1 sampai n apakah merupakan faktor dari n. Parameter n adalah bilangan yang dicari faktornya, sedangkan i adalah nilai pengecekan yang dimulai dari 1.
#### Pada bagian awal, terdapat kondisi if i > n, yang berfungsi sebagai base case (kondisi berhenti). Jika nilai i sudah lebih besar dari n, maka proses rekursif akan dihentikan dengan return.
#### Selanjutnya, program akan mengecek apakah n habis dibagi i menggunakan kondisi n % i == 0. Jika kondisi ini terpenuhi, maka i adalah faktor dari n, sehingga akan dicetak ke layar menggunakan fmt.Print(i, " ").
#### Setelah itu, prosedur akan memanggil dirinya sendiri dengan faktor(n, i + 1) untuk melanjutkan pengecekan ke bilangan berikutnya hingga mencapai batas n.
#### Dengan menggunakan rekursi, program akan mengecek semua kemungkinan faktor secara berurutan dari kecil ke besar, dan mencetak setiap faktor yang memenuhi kondisi.
#### Setelah semua proses selesai, program akan menampilkan seluruh faktor dari n dalam satu baris. Sebagai contoh, jika diinputkan n = 5, maka hasilnya adalah:
#### 1 5

## Unguided 

### 4. Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan tertentu.
### Masukan terdiri dari sebuah bilangan bulat positif N.
### Keluaran terdiri dari barisan bilangan dari N hingga 1 dan kembali ke N.

#### soal4.go

```go
package main
import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    pola(n)
}

func pola(n int) {
    if n == 0 {
        return
    }
    fmt.Print(n, " ")
    pola(n - 1)
    if n != 1 {
        fmt.Print(n, " ")
    }
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 4](https://github.com/frzhwaa/109082500200_ZirlynailaFairuzahwa/blob/main/modul5/output/soal4.png)
#### Program tersebut meminta pengguna untuk menginputkan sebuah bilangan bulat yaitu n. Setelah nilai n dimasukkan, program akan memanggil prosedur pola(n) untuk menampilkan pola bilangan tertentu menggunakan rekursif.
#### Prosedur pola digunakan untuk mencetak pola bilangan dengan cara menurun lalu naik kembali. Prosedur ini memiliki kondisi dasar (base case), yaitu jika n == 0 maka proses dihentikan dengan return.
#### Pada setiap pemanggilan fungsi, program pertama-tama akan mencetak nilai n menggunakan fmt.Print(n, " "). Setelah itu, fungsi akan memanggil dirinya sendiri dengan pola(n - 1) sehingga nilai n akan terus berkurang hingga mencapai 0.
#### Setelah proses rekursif (pemanggilan ke bawah) selesai, program akan melanjutkan ke bagian setelah pemanggilan rekursif. Di sini terdapat kondisi if n != 1, yang berfungsi untuk mencegah angka 1 dicetak dua kali.
#### Jika kondisi tersebut terpenuhi, maka nilai n akan dicetak kembali.Sebagai contoh, jika diinputkan n = 5, maka hasilnya adalah:
#### 5 4 3 2 1 2 3 4 5

## Unguided 

### 5. Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan ganjil.
### Masukan terdiri dari sebuah bilangan bulat positif N.
### Keluaran terdiri dari barisan bilangan ganjil dari 1 hingga N.

#### soal5.go

```go
package main
import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    ganjil(1, n)
}

func ganjil(i, n int) {
    if i > n {
        return
    }
    if i % 2 != 0 {
        fmt.Print(i, " ")
    }
    ganjil(i+1, n)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 5](https://github.com/frzhwaa/109082500200_ZirlynailaFairuzahwa/blob/main/modul5/output/soal5.png)
#### Program tersebut meminta pengguna untuk menginputkan sebuah bilangan bulat yaitu n. Setelah nilai n dimasukkan, program akan memanggil prosedur ganjil(1, n) untuk menampilkan semua bilangan ganjil dari 1 sampai n.
#### Prosedur ganjil digunakan untuk mengecek setiap bilangan dari i = 1 hingga n. Parameter i berfungsi sebagai nilai awal yang akan dicek, sedangkan n adalah batas maksimal.
#### Pada bagian awal terdapat kondisi if i > n, yang merupakan base case (kondisi berhenti). Jika nilai i sudah melebihi n, maka proses rekursi dihentikan dengan return.
#### Selanjutnya, program akan mengecek apakah i adalah bilangan ganjil dengan kondisi i % 2 != 0. Jika kondisi tersebut terpenuhi, maka nilai i akan dicetak ke layar menggunakan fmt.Print(i, " ").
#### Setelah itu, prosedur akan memanggil dirinya sendiri dengan ganjil(i+1, n) untuk melanjutkan pengecekan ke bilangan berikutnya hingga mencapai n.
#### Dengan menggunakan rekursif, program akan menampilkan semua bilangan ganjil secara berurutan dari kecil ke besar.
#### Setelah semua proses selesai, program akan menampilkan hasil dalam satu baris. Sebagai contoh, jika diinputkan n = 5, maka hasilnya adalah:
#### 1 3 5
#### Contoh lain, jika diinputkan n = 20, maka hasilnya adalah:
#### 1 3 5 7 9 11 13 15 17 19

## Unguided 

### 6. Buatlah program yang mengimplementasikan rekursif untuk mencari hasil pangkat dari dua buah bilangan.
### Masukan terdiri dari bilangan bulat x dan y.
### Keluaran terdiri dari hasil x dipangkatkan y.
### Catatan: diperbolehkan menggunakan asterik "*", tapi dilarang menggunakan import "math".

#### soal6.go

```go
package main
import "fmt"

func main() {
    var x, y int
    fmt.Scan(&x, &y)
    fmt.Println(pangkat(x, y))
}

func pangkat(x, y int) int {
    if y == 0 {
        return 1
    }
    return x * pangkat(x, y-1)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 6](https://github.com/frzhwaa/109082500200_ZirlynailaFairuzahwa/blob/main/modul5/output/soal6.png)
#### Program tersebut meminta pengguna untuk menginputkan dua buah bilangan bulat yaitu x dan y. Setelah data dimasukkan, program akan memanggil fungsi pangkat(x, y) untuk menghitung nilai x pangkat y (x^y), kemudian hasilnya akan ditampilkan ke layar.
#### Fungsi pangkat digunakan untuk menghitung perpangkatan secara rekursif. Fungsi ini memiliki kondisi dasar (base case), yaitu jika y == 0 maka hasilnya adalah 1, karena setiap bilangan yang dipangkatkan 0 bernilai 1.
#### Jika nilai y lebih dari 0, maka fungsi akan mengembalikan hasil dari x * pangkat(x, y-1). Artinya, nilai x^y diperoleh dengan mengalikan x dengan hasil perpangkatan sebelumnya, yaitu x^(y-1).
#### Proses rekursif ini akan terus berjalan dengan mengurangi nilai y hingga mencapai 0. Setiap pemanggilan fungsi akan menunggu hasil dari pemanggilan berikutnya sampai semua perhitungan selesai.
#### Setelah semua proses selesai, hasil akhir akan dikembalikan ke fungsi main dan ditampilkan menggunakan fmt.Println. Sebagai contoh, jika diinputkan x = 2 dan y = 2, maka perhitungannya adalah:
#### pangkat(2, 2) = 2 × pangkat(2, 1)
#### pangkat(2, 1) = 2 × pangkat(2, 0)
#### pangkat(2, 0) = 1
#### Sehingga hasil akhirnya adalah 2 × 2 × 1 = 4.
#### Contoh lain, jika diinputkan x = 5 dan y = 3, maka hasilnya adalah 125.