package api

import (
	v1 "github.com/dreadster3/goquotes/api/v1"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	v1Group := router.Group("/v1")
	v1.RegisterRoutes(v1Group)
}
