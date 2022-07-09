package handler

import (
	"github.com/gin-gonic/gin"
	"main.go/pkg/service"
)

type Handler struct {
	services *service.Service
}

func (h *Handler) InitRoutes() *gin.Engine {

	router := gin.New()

	api := router.Group("/api")
	{
		api.GET("/get-translate", h.getTranslate)
		api.GET("/get-all-words", h.getAllWord)
	}

	return router
}

func NewHandler(services *service.Service) *Handler {

	return &Handler{
		services: services,
	}
}
