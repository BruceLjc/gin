package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	r := gin.Default()

	// 处理CORS问题，允许所有来源
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
		}
	})

	r.POST("/register", func(c *gin.Context) {

		user := User{}
		if err := c.ShouldBind(&user); err != nil {
			c.String(400, "give me user")
			log.Fatal("cao")
			return
		}

		fmt.Printf("user: %v\n",  user)
		//........
		if !yanzheng1() {
			c.JSON(400, "failed")
			return
		}
		c.JSON(200, "ok")
	})

	r.POST("/login", func(c *gin.Context) {
		user := User{}
		if err := c.ShouldBind(&user); err != nil {
			c.String(400, "give me user")
			return
		}

		fmt.Println(user)

		//...
		if !yanzheng2() {
			c.JSON(400, "failed")
			return
		}
		c.JSON(200, "ok")
	})
	r.Run(":80")
}

func yanzheng1() bool {
	return false
}

func yanzheng2() bool {
	return true
}
