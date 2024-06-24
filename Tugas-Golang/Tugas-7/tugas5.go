package tugas7

import (
	"fmt"
	"math"
	"strings"
)

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

func (s SegitigaSamaSisi) Luas() int {
	return (s.Alas * s.Tinggi) / 2
}

func (s SegitigaSamaSisi) Keliling() int {
	return 3 * s.Alas
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

func (p PersegiPanjang) Luas() int {
	return p.Panjang * p.Lebar
}

func (p PersegiPanjang) Keliling() int {
	return 2 * (p.Panjang + p.Lebar)
}

type Tabung struct {
	JariJari, Tinggi float64
}

func (t Tabung) Volume() float64 {
	return math.Pi * t.JariJari * t.JariJari * t.Tinggi
}

func (t Tabung) LuasPermukaan() float64 {
	return 2 * math.Pi * t.JariJari * (t.JariJari + t.Tinggi)
}

type Balok struct {
	Panjang, Lebar, Tinggi int
}

func (b Balok) Volume() float64 {
	return float64(b.Panjang * b.Lebar * b.Tinggi)
}

func (b Balok) LuasPermukaan() float64 {
	return 2 * float64(b.Panjang*b.Lebar+b.Panjang*b.Tinggi+b.Lebar*b.Tinggi)
}

type Phone struct {
	Name   string
	Brand  string
	Year   int
	Colors []string
}

func (p Phone) DisplayData() string {
	colorsStr := strings.Join(p.Colors, ", ")
	return fmt.Sprintf("Name: %s\nBrand: %s\nYear: %d\nColors: %s", p.Name, p.Brand, p.Year, colorsStr)
}

type PhoneInterface interface {
	DisplayData() string
}

func LuasPersegi(sisi int, showText bool) interface{} {
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

var Prefix interface{} = "hasil penjumlahan dari "
var KumpulanAngkaPertama interface{} = []int{6, 8}
var KumpulanAngkaKedua interface{} = []int{12, 14}

type Person struct {
	Name   string
	Job    string
	Gender string
	Age    int
}

type Sentence interface {
	Introduction() string
}

var People []Sentence

func (p Person) Introduction() string {
	var honorific string
	if p.Gender == "male" {
		honorific = "Pak"
	} else {
		honorific = "Bu"
	}
	return fmt.Sprintf("%s %s adalah %s yang berusia %d tahun", honorific, p.Name, p.Job, p.Age)
}
