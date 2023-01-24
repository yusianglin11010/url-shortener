package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yusianglin11010/url-shortner/generator"
	"github.com/yusianglin11010/url-shortner/store"
)

type URLCreationReq struct {
	LongUrl string `json:"long_url" blinding:"required"`
}

func CreateShortUrl(c *gin.Context, host, port string) {

	var creationReq URLCreationReq
	if err := c.ShouldBindJSON(&creationReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := generator.GenerateShortLink(creationReq.LongUrl)

	store.SaveUrlMapping(shortUrl, creationReq.LongUrl)
	// TODO: can be replaced with config
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + port + "/" + shortUrl,
	})

}

func HandleShortUrlRedirect(c *gin.Context) {

	shortUrl := c.Param("url")
	initialUrl := store.GetInitialUrl(shortUrl)
	fmt.Println(initialUrl)
	c.Redirect(302, initialUrl)

}
