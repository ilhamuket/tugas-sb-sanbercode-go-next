package controllers

import (
	"net/http"
	"strconv"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/models"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/services"

	"github.com/gin-gonic/gin"
)

type JadwalKuliahController struct {
	service services.JadwalKuliahService
}

func NewJadwalKuliahController(service services.JadwalKuliahService) *JadwalKuliahController {
	return &JadwalKuliahController{service: service}
}

// CreateJadwalKuliah godoc
// @Summary Create a new Jadwal Kuliah
// @Description Create a new Jadwal Kuliah with the input payload
// @Tags jadwal-kuliah
// @Accept json
// @Produce json
// @Param jadwal-kuliah body models.CreateJadwalKuliahInput true "Create Jadwal Kuliah"
// @Success 201 {object} models.JadwalKuliah
// @Failure 400 {object} map[string]interface{}
// @Security BearerAuth
// @Router /jadwal-kuliah [post]
func (ctrl *JadwalKuliahController) CreateJadwalKuliah(c *gin.Context) {
	var input models.CreateJadwalKuliahInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	jadwalKuliah, err := ctrl.service.CreateJadwalKuliah(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, jadwalKuliah)
}

// GetJadwalKuliahByID godoc
// @Summary Get a Jadwal Kuliah by ID
// @Description Get a Jadwal Kuliah by its ID
// @Tags jadwal-kuliah
// @Accept json
// @Produce json
// @Param id path int true "Jadwal Kuliah ID"
// @Success 200 {object} models.JadwalKuliah
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Security BearerAuth
// @Router /jadwal-kuliah/{id} [get]
func (ctrl *JadwalKuliahController) GetJadwalKuliahByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid ID"})
		return
	}

	// Convert id to uint if necessary
	uintID := uint(id)

	jadwalKuliah, err := ctrl.service.GetJadwalKuliahByID(uintID)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, jadwalKuliah)
}

// UpdateJadwalKuliah godoc
// @Summary Update a Jadwal Kuliah
// @Description Update a Jadwal Kuliah with the input payload
// @Tags jadwal-kuliah
// @Accept json
// @Produce json
// @Param id path int true "Jadwal Kuliah ID"
// @Param jadwal-kuliah body models.UpdateJadwalKuliahInput true "Update Jadwal Kuliah"
// @Success 200 {object} models.JadwalKuliah
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Security BearerAuth
// @Router /jadwal-kuliah/{id} [patch]
func (ctrl *JadwalKuliahController) UpdateJadwalKuliah(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid ID"})
		return
	}

	// Convert id to uint if necessary
	uintID := uint(id)

	var input models.UpdateJadwalKuliahInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	jadwalKuliah, err := ctrl.service.UpdateJadwalKuliah(uintID, input)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, jadwalKuliah)
}

// DeleteJadwalKuliah godoc
// @Summary Delete a Jadwal Kuliah
// @Description Delete a Jadwal Kuliah by its ID
// @Tags jadwal-kuliah
// @Accept json
// @Produce json
// @Param id path int true "Jadwal Kuliah ID"
// @Success 204
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Security BearerAuth
// @Router /jadwal-kuliah/{id} [delete]
func (ctrl *JadwalKuliahController) DeleteJadwalKuliah(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid ID"})
		return
	}

	// Convert id to uint if necessary
	uintID := uint(id)

	err = ctrl.service.DeleteJadwalKuliah(uintID)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetAllJadwalKuliah godoc
// @Summary Get All Jadwal Kuliah
// @Description Get All Jadwal Kuliah
// @Tags jadwal-kuliah
// @Accept json
// @Produce json
// @Success 200 {object} models.JadwalKuliah
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /jadwal-kuliah [get]
func (ctrl *JadwalKuliahController) GetAllJadwalKuliah(c *gin.Context) {
	jadwalKuliah, err := ctrl.service.GetAllJadwalKuliah()
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, jadwalKuliah)
}
