package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type NilaiMahasiswa struct {
	ID          uint   `json:"id"`
	Nama        string `json:"nama"`
	MataKuliah  string `json:"mata_kuliah"`
	IndeksNilai string `json:"indeks_nilai"`
	Nilai       uint   `json:"nilai"`
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}

func main() {
	fmt.Println("Starting server on :8080...")
	http.HandleFunc("/nilai-mahasiswa", handleNilaiMahasiswa)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetNilaiMahasiswa(w, r)
	case http.MethodPost:
		handleAddNilaiMahasiswa(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleAddNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {

	contentType := r.Header.Get("Content-Type")
	log.Printf("Received request with Content-Type: %s", contentType)

	username, password, ok := r.BasicAuth()
	if !ok || username != "admin" || password != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorized access")
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	switch contentType {
	case "application/json":

		var input NilaiMahasiswa
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}
		processNilaiMahasiswa(w, input)

	case "application/x-www-form-urlencoded":

		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Invalid form data", http.StatusBadRequest)
			return
		}

		nama := r.Form.Get("nama")
		mataKuliah := r.Form.Get("mata_kuliah")
		nilaiStr := r.Form.Get("nilai")

		nilaiInt, err := strconv.Atoi(nilaiStr)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid nilai dari form data: %s", nilaiStr), http.StatusBadRequest)
			return
		}

		nilai := uint(nilaiInt)

		processNilaiMahasiswa(w, NilaiMahasiswa{
			Nama:       nama,
			MataKuliah: mataKuliah,
			Nilai:      nilai,
		})

	case "multipart/form-data":

		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			http.Error(w, "Invalid multipart form data", http.StatusBadRequest)
			return
		}

		nama := r.FormValue("nama")
		mataKuliah := r.FormValue("mata_kuliah")
		nilaiStr := r.FormValue("nilai")

		nilaiInt, err := strconv.Atoi(nilaiStr)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid nilai dari form data: %s", nilaiStr), http.StatusBadRequest)
			return
		}

		nilai := uint(nilaiInt)

		processNilaiMahasiswa(w, NilaiMahasiswa{
			Nama:       nama,
			MataKuliah: mataKuliah,
			Nilai:      nilai,
		})

	default:
		http.Error(w, "Unsupported Content-Type", http.StatusUnsupportedMediaType)
	}
}

func handleGetNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {

	username, password, ok := r.BasicAuth()
	if !ok || username != "admin" || password != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorized access")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nilaiNilaiMahasiswa)
}

func processNilaiMahasiswa(w http.ResponseWriter, nilai NilaiMahasiswa) {
	if nilai.Nama == "" {
		http.Error(w, "Nama harus diisi", http.StatusBadRequest)
		return
	}
	if nilai.MataKuliah == "" {
		http.Error(w, "Mata Kuliah harus diisi", http.StatusBadRequest)
		return
	}
	if nilai.Nilai < 1 || nilai.Nilai > 100 {
		http.Error(w, "Nilai harus berada di antara 1 dan 100", http.StatusBadRequest)
		return
	}

	var indeksNilai string
	switch {
	case nilai.Nilai >= 80:
		indeksNilai = "A"
	case nilai.Nilai >= 70 && nilai.Nilai < 80:
		indeksNilai = "B"
	case nilai.Nilai >= 60 && nilai.Nilai < 70:
		indeksNilai = "C"
	case nilai.Nilai >= 50 && nilai.Nilai < 60:
		indeksNilai = "D"
	default:
		indeksNilai = "E"
	}

	id := uint(len(nilaiNilaiMahasiswa) + 1)

	nilaiMahasiswa := NilaiMahasiswa{
		ID:          id,
		Nama:        nilai.Nama,
		MataKuliah:  nilai.MataKuliah,
		IndeksNilai: indeksNilai,
		Nilai:       nilai.Nilai,
	}

	nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, nilaiMahasiswa)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(nilaiMahasiswa)
}
