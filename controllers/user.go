package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tiamxu/pecker/models"
)

func Register(c *gin.Context) {

}

func LoginGet(c *gin.Context) {
	// models.InsertUser(username, password)
	username := c.Query("username")
	password := c.Query("password")
	fmt.Println("UserName:", username, "Password:", password)
	users, err := models.QueryUserWithParam(username, password)
	if err != nil {
		fmt.Println("用户不存在")
	}
	fmt.Println("users:", users)

}
func LoginPost(c *gin.Context) {
	// username := c.PostForm("username")
	// password := c.PostForm("password")
	username := c.Param("username")
	password := c.Param("password")
	fmt.Println("UserName:", username, "Password:", password)
	users, _ := models.QueryUserWithParam(username, password)
	fmt.Println("users:", users)

}
