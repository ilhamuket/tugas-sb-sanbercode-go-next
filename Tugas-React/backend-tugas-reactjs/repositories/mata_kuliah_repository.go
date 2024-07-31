package repositories

import (
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/models"

	"gorm.io/gorm"
)

type MataKuliahRepository interface {
	CreateMataKuliah(mataKuliah models.MataKuliah) (models.MataKuliah, error)
	UpdateMataKuliah(mataKuliah models.MataKuliah) (models.MataKuliah, error)
	FindMataKuliahByID(id uint) (models.MataKuliah, error)
	DeleteMataKuliah(mataKuliah models.MataKuliah) error
}

type mataKuliahRepository struct {
	db *gorm.DB
}

func NewMataKuliahRepository(db *gorm.DB) MataKuliahRepository {
	return &mataKuliahRepository{db}
}

func (r *mataKuliahRepository) CreateMataKuliah(mataKuliah models.MataKuliah) (models.MataKuliah, error) {
	err := r.db.Create(&mataKuliah).Error
	return mataKuliah, err
}

func (r *mataKuliahRepository) UpdateMataKuliah(mataKuliah models.MataKuliah) (models.MataKuliah, error) {
	err := r.db.Save(&mataKuliah).Error
	return mataKuliah, err
}

func (r *mataKuliahRepository) FindMataKuliahByID(id uint) (models.MataKuliah, error) {
	var mataKuliah models.MataKuliah
	err := r.db.First(&mataKuliah, id).Error
	return mataKuliah, err
}

func (r *mataKuliahRepository) DeleteMataKuliah(mataKuliah models.MataKuliah) error {
	err := r.db.Delete(&mataKuliah).Error
	return err
}
