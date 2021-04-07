package v1

import (
	"github.com/DrScorpio/ri-blog/pkg/app"
	"github.com/DrScorpio/ri-blog/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

func (A Article) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}
func (A Article) List(c *gin.Context)   {}
func (A Article) Creat(c *gin.Context)  {}
func (A Article) Update(c *gin.Context) {}
func (A Article) Delete(c *gin.Context) {}
