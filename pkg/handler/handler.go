package handler

import (
	"github.com/didsqq/SmartMed_monitoring_system/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		patients := api.Group("patients")
		{
			patients.POST("/", h.createPatients)
			patients.GET("/", h.getAllPatients)
			patients.GET("/:patient_id", h.getPatientsById)

			analysis := patients.Group(":patient_id/analysis")
			{
				analysis.POST("/", h.createAnalysis)
				analysis.GET("/", h.getAllAnalysis)
				analysis.GET("/:analysis_id", h.getAnalysisById)
			}
		}
	}

	// fanzil := router.Group("fanzil")
	// {
	// 	patients := fanzil.Group("patients")
	// 	{
	// 		patients.POST("/", h.createPatients)
	// 		patients.GET("/", h.getAllPatients)
	// 		patients.GET("/:patient_id", h.getPatientsById)

	// 		analysis := patients.Group(":patient_id/analysis")
	// 		{
	// 			analysis.POST("/", h.createAnalysis)
	// 			analysis.GET("/", h.getAllAnalysis)
	// 			analysis.GET("/:analysis_id", h.getAnalysisById)
	// 		}
	// 	}
	// }

	return router
}
