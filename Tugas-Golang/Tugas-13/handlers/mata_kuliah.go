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

// GetMataKuliahs Handler for fetching all mata_kuliah
func GetMataKuliahs(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	rows, err := db.DB.Query("SELECT id, nama, created_at, updated_at FROM mata_kuliah")
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			// Handle or log the error here if necessary
		}
	}(rows)

	var mataKuliahs []models.MataKuliah
	for rows.Next() {
		var mk models.MataKuliah
		var createdAtStr, updatedAtStr string
		if err := rows.Scan(&mk.ID, &mk.Nama, &createdAtStr, &updatedAtStr); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		createdAt, err := utils.ParseTime(createdAtStr)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		mk.CreatedAt = createdAt

		updatedAt, err := utils.ParseTime(updatedAtStr)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		mk.UpdatedAt = updatedAt

		mataKuliahs = append(mataKuliahs, mk)
	}

	utils.RespondWithSuccess(w, http.StatusOK, "Mata Kuliah retrieved successfully", mataKuliahs)
}

// GetMataKuliah Handler for fetching a single mata_kuliah by ID
func GetMataKuliah(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	var mk models.MataKuliah
	var createdAtStr, updatedAtStr string
	err := db.DB.QueryRow("SELECT id, nama, created_at, updated_at FROM mata_kuliah WHERE id = ?", id).Scan(&mk.ID, &mk.Nama, &createdAtStr, &updatedAtStr)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.RespondWithError(w, http.StatusNotFound, "Mata Kuliah not found")
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
	mk.CreatedAt = createdAt

	updatedAt, err := utils.ParseTime(updatedAtStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	mk.UpdatedAt = updatedAt

	utils.RespondWithSuccess(w, http.StatusOK, "Mata Kuliah retrieved successfully", mk)
}

// CreateMataKuliah Handler for creating a new mata_kuliah
func CreateMataKuliah(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var mk models.MataKuliah
	if err := json.NewDecoder(r.Body).Decode(&mk); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	result, err := db.DB.Exec("INSERT INTO mata_kuliah (nama, created_at, updated_at) VALUES (?, ?, ?)", mk.Nama, time.Now(), time.Now())
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	mk.ID = int(id)

	utils.RespondWithSuccess(w, http.StatusCreated, "Mata Kuliah created successfully", mk)
}

// UpdateMataKuliah Handler for updating a mata_kuliah by ID
func UpdateMataKuliah(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	var mk models.MataKuliah
	if err := json.NewDecoder(r.Body).Decode(&mk); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	result, err := db.DB.Exec("UPDATE mata_kuliah SET nama = ?, updated_at = ? WHERE id = ?", mk.Nama, time.Now(), id)
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
		utils.RespondWithError(w, http.StatusNotFound, "Mata Kuliah not found")
		return
	}

	mk.ID, _ = strconv.Atoi(id)
	utils.RespondWithSuccess(w, http.StatusOK, "Mata Kuliah updated successfully", mk)
}

// DeleteMataKuliah Handler for deleting a mata_kuliah by ID
func DeleteMataKuliah(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	result, err := db.DB.Exec("DELETE FROM mata_kuliah WHERE id = ?", id)
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
		utils.RespondWithError(w, http.StatusNotFound, "Mata Kuliah not found")
		return
	}

	utils.RespondWithSuccess(w, http.StatusOK, "Mata Kuliah deleted successfully", map[string]string{"result": "success"})
}
