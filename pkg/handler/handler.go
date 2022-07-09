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

	router.GET("/get-trnslate", h.getTranslate)

	return router
}

func NewHandler(services *service.Service) *Handler {

	return &Handler{
		services: services,
	}
}
