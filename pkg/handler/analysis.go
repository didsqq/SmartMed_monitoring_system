package handler

import (
	"net/http"

	smartmed "github.com/didsqq/SmartMed_monitoring_system"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createAnalysis(c *gin.Context) {
	var userId int
	var input smartmed.Analysis
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Analysis.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllAnalysis(c *gin.Context) {

}

func (h *Handler) getAnalysisById(c *gin.Context) {

}

func (h *Handler) updateAnalysis(c *gin.Context) {

}

func (h *Handler) deleteAnalysis(c *gin.Context) {

}
