package services

import (
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/models"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/repositories"
)

type MataKuliahService interface {
	CreateMataKuliah(input models.CreateMataKuliahInput) (models.MataKuliah, error)
	UpdateMataKuliah(id uint, input models.UpdateMataKuliahInput) (models.MataKuliah, error)
	GetMataKuliahByID(id uint) (models.MataKuliah, error)
	GetAllMataKuliahs() ([]models.MataKuliah, error)
	DeleteMataKuliah(id uint) error
}

type mataKuliahService struct {
	mataKuliahRepository repositories.MataKuliahRepository
}

func NewMataKuliahService(mataKuliahRepository repositories.MataKuliahRepository) MataKuliahService {
	return &mataKuliahService{mataKuliahRepository}
}

func (s *mataKuliahService) CreateMataKuliah(input models.CreateMataKuliahInput) (models.MataKuliah, error) {
	mataKuliah := models.MataKuliah{
		Nama: input.Nama,
	}

	newMataKuliah, err := s.mataKuliahRepository.CreateMataKuliah(mataKuliah)
	return newMataKuliah, err
}

func (s *mataKuliahService) UpdateMataKuliah(id uint, input models.UpdateMataKuliahInput) (models.MataKuliah, error) {
	mataKuliah, err := s.mataKuliahRepository.FindMataKuliahByID(id)
	if err != nil {
		return mataKuliah, err
	}

	if input.Nama != "" {
		mataKuliah.Nama = input.Nama
	}

	updatedMataKuliah, err := s.mataKuliahRepository.UpdateMataKuliah(mataKuliah)
	return updatedMataKuliah, err
}

func (s *mataKuliahService) GetMataKuliahByID(id uint) (models.MataKuliah, error) {
	return s.mataKuliahRepository.FindMataKuliahByID(id)
}

func (s *mataKuliahService) DeleteMataKuliah(id uint) error {
	mataKuliah, err := s.mataKuliahRepository.FindMataKuliahByID(id)
	if err != nil {
		return err
	}

	return s.mataKuliahRepository.DeleteMataKuliah(mataKuliah)
}

func (s *mataKuliahService) GetAllMataKuliahs() ([]models.MataKuliah, error) {
	return s.mataKuliahRepository.GetAllMataKuliahs()
}
