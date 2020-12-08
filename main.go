package main

import (
	"xpertise-go/initialize"

	"xpertise-go/docs"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Xpertise Scholar Golang Backend
// @version 1.0
// @description Xpertise Scholar
// @schemes http https

func main() {
	docs.SwaggerInfo.Title = "Xpertise Scholar"
	docs.SwaggerInfo.Description = "This is Xpertise Scholar's Golang backend."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	err := initialize.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer initialize.Close()

	r := initialize.SetupRouter()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
