package handler

import (
	"net/http"

	smartmed "github.com/didsqq/SmartMed_monitoring_system"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createPatients(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input smartmed.Patient
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Patients.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// type getAllPatientsResponse struct {
// 	Data []smartmed.Patient `json:"data"`
// }

func (h *Handler) getAllPatients(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	patients, err := h.services.Patients.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// c.JSON(http.StatusOK, getAllPatientsResponse{
	// 	patients,
	// })
	c.JSON(http.StatusOK, patients)
}

func (h *Handler) getPatientsById(c *gin.Context) {

}
