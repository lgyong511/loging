package main

import (
	"io"
	"loging"
	"os"
)

var log *loging.Loging

func logDefault() {
	log = loging.Default()
	log.Debug("使用默认配置创建")
}

func logConfig() {

	file, err := os.OpenFile("./dame.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 644)
	if err != nil {
		file = os.Stdout
	}

	var config loging.Config
	config.LogLeve = loging.Debug
	config.LogFormat = loging.Json
	config.TimeFormat = "2006-01-02 15:03:04"
	config.LogCaller = true
	config.LogOutput = []io.Writer{os.Stdout, file}

	log = loging.NewLoging(&config)

	log.Debug("使用配置文件创建")

}

func main() {
	logDefault()
	// logConfig()
}
