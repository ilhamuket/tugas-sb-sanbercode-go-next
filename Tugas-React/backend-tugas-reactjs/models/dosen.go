package models

import (
	"time"
)

type Dosen struct {
	ID           uint   `gorm:"primaryKey"`
	Nama         string `gorm:"type:varchar(255)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	MataKuliahID uint
}

type CreateDosenInput struct {
	Nama         string `json:"nama" binding:"required"`
	MataKuliahID uint   `json:"mata_kuliah_id" binding:"required"`
}

type UpdateDosenInput struct {
	Nama         string `json:"nama"`
	MataKuliahID uint   `json:"mata_kuliah_id"`
}
