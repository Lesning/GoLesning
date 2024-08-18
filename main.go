package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main()  {
	fmt.Println("hello")
	r := gin.Default()
	r.GET("/zhang",func (c *gin.Context)  {
		c.JSON(http.StatusOK, gin.H{
			"messg" : "ping",
		})
	})
	r.Run(":8999")
	
}