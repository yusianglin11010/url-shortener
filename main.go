package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yusianglin11010/url-shortner/config"
	"github.com/yusianglin11010/url-shortner/handler"
	"github.com/yusianglin11010/url-shortner/store"
)

func main() {
	r := gin.Default()
	config := config.NewConfig()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello URL Shortner"})
	})

	r.POST("/url", func(c *gin.Context) {
		handler.CreateShortUrl(c, config.Host, config.Port)
	})

	r.GET("/:url", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	store.InitializeStore()

	if err := r.Run(config.Port); err != nil {
		log.Print(err)
	}
}
