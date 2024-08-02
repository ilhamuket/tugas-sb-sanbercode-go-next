package controllers

import (
	"net/http"
	"strconv"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/models"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/services"

	"github.com/gin-gonic/gin"
)

type NilaiController struct {
	service services.NilaiService
}

func NewNilaiController(service services.NilaiService) *NilaiController {
	return &NilaiController{service: service}
}

// CreateNilai godoc
// @Summary Create a new Nilai
// @Description Create a new Nilai with the input payload
// @Tags nilai
// @Accept json
// @Produce json
// @Param nilai body models.CreateNilaiInput true "Create Nilai"
// @Success 201 {object} models.Nilai
// @Failure 400 {object} map[string]interface{}
// @Security BearerAuth
// @Router /nilai [post]
func (ctrl *NilaiController) CreateNilai(c *gin.Context) {
	var input models.CreateNilaiInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	nilai, err := ctrl.service.CreateNilai(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, nilai)
}

// GetNilaiByID godoc
// @Summary Get a Nilai by ID
// @Description Get a Nilai by its ID
// @Tags nilai
// @Accept json
// @Produce json
// @Param id path int true "Nilai ID"
// @Success 200 {object} models.Nilai
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Security BearerAuth
// @Router /nilai/{id} [get]
func (ctrl *NilaiController) GetNilaiByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid ID"})
		return
	}

	// Convert id to uint if necessary
	uintID := uint(id)

	nilai, err := ctrl.service.GetNilaiByID(uintID)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nilai)
}

// UpdateNilai godoc
// @Summary Update a Nilai
// @Description Update a Nilai with the input payload
// @Tags nilai
// @Accept json
// @Produce json
// @Param id path int true "Nilai ID"
// @Param nilai body models.UpdateNilaiInput true "Update Nilai"
// @Success 200 {object} models.Nilai
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Security BearerAuth
// @Router /nilai/{id} [patch]
func (ctrl *NilaiController) UpdateNilai(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid ID"})
		return
	}

	// Convert id to uint if necessary
	uintID := uint(id)

	var input models.UpdateNilaiInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	nilai, err := ctrl.service.UpdateNilai(uintID, input)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nilai)
}

// DeleteNilai godoc
// @Summary Delete a Nilai
// @Description Delete a Nilai by its ID
// @Tags nilai
// @Accept json
// @Produce json
// @Param id path int true "Nilai ID"
// @Success 204
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Security BearerAuth
// @Router /nilai/{id} [delete]
func (ctrl *NilaiController) DeleteNilai(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid ID"})
		return
	}

	// Convert id to uint if necessary
	uintID := uint(id)

	err = ctrl.service.DeleteNilai(uintID)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetAllNilai godoc
// @Summary Get All Nilai
// @Description Get All Nilai
// @Tags nilai
// @Accept json
// @Produce json
// @Success 200 {object} models.Nilai
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /nilai [get]
func (ctrl *NilaiController) GetAllNilai(c *gin.Context) {
	nilai, err := ctrl.service.GetAllNilai()
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nilai)
}
