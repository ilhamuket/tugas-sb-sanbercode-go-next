package services

import (
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/models"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/repositories"
)

type MahasiswaService interface {
	CreateMahasiswa(input models.CreateMahasiswaInput) (models.Mahasiswa, error)
	UpdateMahasiswa(id uint, input models.UpdateMahasiswaInput) (models.Mahasiswa, error)
	GetMahasiswaByID(id uint) (models.Mahasiswa, error)
	DeleteMahasiswa(id uint) error
}

type mahasiswaService struct {
	mahasiswaRepository repositories.MahasiswaRepository
}

func NewMahasiswaService(mahasiswaRepository repositories.MahasiswaRepository) MahasiswaService {
	return &mahasiswaService{mahasiswaRepository}
}

func (s *mahasiswaService) CreateMahasiswa(input models.CreateMahasiswaInput) (models.Mahasiswa, error) {
	mahasiswa := models.Mahasiswa{
		Nama: input.Nama,
	}

	newMahasiswa, err := s.mahasiswaRepository.CreateMahasiswa(mahasiswa)
	return newMahasiswa, err
}

func (s *mahasiswaService) UpdateMahasiswa(id uint, input models.UpdateMahasiswaInput) (models.Mahasiswa, error) {
	mahasiswa, err := s.mahasiswaRepository.FindMahasiswaByID(id)
	if err != nil {
		return mahasiswa, err
	}

	if input.Nama != "" {
		mahasiswa.Nama = input.Nama
	}

	updatedMahasiswa, err := s.mahasiswaRepository.UpdateMahasiswa(mahasiswa)
	return updatedMahasiswa, err
}

func (s *mahasiswaService) GetMahasiswaByID(id uint) (models.Mahasiswa, error) {
	return s.mahasiswaRepository.FindMahasiswaByID(id)
}

func (s *mahasiswaService) DeleteMahasiswa(id uint) error {
	mahasiswa, err := s.mahasiswaRepository.FindMahasiswaByID(id)
	if err != nil {
		return err
	}

	return s.mahasiswaRepository.DeleteMahasiswa(mahasiswa)
}
