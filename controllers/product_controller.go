package controllers

import (
	"Shopping-cart/models"
	"Shopping-cart/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service services.ProductService
}

func NewProductController(service services.ProductService) *ProductController {
	return &ProductController{service}
}

func (ctrl *ProductController) CreateProduct(c *gin.Context) {
	_, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusBadGateway, gin.H{"err": "Unauthorized"})
		return
	}
	var input models.Product
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err := ctrl.service.CreateProduct(&input)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{"message": "tao san pham thanh cong"})
		return
	}
}

func (ctrl *ProductController) GetByID(c *gin.Context) {
	_, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusBadGateway, gin.H{"err": "Unauthorized"})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	result, err := ctrl.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"err": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": result})
		return
	}
}

func (ctrl *ProductController) UpdateProduct(c *gin.Context) {
	_, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusBadGateway, gin.H{"err": "Unauthorized"})
		return
	}
	var update models.CreateProduct
	id, _ := strconv.Atoi(c.Param("id"))
	if err := c.ShouldBind(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	err := ctrl.service.UpdateProduct(id, &update)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "cap nhat san pham thanh cong"})
		return
	}
}

func (ctrl *ProductController) DeleteProduct(c *gin.Context) {
	_, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusBadGateway, gin.H{"err": "Unauthorized"})
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	err := ctrl.service.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "xoa san pham thanh cong"})
		return
	}
}

func (ctrl *ProductController) ListProduct(c *gin.Context) {
	_, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusBadGateway, gin.H{"err": "Unauthorized"})
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	items, total, err := ctrl.service.ListProduct(page, limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data":  items,
			"total": total,
		})
		return
	}
}

func (ctrl *ProductController) DeleteMany(c *gin.Context) {
	_, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusBadGateway, gin.H{"err": "Unauthorized"})
		return
	}
	var input models.Ids
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"err": err.Error()})
		return
	}
	err := ctrl.service.DeleteMany(input.Ids)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"err": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "xoa san pham thanh cong"})
		return
	}

}
