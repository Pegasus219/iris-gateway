package main

import (
	"fmt"
	"github.com/kataras/iris"
	"iris-gateway/common"
	"iris-gateway/config"
	"iris-gateway/web/middlewares"
	"iris-gateway/web/routers"
)

func main() {
	localIp := "0.0.0.0"
	listenPort := config.HttpPort
	serverAddr := fmt.Sprintf("%s:%s", localIp, listenPort)

	app := iris.New()
	common.InitLogger(app)

	app.Use(middlewares.PanicHandle())
	routers.Register(app)

	err := app.Run(iris.Addr(serverAddr), iris.WithOptimizations)
	if err != nil {
		panic(fmt.Sprintf("Server Start Error：%v", err))
	}
}
