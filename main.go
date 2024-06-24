package main

import (
	"fmt"
	tugas7 "my-app/Tugas-Golang/Tugas-7"
)

func main() {
	fmt.Println("---------Soal 1-----------")

	var segitiga tugas7.SegitigaSamaSisi = tugas7.SegitigaSamaSisi{Alas: 5, Tinggi: 4}
	fmt.Println("Luas segitiga:", segitiga.Luas())
	fmt.Println("Keliling segitiga:", segitiga.Keliling())

	var persegi tugas7.PersegiPanjang = tugas7.PersegiPanjang{Panjang: 6, Lebar: 3}
	fmt.Println("Luas persegi panjang:", persegi.Luas())
	fmt.Println("Keliling persegi panjang:", persegi.Keliling())

	var silinder tugas7.Tabung = tugas7.Tabung{JariJari: 2.5, Tinggi: 7.0}
	fmt.Println("Volume tabung:", silinder.Volume())
	fmt.Println("Luas permukaan tabung:", silinder.LuasPermukaan())

	var kotak tugas7.Balok = tugas7.Balok{Panjang: 3, Lebar: 4, Tinggi: 5}
	fmt.Println("Volume balok:", kotak.Volume())
	fmt.Println("Luas permukaan balok:", kotak.LuasPermukaan())

	var hitungDatar tugas7.HitungBangunDatar = segitiga
	fmt.Println("Luas segitiga via interface:", hitungDatar.Luas())
	fmt.Println("Keliling segitiga via interface:", hitungDatar.Keliling())

	var hitungRuang tugas7.HitungBangunRuang = silinder
	fmt.Println("Volume tabung via interface:", hitungRuang.Volume())
	fmt.Println("Luas permukaan tabung via interface:", hitungRuang.LuasPermukaan())

	fmt.Println("\n---------Soal 2-----------")

	myPhone := tugas7.Phone{
		Name:   "Samsung Galaxy Note 20",
		Brand:  "Samsung",
		Year:   2020,
		Colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}

	fmt.Println(myPhone.DisplayData())

	var myPhoneInterface tugas7.PhoneInterface = myPhone
	fmt.Println("Data phone via interface:")
	fmt.Println(myPhoneInterface.DisplayData())

	fmt.Println("---------Soal 3-----------")
	fmt.Println(tugas7.LuasPersegi(4, true))
	fmt.Println(tugas7.LuasPersegi(8, false))
	fmt.Println(tugas7.LuasPersegi(0, true))
	fmt.Println(tugas7.LuasPersegi(0, false))

	fmt.Println("\n---------Soal 4-----------")

	prefixString := tugas7.Prefix.(string)
	angkaPertama := tugas7.KumpulanAngkaPertama.([]int)
	angkaKedua := tugas7.KumpulanAngkaKedua.([]int)

	total := 0
	for _, angka := range angkaPertama {
		total += angka
	}
	for _, angka := range angkaKedua {
		total += angka
	}

	fmt.Printf("%s%d + %d + %d + %d = %d\n", prefixString, angkaPertama[0], angkaPertama[1], angkaKedua[0], angkaKedua[1], total)

	fmt.Println("\n---------Soal 5-----------")

	var people []tugas7.Sentence

	people = append(people,
		tugas7.Person{Name: "John", Job: "Programmer", Gender: "male", Age: 30},
		tugas7.Person{Name: "Sarah", Job: "Model", Gender: "female", Age: 27},
		tugas7.Person{Name: "Jack", Job: "Engineer", Gender: "male", Age: 25},
		tugas7.Person{Name: "Ellie", Job: "Designer", Gender: "female", Age: 35},
		tugas7.Person{Name: "Danny", Job: "Footballer", Gender: "male", Age: 31},
	)

	for _, p := range people {
		fmt.Println(p.Introduction())
	}
}
