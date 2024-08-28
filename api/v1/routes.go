package v1

import (
	"github.com/dreadster3/goquotes/api/v1/quotes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	quotesGroup := router.Group("/quotes")
	quotes.RegisterRoutes(quotesGroup)
}
