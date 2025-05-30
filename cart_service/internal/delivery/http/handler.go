package http

import (
	"net/http"
	"strconv"

	"github.com/da-er-gaultier/cart_service/internal/domain"
	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	usecase domain.CartUsecase
}

func NewCartHandler(r *gin.Engine, uc domain.CartUsecase) {
	handler := &CartHandler{usecase: uc}

	r.POST("/cart", handler.Add)
	r.GET("/cart/:user_id", handler.Get)
	r.DELETE("/cart/:id", handler.Delete)
}

func (h *CartHandler) Add(c *gin.Context) {
	var item domain.CartItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.usecase.AddToCart(&item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, item)
}

func (h *CartHandler) Get(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
		return
	}

	items, err := h.usecase.GetCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if items == nil {
		items = []domain.CartItem{}
	}

	c.JSON(http.StatusOK, items)
}

func (h *CartHandler) Delete(c *gin.Context) {
	itemID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid item id"})
		return
	}
	if err := h.usecase.RemoveFromCart(itemID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
