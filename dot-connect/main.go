package main

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.POST("/upload", func(c *gin.Context) {
		location := c.PostForm("location")
		content := c.PostForm("content")

		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, "get form err: %s", err.Error())
			return
		}

		// filename := filepath.Base(file.Filename)
		filename := filepath.Join("media", file.Filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
			return
		}

		c.String(http.StatusOK, "upload file success: %s, %s, %s", location, content, file.Filename)
	})

	r.GET("/my-report/:userId", func(c *gin.Context) {
		userId := c.Param("userId")
		// user id 로 db 조회
		c.String(http.StatusOK, "userId: %s", userId)
	})

	r.GET("/report/:reportId", func(c *gin.Context) {
		reportId := c.Param("reportId")
		// report id 로 db 조회
		c.String(http.StatusOK, "reportId: %s", reportId)
	})

	r.Run(":8080") // listen and serve on
}