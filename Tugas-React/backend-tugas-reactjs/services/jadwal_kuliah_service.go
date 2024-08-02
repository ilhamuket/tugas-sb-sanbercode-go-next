package services

import (
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/models"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/repositories"
)

type JadwalKuliahService interface {
	CreateJadwalKuliah(input models.CreateJadwalKuliahInput) (models.JadwalKuliah, error)
	UpdateJadwalKuliah(id uint, input models.UpdateJadwalKuliahInput) (models.JadwalKuliah, error)
	GetJadwalKuliahByID(id uint) (models.JadwalKuliah, error)
	GetAllJadwalKuliah() ([]models.JadwalKuliah, error)
	DeleteJadwalKuliah(id uint) error
}

type jadwalKuliahService struct {
	jadwalKuliahRepository repositories.JadwalKuliahRepository
}

func NewJadwalKuliahService(jadwalKuliahRepository repositories.JadwalKuliahRepository) JadwalKuliahService {
	return &jadwalKuliahService{jadwalKuliahRepository}
}

func (s *jadwalKuliahService) CreateJadwalKuliah(input models.CreateJadwalKuliahInput) (models.JadwalKuliah, error) {
	jadwalKuliah := models.JadwalKuliah{
		DosenID:     input.DosenID,
		MahasiswaID: input.MahasiswaID,
		Hari:        input.Hari,
		JamMulai:    input.JamMulai,
		JamSelesai:  input.JamSelesai,
	}

	newJadwalKuliah, err := s.jadwalKuliahRepository.CreateJadwalKuliah(jadwalKuliah)
	return newJadwalKuliah, err
}

func (s *jadwalKuliahService) UpdateJadwalKuliah(id uint, input models.UpdateJadwalKuliahInput) (models.JadwalKuliah, error) {
	jadwalKuliah, err := s.jadwalKuliahRepository.FindJadwalKuliahByID(id)
	if err != nil {
		return jadwalKuliah, err
	}

	if input.DosenID != 0 {
		jadwalKuliah.DosenID = input.DosenID
	}

	if input.MahasiswaID != 0 {
		jadwalKuliah.MahasiswaID = input.MahasiswaID
	}

	if input.Hari != "" {
		jadwalKuliah.Hari = input.Hari
	}

	if !input.JamMulai.IsZero() {
		jadwalKuliah.JamMulai = input.JamMulai
	}

	if !input.JamSelesai.IsZero() {
		jadwalKuliah.JamSelesai = input.JamSelesai
	}

	updatedJadwalKuliah, err := s.jadwalKuliahRepository.UpdateJadwalKuliah(jadwalKuliah)
	return updatedJadwalKuliah, err
}

func (s *jadwalKuliahService) GetJadwalKuliahByID(id uint) (models.JadwalKuliah, error) {
	return s.jadwalKuliahRepository.FindJadwalKuliahByID(id)
}

func (s *jadwalKuliahService) GetAllJadwalKuliah() ([]models.JadwalKuliah, error) {
	return s.jadwalKuliahRepository.GetAllJadwalKuliah()
}

func (s *jadwalKuliahService) DeleteJadwalKuliah(id uint) error {
	jadwalKuliah, err := s.jadwalKuliahRepository.FindJadwalKuliahByID(id)
	if err != nil {
		return err
	}

	return s.jadwalKuliahRepository.DeleteJadwalKuliah(jadwalKuliah)
}
