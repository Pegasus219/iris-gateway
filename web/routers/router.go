package routers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"iris-gateway/web/controllers"
	"iris-gateway/web/controllers/demo"
)

func Register(app *iris.Application) {
	//基类控制器
	baseController := new(controllers.BaseController)

	// demo控制器模块
	webApp := mvc.New(app.Party("/"))
	webApp.Register(baseController)
	webApp.Handle(new(demo.IndexController))
}
