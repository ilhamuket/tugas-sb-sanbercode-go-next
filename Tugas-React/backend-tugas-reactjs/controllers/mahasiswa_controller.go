package controllers

import (
	"net/http"
	"strconv"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/models"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/services"

	"github.com/gin-gonic/gin"
)

type MahasiswaController struct {
	service services.MahasiswaService
}

func NewMahasiswaController(service services.MahasiswaService) *MahasiswaController {
	return &MahasiswaController{service: service}
}

// CreateMahasiswa godoc
// @Summary Create a new Mahasiswa
// @Description Create a new Mahasiswa with the input payload
// @Tags mahasiswa
// @Accept json
// @Produce json
// @Param mahasiswa body models.CreateMahasiswaInput true "Create Mahasiswa"
// @Success 201 {object} models.Mahasiswa
// @Failure 400 {object} map[string]interface{}
// @Router /mahasiswa [post]
func (ctrl *MahasiswaController) CreateMahasiswa(c *gin.Context) {
	var input models.CreateMahasiswaInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	mahasiswa, err := ctrl.service.CreateMahasiswa(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, mahasiswa)
}

// GetMahasiswaByID godoc
// @Summary Get a Mahasiswa by ID
// @Description Get a Mahasiswa by its ID
// @Tags mahasiswa
// @Accept json
// @Produce json
// @Param id path int true "Mahasiswa ID"
// @Success 200 {object} models.Mahasiswa
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /mahasiswa/{id} [get]
func (ctrl *MahasiswaController) GetMahasiswaByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid ID"})
		return
	}

	// Convert id to uint if necessary
	uintID := uint(id)

	mahasiswa, err := ctrl.service.GetMahasiswaByID(uintID)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mahasiswa)
}

// UpdateMahasiswa godoc
// @Summary Update a Mahasiswa
// @Description Update a Mahasiswa with the input payload
// @Tags mahasiswa
// @Accept json
// @Produce json
// @Param id path int true "Mahasiswa ID"
// @Param mahasiswa body models.UpdateMahasiswaInput true "Update Mahasiswa"
// @Success 200 {object} models.Mahasiswa
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /mahasiswa/{id} [patch]
func (ctrl *MahasiswaController) UpdateMahasiswa(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid ID"})
		return
	}

	// Convert id to uint if necessary
	uintID := uint(id)

	var input models.UpdateMahasiswaInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	mahasiswa, err := ctrl.service.UpdateMahasiswa(uintID, input)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mahasiswa)
}

// DeleteMahasiswa godoc
// @Summary Delete a Mahasiswa
// @Description Delete a Mahasiswa by its ID
// @Tags mahasiswa
// @Accept json
// @Produce json
// @Param id path int true "Mahasiswa ID"
// @Success 204
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /mahasiswa/{id} [delete]
func (ctrl *MahasiswaController) DeleteMahasiswa(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid ID"})
		return
	}

	// Convert id to uint if necessary
	uintID := uint(id)

	err = ctrl.service.DeleteMahasiswa(uintID)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
