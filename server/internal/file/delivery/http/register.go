package http

import (
	"github.com/22Fariz22/mycloud/server/internal/file"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(r *gin.RouterGroup, uc file.UseCase) {
	h := NewHandler(uc)

	files := r.Group("/file")
	{
		files.POST("", h.Upload)
		files.GET("", h.Download)
		files.GET("", h.GetList)
		files.GET("", h.Share)
		files.DELETE("", h.Delete)
	}
}
