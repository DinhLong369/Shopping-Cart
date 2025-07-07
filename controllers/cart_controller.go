package controllers

import (
	"Shopping-cart/models"
	"Shopping-cart/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CartController struct {
	service services.CartService
}

func NewCartController(service services.CartService) *CartController {
	return &CartController{service}
}

func (ctrl *CartController) AddToCart(c *gin.Context) {
	idUser, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusBadGateway, gin.H{"err": "Unauthorized"})
		return
	}
	userID := idUser.(uint)
	var input models.Request
	if err := c.ShouldBind(&input); err != nil || input.Quantity == 0 {
		c.JSON(http.StatusBadGateway, gin.H{"err": err.Error()})
	}
	if err := ctrl.service.AddToCart(userID, input.ProductID, input.Quantity); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"err": err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{"message": "da them vao gio hang thanh cong"})
	}
}

func (ctrl *CartController) ListCart(c *gin.Context) {
	idUser, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusBadGateway, gin.H{"err": "Unauthorized"})
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	userID := idUser.(uint)
	list_cart, err := ctrl.service.ListItems(userID, page, limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"list_cart": list_cart,
		})
	}
}

func (ctrl *CartController) DeleteItem(c *gin.Context) {
	_, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusBadGateway, gin.H{"err": "Unauthorized"})
		return
	}
	productID, _ := strconv.Atoi(c.Param("product_id"))
	if err := ctrl.service.DeleteItem(uint(productID)); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "xoa san pham trong gio hang thanh cong"})
}

func (ctrl *CartController) DeleteMany(c *gin.Context) {
	_, exist := c.Get("userID")
	logrus.Println("============================")
	if !exist {
		c.JSON(http.StatusBadGateway, gin.H{"err": "Unauthorized"})
		return
	}

	var input models.IdsProduct
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"err": err.Error()})
		return
	}
	logrus.Println("===============lay gia tri IDS:", input.IdsProduct)
	err := ctrl.service.DeleteMany(input.IdsProduct)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"err": err.Error()})
	}
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "xoa san pham trong gio hang thanh cong"})
	}
}
