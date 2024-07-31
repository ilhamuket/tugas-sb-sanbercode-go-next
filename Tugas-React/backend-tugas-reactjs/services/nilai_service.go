package services

import (
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/models"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/repositories"
)

type NilaiService interface {
	CreateNilai(input models.CreateNilaiInput) (models.Nilai, error)
	UpdateNilai(id uint, input models.UpdateNilaiInput) (models.Nilai, error)
	GetNilaiByID(id uint) (models.Nilai, error)
	DeleteNilai(id uint) error
}

type nilaiService struct {
	nilaiRepository repositories.NilaiRepository
}

func NewNilaiService(nilaiRepository repositories.NilaiRepository) NilaiService {
	return &nilaiService{nilaiRepository}
}

func (s *nilaiService) CreateNilai(input models.CreateNilaiInput) (models.Nilai, error) {
	nilai := models.Nilai{
		Indeks:       input.Indeks,
		Skor:         input.Skor,
		MahasiswaID:  input.MahasiswaID,
		MataKuliahID: input.MataKuliahID,
		UsersID:      input.UsersID,
	}

	newNilai, err := s.nilaiRepository.CreateNilai(nilai)
	return newNilai, err
}

func (s *nilaiService) UpdateNilai(id uint, input models.UpdateNilaiInput) (models.Nilai, error) {
	nilai, err := s.nilaiRepository.FindNilaiByID(id)
	if err != nil {
		return nilai, err
	}

	if input.Indeks != "" {
		nilai.Indeks = input.Indeks
	}

	if input.Skor != 0 {
		nilai.Skor = input.Skor
	}

	if input.MahasiswaID != 0 {
		nilai.MahasiswaID = input.MahasiswaID
	}

	if input.MataKuliahID != 0 {
		nilai.MataKuliahID = input.MataKuliahID
	}

	if input.UsersID != 0 {
		nilai.UsersID = input.UsersID
	}

	updatedNilai, err := s.nilaiRepository.UpdateNilai(nilai)
	return updatedNilai, err
}

func (s *nilaiService) GetNilaiByID(id uint) (models.Nilai, error) {
	return s.nilaiRepository.FindNilaiByID(id)
}

func (s *nilaiService) DeleteNilai(id uint) error {
	nilai, err := s.nilaiRepository.FindNilaiByID(id)
	if err != nil {
		return err
	}

	return s.nilaiRepository.DeleteNilai(nilai)
}
