package handler

import (
	"net/http"
	"strconv"

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
	analysis, err := h.services.Analysis.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, analysis)
}

func (h *Handler) getAnalysisById(c *gin.Context) {

	analysisId, err := strconv.Atoi(c.Param("analysis_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid analysis id param")
	}

	analysis, err := h.services.Analysis.GetById(analysisId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, analysis)
}
