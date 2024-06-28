package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"my-app/Tugas-Golang/Tugas-13/db"
	"my-app/Tugas-Golang/Tugas-13/models"
	"my-app/Tugas-Golang/Tugas-13/utils"
)

func GetMahasiswas(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	rows, err := db.DB.Query("SELECT id, nama, created_at, updated_at FROM mahasiswa")
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var mahasiswas []models.Mahasiswa
	for rows.Next() {
		var m models.Mahasiswa
		var createdAtStr, updatedAtStr string
		if err := rows.Scan(&m.ID, &m.Nama, &createdAtStr, &updatedAtStr); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		createdAt, err := utils.ParseTime(createdAtStr)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		m.CreatedAt = createdAt

		updatedAt, err := utils.ParseTime(updatedAtStr)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		m.UpdatedAt = updatedAt

		mahasiswas = append(mahasiswas, m)
	}

	utils.RespondWithSuccess(w, http.StatusOK, "Mahasiswas retrieved successfully", mahasiswas)
}

func GetMahasiswa(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	var m models.Mahasiswa
	var createdAtStr, updatedAtStr string
	err := db.DB.QueryRow("SELECT id, nama, created_at, updated_at FROM mahasiswa WHERE id = ?", id).Scan(&m.ID, &m.Nama, &createdAtStr, &updatedAtStr)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.RespondWithError(w, http.StatusNotFound, "Mahasiswa not found")
		} else {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	createdAt, err := utils.ParseTime(createdAtStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	m.CreatedAt = createdAt

	updatedAt, err := utils.ParseTime(updatedAtStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	m.UpdatedAt = updatedAt

	utils.RespondWithSuccess(w, http.StatusOK, "Mahasiswa retrieved successfully", m)
}

func CreateMahasiswa(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var m models.Mahasiswa
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	result, err := db.DB.Exec("INSERT INTO mahasiswa (nama, created_at, updated_at) VALUES (?, ?, ?)", m.Nama, time.Now(), time.Now())
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	m.ID = int(id)

	utils.RespondWithSuccess(w, http.StatusCreated, "Mahasiswa created successfully", m)
}

func UpdateMahasiswa(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	var m models.Mahasiswa
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	result, err := db.DB.Exec("UPDATE mahasiswa SET nama = ?, updated_at = ? WHERE id = ?", m.Nama, time.Now(), id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected == 0 {
		utils.RespondWithError(w, http.StatusNotFound, "Mahasiswa not found")
		return
	}

	m.ID, _ = strconv.Atoi(id)
	utils.RespondWithSuccess(w, http.StatusOK, "Mahasiswa updated successfully", m)
}

func DeleteMahasiswa(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	result, err := db.DB.Exec("DELETE FROM mahasiswa WHERE id = ?", id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected == 0 {
		utils.RespondWithError(w, http.StatusNotFound, "Mahasiswa not found")
		return
	}

	utils.RespondWithSuccess(w, http.StatusOK, "Mahasiswa deleted successfully", map[string]string{"result": "success"})
}
