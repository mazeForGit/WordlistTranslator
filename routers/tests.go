package routers

import (
	"github.com/gin-gonic/gin"
	data "github.com/mazeForGit/WordlistExtractor/data"
)

func TestsGET(c *gin.Context) {

	var vars map[string][]string
	vars = c.Request.URL.Query()
	if _, ok := vars["name"]; ok {
		// different format
		name := c.Request.URL.Query().Get("name")

		var list []string
		list = data.GetTestsList(name)

		if len(list) == 0 {
			c.JSON(422, gin.H{
				"status": "unprocessable entity name = " + name,
			})
			return
		}

		c.JSON(200, list)
		return
	}

	c.JSON(200, data.GlobalWordList.Tests)
}
