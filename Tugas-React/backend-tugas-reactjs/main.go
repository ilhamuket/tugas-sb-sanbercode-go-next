package main

import (
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/config"
	_ "tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/docs"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/routes"
)

// @title Book API
// @version 1.0
// @description This is a sample server for a book store.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

// @securityDefinitions.basic BasicAuth

func main() {
	// Initialize database connection and auto migrate
	config.InitDB()

	// Setup router and start the server
	r := routes.SetupRouter()
	err := r.Run(":8080")
	if err != nil {
		return
	} // Menggunakan port 8080 untuk menjalankan server
}
