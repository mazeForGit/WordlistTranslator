package routers

import (
	//"strconv"
	
	"github.com/gin-gonic/gin"
	
	data "github.com/mazeForGit/WordlistTranslator/model"
)
func ConfigGET(c *gin.Context) {
	c.JSON(200, model.GlobalConfig)
}
func ConfigPUT(c *gin.Context) {
	var s model.Status
	
	err := c.BindJSON(&model.GlobalConfig)
	if err != nil {
		s = model.Status{Code: 422, Text: "unprocessable entity"}
		c.JSON(422, s)
		return
	}
	
	s = model.Status{Code: 200, Text: "entity added"}
	c.JSON(200, s)
}
func ConfigPOST(c *gin.Context) {
	var s model.Status
	var vars map[string][]string
	vars = c.Request.URL.Query()
	var execution string = ""

	if _, ok := vars["execution"]; ok {
		execution = c.Request.URL.Query().Get("execution")
	} 
	
	//fmt.Println("execution = " + execution)
	if execution == "true" {
		
		if (model.GlobalConfig.WordToStartWith != "" && model.GlobalConfig.WordToStartWithNext == "") {
			model.GlobalConfig.WordToStartWithNext = model.GlobalConfig.WordToStartWith
		}
		if (model.GlobalConfig.WordToStartWith != "" && model.GlobalConfig.WordToStartWithNext != "") {
			model.GlobalConfig.RequestExecution = true
		} else {
			s = model.Status{Code: 422, Text: "missing data"}
			c.JSON(200, s)
			return
		}
		
		s = model.Status{Code: 200, Text: "start execution"}
		c.JSON(200, s)
	} else if execution == "false" {
		model.GlobalConfig.RequestExecution = false
		
		s = model.Status{Code: 200, Text: "stop execution"}
		c.JSON(200, s)
	} else {
		s = model.Status{Code: 422, Text: "unknown request"}
		c.JSON(422, s)
	}
}
