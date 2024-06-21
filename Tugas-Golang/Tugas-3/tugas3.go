package main

import (
	"fmt"
	"math"
	"strings"
)

type Phone struct {
	name  string
	brand string
	color string
	year  string
}

func descriptionPhone(name, brand, color, year string) string {
	phone := Phone{
		name:  name,
		brand: brand,
		color: color,
		year:  year,
	}
	return introducePhone(phone)
}

func introducePhone(phone Phone) string {
	return fmt.Sprintf("%s adalah smartphone yang dirilis oleh %s pada tahun %s dan memiliki varian warna %s", phone.name, phone.brand, phone.year, phone.color)
}

func luasPersegiPanjang(panjang, lebar int) int {
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang, lebar int) int {
	return 2 * (panjang + lebar)
}

func volumeBalok(panjang, lebar, tinggi int) int {
	return panjang * lebar * tinggi
}

func buahFavorit(nama string, buah ...string) string {
	buahString := strings.Join(buah, ", ")
	return fmt.Sprintf("halo nama saya %s dan buah favorit saya adalah %s", nama, buahString)
}

var dataBuku = []map[string]string{}

func tambahDataBuku() func(string, string, string, string) {
	return func(title, author, category, year string) {
		buku := map[string]string{
			"title":    title,
			"author":   author,
			"category": category,
			"year":     year,
		}
		dataBuku = append(dataBuku, buku)
	}
}

var luasLingkaran float64
var kelilingLingkaran float64

func hitungLingkaran(radius float64, luas *float64, keliling *float64) {
	*luas = math.Pi * radius * radius
	*keliling = 2 * math.Pi * radius
}

func introduce(sentence *string, name, gender, job, age string) {
	if gender == "laki-laki" {
		*sentence = fmt.Sprintf("Pak %s adalah seorang %s yang berusia %s tahun", name, job, age)
	} else if gender == "perempuan" {
		*sentence = fmt.Sprintf("Bu %s adalah seorang %s yang berusia %s tahun", name, job, age)
	}
}

func tambahBuah(buah *[]string, namaBuah ...string) {
	*buah = append(*buah, namaBuah...)
}

func tambahDataFilm(title, duration, genre, year string, dataFilm *[]map[string]string) {
	film := map[string]string{
		"title":    title,
		"duration": duration,
		"genre":    genre,
		"year":     year,
	}
	*dataFilm = append(*dataFilm, film)
}

func main() {
	fmt.Println("---------Soal 1-----------")
	samsung := descriptionPhone("Samsung Galaxy Note 20", "Samsung", "Bronze", "2021")
	fmt.Println(samsung)

	xiaomi := descriptionPhone("Redmi Note 10 Pro", "Xiaomi", "Black", "2021")
	fmt.Println(xiaomi)

	fmt.Println("---------Soal 2-----------")
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println("Luas Persegi Panjang:", luas)
	fmt.Println("Keliling Persegi Panjang:", keliling)
	fmt.Println("Volume Balok:", volume)

	fmt.Println("---------Soal 3-----------")
	var buahArray = []string{"semangka", "jeruk", "melon", "pepaya"}
	buahFavoritJohn := buahFavorit("John", buahArray...)
	fmt.Println(buahFavoritJohn)

	fmt.Println("---------Soal 4-----------")
	tambahData := tambahDataBuku()

	tambahData("Harry Potter", "J.K Rowling", "Novel", "1997")
	tambahData("Dracula", "Bram Stoker", "Novel", "2019")
	tambahData("Algoritma Dan Pemrograman", "Rinaldi Munnir", "Pelajaran", "2010")
	tambahData("Matematika Diskrit", "Rinaldi Munir", "Pelajaran", "2010")

	for _, buku := range dataBuku {
		fmt.Println(buku)
	}

	fmt.Println("---------Soal 5-----------")
	radius := 7.0
	hitungLingkaran(radius, &luasLingkaran, &kelilingLingkaran)
	fmt.Printf("Dengan jari-jari %.2f, luas lingkaran adalah %.2f dan keliling lingkaran adalah %.2f\n", radius, luasLingkaran, kelilingLingkaran)

	fmt.Println("---------Soal 6-----------")
	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence)

	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence)

	fmt.Println("---------Soal 7-----------")
	var buah = []string{}
	tambahBuah(&buah, "Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat")

	for i, v := range buah {
		fmt.Printf("%d. %s\n", i+1, v)
	}

	fmt.Println("---------Soal 8-----------")
	var dataFilm = []map[string]string{}

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	for i, film := range dataFilm {
		fmt.Printf("%d. title: %s \nduration: %s \ngenre: %s \nyear: %s\n", i+1, film["title"], film["duration"], film["genre"], film["year"])
	}
}
