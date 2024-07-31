package models

import (
	"time"
)

type Mahasiswa struct {
	ID        uint   `gorm:"primaryKey"`
	Nama      string `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateMahasiswaInput struct {
	Nama string `json:"nama" binding:"required"`
}

type UpdateMahasiswaInput struct {
	Nama string `json:"nama"`
}
