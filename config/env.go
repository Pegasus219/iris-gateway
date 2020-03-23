package config

import "iris-gateway/common/utils"

var (
	//web端口号
	HttpPort string

	//是否开启debug日志
	IsDebugOpen bool

	//本地日志文件路径
	LogFilePath string
)

func init() {
	HttpPort = utils.GetEnv("IRIS_HTTP_PORT", "80")
	LogFilePath = utils.GetEnv("IRIS_LOG_PATH", "")
	IsDebugOpen = getIsDebug()
}

func getIsDebug() bool {
	debugOpen, err := utils.GetIntEnv("IRIS_DEBUG", 0)
	if err != nil {
		panic(err)
	}
	if debugOpen > 0 {
		return true
	}
	return false
}
