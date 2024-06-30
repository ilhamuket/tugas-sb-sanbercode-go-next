package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"my-app/Tugas-Golang/Tugas-13/db"
	"my-app/Tugas-Golang/Tugas-13/models"
	"my-app/Tugas-Golang/Tugas-13/utils"
)

// GetNilais Handler for fetching all nilai
func GetNilais(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	rows, err := db.DB.Query("SELECT id, indeks, skor, created_at, updated_at, mata_kuliah_id, mahasiswa_id FROM nilai")
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	var nilais []models.Nilai
	for rows.Next() {
		var n models.Nilai
		var createdAtStr, updatedAtStr string
		if err := rows.Scan(&n.ID, &n.Indeks, &n.Skor, &createdAtStr, &updatedAtStr, &n.MataKuliahID, &n.MahasiswaID); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		createdAt, err := utils.ParseTime(createdAtStr)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		n.CreatedAt = createdAt

		updatedAt, err := utils.ParseTime(updatedAtStr)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		n.UpdatedAt = updatedAt

		nilais = append(nilais, n)
	}

	utils.RespondWithSuccess(w, http.StatusOK, "Nilai retrieved successfully", nilais)
}

// GetNilai Handler for fetching a single nilai by ID
func GetNilai(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	var n models.Nilai
	var createdAtStr, updatedAtStr string
	err := db.DB.QueryRow("SELECT id, indeks, skor, created_at, updated_at, mata_kuliah_id, mahasiswa_id FROM nilai WHERE id = ?", id).Scan(&n.ID, &n.Indeks, &n.Skor, &createdAtStr, &updatedAtStr, &n.MataKuliahID, &n.MahasiswaID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.RespondWithError(w, http.StatusNotFound, "Nilai not found")
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
	n.CreatedAt = createdAt

	updatedAt, err := utils.ParseTime(updatedAtStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	n.UpdatedAt = updatedAt

	utils.RespondWithSuccess(w, http.StatusOK, "Nilai retrieved successfully", n)
}

// CreateNilai Handler for creating a new nilai
func CreateNilai(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var n models.Nilai
	if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Check if mata_kuliah_id exists
	err := db.DB.QueryRow("SELECT id FROM mata_kuliah WHERE id = ?", n.MataKuliahID).Scan(&n.MataKuliahID)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid mata_kuliah_id")
		return
	}

	// Check if mahasiswa_id exists
	err = db.DB.QueryRow("SELECT id FROM mahasiswa WHERE id = ?", n.MahasiswaID).Scan(&n.MahasiswaID)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid mahasiswa_id")
		return
	}

	// Set indeks based on skor
	n.Indeks = calculateIndeks(n.Skor)

	result, err := db.DB.Exec("INSERT INTO nilai (indeks, skor, created_at, updated_at, mata_kuliah_id, mahasiswa_id) VALUES (?, ?, ?, ?, ?, ?)", n.Indeks, n.Skor, time.Now(), time.Now(), n.MataKuliahID, n.MahasiswaID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	n.ID = int(id)

	utils.RespondWithSuccess(w, http.StatusCreated, "Nilai created successfully", n)
}

// UpdateNilai asndland'ad Handler for updating an existing nilai
func UpdateNilai(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	var n models.Nilai
	if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Pastikan Anda mendapatkan skor yang baru dari input
	newSkor := n.Skor

	// Check if mata_kuliah_id exists
	err := db.DB.QueryRow("SELECT id FROM mata_kuliah WHERE id = ?", n.MataKuliahID).Scan(&n.MataKuliahID)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid mata_kuliah_id")
		return
	}

	// Check if mahasiswa_id exists
	err = db.DB.QueryRow("SELECT id FROM mahasiswa WHERE id = ?", n.MahasiswaID).Scan(&n.MahasiswaID)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid mahasiswa_id")
		return
	}

	// Update nilai with new Skor and Indeks
	n.Indeks = calculateIndeks(newSkor) // Recalculate Indeks based on new Skor

	_, err = db.DB.Exec("UPDATE nilai SET indeks = ?, skor = ?, updated_at = ?, mata_kuliah_id = ?, mahasiswa_id = ? WHERE id = ?", n.Indeks, newSkor, time.Now(), n.MataKuliahID, n.MahasiswaID, id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(w, http.StatusOK, "Nilai updated successfully", n)
}

// DeleteNilai Handler for deleting a nilai by ID
func DeleteNilai(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	result, err := db.DB.Exec("DELETE FROM nilai WHERE id = ?", id)
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
		utils.RespondWithError(w, http.StatusNotFound, "Nilai not found")
		return
	}

	utils.RespondWithSuccess(w, http.StatusOK, "Nilai deleted successfully", map[string]string{"result": "success"})
}

// Helper function to calculate indeks based on skor
func calculateIndeks(skor int) string {
	switch {
	case skor >= 80:
		return "A"
	case skor >= 70:
		return "B"
	case skor >= 60:
		return "C"
	case skor >= 50:
		return "D"
	default:
		return "E"
	}
}
