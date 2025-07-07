package routes

import (
	"Shopping-cart/config"
	"Shopping-cart/controllers"
	"Shopping-cart/middleware"
	"Shopping-cart/repositories"
	"Shopping-cart/services"

	"github.com/gin-gonic/gin"
)

func SetupRoute(router *gin.RouterGroup) {
	repoUser := repositories.NewUserRepo(config.DB)
	serviceUser := services.NewUserService(repoUser)
	ctrlUser := controllers.NewUserController(serviceUser)

	user := router.Group("/user")
	{
		user.POST("/signup", ctrlUser.SignUp)
		user.POST("/login", ctrlUser.Login)
	}

	repoPro := repositories.NewProductRepo(config.DB)
	servicePro := services.NewProductService(repoPro)
	ctrlPro := controllers.NewProductController(servicePro)
	pro := router.Group("/product")
	pro.Use(middleware.JWTAuthMiddleware())
	{
		pro.POST("/create", ctrlPro.CreateProduct)
		pro.GET("/detail/:id", ctrlPro.GetByID)
		pro.PATCH("/edit/:id", ctrlPro.UpdateProduct)
		pro.DELETE("/delete/:id", ctrlPro.DeleteProduct)
		pro.DELETE("/delete/many", ctrlPro.DeleteMany)
		pro.GET("/list", ctrlPro.ListProduct)
	}

	repoCart := repositories.NewCartRepo(config.DB)
	serviceCart := services.NewCartService(repoCart)
	ctrlCart := controllers.NewCartController(serviceCart)
	cart := router.Group("/cart")
	cart.Use(middleware.JWTAuthMiddleware())
	{
		cart.POST("/add", ctrlCart.AddToCart)
		cart.GET("/list", ctrlCart.ListCart)
		cart.DELETE("/delete/:product_id", ctrlCart.DeleteItem)
		cart.DELETE("/delete/many", ctrlCart.DeleteMany)
	}
}
