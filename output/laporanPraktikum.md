# <h1 align="center">Laporan Praktikum Modul 2 - REVIEW ALGORITMA DAN PEMROGRAMAN 1 </h1>
<p align="center">ZIRLYNAILA FAIRUZAHWA - 109082500200</p>

## Unguided 

### 1. Telusuri program berikut dengan cara mengkompilasi dan mengeksekusi program. Silakan masukan data yang sesuai sebanyak yang diminta program. Perhatikan keluaran yang diperoleh. Coba terangkan apa sebenarnya yang dilakukan program tersebut?
#### soal1.go

```go
package main
import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp string
	)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1](https://github.com/frzhwaa/109082500200_ZirlynailaFairuzahwa/blob/main/output/output-soal1.png)
Program tersebut meminta pengguna untuk menginputkan tiga buah string. Setelah ketiga data dimasukkan, program akan menampilkan urutan data awal yang telah diinputkan. Selanjutnya program melakukan pertukaran nilai pada ketiga variabel string dengan bantuan variabel temp. Prosesnya dimulai dengan menyimpan nilai variabel pertama ke dalam variabel temp, kemudian nilai variabel kedua dipindahkan ke variabel pertama, nilai variabel ketiga dipindahkan ke variabel kedua, dan terakhir nilai yang ada pada variabel temp dipindahkan ke variabel ketiga. Dengan cara ini urutan data berubah dari satu dua tiga menjadi dua tiga satu.
Sebagai contoh, ketika saya menginputkan nama, saya, dan zahwa, program akan menampilkan output awal yaitu "Output awal = nama saya zahwa". Setelah proses pertukaran nilai dilakukan, urutan data berubah menjadi "Output akhir = saya zahwa nama". Program melakukan pertukaran posisi data string, di mana data pertama berpindah ke posisi terakhir, sedangkan data kedua dan ketiga bergeser satu posisi ke depan.

## Unguided 

### 2. Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakan praktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunan warna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah ‘merah’, ‘kuning’, ‘hijau’, dan ‘ungu’ selama 5 kali percobaan berulang.
### Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan warna sesuai dengan informasi yang diberikan pada paragraf sebelumnya, dan false untuk urutan warna lainnya.
### Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah adalah input/read):
### Percobaan 1: merah kuning hijau ungu
### Percobaan 2: merah kuning hijau ungu
### Percobaan 3: merah kuning hijau ungu
### Percobaan 4: merah kuning hijau ungu
### Percobaan 5: merah kuning hijau ungu
### BERHASIL: true

### Percobaan 1: merah kuning hijau ungu
### Percobaan 2: merah kuning hijau ungu
### Percobaan 3: merah kuning hijau ungu
### Percobaan 4: ungu kuning hijau merah
### Percobaan 5: merah kuning hijau ungu
### BERHASIL: false
#### soal2.go

```go
package main
import "fmt"

func main() {
	var gelas1, gelas2, gelas3, gelas4 string
	berhasil := true
	for i := 1; i <= 5; i++ {
		fmt.Print("Percobaan ke ", i, ": ")
		fmt.Scan(&gelas1, &gelas2, &gelas3, &gelas4)
		if !(gelas1 == "merah" && gelas2 == "kuning" && gelas3 == "hijau" && gelas4 == "ungu") {
			berhasil = false
		}
	}
	fmt.Print("BERHASIL: ", berhasil)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 2](https://github.com/frzhwaa/109082500200_ZirlynailaFairuzahwa/blob/main/output/output-soal2.png)
Program tersebut digunakan untuk memeriksa keberhasilan percobaan praktikum kimia berdasarkan susunan warna cairan pada empat tabung reaksi yang dimasukkan oleh pengguna sebanyak lima kali percobaan. Pada setiap percobaan, pengguna menginputkan empat warna sekaligus yang mewakili isi dari gelas 1 hingga gelas 4. Program kemudian membandingkan urutan warna yang dimasukkan dengan urutan yang telah ditentukan. Pengecekan dilakukan menggunakan kondisi if !(gelas1 == "merah" && gelas2 == "kuning" && gelas3 == "hijau" && gelas4 == "ungu"). Tanda ! berarti tidak, sehingga jika pada salah satu percobaan ada urutan warna yang dimasukkan tidak sama dengan urutan tersebut, maka variabel berhasil akan diubah menjadi false.
Pada saat saya melakukan input di mana setiap percobaan yang diisi dengan urutan merah kuning hijau ungu, maka program menampilkan output "BERHASIL: true" karena semua percobaan sesuai dengan ketentuan. Namun ketika pada salah satu percobaan, misalnya percobaan ke-4, saya memasukkan urutan ungu kuning hijau merah, maka urutan tersebut tidak sesuai dengan yang ditentukan sehingga program menampilkan "BERHASIL: false".

## Unguided 

### 3. PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka, buatlah program BiayaPos untuk menghitung biaya pengiriman tersebut dengan ketentuan sebagai berikut!
# Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam gram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500 gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500 gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat (yang kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.
### Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah adalah input/read):
### Contoh 1
### Berat parsel (gram): 8500
### Detail berat: 8 kg + 500 gr
### Detail biaya: Rp. 80000 + Rp. 2500
### Total biaya: Rp. 82500

### Contoh 2
### Berat parsel (gram): 9250
### Detail berat: 9 kg + 250 gr
### Detail biaya: Rp. 90000 + Rp. 3750
### Total biaya: Rp. 93750

### Contoh 3
### Berat parsel (gram): 11750
### Detail berat: 11 kg + 750 gr
### Detail biaya: Rp. 110000 + Rp. 3750
### Total biaya: Rp. 110000
#### soal3.go

```go
package main
import "fmt"

func main() {
	var berat int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&berat)
	kg := berat / 1000
	gr := berat % 1000
	biayaKg := kg * 10000
	biayaGr := 0
	if gr >= 500 {
		biayaGr = gr * 5
	} else {
		biayaGr = gr * 15
	}
	total := 0
	if kg >= 10 {
		total = biayaKg
	} else {
		total = biayaKg + biayaGr
	}
	fmt.Println("Detail berat: ", kg, " kg + ", gr, " gr")
	fmt.Println("Detail biaya: Rp. ", biayaKg, " + Rp. ", biayaGr)
	fmt.Println("Total biaya: Rp. ", total)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 3](https://github.com/frzhwaa/109082500200_ZirlynailaFairuzahwa/blob/main/output/output-soal3.png)
Program tersebut digunakan untuk menghitung biaya pengiriman parsel di PT POS berdasarkan berat parsel. Program terlebih dahulu meminta pengguna memasukkan berat parsel dalam gram, kemudian menghitungnya dalam kilogram (kg) dan sisanya dalam gram (gr) menggunakan operasi pembagian dan modulus. Setelah itu program menghitung biaya pengiriman di mana setiap 1 kg dikenakan biaya Rp. 10.000. Untuk sisa berat dalam gram, jika sisa berat lebih dari atau sama dengan 500 gram maka biaya tambahannya adalah Rp. 5 per gram, sedangkan jika kurang dari 500 gram maka biaya tambahan adalah Rp. 15 per gram. Namun jika total berat parsel lebih dari 10 kg, maka biaya untuk sisa gram tidak dihitung (gratis) sehingga hanya biaya per kilogram saja yang dibayarkan.
Sebagai contoh saat saya menginputkan 8500 gram, program menampilkan detail berat 8 kg + 500 gr dengan biaya Rp. 80000 + Rp. 2500, sehingga total biaya menjadi Rp. 82500. Ketika saya menginputkan 9250 gram, program menampilkan detail berat 9 kg + 250 gr dengan biaya Rp. 90000 + Rp. 3750, sehingga total biaya menjadi Rp. 93750. Sedangkan ketika saya menginputkan 11750 gram, beratnya menjadi 11 kg + 750 gr, tetapi karena beratnya lebih dari 10 kg, maka biaya sisa gram digratiskan sehingga total biaya yang harus dibayar hanya Rp. 110000.