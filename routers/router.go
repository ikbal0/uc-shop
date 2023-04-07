package routers

import (
	"uc-shop/controllers"
	"uc-shop/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.GET("/", controllers.GetUser)
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middleware.Authentication())
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.PUT("/:productID", middleware.ProductAuthorization(), controllers.UpdateProduct)
	}

	roleRouter := r.Group("/role")
	{
		roleRouter.Use(middleware.Authentication())
		roleRouter.POST("/", middleware.IsRoleAdmin(), controllers.AddRole)
		roleRouter.POST("/set", controllers.SetRole)
	}

	return r
}
