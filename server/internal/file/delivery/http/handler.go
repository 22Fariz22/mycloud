package http

import (
	"github.com/22Fariz22/mycloud/server/internal/auth/usecase"
	"github.com/22Fariz22/mycloud/server/internal/file"
	"github.com/gin-gonic/gin"
)

type File struct {
	ID    string `json:"id"`
	File  []byte `json:"file"`
	Title string `json:"title"`
}

type Handler struct {
	uc file.UseCase
}

func NewHandler(uc file.UseCase) *Handler {
	return &Handler{
		uc: uc,
	}
}

type uploadInput struct {
	File  []byte `json:"file"`
	Title string `json:"title"`
}

func (h *Handler) Upload(c *gin.Context) {

}

func (h *Handler) Download(c *gin.Context) {}

func (h *Handler) GetList(c *gin.Context) {}

func (h *Handler) Share(c *gin.Context) {}

func (h *Handler) Delete(c *gin.Context) {}
