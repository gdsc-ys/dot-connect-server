package handler

import (
	"net/http"
	// "path/filepath"

	"github.com/gin-gonic/gin"
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
func GetUserReport(c *gin.Context) {
	userId := c.Param("userId")
	// user id 로 db 조회
	c.String(http.StatusOK, "userId: %s", userId)
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
