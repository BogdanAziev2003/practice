package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/pkg/models"
)

func (h *Handler) getTranslate(c *gin.Context) {
	var input models.Word
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	translate, err := h.services.Translate(input.Origin)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"origin":    input.Origin,
		"translate": translate,
	})
}
