package main

import (
	docs "dot-connect/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"dot-connect/handler"
)

func main() {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	docs.SwaggerInfo.Title = "Dot Connect API"
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.POST("/upload", handler.PostReport)

	r.GET("/my-report/:userId", handler.GetUserReport)

	r.GET("/report/:reportId", handler.GetDetailReport)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(":8080") // listen and serve on
}