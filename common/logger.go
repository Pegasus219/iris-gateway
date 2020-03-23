package common

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"iris-gateway/common/utils"
	"log"
	"os"
)

func InitLogger(app *iris.Application) {
	golog.SetTimeFormat("2006/01/02 15:04:05.000")
	if debug() {
		log.SetFlags(log.Lshortfile | log.Lmicroseconds)
		golog.SetLevel("debug")
		golog.SetOutput(logPath())
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

func debug() bool {
	debugOpen, err := utils.GetIntEnv("IRIS_DEBUG", 0)
	if err != nil {
		panic(err)
	}
	if debugOpen > 0 {
		return true
	}
	return false
}

func logPath() *os.File {
	logPath := utils.GetEnv("IRIS_LOG_PATH", "")
	if logPath == "" {
		return os.Stdout
	}
	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	return f
}
