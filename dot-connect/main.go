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
	"log"
  
	firebase "firebase.google.com/go"
	// "google.golang.org/api/option"
	"context"
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
		AllowOrigins:     []string{"https://dot-connect-r6ge.vercel.app", "http://localhost:3000"},
		AllowMethods:     []string{"POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	  }))

	// r.GET("/firebase", handler.GetFirebaseToken)
	ctx := context.Background() 
	conf := &firebase.Config{ProjectID: "dot-connect-374203"}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
	log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
	log.Fatalln(err)
	}
	defer client.Close()

	reports := r.Group("/reports")
	{
		reports.POST("/upload", handler.PostReport)
		reports.GET("/my/:userId", handler.GetUserReport(client))
		reports.GET("/:reportId", handler.GetDetailReport)
	}
	r.GET("/translation", handler.TranslateToBraille)
	r.POST("/translation/braille", handler.TranslateToKorean)
	setupSwagger(r)

	r.Run(":8080") // listen and serve on
}