package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	go func() {
		for {
			// Replace this block with data write to central-db
			fmt.Println("RAPIDSERVICE EMITTING DATA!!")
			time.Sleep(5 * time.Second)
		}
	}()

	server.Run(":8011")
}
