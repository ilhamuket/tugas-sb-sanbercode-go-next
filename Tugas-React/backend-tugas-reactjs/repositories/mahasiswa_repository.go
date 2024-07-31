package repositories

import (
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/models"

	"gorm.io/gorm"
)

type MahasiswaRepository interface {
	CreateMahasiswa(mahasiswa models.Mahasiswa) (models.Mahasiswa, error)
	UpdateMahasiswa(mahasiswa models.Mahasiswa) (models.Mahasiswa, error)
	FindMahasiswaByID(id uint) (models.Mahasiswa, error)
	DeleteMahasiswa(mahasiswa models.Mahasiswa) error
}

type mahasiswaRepository struct {
	db *gorm.DB
}

func NewMahasiswaRepository(db *gorm.DB) MahasiswaRepository {
	return &mahasiswaRepository{db}
}

func (r *mahasiswaRepository) CreateMahasiswa(mahasiswa models.Mahasiswa) (models.Mahasiswa, error) {
	err := r.db.Create(&mahasiswa).Error
	return mahasiswa, err
}

func (r *mahasiswaRepository) UpdateMahasiswa(mahasiswa models.Mahasiswa) (models.Mahasiswa, error) {
	err := r.db.Save(&mahasiswa).Error
	return mahasiswa, err
}

func (r *mahasiswaRepository) FindMahasiswaByID(id uint) (models.Mahasiswa, error) {
	var mahasiswa models.Mahasiswa
	err := r.db.First(&mahasiswa, id).Error
	return mahasiswa, err
}

func (r *mahasiswaRepository) DeleteMahasiswa(mahasiswa models.Mahasiswa) error {
	err := r.db.Delete(&mahasiswa).Error
	return err
}
