package v1

import (
	"github.com/gin-gonic/gin"
	"go-depot/global"
	"go-depot/pkg/app"
	"go-depot/pkg/errcode"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {}

/**

 */
func (t Tag) List(c *gin.Context) {
	param := struct {
		Name  string `form:"name" binding:"max=100"`
		State uint8  `form:"state,default=1" binding:"oneof=0 1"`
	}{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	response.ToResponse(gin.H{})
}
func (t Tag) Create(c *gin.Context) {}
func (t Tag) Update(c *gin.Context) {}
func (t Tag) Delete(c *gin.Context) {}
