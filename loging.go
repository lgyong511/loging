package loging

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"runtime"
	"time"
)

// 配置信息
type Config struct {
	LogLeve    Level       //日志输出级别
	TimeFormat string      //日志输出时间格式
	LogFormat  Format      //日志输出格式
	LogOutput  []io.Writer //日志输出目标
	LogCaller  bool        // 文件名、行号、函数名
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
		LogCaller:  false,
	}}
}

func NewLoging(config *Config) *Loging {
	return &Loging{*config}
}

func (l *Loging) Debug(msg string) {
	if l.LogLeve > Debug {
		//获取日志字符串
		// 根据日志输出格式格式化日志
		// 调用输出目标输出日志
		l.logOutput(l.format(msg))
	}

}

// 根据配置中的时间格式获取当前时间
func (l *Loging) getTime() string {
	return time.Now().Format(l.TimeFormat)
}

// 获取日志所在文件名、行号、函数名。
func (l *Loging) getLogCaller() string {
	// 获取调用栈标识符、带路径的完整文件名、该调用在文件中的行号。如果无法获得信息，ok会被设为false。
	pc, file, line, ok := runtime.Caller(3)
	if !ok {
		return ""
	}
	funcName := runtime.FuncForPC(pc).Name()
	return fmt.Sprintf("%s %d %s\n", file, line, funcName)
}

// 返回对应日志级别的字符串
func (l *Loging) getLevel() string {
	switch l.LogLeve {
	case All:
		return "all"
	case Trace:
		return "tracd"
	case Debug:
		return "debug"
	case Info:
		return "info"
	case Warn:
		return "warn"
	case Error:
		return "error"
	case Fatal:
		return "fatal"
	default:
		return ""
	}
}

// 根据配置拼接日志字符串
func (l *Loging) format(msg string) string {
	//map
	log := make(map[string]string)
	//把日志级别保存到map
	log["level"] = l.getLevel()
	// 判断是否需要文件名、行号、函数名
	if l.LogCaller {
		// 将文件名、行号、函数名保存到map
		log["caller"] = l.getLogCaller()
	}
	// 将消息内容保存到map
	log["msg"] = msg
	// 将时间保存到map
	log["time"] = l.getTime()

	// 返回json串
	if l.LogFormat == Json {
		// 将map序列化成json串
		b, err := json.Marshal(log)
		if err != nil {
			return ""
		}
		return string(b)
	} else { //返回text串
		str := ""
		for k, v := range log {
			str += k + "=" + v
			str += " "
		}
		return str
	}
}

// 输出日志
func (l *Loging) logOutput(log string) {
	for _, w := range l.LogOutput {
		w.Write([]byte(log))
	}
}
