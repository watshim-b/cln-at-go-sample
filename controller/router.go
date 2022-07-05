package controller

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.GET("ping", ping)
	r.GET("login", Login)
	r.GET("user", GetUser)
	return r
}
