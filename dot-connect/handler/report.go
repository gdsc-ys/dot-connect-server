package handler

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)


func PostReport(c *gin.Context) {
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
}

func GetUserReport(c *gin.Context) {
	userId := c.Param("userId")
	// user id 로 db 조회
	c.String(http.StatusOK, "userId: %s", userId)
}

func GetDetailReport(c *gin.Context){
	reportId := c.Param("reportId")
	// report id 로 db 조회
	c.String(http.StatusOK, "reportId: %s", reportId)
}
