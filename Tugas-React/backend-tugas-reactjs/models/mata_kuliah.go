package models

import (
	"time"
)

type MataKuliah struct {
	ID        uint   `gorm:"primaryKey"`
	Nama      string `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateMataKuliahInput struct {
	Nama string `json:"nama" binding:"required"`
}

type UpdateMataKuliahInput struct {
	Nama string `json:"nama"`
}
