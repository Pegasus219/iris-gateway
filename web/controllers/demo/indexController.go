package demo

import (
	"github.com/kataras/iris"
	"iris-gateway/web/controllers"
)

type (
	IndexController struct {
		Ctx iris.Context
		*controllers.BaseController
	}

	rspType struct {
		Get  string `json:"getParam"`
		Post string `json:"postParam"`
	}
)

//测试GET
func (c *IndexController) GetIndex() {
	param := c.Ctx.URLParamDefault("param", "")
	rsp := &rspType{
		Get: param,
	}
	c.RenderSuccess(&c.Ctx, rsp)
}

//测试POST
func (c *IndexController) PostIndex() {
	param := c.Ctx.PostValueDefault("param", "")
	rsp := rspType{
		Post: param,
	}
	c.RenderSuccess(&c.Ctx, rsp)
}

//测试panic
func (c *IndexController) GetPanic() {
	panic("测试panic恢复")
}
