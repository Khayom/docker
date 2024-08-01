package controller

import "github.com/gin-gonic/gin"

func Add(c *gin.Context){
	c.JSON(200, "Add Hello")
}

func Get(c *gin.Context){
	c.JSON(200, "Get Hello")
}