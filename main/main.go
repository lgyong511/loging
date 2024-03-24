package main

import (
	"io"
	"loging"
	"os"
)

var log *loging.Loging

func logDefault() {
	log = loging.Default()

	log.Trace("使用默认配置创建的Trace")
	log.Debug("使用默认配置创建的Debug")
	log.Info("使用默认配置创建的Info")
	log.Warn("使用默认配置创建的Warn")
	log.Error("使用默认配置创建的Error")
	// log.Fatal("使用默认配置创建的Fatal")

	log.WithField("field", "用field添加").Debug("测试")
	log.WithFields(map[string]string{"fields1": "fields添加", "fields2": "fields添加"}).Info("fields")
}

func logConfig() {

	file, err := os.OpenFile("./dame.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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

	log.Trace("使用配置文件创建的Trace")
	log.Debug("使用配置文件创建的Debug")
	log.Info("使用配置文件创建的Info")
	log.Warn("使用配置文件创建的Warn")
	log.Error("使用配置文件创建的Error")
	log.Fatal("使用配置文件创建的Fatal")

}

func main() {
	logDefault()
	// logConfig()
}
