package main

import (
	"log"
	"net/http"

	"my-app/Tugas-Golang/Tugas-13/auth" // import package auth
	"my-app/Tugas-Golang/Tugas-13/db"
	"my-app/Tugas-Golang/Tugas-13/handlers"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// Initialize database connection
	db.InitDB()
	defer func() {
		err := db.DB.Close()
		if err != nil {
			log.Fatalf("Error closing database connection: %v", err)
		}
	}()

	router := httprouter.New()

	// Middleware BasicAuth untuk semua rute
	router.GET("/mahasiswa", auth.BasicAuth(handlers.GetMahasiswas))
	router.GET("/mahasiswa/:id", auth.BasicAuth(handlers.GetMahasiswa))
	router.POST("/mahasiswa", auth.BasicAuth(handlers.CreateMahasiswa))
	router.PUT("/mahasiswa/:id", auth.BasicAuth(handlers.UpdateMahasiswa))
	router.DELETE("/mahasiswa/:id", auth.BasicAuth(handlers.DeleteMahasiswa))

	router.GET("/nilai", auth.BasicAuth(handlers.GetNilais))
	router.GET("/nilai/:id", auth.BasicAuth(handlers.GetNilai))
	router.POST("/nilai", auth.BasicAuth(handlers.CreateNilai))
	router.PUT("/nilai/:id", auth.BasicAuth(handlers.UpdateNilai))
	router.DELETE("/nilai/:id", auth.BasicAuth(handlers.DeleteNilai))

	router.GET("/mata-kuliah", auth.BasicAuth(handlers.GetMataKuliahs))
	router.GET("/mata-kuliah/:id", auth.BasicAuth(handlers.GetMataKuliah))
	router.POST("/mata-kuliah", auth.BasicAuth(handlers.CreateMataKuliah))
	router.PUT("/mata-kuliah/:id", auth.BasicAuth(handlers.UpdateMataKuliah))
	router.DELETE("/mata-kuliah/:id", auth.BasicAuth(handlers.DeleteMataKuliah))

	// Start HTTP server
	log.Fatal(http.ListenAndServe(":8080", router))
}
