package controllers

import (
	"github.com/gin-gonic/gin"
	"go-server/services"
)

type NewController struct {

}

var News NewController

func (self NewController) News(c *gin.Context){
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	//test := services.TestService{}

	//test := services.Services.Test

	//c.JSON(200, services.Services.Test)

	//services.Services.Test()

	// 서비스에서 제이슨 형식 데이터 받아와서 뿌려주기
	//c.JSON(200, gin.H{
	//	//"message": "new",
	//	test
	//})
	//c.JSON(200, services.Services.Test)
	c.JSON(200, gin.H{
		"message" : services.RandString(),
	})
}
