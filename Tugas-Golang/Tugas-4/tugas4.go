package main

import (
	"fmt"
)

type Buah struct {
	nama       string
	warna      string
	adaBijinya bool
	harga      int
}

type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

func (s segitiga) luas() float64 {
	return 0.5 * float64(s.alas) * float64(s.tinggi)
}

func (p persegi) luas() int {
	return p.sisi * p.sisi
}

func (pp persegiPanjang) luas() int {
	return pp.panjang * pp.lebar
}

type phone struct {
	name, brand string
	year        int
	colors      []string
}

func (p *phone) addColors(colors ...string) {
	p.colors = append(p.colors, colors...)
}

type movie struct {
	title, genre   string
	duration, year int
}

var dataFilm = []movie{}

func tambahDataFilm(title string, duration int, genre string, year int, dataFilm *[]movie) {
	m := movie{
		title:    title,
		duration: duration,
		genre:    genre,
		year:     year,
	}
	*dataFilm = append(*dataFilm, m)
}

type person struct {
	name, job, gender string
	age               int
}

func main() {

	nanas := Buah{
		nama:       "Nanas",
		warna:      "Kuning",
		adaBijinya: false,
		harga:      9000,
	}

	jeruk := Buah{
		nama:       "Jeruk",
		warna:      "Oranye",
		adaBijinya: true,
		harga:      8000,
	}

	semangka := Buah{
		nama:       "Semangka",
		warna:      "Hijau & Merah",
		adaBijinya: true,
		harga:      10000,
	}

	pisang := Buah{
		nama:       "Pisang",
		warna:      "Kuning",
		adaBijinya: false,
		harga:      5000,
	}

	fmt.Println("---------Soal 1-----------")
	fmt.Println(nanas)
	fmt.Println(jeruk)
	fmt.Println(semangka)
	fmt.Println(pisang)

	fmt.Println("---------Soal 2-----------")

	s := segitiga{
		alas:   10,
		tinggi: 5,
	}

	p := persegi{
		sisi: 4,
	}

	pp := persegiPanjang{
		panjang: 6,
		lebar:   3,
	}

	fmt.Printf("Luas segitiga: %.2f\n", s.luas())
	fmt.Printf("Luas persegi: %d\n", p.luas())
	fmt.Printf("Luas persegi panjang: %d\n", pp.luas())

	fmt.Println("---------Soal 3-----------")

	samsung := phone{name: "Samsung Galaxy Note 20", brand: "Samsung", year: 2020}

	fmt.Println(samsung)
	samsung.addColors("Black", "Bronze", "Silver")
	fmt.Println(samsung)

	fmt.Println("---------Soal 4-----------")

	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	for i, film := range dataFilm {
		fmt.Printf("%d. title: %s\n", i+1, film.title)
		fmt.Printf("   duration: %d jam\n", film.duration)
		fmt.Printf("   genre: %s\n", film.genre)
		fmt.Printf("   year: %d\n", film.year)
	}

	fmt.Println("---------Soal 5-----------")

	// Menambahkan data people
	people := []person{
		{name: "John", job: "Programmer", gender: "male", age: 30},
		{name: "Sarah", job: "Model", gender: "female", age: 27},
		{name: "Jack", job: "Engineer", gender: "male", age: 25},
		{name: "Ellie", job: "Designer", gender: "female", age: 35},
		{name: "Danny", job: "Footballer", gender: "male", age: 31},
	}

	// Print keseluruhan data people
	for i, p := range people {
		fmt.Printf("%d. Name: %s\n", i+1, p.name)
		fmt.Printf("   Job: %s\n", p.job)
		fmt.Printf("   Gender: %s\n", p.gender)
		fmt.Printf("   Age: %d\n", p.age)
	}

	fmt.Println("---------Soal 6-----------")
	count := 0
	for _, p := range people {
		if p.age > 29 {
			count++
			fmt.Printf("%d. %s\n", count, p.name)
		}
	}

	fmt.Println("---------Soal 7-----------")
	countFemales := 0
	for _, p := range people {
		if p.gender == "female" {
			countFemales++
			fmt.Printf("%d. %s\n", countFemales, p.name)
		}
	}
}
