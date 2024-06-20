package main

import (
	"fmt"
)

func main() {
	// no 1
	fmt.Println("LOOPING PERTAMA")
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Printf("%d - I love coding\n", i)
		}
	}

	fmt.Println("LOOPING KEDUA")
	for i := 20; i >= 1; i-- {
		if i%2 == 0 {
			fmt.Printf("%d - I will become a fullstack developer\n", i)
		}
	}

	// no 2
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			if i%3 == 0 {
				fmt.Printf("%d - Super Bootcamp\n", i)
			} else {
				fmt.Printf("%d - Berkualitas\n", i)
			}
		} else {
			if i%3 == 0 {
				fmt.Printf("%d - I Love Coding\n", i)
			} else {
				fmt.Printf("%d - Santai\n", i)
			}
		}
	}

	// no 3
	for i := 1; i <= 7; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}

	// no 4
	for i := 1; i <= 7; i++ {
		for j := 7; j > i; j-- {
			fmt.Print(" ")
		}
		for k := 1; k <= i; k++ {
			fmt.Print("#")
		}
		fmt.Println()
	}

	// no 5
	kalimat := [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}

	hasil := kalimat[2:]

	output := fmt.Sprintf("[%s]", fmt.Sprint(hasil))

	fmt.Println(output)

	// no 6
	var sayuran = []string{}

	sayuran = append(sayuran, "Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun")

	for i, v := range sayuran {
		fmt.Printf("%d. %s\n", i+1, v)
	}

	// no 7
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}

	for key, value := range satuan {
		fmt.Printf("%s = %d\n", key, value)
	}

	volume := satuan["panjang"] * satuan["lebar"] * satuan["tinggi"]

	fmt.Printf("volume balok = %d\n", volume)

	// no 8
	var word = "car"

	var combinations []string

	for i := 0; i < len(word); i++ {
		for j := i + 1; j <= len(word); j++ {
			combinations = append(combinations, word[i:j])
		}
	}

	fmt.Println(combinations)

	// no 9
	kumpulanAngkaBerurut := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, angka := range kumpulanAngkaBerurut {
		if angka%2 == 0 {
			fmt.Println(angka)
		}
	}

	// no 10
	kumpulanAngkaDuaDimensi := [][]int{
		{1, 3, 5, 7, 8, 9},
		{4, 5, 6, 2, 3, 1},
		{6, 7, 8, 1, 3, 5},
	}

	var hasilJumlah []int

	for _, baris := range kumpulanAngkaDuaDimensi {
		sum := 0
		for _, angka := range baris {
			sum += angka
		}
		hasilJumlah = append(hasilJumlah, sum)
	}

	fmt.Println(hasilJumlah)

}
