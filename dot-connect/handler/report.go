package handler

import (
	"context"
	"dot-connect/model"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

// PostReport
// @Summary      에러 신고 업로드
// @Description  에러 신고에 이미지, 장소, 내용이 담긴다.
// @Tags         reports
//	@Param		message	formData   model.ReportForm	true	"Report Info"
// @Produce      json
// @Success      200   {object} model.ReportForm
// @Router       /reports/upload [post]
func PostReport(client *firestore.Client) gin.HandlerFunc{
	fn := func(c *gin.Context) {
		ctx := context.Background()
		r := model.ReportForm{}
		
		latitude, _ := strconv.ParseFloat(c.PostForm("latitude"), 64)
		longtitude, _ := strconv.ParseFloat(c.PostForm("longtitude"), 64)
		content := c.PostForm("content")
		r.CONTENT = content
		r.LATITUDE = latitude
		r.LONGTITUDE = longtitude

		_, _, err := client.Collection("reports").Add(ctx, map[string]interface{}{
			"latitude": latitude,
			"longitude": longtitude,
			"content": content,
		})
		if err != nil {
			// Handle any errors in an appropriate way, such as returning them.
			log.Printf("An error has occurred: %s", err)
		}

	c.JSON(http.StatusOK, r)
	}
	return gin.HandlerFunc(fn)
}

// ListMyReport godoc
//
//	@Summary		List reports
//	@Description	get MyReports
//	@Tags			reports
//	@Produce		json
//	@Success		200	{object}	model.ReportForm
//	@Router			/reports [get]
func GetReports(client *firestore.Client) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		ctx := context.Background()
		iter := client.Collection("reports").Documents(ctx)
		
		r := []model.ReportForm{}
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Fatalf("Failed to iterate: %v", err)
			}
			mapData := doc.Data()
			report := model.ReportForm{}

			report.CONTENT = mapData["content"].(string)
			report.LATITUDE = mapData["latitude"].(float64)
			report.LONGTITUDE = mapData["longitude"].(float64) 
			r = append(r, report)
		}
		c.JSON(http.StatusOK, r)
	}
	return gin.HandlerFunc(fn)
}


func GetFirebaseToken(c *gin.Context) {

	c.String(http.StatusOK, "firebase token")
}
