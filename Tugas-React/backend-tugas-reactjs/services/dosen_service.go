package services

import (
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/models"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/repositories"
)

type DosenService interface {
	CreateDosen(input models.CreateDosenInput) (models.Dosen, error)
	UpdateDosen(id uint, input models.UpdateDosenInput) (models.Dosen, error)
	GetDosenByID(id uint) (models.Dosen, error)
	DeleteDosen(id uint) error
}

type dosenService struct {
	dosenRepository repositories.DosenRepository
}

func NewDosenService(dosenRepository repositories.DosenRepository) DosenService {
	return &dosenService{dosenRepository}
}

func (s *dosenService) CreateDosen(input models.CreateDosenInput) (models.Dosen, error) {
	dosen := models.Dosen{
		Nama:         input.Nama,
		MataKuliahID: input.MataKuliahID,
	}

	newDosen, err := s.dosenRepository.CreateDosen(dosen)
	return newDosen, err
}

func (s *dosenService) UpdateDosen(id uint, input models.UpdateDosenInput) (models.Dosen, error) {
	dosen, err := s.dosenRepository.FindDosenByID(id)
	if err != nil {
		return dosen, err
	}

	if input.Nama != "" {
		dosen.Nama = input.Nama
	}

	if input.MataKuliahID != 0 {
		dosen.MataKuliahID = input.MataKuliahID
	}

	updatedDosen, err := s.dosenRepository.UpdateDosen(dosen)
	return updatedDosen, err
}

func (s *dosenService) GetDosenByID(id uint) (models.Dosen, error) {
	return s.dosenRepository.FindDosenByID(id)
}

func (s *dosenService) DeleteDosen(id uint) error {
	dosen, err := s.dosenRepository.FindDosenByID(id)
	if err != nil {
		return err
	}

	return s.dosenRepository.DeleteDosen(dosen)
}
