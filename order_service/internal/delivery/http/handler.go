package http

import (
	"net/http"
	"strconv"

	"github.com/da-er-gaultier/order_service/internal/domain"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	usecase domain.OrderUsecase
}

func NewOrderHandler(r *gin.Engine, uc domain.OrderUsecase) {
	h := &OrderHandler{usecase: uc}

	r.POST("/orders", h.Create)
	r.GET("/orders/:id", h.Get)
	r.GET("/orders/user/:id", h.GetByUser)
	r.DELETE("/orders/:id", h.Delete)
}

func (h *OrderHandler) Create(c *gin.Context) {
	var order domain.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order.Status = "created"
	if err := h.usecase.Create(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, order)
}

func (h *OrderHandler) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	order, err := h.usecase.Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if order == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) GetByUser(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))
	orders, err := h.usecase.GetByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.usecase.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
