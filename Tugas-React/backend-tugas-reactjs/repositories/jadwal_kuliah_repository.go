package repositories

import (
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/models"

	"gorm.io/gorm"
)

type JadwalKuliahRepository interface {
	CreateJadwalKuliah(jadwalKuliah models.JadwalKuliah) (models.JadwalKuliah, error)
	UpdateJadwalKuliah(jadwalKuliah models.JadwalKuliah) (models.JadwalKuliah, error)
	FindJadwalKuliahByID(id uint) (models.JadwalKuliah, error)
	GetAllJadwalKuliah() ([]models.JadwalKuliah, error)
	DeleteJadwalKuliah(jadwalKuliah models.JadwalKuliah) error
}

type jadwalKuliahRepository struct {
	db *gorm.DB
}

func NewJadwalKuliahRepository(db *gorm.DB) JadwalKuliahRepository {
	return &jadwalKuliahRepository{db}
}

func (r *jadwalKuliahRepository) CreateJadwalKuliah(jadwalKuliah models.JadwalKuliah) (models.JadwalKuliah, error) {
	err := r.db.Create(&jadwalKuliah).Error
	return jadwalKuliah, err
}

func (r *jadwalKuliahRepository) UpdateJadwalKuliah(jadwalKuliah models.JadwalKuliah) (models.JadwalKuliah, error) {
	err := r.db.Save(&jadwalKuliah).Error
	return jadwalKuliah, err
}

func (r *jadwalKuliahRepository) FindJadwalKuliahByID(id uint) (models.JadwalKuliah, error) {
	var jadwalKuliah models.JadwalKuliah
	err := r.db.First(&jadwalKuliah, id).Error
	return jadwalKuliah, err
}

func (r *jadwalKuliahRepository) GetAllJadwalKuliah() ([]models.JadwalKuliah, error) {
	var jadwalKuliah []models.JadwalKuliah
	err := r.db.Preload("Dosen").Preload("Mahasiswa").Find(&jadwalKuliah).Error
	if err != nil {
		return nil, err
	}
	return jadwalKuliah, nil
}

func (r *jadwalKuliahRepository) DeleteJadwalKuliah(jadwalKuliah models.JadwalKuliah) error {
	err := r.db.Delete(&jadwalKuliah).Error
	return err
}
