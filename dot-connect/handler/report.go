package handler

import (
	"context"
	"dot-connect/model"
	"log"
	"net/http"
	"strconv"

	// "path/filepath"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
	"google.golang.org/genproto/googleapis/type/latlng"
)

// PostReport
// @Summary      에러 신고 업로드
// @Description  에러 신고에 이미지, 장소, 내용이 담긴다.
// @Tags         reports
//	@Param		message	formData   model.Report	true	"Report Info"
// @Produce      json
// @Success      200   {object} model.Report
// @Router       /reports/upload [post]
func PostReport(client *firestore.Client) gin.HandlerFunc{
	fn := func(c *gin.Context) {
		ctx := context.Background()
		r := model.Report{}
		
		latitude, _ := strconv.ParseFloat(c.PostForm("Latitude"), 64)
		longtitude, _ := strconv.ParseFloat(c.PostForm("Longtitude"), 64)
		content := c.PostForm("content")
		r.CONTENT = content
		r.LATITUDE = latitude
		r.LONGTITUDE = longtitude
		_, _, err := client.Collection("reports").Add(ctx, map[string]interface{}{
			"location": latlng.LatLng{Latitude: latitude, Longitude: longtitude},
			"content": content,
		})
		if err != nil {
			// Handle any errors in an appropriate way, such as returning them.
			log.Printf("An error has occurred: %s", err)
		}

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
//	@Success		200	{object}	model.Report
//	@Router			/reports [get]
func GetReports(client *firestore.Client) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		ctx := context.Background()
		iter := client.Collection("reports").Documents(ctx)
		
		r := []model.Report{}
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Fatalf("Failed to iterate: %v", err)
			}
			mapData := doc.Data()
			report := model.Report{}
			location := mapData["location"].(*latlng.LatLng)

			report.CONTENT = mapData["content"].(string)
			report.LATITUDE =  location.GetLatitude()
			report.LONGTITUDE = location.GetLongitude()
			r = append(r, report)
		}
		c.JSON(http.StatusOK, r)
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
func GetDetailReport(c *gin.Context) {
	reportId := c.Param("reportId")
	// report id 로 db 조회
	c.String(http.StatusOK, "reportId: %s", reportId)
}

func GetFirebaseToken(c *gin.Context) {

	c.String(http.StatusOK, "firebase token")
}
