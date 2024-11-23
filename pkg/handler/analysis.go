package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	smartmed "github.com/didsqq/SmartMed_monitoring_system"
	"github.com/gin-gonic/gin"
)

type Message struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

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

	message := fmt.Sprintf("pulse: %d,\nRespiratoryRate: %d,\nOxygenSaturation: %f,\nSystolicBloodPressure: %d,\nDiastolicBloodPressure: %d,\nHeartRate: %d,\nAnalysisTimestamp: %s", input.Pulse, input.RespiratoryRate, input.OxygenSaturation, input.SystolicBloodPressure, input.DiastolicBloodPressure, input.HeartRate, input.AnalysisTimestamp)
	// Отправка POST-запроса на bot
	botURL := "https://api.telegram.org/bot7951143788:AAFPuqecSG-VVeM6IwavsdWmU9oV5W7-wKg/sendMessage"
	msg := Message{
		ChatID: "1357553243",
		Text:   message,
	}
	jsonData, err := json.Marshal(msg)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to marshal message:")
		return
	}
	resp, err := http.Post(botURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Failed to send data to Flask: "+err.Error())
		return
	}
	defer resp.Body.Close()

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
