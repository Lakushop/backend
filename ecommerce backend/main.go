package main

import (
	"log"
	"os"

	"github.com/lakushop/backend/controllers"
	"github.com/lakushop/backend/middleware"
	"github.com/lakushop/backend/routes"
	"github.com/lakushop/backend/database"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Gerenv("PORT")
	if port == "" {
		port = "9090"
	}

	app := controllers.NewApp(database.productdata(database.client, "products"), database.userdata(database.client, "users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes()(router)
	router.Use(middleware.Authenticate())

	roter.Get("/admin/addproduct", app.AddToCart())
	router.Get("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.Paycart())
	router.GET("/instantbuy", app.instantbuy())

	log.Fatal(router.Run(":" + port))
}
