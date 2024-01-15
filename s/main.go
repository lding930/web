package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/connect", func(c *gin.Context) {
		address := c.PostForm("address")
		fmt.Println("Connected Address:", address)
		c.JSON(http.StatusOK, gin.H{"address": address})
	})

	router.GET("/switch-network", func(c *gin.Context) {
		// You can add logic here to initiate network switch in MetaMask
		// For demonstration purposes, we'll return a message.
		c.JSON(http.StatusOK, gin.H{"message": "Switch network request sent to MetaMask"})
	})

	fmt.Println("Server running on port 80")
	router.Run(":80")
}

