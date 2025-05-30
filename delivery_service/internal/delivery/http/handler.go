package http

import (
	"net/http"
	"strconv"

	"github.com/da-er-gaultier/delivery_service/internal/domain"
	"github.com/gin-gonic/gin"
)

type DeliveryHandler struct {
	usecase domain.DeliveryUsecase
}

func NewDeliveryHandler(r *gin.Engine, uc domain.DeliveryUsecase) {
	h := &DeliveryHandler{usecase: uc}

	r.POST("/deliveries", h.Create)
	r.PATCH("/deliveries/:id", h.UpdateStatus)
	r.GET("/deliveries/:id", h.GetByID)
	r.GET("/deliveries/order/:id", h.GetByOrder)
}

func (h *DeliveryHandler) Create(c *gin.Context) {
	var d domain.Delivery
	if err := c.ShouldBindJSON(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	d.Status = "packed"
	if err := h.usecase.Create(&d); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, d)
}

func (h *DeliveryHandler) UpdateStatus(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.usecase.UpdateStatus(id, input.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (h *DeliveryHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	d, err := h.usecase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if d == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, d)
}

func (h *DeliveryHandler) GetByOrder(c *gin.Context) {
	orderID, _ := strconv.Atoi(c.Param("id"))
	d, err := h.usecase.GetByOrderID(orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if d == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, d)
}
