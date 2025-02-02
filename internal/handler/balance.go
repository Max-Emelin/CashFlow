package handler

import (
	"cashflow/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) deposit(c *gin.Context) {
	var input model.TransactionInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoItem.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
