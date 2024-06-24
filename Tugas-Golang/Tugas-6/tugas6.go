package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"sort"
	"time"
)

var phones = []string{}

func main() {
	fmt.Println("---------Soal 1-----------")
	printSentence("Golang Backend Development", 2021)

	fmt.Println("---------Soal 2-----------")
	fmt.Println(kelilingSegitigaSamaSisi(4, true))
	fmt.Println(kelilingSegitigaSamaSisi(8, false))
	fmt.Println(kelilingSegitigaSamaSisi(0, true))
	fmt.Println(kelilingSegitigaSamaSisi(0, false))

	fmt.Println("---------Soal 3-----------")
	angka := 1
	defer cetakAngka(&angka)

	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)

	fmt.Println("---------Soal 4-----------")
	tambahDataPhones(&phones)
	sort.Strings(phones)

	for i, phone := range phones {
		time.Sleep(time.Second)
		fmt.Printf("%d. %s\n", i+1, phone)
	}

	fmt.Println("---------Soal 5-----------")
	fmt.Println("Jari-jari 7 cm:")
	luas, keliling := hitungLingkaran(7)
	fmt.Printf("Luas lingkaran: %.2f\n", luas)
	fmt.Printf("Keliling lingkaran: %.2f\n", keliling)

	fmt.Println("\nJari-jari 10 cm:")
	luas, keliling = hitungLingkaran(10)
	fmt.Printf("Luas lingkaran: %.2f\n", luas)
	fmt.Printf("Keliling lingkaran: %.2f\n", keliling)

	fmt.Println("\nJari-jari 15 cm:")
	luas, keliling = hitungLingkaran(15)
	fmt.Printf("Luas lingkaran: %.2f\n", luas)
	fmt.Printf("Keliling lingkaran: %.2f\n", keliling)

	fmt.Println("---------Soal 6-----------")
	// Menggunakan package flag untuk mengambil panjang dan lebar dari persegi panjang
	panjangPtr := flag.Int("panjang", 0, "panjang dari persegi panjang")
	lebarPtr := flag.Int("lebar", 0, "lebar dari persegi panjang")
	flag.Parse()

	if *panjangPtr == 0 || *lebarPtr == 0 {
		fmt.Println("Maaf, panjang dan lebar harus diisi")
		return
	}

	luasPersegi := luasPersegiPanjang(*panjangPtr, *lebarPtr)
	kelilingPersegi := kelilingPersegiPanjang(*panjangPtr, *lebarPtr)

	fmt.Printf("Luas persegi panjang dengan panjang %d dan lebar %d adalah %d\n", *panjangPtr, *lebarPtr, luasPersegi)
	fmt.Printf("Keliling persegi panjang dengan panjang %d dan lebar %d adalah %d\n", *panjangPtr, *lebarPtr, kelilingPersegi)
}

func printSentence(kalimat string, tahun int) {
	defer func() {
		fmt.Printf("%s %d\n", kalimat, tahun)
	}()
}

func kelilingSegitigaSamaSisi(sisi int, showText bool) (string, error) {
	if sisi != 0 {
		keliling := 3 * sisi
		if showText {
			return fmt.Sprintf("keliling segitiga sama sisinya dengan sisi %d cm adalah %d cm", sisi, keliling), nil
		} else {
			return fmt.Sprintf("%d", keliling), nil
		}
	} else {
		if showText {
			return "", errors.New("maaf anda belum menginput sisi dari segitiga sama sisi")
		} else {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered from panic:", r)
				}
			}()
			panic(errors.New("maaf anda belum menginput sisi dari segitiga sama sisi"))
		}
	}
}

func tambahAngka(nilai int, total *int) {
	*total += nilai
}

func cetakAngka(total *int) {
	fmt.Println("Total angka adalah:", *total)
}

func tambahDataPhones(phones *[]string) {
	*phones = append(*phones,
		"Xiaomi",
		"Asus",
		"IPhone",
		"Samsung",
		"Oppo",
		"Realme",
		"Vivo",
	)
}

func hitungLingkaran(jariJari float64) (float64, float64) {
	luas := math.Pi * math.Pow(jariJari, 2)
	keliling := 2 * math.Pi * jariJari
	return math.Round(luas*100) / 100, math.Round(keliling*100) / 100
}

func luasPersegiPanjang(panjang, lebar int) int {
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang, lebar int) int {
	return 2 * (panjang + lebar)
}
