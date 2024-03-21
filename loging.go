package loging

import (
	"io"
	"os"
)

// 配置信息
type Config struct {
	LogLeve    Level       //日志输出级别
	TimeFormat string      //日志输出时间格式
	LogFormat  Format      //日志输出格式
	LogOutput  []io.Writer //日志输出目标
	FileName   bool        // 是否输出文件名称
	FuncName   bool        //是否输出函数名称

}

type Loging struct {
	Config
}

func Default() *Loging {
	return &Loging{Config{
		LogLeve:    Info,
		TimeFormat: "2006-01-02 15:03:04",
		LogFormat:  Json,
		LogOutput:  []io.Writer{os.Stdout},
		FileName:   false,
		FuncName:   false,
	}}
}

func NewLoging(config *Config) *Loging {
	return &Loging{*config}
}

func (l *Loging) Debug(msg string) {
	//获取日志字符串
	// 根据日志输出格式格式化日志
	// 调用输出目标输出日志
}
