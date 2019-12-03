package routers

import (
	"strconv"
	"github.com/gin-gonic/gin"
	data "github.com/mazeForGit/WordlistExtractor/data"
)
func WordListDELETE(c *gin.Context) {
	var s data.Status

	outputstring := "removed .."
	w, err := strconv.ParseBool(c.Request.URL.Query().Get("words"))
	if err != nil {
		s = data.Status{Code: 422, Text: "unprocessable entity, expecting words=true/false"}
		c.JSON(422, s)
		return
	}
	
	t, err := strconv.ParseBool(c.Request.URL.Query().Get("tests")) 
	if err != nil {
		s = data.Status{Code: 422, Text: "unprocessable entity, expecting tests=true/false"}
		c.JSON(422, s)
		return
	}
	
	if w {
		removeWordsCount := strconv.Itoa(len(data.GlobalWordList.Words))
		outputstring += " words=" + removeWordsCount
	}
	if t {
		removeTestsCount := strconv.Itoa(len(data.GlobalWordList.Tests))
		outputstring += " tests=" + removeTestsCount
	}
	
	data.GlobalWordList = data.Clear(data.GlobalWordList, w, t)
	s = data.Status{Code: 200, Text: outputstring}
	c.JSON(200, s)
}
func WordListGET(c *gin.Context) {
	c.JSON(200, data.GlobalWordList)
}
func WordListPUT(c *gin.Context) {
	var s data.Status

	removeCount := strconv.Itoa(len(data.GlobalWordList.Words))
	err := c.Bind(&data.GlobalWordList)
	if err != nil {
		s = data.Status{Code: 422, Text: "unprocessable entity"}
		c.JSON(422, s)
		return
	}
	data.GlobalWordList = data.RebuildTestIndex(data.GlobalWordList)
	replaceCount := strconv.Itoa(len(data.GlobalWordList.Words))
	s = data.Status{Code: 200, Text: "replaced items = " + removeCount + " by " + replaceCount}
	c.JSON(200, s)
}
