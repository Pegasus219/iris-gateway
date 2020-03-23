package common

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"iris-gateway/config"
	"log"
	"os"
)

func InitLogger(app *iris.Application) {
	golog.SetTimeFormat("2006/01/02 15:04:05.000")
	if config.IsDebugOpen {
		log.SetFlags(log.Lshortfile | log.Lmicroseconds)
		golog.SetLevel("debug")
		golog.SetOutput(logFile())
		//打开请求日志
		customLogger := logger.New(logger.Config{
			Status:  true,
			IP:      true,
			Method:  true,
			Path:    true,
			Query:   true,
			Columns: false,
		})
		app.Use(customLogger)
	} else {
		golog.SetLevel("warn")
	}
}

func logFile() *os.File {
	logPath := config.LogFilePath
	if logPath == "" {
		return os.Stdout
	}
	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	return f
}
