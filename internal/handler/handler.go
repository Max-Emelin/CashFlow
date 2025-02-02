package handler

import (
	"cashflow/internal/service"

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

	api := router.Group("/api")
	{
		lists := api.Group("balance")
		{
			lists.POST("/deposit", h.deposit)
			lists.POST("/transfer", h.transfer)
		}
		items := api.Group("transactions")
		{
			items.GET("/", h.getLastTransactions)
		}
	}

	return router
}
