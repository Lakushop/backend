package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lakushop/backend/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controllers.Signup())
	incomingRoutes.POST("users/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.productviewadmin())
	incomingRoutes.GET("/users/productview", controllers.productview())
	incomingRoutes.GET("/users/search", controllers.searchproductbyquery())
}
