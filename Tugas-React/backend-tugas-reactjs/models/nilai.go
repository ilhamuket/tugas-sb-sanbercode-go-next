package models

import (
	"time"
)

type Nilai struct {
	ID           uint   `gorm:"primaryKey"`
	Indeks       string `gorm:"type:varchar(255)"`
	Skor         int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	MahasiswaID  uint
	Mahasiswa    Mahasiswa `gorm:"foreignKey:MahasiswaID"`
	MataKuliahID uint
	MataKuliah   MataKuliah `gorm:"foreignKey:MataKuliahID"`
	UsersID      uint
	User         User `gorm:"foreignKey:UsersID"`
}

type CreateNilaiInput struct {
	Indeks       string `json:"indeks" binding:"required"`
	Skor         int    `json:"skor" binding:"required"`
	MahasiswaID  uint   `json:"mahasiswa_id" binding:"required"`
	MataKuliahID uint   `json:"mata_kuliah_id" binding:"required"`
	UsersID      uint   `json:"users_id" binding:"required"`
}

type UpdateNilaiInput struct {
	Indeks       string `json:"indeks"`
	Skor         int    `json:"skor"`
	MahasiswaID  uint   `json:"mahasiswa_id"`
	MataKuliahID uint   `json:"mata_kuliah_id"`
	UsersID      uint   `json:"users_id"`
}
