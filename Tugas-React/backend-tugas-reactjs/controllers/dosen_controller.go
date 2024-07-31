package controllers

import (
	"net/http"
	"strconv"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/models"
	"tugas-sb-sanbercode-go-next-2024/Tugas-React/backend-tugas-reactjs/services"

	"github.com/gin-gonic/gin"
)

type DosenController struct {
	service services.DosenService
}

func NewDosenController(service services.DosenService) *DosenController {
	return &DosenController{service: service}
}

// CreateDosen godoc
// @Summary Create a new Dosen
// @Description Create a new Dosen with the input payload
// @Tags dosen
// @Accept json
// @Produce json
// @Param dosen body models.CreateDosenInput true "Create Dosen"
// @Success 201 {object} models.Dosen
// @Failure 400 {object} map[string]interface{}
// @Router /dosen [post]
func (ctrl *DosenController) CreateDosen(c *gin.Context) {
	var input models.CreateDosenInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	dosen, err := ctrl.service.CreateDosen(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dosen)
}

// GetDosenByID godoc
// @Summary Get a Dosen by ID
// @Description Get a Dosen by its ID
// @Tags dosen
// @Accept json
// @Produce json
// @Param id path int true "Dosen ID"
// @Success 200 {object} models.Dosen
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /dosen/{id} [get]
func (ctrl *DosenController) GetDosenByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid ID"})
		return
	}

	// Convert id to uint if necessary
	uintID := uint(id)

	dosen, err := ctrl.service.GetDosenByID(uintID)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dosen)
}

// UpdateDosen godoc
// @Summary Update a Dosen
// @Description Update a Dosen with the input payload
// @Tags dosen
// @Accept json
// @Produce json
// @Param id path int true "Dosen ID"
// @Param dosen body models.UpdateDosenInput true "Update Dosen"
// @Success 200 {object} models.Dosen
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /dosen/{id} [patch]
func (ctrl *DosenController) UpdateDosen(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid ID"})
		return
	}

	// Convert id to uint if necessary
	uintID := uint(id)

	var input models.UpdateDosenInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	dosen, err := ctrl.service.UpdateDosen(uintID, input)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dosen)
}

// DeleteDosen godoc
// @Summary Delete a Dosen
// @Description Delete a Dosen by its ID
// @Tags dosen
// @Accept json
// @Produce json
// @Param id path int true "Dosen ID"
// @Success 204
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /dosen/{id} [delete]
func (ctrl *DosenController) DeleteDosen(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid ID"})
		return
	}

	// Convert id to uint if necessary
	uintID := uint(id)

	err = ctrl.service.DeleteDosen(uintID)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
