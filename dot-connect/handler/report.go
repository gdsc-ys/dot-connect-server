package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	// "path/filepath"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

// PostReport
// @Summary      에러 신고 업로드
// @Description  에러 신고에 이미지, 장소, 내용이 담긴다.
// @Tags         reports
//	@Param			message	formData		model.PostReport	true	"Report Info"
// @Produce      json
// @Success      200   {object} model.PostReport
// @Router       /reports/upload [post]
func PostReport(c *gin.Context) {
	location := c.PostForm("location")
	locationDetail := c.PostForm("locationDetail")
	content := c.PostForm("content")

	// file, err := c.FormFile("file")
	// if err != nil {
	// 	c.String(http.StatusBadRequest, "get form err: %s", err.Error())
	// 	return
	// }

	// // filename := filepath.Base(file.Filename)
	// filename := filepath.Join("media", file.Filename)
	// if err := c.SaveUploadedFile(file, filename); err != nil {
	// 	c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
	// 	return
	// }

	// c.String(http.StatusOK, "upload file success: %s, %s, %s, %s", location, locationDetail, content, file.Filename)
	c.String(http.StatusOK, "upload file success: %s, %s, %s", location, locationDetail, content)
}

// ListMyReport godoc
//
//	@Summary		List reports
//	@Description	get MyReports
//	@Tags			reports
//	@Param			userId	path		int	true	"user ID"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.Report
//	@Router			/reports/my/{userId} [get]
func GetUserReport(client *firestore.Client) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		userId := c.Param("userId")
		ctx := context.Background()
		iter := client.Collection("reports").Documents(ctx);

		for {
        doc, err := iter.Next()
        if err == iterator.Done {
                break
        }
        if err != nil {
                log.Fatalf("Failed to iterate: %v", err)
        }
        fmt.Println(doc.Data())
}
		// user id 로 db 조회
		c.String(http.StatusOK, "userId: %s", userId)
	}
	return gin.HandlerFunc(fn)
}

// ListMyReport godoc
//
//	@Summary		get report detail	
//	@Description	get one of report
//	@Tags			reports
//	@Accept			json
//	@Produce		json
//	@Param			reportId	path		int	true	"report ID"
//	@Success		200	{object}	model.Report	
//	@Router			/reports/{reportId} [get]
func GetDetailReport(c *gin.Context){
	reportId := c.Param("reportId")
	// report id 로 db 조회
	c.String(http.StatusOK, "reportId: %s", reportId)
}

func GetFirebaseToken(c *gin.Context) {
	
	c.String(http.StatusOK, "firebase token")
}