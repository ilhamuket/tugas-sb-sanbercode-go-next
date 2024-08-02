package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"time"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/controllers"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/middlewares"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/repositories"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/services"
)

func SetupRouter(db *gorm.DB, app *gin.Engine) {

	// Middleware to set db in context
	app.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Setup CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Content-Type", "Authorization"}
	corsConfig.ExposeHeaders = []string{"Content-Length"}
	corsConfig.AllowCredentials = true
	corsConfig.MaxAge = 12 * time.Hour

	// Enable CORS with the configured settings
	app.Use(cors.New(corsConfig))

	// Initialize repositories
	dosenRepository := repositories.NewDosenRepository(db)
	mataKuliahRepository := repositories.NewMataKuliahRepository(db)
	mahasiswaRepository := repositories.NewMahasiswaRepository(db)
	nilaiRepository := repositories.NewNilaiRepository(db)
	jadwalKuliahRepository := repositories.NewJadwalKuliahRepository(db)
	userRepository := repositories.NewUserRepository(db)

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

	// Routes for books (GET routes don't require auth middleware)
	app.GET("/books", controllers.GetBooks)
	app.GET("/books/:id", controllers.GetBook)

	// Auth middleware for POST, PATCH, DELETE routes
	authRoutes := app.Group("")
	authRoutes.Use(middlewares.AuthMiddleware())

	authRoutes.POST("/books", controllers.CreateBook)
	authRoutes.PATCH("/books/:id", controllers.UpdateBook)
	authRoutes.DELETE("/books/:id", controllers.DeleteBook)

	// Routes for dosen
	authRoutes.POST("/dosen", dosenController.CreateDosen)
	authRoutes.PATCH("/dosen/:id", dosenController.UpdateDosen)
	authRoutes.DELETE("/dosen/:id", dosenController.DeleteDosen)
	app.GET("/dosen/:id", dosenController.GetDosenByID)
	app.GET("/dosen", dosenController.GetAllDosens)

	// Routes for mata kuliah
	authRoutes.POST("/mata-kuliah", mataKuliahController.CreateMataKuliah)
	authRoutes.PATCH("/mata-kuliah/:id", mataKuliahController.UpdateMataKuliah)
	authRoutes.DELETE("/mata-kuliah/:id", mataKuliahController.DeleteMataKuliah)
	app.GET("/mata-kuliah/:id", mataKuliahController.GetMataKuliahByID)
	app.GET("/mata-kuliah", mataKuliahController.GetAllMataKuliahs)

	// Routes for mahasiswa
	authRoutes.POST("/mahasiswa", mahasiswaController.CreateMahasiswa)
	authRoutes.PATCH("/mahasiswa/:id", mahasiswaController.UpdateMahasiswa)
	authRoutes.DELETE("/mahasiswa/:id", mahasiswaController.DeleteMahasiswa)
	app.GET("/mahasiswa/:id", mahasiswaController.GetMahasiswaByID)
	app.GET("/mahasiswa", mahasiswaController.GetAllMahasiswa)

	// Routes for nilai
	authRoutes.POST("/nilai", nilaiController.CreateNilai)
	authRoutes.PATCH("/nilai/:id", nilaiController.UpdateNilai)
	authRoutes.DELETE("/nilai/:id", nilaiController.DeleteNilai)
	app.GET("/nilai/:id", nilaiController.GetNilaiByID)
	app.GET("/nilai", nilaiController.GetAllNilai)

	// Routes for jadwal kuliah
	authRoutes.POST("/jadwal-kuliah", jadwalKuliahController.CreateJadwalKuliah)
	authRoutes.PATCH("/jadwal-kuliah/:id", jadwalKuliahController.UpdateJadwalKuliah)
	authRoutes.DELETE("/jadwal-kuliah/:id", jadwalKuliahController.DeleteJadwalKuliah)
	app.GET("/jadwal-kuliah/:id", jadwalKuliahController.GetJadwalKuliahByID)
	app.GET("/jadwal-kuliah", jadwalKuliahController.GetAllJadwalKuliah)

	// Routes for user
	authRoutes.POST("/user", userController.CreateUser)
	authRoutes.PATCH("/user/:id", userController.UpdateUser)
	authRoutes.DELETE("/user/:id", userController.DeleteUser)
	app.GET("/user/:id", userController.GetUserByID)

	// Auth routes
	app.POST("/register", userController.RegisterUser)
	app.POST("/login", userController.LoginUser)

	// Swagger endpoint
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
