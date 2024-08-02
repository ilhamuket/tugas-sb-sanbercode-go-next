package models

import (
	"time"
)

type JadwalKuliah struct {
	ID          uint `gorm:"primaryKey"`
	DosenID     uint
	Dosen       Dosen `gorm:"foreignKey:DosenID"`
	MahasiswaID uint
	Mahasiswa   Mahasiswa `gorm:"foreignKey:MahasiswaID"`
	Hari        string    `gorm:"type:varchar(255)"`
	JamMulai    time.Time
	JamSelesai  time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CreateJadwalKuliahInput struct {
	DosenID     uint      `json:"dosen_id" binding:"required"`
	MahasiswaID uint      `json:"mahasiswa_id" binding:"required"`
	Hari        string    `json:"hari" binding:"required"`
	JamMulai    time.Time `json:"jam_mulai" binding:"required"`
	JamSelesai  time.Time `json:"jam_selesai" binding:"required"`
}

type UpdateJadwalKuliahInput struct {
	DosenID     uint      `json:"dosen_id"`
	MahasiswaID uint      `json:"mahasiswa_id"`
	Hari        string    `json:"hari"`
	JamMulai    time.Time `json:"jam_mulai"`
	JamSelesai  time.Time `json:"jam_selesai"`
}
