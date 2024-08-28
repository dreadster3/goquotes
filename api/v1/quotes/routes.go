package quotes

import (
	"github.com/dreadster3/goquotes/pkg/services"
	"github.com/gin-gonic/gin"
)

func GetQuotes(c *gin.Context) {
	quotes, err := services.Quote.GetQuotes()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"quotes": quotes})
}

func GetRandomQuote(c *gin.Context) {
	quote, err := services.Quote.GetRandomQuote()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"quote": quote})
}

func CreateQuote(c *gin.Context) {
	var req struct {
		Quote string `json:"quote"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	quote, err := services.Quote.CreateQuote(req.Quote)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"quote": quote})
}

func CreateQuotes(c *gin.Context) {
	var req struct {
		Quotes []string `json:"quotes"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	for _, quote := range req.Quotes {
		_, err := services.Quote.CreateQuote(quote)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(200, gin.H{"message": "Quotes created successfully"})
}

func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("", GetQuotes)
	router.GET("/random", GetRandomQuote)
	router.POST("", CreateQuote)
	router.PATCH("", CreateQuotes)
}
