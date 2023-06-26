package main

import (
	"dounis/appdata"
	"dounis/auth"
	"dounis/docs"
	"dounis/users"
	"fmt"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Dounis API
// @version 1.0
// @description Protech

// @contact.name Marcel
// @contact.email marceltournesol@gmail.com

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	r := gin.Default()

	// Swagger for ez api documentation
	docs.SwaggerInfo.BasePath = "/"

	// Shared data for all of the controllers (eg. DB, sockets, ...)
	app_data := appdata.New()

	// Controllers
	auth_controller := auth.New(app_data)
	users_controller := users.New(app_data)

	auth_controller.Attach(r)
	users_controller.Attach(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	if err := r.Run(":6942"); err != nil {
		fmt.Println(err)
	}
}
