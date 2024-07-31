package main

import (
	"log"
	"net/http"
	"os"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/config"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/controllers"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/docs"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/repositories"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var app *gin.Engine

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Initialize Gin engine
	app = gin.Default()

	// Setup CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Content-Type", "Accept", "Origin"}

	// Enable CORS with the configured settings
	app.Use(cors.New(corsConfig))

	// Initialize database
	config.InitDB()

	// Configure Swagger documentation
	docs.SwaggerInfo.Title = "Book REST API"
	docs.SwaggerInfo.Description = "This is a REST API for managing books."
	docs.SwaggerInfo.Version = "1.0"

	environment := os.Getenv("ENVIRONMENT")
	if environment == "development" {
		docs.SwaggerInfo.Host = "localhost:8080"
		docs.SwaggerInfo.Schemes = []string{"http"}
	} else {
		docs.SwaggerInfo.Host = os.Getenv("VERCEL_URL")
		docs.SwaggerInfo.Schemes = []string{"https"}
	}

	// Setup routes
	setupRouter()
}

func setupRouter() {
	// Middleware to set db in context
	app.Use(func(c *gin.Context) {
		c.Set("db", config.GetDB())
		c.Next()
	})

	// Routes for books
	app.POST("/books", controllers.CreateBook)
	app.GET("/books", controllers.GetBooks)
	app.GET("/books/:id", controllers.GetBook)
	app.PATCH("/books/:id", controllers.UpdateBook)
	app.DELETE("/books/:id", controllers.DeleteBook)

	// Initialize repositories
	dosenRepository := repositories.NewDosenRepository(config.GetDB())
	mataKuliahRepository := repositories.NewMataKuliahRepository(config.GetDB())
	mahasiswaRepository := repositories.NewMahasiswaRepository(config.GetDB())
	nilaiRepository := repositories.NewNilaiRepository(config.GetDB())
	jadwalKuliahRepository := repositories.NewJadwalKuliahRepository(config.GetDB())
	userRepository := repositories.NewUserRepository(config.GetDB())

	// Create service instances
	dosenService := services.NewDosenService(dosenRepository)
	mataKuliahService := services.NewMataKuliahService(mataKuliahRepository)
	mahasiswaService := services.NewMahasiswaService(mahasiswaRepository)
	nilaiService := services.NewNilaiService(nilaiRepository)
	jadwalKuliahService := services.NewJadwalKuliahService(jadwalKuliahRepository)
	userService := services.NewUserService(userRepository)

	// Create controller instances
	dosenController := controllers.NewDosenController(dosenService)
	mataKuliahController := controllers.NewMataKuliahController(mataKuliahService)
	mahasiswaController := controllers.NewMahasiswaController(mahasiswaService)
	nilaiController := controllers.NewNilaiController(nilaiService)
	jadwalKuliahController := controllers.NewJadwalKuliahController(jadwalKuliahService)
	userController := controllers.NewUserController(userService)

	// Routes for dosen
	app.POST("/dosen", dosenController.CreateDosen)
	app.GET("/dosen/:id", dosenController.GetDosenByID)
	app.PATCH("/dosen/:id", dosenController.UpdateDosen)
	app.DELETE("/dosen/:id", dosenController.DeleteDosen)

	// Routes for mata kuliah
	app.POST("/mata-kuliah", mataKuliahController.CreateMataKuliah)
	app.GET("/mata-kuliah/:id", mataKuliahController.GetMataKuliahByID)
	app.PATCH("/mata-kuliah/:id", mataKuliahController.UpdateMataKuliah)
	app.DELETE("/mata-kuliah/:id", mataKuliahController.DeleteMataKuliah)

	// Routes for mahasiswa
	app.POST("/mahasiswa", mahasiswaController.CreateMahasiswa)
	app.GET("/mahasiswa/:id", mahasiswaController.GetMahasiswaByID)
	app.PATCH("/mahasiswa/:id", mahasiswaController.UpdateMahasiswa)
	app.DELETE("/mahasiswa/:id", mahasiswaController.DeleteMahasiswa)

	// Routes for nilai
	app.POST("/nilai", nilaiController.CreateNilai)
	app.GET("/nilai/:id", nilaiController.GetNilaiByID)
	app.PATCH("/nilai/:id", nilaiController.UpdateNilai)
	app.DELETE("/nilai/:id", nilaiController.DeleteNilai)

	// Routes for jadwal kuliah
	app.POST("/jadwal-kuliah", jadwalKuliahController.CreateJadwalKuliah)
	app.GET("/jadwal-kuliah/:id", jadwalKuliahController.GetJadwalKuliahByID)
	app.PATCH("/jadwal-kuliah/:id", jadwalKuliahController.UpdateJadwalKuliah)
	app.DELETE("/jadwal-kuliah/:id", jadwalKuliahController.DeleteJadwalKuliah)

	// Routes for user
	app.POST("/user", userController.CreateUser)
	app.GET("/user/:id", userController.GetUserByID)
	app.PATCH("/user/:id", userController.UpdateUser)
	app.DELETE("/user/:id", userController.DeleteUser)

	// Swagger endpoint
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func main() {
	// Run the server
	app.Run(":8080")
}

// Handler function to handle HTTP requests
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
