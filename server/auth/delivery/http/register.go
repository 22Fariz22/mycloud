package http

import (
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine,uc auth.UseCase){
  h:=NewHandler(uc)
	authEndpoints := router.Group("/auth")
	{
		authEndpoints.POST("/sign-up",h.Signup)
		authEndpoints.POST("/sign-in",h.Signin)

	}
}