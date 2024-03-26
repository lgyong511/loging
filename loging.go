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

// 日志信息
type Loging struct {
	config  *Config
	logMap  map[string]string //保存日志信息
	logByte []byte            //根据配置序列化后的日志信息
	skip    int               //记录函数栈调用次数
}

//---------------------------可以导出函数开始---------------------------

// 创建一个默认配置的Loging实例
func Default() *Loging {
	return &Loging{&Config{
		LogLeve:    All,
		TimeFormat: "2006-01-02 15:04:05",
		LogFormat:  Json,
		LogOutput:  []io.Writer{os.Stdout},
		LogCaller:  true,
	}, nil, nil, 0}
}

// 创建一个指定配置信息的Loging实例
func NewLoging(config *Config) *Loging {
	return &Loging{config, nil, nil, 0}
}

// 向日志消息自定义一个k/v
// 如果k和预设的k相同（file,func,level,msg,time），会被覆盖。
func (l *Loging) WithField(key, value string) *Loging {
	return l.initMap().addMap(key, value)
}

// 向日志消息自定义多个k/v
// 如果k和预设的k相同（file,func,level,msg,time），会被覆盖。
func (l *Loging) WithFields(m map[string]string) *Loging {
	l.initMap()
	for k, v := range m {
		l.addMap(k, v)
	}

	return l
}

// Trace级别日志
func (l *Loging) Trace(msg string) {
	defer l.clear()

	l.skip++
	if l.config.LogLeve <= Trace {
		l.initMap().format(Trace, msg).logOutput()
	}

}

// Debug级别日志
func (l *Loging) Debug(msg string) {
	defer l.clear()

	l.skip++
	if l.config.LogLeve <= Debug {
		l.initMap().format(Debug, msg).logOutput()
	}

}

// Info级别日志
func (l *Loging) Info(msg string) {
	defer l.clear()

	l.skip++
	if l.config.LogLeve <= Info {
		l.initMap().format(Info, msg).logOutput()
	}

}

// Warn级别日志
func (l *Loging) Warn(msg string) {
	defer l.clear()

	l.skip++
	if l.config.LogLeve <= Warn {
		l.initMap().format(Warn, msg).logOutput()
	}

}

// Error级别日志
func (l *Loging) Error(msg string) {
	defer l.clear()

	l.skip++
	if l.config.LogLeve <= Error {
		l.initMap().format(Error, msg).logOutput()
	}

}

// Fatal级别日志，程序退出返回状态码1
func (l *Loging) Fatal(msg string) {
	defer os.Exit(1)

	l.skip++
	if l.config.LogLeve <= Fatal {
		l.initMap().format(Fatal, msg).logOutput()
	}

	l.clear()
}

//---------------------------可以导出函数结束---------------------------

//---------------------------不可导出函数开始---------------------------

// 清理日志信息
func (l *Loging) clear() {
	// 重置日志信息
	l.logByte = nil
	l.logMap = nil
	l.skip = 0
}

// 向日志添加一个k/v
func (l *Loging) addMap(key, value string) *Loging {
	l.logMap[key] = value
	return l
}

// 根据配置中的时间格式获取当前时间
func (l *Loging) getTime() string {
	return time.Now().Format(l.config.TimeFormat)
}

// 获取日志所在文件名、行号、函数名。
func (l *Loging) getLogCaller() {
	l.skip++

	// 判断是否需要文件名、行号、函数名
	if l.config.LogCaller {
		// 将文件名、行号、函数名保存到map
		// 获取调用栈标识符、带路径的完整文件名、该调用在文件中的行号。如果无法获得信息，ok会被设为false。
		pc, file, line, ok := runtime.Caller(l.skip)
		if ok {
			funcName := runtime.FuncForPC(pc).Name()
			l.addMap("file", fmt.Sprintf("%s:%d", file, line))
			l.addMap("func", funcName)
		}
	}

}

// 返回对应日志级别的字符串
func (l *Loging) getLevel(level Level) string {
	switch level {
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
func (l *Loging) format(level Level, msg string) *Loging {
	l.skip++
	//把日志级别保存到map
	l.addMap("level", l.getLevel(level))

	l.getLogCaller()

	// 将消息内容保存到map
	l.addMap("msg", msg)
	// 将时间保存到map
	l.addMap("time", l.getTime())

	// 返回json串
	if l.config.LogFormat == Json {
		// 将map序列化成json串
		var err error
		l.logByte, err = json.Marshal(l.logMap)
		if err != nil {
			return l
		}
		// 追加换行
		b := []byte("\n")
		l.logByte = append(l.logByte, b...)
		return l
	} else { //返回text串
		str := ""
		for k, v := range l.logMap {
			str += k + "=" + `"` + v + `"`
			str += " "
		}
		// 追加换行
		str += "\n"
		l.logByte = []byte(str)
		return l
	}

}

// 输出日志
func (l *Loging) logOutput() {
	for _, w := range l.config.LogOutput {
		w.Write(l.logByte)
	}
}

// 初始化map
func (l *Loging) initMap() *Loging {
	// 判断map是否初始化
	if l.logMap == nil {
		// 初始化map
		l.logMap = make(map[string]string)
	}
	return l
}

//---------------------------不可导出函数结束---------------------------
