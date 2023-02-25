package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"dot-connect/handler"
	"dot-connect/docs"
)

func main() {
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"

	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	// docs.SwaggerInfo.Title = "Dot Connect API"
	v1 := r.Group("/api/v1")
	{
		reports := v1.Group("/reports")
		{
			reports.POST("/upload", handler.PostReport)
			reports.GET("/my/:userId", handler.GetUserReport)
			reports.GET("/:reportId", handler.GetDetailReport)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(":8080") // listen and serve on
}