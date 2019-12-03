package routers

import (
	"strconv"
	"github.com/gin-gonic/gin"
	data "github.com/mazeForGit/WordlistExtractor/data"
	//"fmt"
	//"io/ioutil"
)

func WordsByIdDELETE(c *gin.Context) {
	var s data.Status

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		s = data.Status{Code: 422, Text: "unprocessable entity id = " + strconv.Itoa(id)}
		c.JSON(422, s)
		return
	}

	data.GlobalWordList, err = data.GlobalWordList.DeleteWordById(id)
	if (err != nil) {
		s = data.Status{Code: 200, Text: "not found id = " + strconv.Itoa(id)}
		c.JSON(200, s)
		return
	}
	s = data.Status{Code: 200, Text: "entity deleted"}
	c.JSON(200, s)
}
func WordsByIdGET(c *gin.Context) {
	var s data.Status

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		s = data.Status{Code: 422, Text: "unprocessable entity id = " + strconv.Itoa(id)}
		c.JSON(422, s)
		return
	}

	w, err := data.GlobalWordList.GetWordById(id)
	if (err != nil) {
		s = data.Status{Code: 200, Text: "not found id = " + strconv.Itoa(id)}
		c.JSON(200, s)
		return
	}
	c.JSON(200, w)
}
func WordsGET(c *gin.Context) {
	var s data.Status
	var vars map[string][]string

	vars = c.Request.URL.Query()
	if _, ok := vars["format"]; ok {
		// different format
		format := c.Request.URL.Query().Get("format")
		
		if format == "json" {
			c.JSON(200, data.GlobalWordList.Words)
		} else if format == "csv" {
			c.String(200, data.GetWordsListAsCsv())
		} else {
			s = data.Status{Code: 422, Text: "unknown format = " + format}
			c.JSON(422, s)
		}
	} else {
		// default is json
		c.JSON(200, data.GlobalWordList.Words)
	}
}
func WordsPOST(c *gin.Context) {
	var s data.Status
	var wrd data.Word
	
	err := c.BindJSON(&wrd)
	if err != nil || wrd.Id != 0 || len(wrd.Name) == 0 || wrd.New == false {
		s = data.Status{Code: 422, Text: "unprocessable entity"}
		c.JSON(422, s)
		return
	}
	
	data.GlobalWordList, err = data.AddWordToList(data.GlobalWordList, wrd.Name)
	if err != nil {
		s = data.Status{Code: 409, Text: "entity already exists"}
		c.JSON(409, s)
		return
	}

	s = data.Status{Code: 200, Text: "entity added"}
	c.JSON(200, s)
}
