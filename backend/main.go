package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func TimeNow() string {
	timestr := time.Now()
	newtime := timestr.Format(time.RFC3339)[0:10] + " " + timestr.Format(time.RFC3339)[11:19]
	return newtime
}
func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/api/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": TimeNow(),
		})
		fmt.Println("Button Clicked")
	})

	r.Run(":8000")
}
