package controllers

import (
	"github.com/gin-gonic/gin"
)

type NewController struct {

}

var News NewController

func (self NewController) News(c *gin.Context){
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	c.JSON(200, gin.H{
		"message": "new",
	})
}
