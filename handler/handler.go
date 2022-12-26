package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yusianglin11010/url-shortner/generator"
	"github.com/yusianglin11010/url-shortner/store"
)

type URLCreationReq struct {
	LongUrl string `json:"long_url" blinding:"required"`
	UserID  string `json:"user_id" blinding:"required"`
}

func CreateShortUrl(c *gin.Context) {

	var creationReq URLCreationReq
	if err := c.ShouldBindJSON(&creationReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := generator.GenerateShortLink(creationReq.LongUrl, creationReq.UserID)
	store.SaveUrlMapping(shortUrl, creationReq.LongUrl, creationReq.UserID)

	// TODO: can be replaced with config
	host := "http://localhost:9808/"
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})

}

func HandleShortUrlRedirect(c *gin.Context) {

	shortUrl := c.Param("url")
	initialUrl := store.GetInitialUrl(shortUrl)
	c.Redirect(302, initialUrl)

}
