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

	chatid, err := h.services.Patients.GetChatId(input.PatientId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("failed getchatid%s", err.Error()))
	}
	var message string
	if input.Pulse == 0 {
		message = fmt.Sprintf("pulse: %d,\nRespiratoryRate: %d,\nOxygenSaturation: %f,\nSystolicBloodPressure: %d,\nDiastolicBloodPressure: %d,\nHeartRate: %d,\nAnalysisTimestamp: %s\n\n\nВы умираете, вызовите скорую ", input.Pulse, input.RespiratoryRate, input.OxygenSaturation, input.SystolicBloodPressure, input.DiastolicBloodPressure, input.HeartRate, input.AnalysisTimestamp)
	} else {
		message = fmt.Sprintf("pulse: %d,\nRespiratoryRate: %d,\nOxygenSaturation: %f,\nSystolicBloodPressure: %d,\nDiastolicBloodPressure: %d,\nHeartRate: %d,\nAnalysisTimestamp: %s\n\n\nВаш показатель находится в нормальном диапазоне. Это говорит о хорошем состоянии сердечно-сосудистой системы в состоянии покоя.\nВаш показатель находится в пределах нормы. Это свидетельствует о нормальной функции легких.\nВаш показатель превосходный. Это указывает на хорошее насыщение крови кислородом.\nВаше кровяное давление идеально и соответствует нормам.Рекомендация: Поддерживайте оптимальный вес, избегайте чрезмерного потребления соли, алкоголя и жирной пищи.\nВаш показатель в норме, что указывает на стабильную работу сердечной мышцы.Рекомендация: Сохраняйте умеренную физическую активность, например, прогулки или плавание.", input.Pulse, input.RespiratoryRate, input.OxygenSaturation, input.SystolicBloodPressure, input.DiastolicBloodPressure, input.HeartRate, input.AnalysisTimestamp)
	}

	botURL := "https://api.telegram.org/bot7951143788:AAFPuqecSG-VVeM6IwavsdWmU9oV5W7-wKg/sendMessage"
	msg := Message{
		ChatID: chatid, // Преобразуем int в string
		Text:   message,
	}
	jsonData, err := json.Marshal(msg)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to marshal message:")
		return
	}
	resp, err := http.Post(botURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Failed to send data to bot: "+err.Error())
		return
	}
	defer resp.Body.Close()

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllAnalysis(c *gin.Context) {
	patientId, err := strconv.Atoi(c.Param("patient_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid patientId id param")
		return
	}
	analysis, err := h.services.Analysis.GetAll(patientId)
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
