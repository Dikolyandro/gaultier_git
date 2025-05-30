package http

import (
	"net/http"
	"strconv"

	"github.com/da-er-gaultier/admin_service/internal/domain"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	usecase domain.AdminUsecase
}

func New(r *gin.Engine, u domain.AdminUsecase) {
	h := &AdminHandler{usecase: u}

	r.GET("/admin/users", h.GetUsers)
	r.DELETE("/admin/users/:id", h.DeleteUser)

	r.GET("/admin/products", h.GetProducts)
	r.DELETE("/admin/products/:id", h.DeleteProduct)

	r.GET("/admin/orders", h.GetOrders)
	r.PATCH("/admin/orders/:id", h.UpdateOrderStatus)
}

func (h *AdminHandler) GetUsers(c *gin.Context) {
	users, err := h.usecase.ListUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *AdminHandler) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.usecase.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete user"})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *AdminHandler) GetProducts(c *gin.Context) {
	products, err := h.usecase.ListProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *AdminHandler) DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.usecase.DeleteProduct(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete product"})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *AdminHandler) GetOrders(c *gin.Context) {
	orders, err := h.usecase.ListOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch orders"})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (h *AdminHandler) UpdateOrderStatus(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	if err := h.usecase.UpdateOrderStatus(id, input.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not update order status"})
		return
	}
	c.Status(http.StatusOK)
}
