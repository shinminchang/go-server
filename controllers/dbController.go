package controllers

import (
	"github.com/gin-gonic/gin"
	"go-server/services"
)

type DbContoller struct {

}

var Dbtest DbContoller

func (self DbContoller) Test(c *gin.Context){
	c.JSON(200, gin.H{
		"message" : services.MemberSelect(),
	})
	//services.MemberSelect()
}