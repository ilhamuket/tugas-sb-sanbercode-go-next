package main

import (
	"fmt"
	"math"
	"strings"
)

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

type segitigaSamaSisi struct {
	alas, tinggi int
}

func (s segitigaSamaSisi) luas() int {
	return (s.alas * s.tinggi) / 2
}

func (s segitigaSamaSisi) keliling() int {
	return 3 * s.alas
}

type persegiPanjang struct {
	panjang, lebar int
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

type tabung struct {
	jariJari, tinggi float64
}

func (t tabung) volume() float64 {
	return math.Pi * t.jariJari * t.jariJari * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return 2 * math.Pi * t.jariJari * (t.jariJari + t.tinggi)
}

type balok struct {
	panjang, lebar, tinggi int
}

func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return 2 * float64(b.panjang*b.lebar+b.panjang*b.tinggi+b.lebar*b.tinggi)
}

type phone struct {
	name   string
	brand  string
	year   int
	colors []string
}

func (p phone) displayData() string {
	colorsStr := strings.Join(p.colors, ", ")
	return fmt.Sprintf("name: %s\nbrand: %s\nyear: %d\ncolors: %s", p.name, p.brand, p.year, colorsStr)
}

type phoneInterface interface {
	displayData() string
}

func luasPersegi(sisi int, showText bool) interface{} {
	if sisi != 0 {
		luas := sisi * sisi
		if showText {
			return fmt.Sprintf("luas persegi dengan sisi %d cm adalah %d cm", sisi, luas)
		} else {
			return luas
		}
	} else {
		if showText {
			return "Maaf anda belum menginput sisi dari persegi"
		} else {
			return nil
		}
	}
}

var prefix interface{} = "hasil penjumlahan dari "
var kumpulanAngkaPertama interface{} = []int{6, 8}
var kumpulanAngkaKedua interface{} = []int{12, 14}

type person struct {
	name   string
	job    string
	gender string
	age    int
}

type sentence interface {
	introduction() string
}

var people []sentence

func (p person) introduction() string {
	var honorific string
	if p.gender == "male" {
		honorific = "Pak"
	} else {
		honorific = "Bu"
	}
	return fmt.Sprintf("%d. %s %s adalah %s yang berusia %d tahun", len(people)+1, honorific, p.name, p.job, p.age)
}

func main() {
	fmt.Println("---------Soal 1-----------")

	var segitiga segitigaSamaSisi = segitigaSamaSisi{alas: 5, tinggi: 4}
	fmt.Println("Luas segitiga:", segitiga.luas())
	fmt.Println("Keliling segitiga:", segitiga.keliling())

	var persegi persegiPanjang = persegiPanjang{panjang: 6, lebar: 3}
	fmt.Println("Luas persegi panjang:", persegi.luas())
	fmt.Println("Keliling persegi panjang:", persegi.keliling())

	var silinder tabung = tabung{jariJari: 2.5, tinggi: 7.0}
	fmt.Println("Volume tabung:", silinder.volume())
	fmt.Println("Luas permukaan tabung:", silinder.luasPermukaan())

	var kotak balok = balok{panjang: 3, lebar: 4, tinggi: 5}
	fmt.Println("Volume balok:", kotak.volume())
	fmt.Println("Luas permukaan balok:", kotak.luasPermukaan())

	var hitungDatar hitungBangunDatar = segitiga
	fmt.Println("Luas segitiga via interface:", hitungDatar.luas())
	fmt.Println("Keliling segitiga via interface:", hitungDatar.keliling())

	var hitungRuang hitungBangunRuang = silinder
	fmt.Println("Volume tabung via interface:", hitungRuang.volume())
	fmt.Println("Luas permukaan tabung via interface:", hitungRuang.luasPermukaan())

	fmt.Println("\n---------Soal 2-----------")

	myPhone := phone{
		name:   "Samsung Galaxy Note 20",
		brand:  "Samsung",
		year:   2020,
		colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}

	fmt.Println(myPhone.displayData())

	var myPhoneInterface phoneInterface = myPhone
	fmt.Println("Data phone via interface:")
	fmt.Println(myPhoneInterface.displayData())

	fmt.Println("---------Soal 3-----------")
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))

	fmt.Println("\n---------Soal 4-----------")

	prefixString := prefix.(string)
	angkaPertama := kumpulanAngkaPertama.([]int)
	angkaKedua := kumpulanAngkaKedua.([]int)

	total := 0
	for _, angka := range angkaPertama {
		total += angka
	}
	for _, angka := range angkaKedua {
		total += angka
	}

	fmt.Printf("%s%d + %d + %d + %d = %d\n", prefixString, angkaPertama[0], angkaPertama[1], angkaKedua[0], angkaKedua[1], total)

	fmt.Println("\n---------Soal 5-----------")

	people = append(people,
		person{name: "John", job: "Programmer", gender: "male", age: 30},
		person{name: "Sarah", job: "Model", gender: "female", age: 27},
		person{name: "Jack", job: "Engineer", gender: "male", age: 25},
		person{name: "Ellie", job: "Designer", gender: "female", age: 35},
		person{name: "Danny", job: "Footballer", gender: "male", age: 31},
	)

	for _, p := range people {
		fmt.Println(p.introduction())
	}
}
