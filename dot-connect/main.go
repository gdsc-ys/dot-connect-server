package main

import (
	"dot-connect/docs"
	"dot-connect/handler"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupSwagger(r *gin.Engine) {
    r.GET("/", func(c *gin.Context) {
        c.Redirect(http.StatusFound, "/swagger/index.html")
    })

    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func main() {
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"

	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	// docs.SwaggerInfo.Title = "Dot Connect API"

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://dot-connect-r6ge.vercel.app"},
		AllowMethods:     []string{"POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	  }))


	reports := r.Group("/reports")
	{
		reports.POST("/upload", handler.PostReport)
		reports.GET("/my/:userId", handler.GetUserReport)
		reports.GET("/:reportId", handler.GetDetailReport)
	}
	r.GET("/translation", handler.TranslateToBraille)
	r.POST("/translation/braille", handler.TranslateToKorean)
	setupSwagger(r)

	r.Run(":8080") // listen and serve on
}