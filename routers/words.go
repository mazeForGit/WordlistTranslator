package routers

import (
	"strconv"
	"github.com/gin-gonic/gin"
	data "github.com/mazeForGit/WordlistExtractor/data"
	"fmt"
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
	var format string = ""
	var name string = ""
	var testOnly bool = false
	var newOnly bool = false
	
	if _, ok := vars["format"]; ok {
		format = c.Request.URL.Query().Get("format")
	} 
	if _, ok := vars["name"]; ok {
		name = c.Request.URL.Query().Get("name")
	}
	if _, ok := vars["testOnly"]; ok {
		t := c.Request.URL.Query().Get("testOnly")
		if t == "true" {
			testOnly = true
		} 
	}
	if _, ok := vars["newOnly"]; ok {
		t := c.Request.URL.Query().Get("newOnly")
		if t == "true" {
			newOnly = true
		} 
	}
	
	fmt.Println("parsing vars .. format = " + format + ", name = " + name + ", testOnly = " + strconv.FormatBool(testOnly) + ", newOnly = " + strconv.FormatBool(newOnly))
		
	if format != "" && name == "" {
		// complete list by format
		if format == "json" {
			c.JSON(200, data.GlobalWordList.Words)
		} else if format == "csv" {
			c.String(200, data.GetWordsListAsCsv("", testOnly, newOnly))
		} else {
			s = data.Status{Code: 422, Text: "unknown format = " + format}
			c.JSON(422, s)
		}
	} else if format == "" && name != "" {
		// look up by name
		w, err := data.GlobalWordList.GetWordByName(name)
		if (err != nil) {
			s = data.Status{Code: 200, Text: "not found name = " + name}
			c.JSON(200, s)
			return
		}
		c.JSON(200, w)
	} else if format != "" && name != "" {
		// complete list by format and name
		if format == "json" {
			c.JSON(200, data.GetWordsList(name, testOnly, newOnly))
		} else if format == "csv" {
			c.String(200, data.GetWordsListAsCsv(name, testOnly, newOnly))
		}else {
			s = data.Status{Code: 422, Text: "unknown format = " + format}
			c.JSON(422, s)
		}
	} else {
		// default is json
		c.JSON(200, data.GetWordsList(name, testOnly, newOnly))
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
