package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// no 1
	var nama, kursus, tahun, lokasi, waktu, pesan string
	nama = "Sanbercode"
	kursus = "Golang Nextjs"
	tahun = "2024"
	lokasi = "Jakarta"
	waktu = "Pagi"
	pesan = "Super Bootcamp"
	result := fmt.Sprintf("%s %s %s %s %s %s", nama, kursus, tahun, lokasi, waktu, pesan)
	fmt.Println(result)

	// no 2
	halo := "Halo Dunia"
	hasil := strings.Replace(halo, "Dunia", "Golang", 1)
	fmt.Println(hasil)

	// no 3
	var kataPertama, kataKedua, kataKetiga, kataKeempat string
	kataPertama = "saya"
	kataKedua = "senang"
	kataKetiga = "belajar"
	kataKeempat = "golang"

	kataKedua = strings.Title(kataKedua)
	kataKeempat = strings.ToUpper(kataKeempat)
	kataKetiga = strings.Replace(kataKetiga, "belajar", "belajaR", 1)
	result3 := kataPertama + " " + kataKedua + " " + kataKetiga + " " + kataKeempat
	fmt.Println(result3)

	//no4
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"
	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	panjang, _ := strconv.Atoi(panjangPersegiPanjang)
	lebar, _ := strconv.Atoi(lebarPersegiPanjang)
	alas, _ := strconv.Atoi(alasSegitiga)
	tinggi, _ := strconv.Atoi(tinggiSegitiga)

	var luasPersegiPanjang = panjang * lebar
	var kelilingPersegiPanjang = 2 * (panjang + lebar)
	var luasSegitiga = (alas * tinggi) / 2

	fmt.Println("Luas Persegi Panjang:", luasPersegiPanjang)
	fmt.Println("Keliling Persegi Panjang:", kelilingPersegiPanjang)
	fmt.Println("Luas Segitiga:", luasSegitiga)

	//5
	kalimat := "halo halo bandung"
	angka := 2024

	words := strings.Fields(kalimat)

	var modifiedWords []string
	for _, word := range words {
		if word == "halo" {
			modifiedWords = append(modifiedWords, "Hi")
		} else {
			modifiedWords = append(modifiedWords, word)
		}
	}

	modifiedKalimat := strings.Join(modifiedWords, " ")
	output := fmt.Sprintf("%s - %d", modifiedKalimat, angka)
	fmt.Println(output)

	//6
	var sentence = "Saya Sangat Senang Sekali Belajar Programming dan Saya Juga Senang Belajar Javascript"

	length := len(sentence)

	var category string
	switch {
	case length < 10:
		category = "Pendek"
	case length >= 10 && length <= 30:
		category = "Sedang"
	case length > 30:
		category = "Panjang"
	}

	fmt.Printf("Panjang string: %d karakter\nKategori: %s\n", length, category)

	//7
	var nilaiJohn = 80
	var nilaiDoe = 50

	var indexJohn string
	switch {
	case nilaiJohn >= 80:
		indexJohn = "A"
	case nilaiJohn >= 70 && nilaiJohn < 80:
		indexJohn = "B"
	case nilaiJohn >= 60 && nilaiJohn < 70:
		indexJohn = "C"
	case nilaiJohn >= 50 && nilaiJohn < 60:
		indexJohn = "D"
	case nilaiJohn < 50:
		indexJohn = "E"
	}

	var indexDoe string
	switch {
	case nilaiDoe >= 80:
		indexDoe = "A"
	case nilaiDoe >= 70 && nilaiDoe < 80:
		indexDoe = "B"
	case nilaiDoe >= 60 && nilaiDoe < 70:
		indexDoe = "C"
	case nilaiDoe >= 50 && nilaiDoe < 60:
		indexDoe = "D"
	case nilaiDoe < 50:
		indexDoe = "E"
	}

	fmt.Println("Indeks Nilai John:", indexJohn)
	fmt.Println("Indeks Nilai Doe:", indexDoe)

	//8
	var tanggal2 = 17
	var bulan2 = 8
	var tahun2 = 1945

	var namaBulan string
	switch bulan2 {
	case 1:
		namaBulan = "Januari"
	case 2:
		namaBulan = "Februari"
	case 3:
		namaBulan = "Maret"
	case 4:
		namaBulan = "April"
	case 5:
		namaBulan = "Mei"
	case 6:
		namaBulan = "Juni"
	case 7:
		namaBulan = "Juli"
	case 8:
		namaBulan = "Agustus"
	case 9:
		namaBulan = "September"
	case 10:
		namaBulan = "Oktober"
	case 11:
		namaBulan = "November"
	case 12:
		namaBulan = "Desember"
	default:
		namaBulan = "Bulan tidak valid"
	}

	fmt.Printf("%d %s %d\n", tanggal2, namaBulan, tahun2)

	//9
	var tahunKelahiran = 1990

	var generasi string
	if tahunKelahiran >= 1944 && tahunKelahiran <= 1964 {
		generasi = "Baby boomer"
	} else if tahunKelahiran >= 1965 && tahunKelahiran <= 1979 {
		generasi = "Generasi X"
	} else if tahunKelahiran >= 1980 && tahunKelahiran <= 1994 {
		generasi = "Generasi Y (Millenials)"
	} else if tahunKelahiran >= 1995 && tahunKelahiran <= 2015 {
		generasi = "Generasi Z"
	} else {
		generasi = "Tidak termasuk dalam kategori yang disebutkan"
	}

	fmt.Printf("Anda termasuk dalam generasi: %s\n", generasi)

	//10
	var penjualan = 2500000
	var uangJasa int
	var uangKomisi float64

	if penjualan <= 2000000 {
		uangJasa = 100000
		uangKomisi = float64(penjualan) * 0.1
	} else if penjualan > 2000000 && penjualan <= 5000000 {
		uangJasa = 200000
		uangKomisi = float64(penjualan) * 0.15
	} else if penjualan > 5000000 {
		uangJasa = 300000
		uangKomisi = float64(penjualan) * 0.2
	}

	totalPendapatan := uangJasa + int(uangKomisi)

	fmt.Printf("Penjualan: Rp %d\n", penjualan)
	fmt.Printf("Uang Jasa: Rp %d\n", uangJasa)
	fmt.Printf("Uang Komisi: Rp %.2f\n", uangKomisi)
	fmt.Printf("Total Pendapatan: Rp %d\n", totalPendapatan)
}
