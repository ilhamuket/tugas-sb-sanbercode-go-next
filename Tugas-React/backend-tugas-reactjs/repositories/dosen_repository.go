package repositories

import (
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/models"

	"gorm.io/gorm"
)

type DosenRepository interface {
	CreateDosen(dosen models.Dosen) (models.Dosen, error)
	UpdateDosen(dosen models.Dosen) (models.Dosen, error)
	FindDosenByID(id uint) (models.Dosen, error)
	DeleteDosen(dosen models.Dosen) error
}

type dosenRepository struct {
	db *gorm.DB
}

func NewDosenRepository(db *gorm.DB) DosenRepository {
	return &dosenRepository{db}
}

func (r *dosenRepository) CreateDosen(dosen models.Dosen) (models.Dosen, error) {
	err := r.db.Create(&dosen).Error
	return dosen, err
}

func (r *dosenRepository) UpdateDosen(dosen models.Dosen) (models.Dosen, error) {
	err := r.db.Save(&dosen).Error
	return dosen, err
}

func (r *dosenRepository) FindDosenByID(id uint) (models.Dosen, error) {
	var dosen models.Dosen
	err := r.db.First(&dosen, id).Error
	return dosen, err
}

func (r *dosenRepository) DeleteDosen(dosen models.Dosen) error {
	err := r.db.Delete(&dosen).Error
	return err
}
