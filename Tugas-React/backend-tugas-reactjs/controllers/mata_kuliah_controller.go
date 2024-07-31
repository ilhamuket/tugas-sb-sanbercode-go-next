package controllers

import (
	"net/http"
	"strconv"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/models"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/services"

	"github.com/gin-gonic/gin"
)

type MataKuliahController struct {
	service services.MataKuliahService
}

func NewMataKuliahController(service services.MataKuliahService) *MataKuliahController {
	return &MataKuliahController{service: service}
}

// CreateMataKuliah godoc
// @Summary Create a new Mata Kuliah
// @Description Create a new Mata Kuliah with the input payload
// @Tags mata-kuliah
// @Accept json
// @Produce json
// @Param mata_kuliah body models.CreateMataKuliahInput true "Create Mata Kuliah"
// @Success 201 {object} models.MataKuliah
// @Failure 400 {object} map[string]interface{}
// @Router /mata-kuliah [post]
func (ctrl *MataKuliahController) CreateMataKuliah(c *gin.Context) {
	var input models.CreateMataKuliahInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	mataKuliah, err := ctrl.service.CreateMataKuliah(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, mataKuliah)
}

// GetMataKuliahByID godoc
// @Summary Get a Mata Kuliah by ID
// @Description Get a Mata Kuliah by its ID
// @Tags mata-kuliah
// @Accept json
// @Produce json
// @Param id path int true "Mata Kuliah ID"
// @Success 200 {object} models.MataKuliah
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /mata-kuliah/{id} [get]
func (ctrl *MataKuliahController) GetMataKuliahByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid ID"})
		return
	}

	// Convert id to uint if necessary
	uintID := uint(id)

	mataKuliah, err := ctrl.service.GetMataKuliahByID(uintID)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mataKuliah)
}

// UpdateMataKuliah godoc
// @Summary Update a Mata Kuliah
// @Description Update a Mata Kuliah with the input payload
// @Tags mata-kuliah
// @Accept json
// @Produce json
// @Param id path int true "Mata Kuliah ID"
// @Param mata_kuliah body models.UpdateMataKuliahInput true "Update Mata Kuliah"
// @Success 200 {object} models.MataKuliah
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /mata-kuliah/{id} [patch]
func (ctrl *MataKuliahController) UpdateMataKuliah(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid ID"})
		return
	}

	// Convert id to uint if necessary
	uintID := uint(id)

	var input models.UpdateMataKuliahInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	mataKuliah, err := ctrl.service.UpdateMataKuliah(uintID, input)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mataKuliah)
}

// DeleteMataKuliah godoc
// @Summary Delete a Mata Kuliah
// @Description Delete a Mata Kuliah by its ID
// @Tags mata-kuliah
// @Accept json
// @Produce json
// @Param id path int true "Mata Kuliah ID"
// @Success 204
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /mata-kuliah/{id} [delete]
func (ctrl *MataKuliahController) DeleteMataKuliah(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid ID"})
		return
	}

	// Convert id to uint if necessary
	uintID := uint(id)

	err = ctrl.service.DeleteMataKuliah(uintID)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
