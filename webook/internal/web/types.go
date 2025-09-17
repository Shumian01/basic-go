package web

import "github.com/gin-gonic/gin"

type headler interface {
	RegisterRoutes(server *gin.Engine)
}
