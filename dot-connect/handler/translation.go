package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BrailleMsg struct {
	Str string `json:"str"`
	Braille []string `json:"braille"`
}

type RawBrailleBody struct {
	Braille string `json:"braille"`
}

// translation-kor             
// @Summary      한글 -> 점자 점역 
// @Description  한글를 점자로 점역하는 API 
// @Tags         braille 
// @Produce      json
// @Router       /translation [get] 
func TranslateToBraille(c *gin.Context) {
	kor_word := c.DefaultQuery("kor", "안녕")

	braille_hexcodes := [][]int{{0x2823, 0x2812}, {0x2809, 0x283B}}

	msg := map[string][]int{}	

	for i, kor_letter := range kor_word {
		if i == len(kor_word) {
			break
		}
		array_index := int(i / 3)
		msg[string(kor_letter)] = braille_hexcodes[array_index]
	}

	c.JSON(http.StatusOK, msg)
}


// translation-braille             
// @Summary      점자 -> 한글 역점역 
// @Description  점자를 점자로 역점역하는 API 
// @Tags         braille 
// @Produce      json
// @Router       /translation/braille [post] 
func TranslateToKorean(c *gin.Context) {
	var requestBody RawBrailleBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	braille := requestBody.Braille

	fmt.Println(braille)

	// trnalate braile to korean	

	kor_word := "안녕"

	braille_hexcodes := [][]int{{0x2823, 0x2812}, {0x2809, 0x283B}}

	msg := map[string][]int{}	

	for i, kor_letter := range kor_word {
		if i == len(kor_word) {
			break
		}
		array_index := int(i / 3)
		msg[string(kor_letter)] = braille_hexcodes[array_index]
	}

	c.JSON(http.StatusOK, msg)
}