package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func main() {
	println("---------Soal 1-----------")
	sortPhones()

	println("\n---------Soal 2-----------")
	var movies = []string{"Harry Potter", "LOTR", "Spiderman", "Logan", "Avengers", "Insidious", "Toy Story"}

	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies...)

	printMovies(moviesChannel)

	println("\n---------Soal 3-----------")
	calculateShapes()

	println("\n---------Soal 4-----------")
	calculateRectangleAndCuboid()
}

func sortPhones() {
	phones := []string{"Xiaomi", "Asus", "IPhone", "Samsung", "Oppo", "Realme", "Vivo"}

	var wg sync.WaitGroup
	wg.Add(len(phones))

	for i, phone := range phones {
		go func(i int, phone string) {
			defer wg.Done()
			time.Sleep(time.Duration(i) * time.Second)
			fmt.Printf("%d. %s\n", i+1, phone)
		}(i, phone)
	}
	wg.Wait()
}

func getMovies(moviesChannel chan string, movies ...string) {
	for _, movie := range movies {
		moviesChannel <- movie
	}
	close(moviesChannel)
}

func printMovies(moviesChannel chan string) {
	i := 1
	for movie := range moviesChannel {
		fmt.Printf("%d. %s\n", i, movie)
		i++
	}
}

func calculateShapes() {

	jariJari := []float64{8, 14, 20}
	tinggi := 10

	resultChannel := make(chan string)

	var wg sync.WaitGroup

	for _, jari := range jariJari {
		wg.Add(2)
		go func(jari float64) {
			defer wg.Done()
			area := math.Pi * math.Pow(jari, 2)
			circumFerence := 2 * math.Pi * jari
			resultChannel <- fmt.Sprintf("Jari-jari: %.1f, Luas: %.2f, Keliling: %.2f", jari, area, circumFerence)
		}(jari)

		go func(jari float64, tinggi int) {
			defer wg.Done()
			volume := math.Pi * math.Pow(jari, 2) * float64(tinggi)
			resultChannel <- fmt.Sprintf("Jari-jari: %.1f, Tinggi: %d, Volume: %.2f", jari, tinggi, volume)
		}(jari, tinggi)
	}

	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	for result := range resultChannel {
		fmt.Println(result)
	}
}

func calculateRectangleAndCuboid() {
	panjang := 10.0
	lebar := 5.0
	tinggi := 8.0

	areaChannel := make(chan string)
	perimeterChannel := make(chan string)
	volumeChannel := make(chan string)
	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()
		area := panjang * lebar
		areaChannel <- fmt.Sprintf("Luas Persegi Panjang: %.2f", area)
	}()

	go func() {
		defer wg.Done()
		perimeter := 2 * (panjang + lebar)
		perimeterChannel <- fmt.Sprintf("Keliling Persegi Panjang: %.2f", perimeter)
	}()

	go func() {
		defer wg.Done()
		volume := panjang * lebar * tinggi
		volumeChannel <- fmt.Sprintf("Volume Balok: %.2f", volume)
	}()

	go func() {
		wg.Wait()
		close(areaChannel)
		close(perimeterChannel)
		close(volumeChannel)
	}()

	for {
		select {
		case result, ok := <-areaChannel:
			if !ok {
				areaChannel = nil
			} else {
				fmt.Println(result)
			}
		case result, ok := <-perimeterChannel:
			if !ok {
				perimeterChannel = nil
			} else {
				fmt.Println(result)
			}
		case result, ok := <-volumeChannel:
			if !ok {
				volumeChannel = nil
			} else {
				fmt.Println(result)
			}
		default:
			return
		}
	}
}
