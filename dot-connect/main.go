package main

import (
	"dot-connect/docs"
	"dot-connect/handler"
	"net/http"
	"os"
	"time"

	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	firebase "firebase.google.com/go"
	"context"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
  
	return os.Getenv(key)
  }

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
		MaxAge:           12 * time.Hour,
	}))

	// r.GET("/firebase", handler.GetFirebaseToken)
	ctx := context.Background()
	project_id := goDotEnvVariable("FIRE_BASE_PROJECT_ID")
	conf := &firebase.Config{ProjectID: project_id}
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
		reports.POST("/upload", handler.PostReport(client))
		reports.GET("/", handler.GetReports(client))
	}
	setupSwagger(r)

	r.Run(":8080") // listen and serve on
}
