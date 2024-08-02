package repositories

import (
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/models"

	"gorm.io/gorm"
)

type NilaiRepository interface {
	CreateNilai(nilai models.Nilai) (models.Nilai, error)
	UpdateNilai(nilai models.Nilai) (models.Nilai, error)
	FindNilaiByID(id uint) (models.Nilai, error)
	GetAllNilai() ([]models.Nilai, error)
	DeleteNilai(nilai models.Nilai) error
}

type nilaiRepository struct {
	db *gorm.DB
}

func NewNilaiRepository(db *gorm.DB) NilaiRepository {
	return &nilaiRepository{db}
}

func (r *nilaiRepository) CreateNilai(nilai models.Nilai) (models.Nilai, error) {
	err := r.db.Create(&nilai).Error
	return nilai, err
}

func (r *nilaiRepository) UpdateNilai(nilai models.Nilai) (models.Nilai, error) {
	err := r.db.Save(&nilai).Error
	return nilai, err
}

func (r *nilaiRepository) FindNilaiByID(id uint) (models.Nilai, error) {
	var nilai models.Nilai
	err := r.db.First(&nilai, id).Error
	return nilai, err
}

func (r *nilaiRepository) DeleteNilai(nilai models.Nilai) error {
	err := r.db.Delete(&nilai).Error
	return err
}

func (r *nilaiRepository) GetAllNilai() ([]models.Nilai, error) {
	var nilai []models.Nilai
	err := r.db.Preload("Mahasiswa").Preload("MataKuliah").Preload("User").Find(&nilai).Error
	if err != nil {
		return nil, err
	}
	return nilai, nil
}
