package loging

// 日志级别类型
type Level int

// 日志级别常量
const (
	All Level = iota
	Trace
	Debug
	Info
	Warn
	Error
	Fatal
)

// 日志输出格式
type Format int

// json和text格式
// json：{"file":"E:/git管理代码仓库/github/loging/main/main.go:12","func":"main.initF","level":"debug","msg":"初始化函数开始工作了","time":"2024-03-23 20:08:34"}
//text：level="debug" file="E:/git管理代码仓库/github/loging/main/main.go:12" func="main.initF" msg="初始化函数开始工作了" time="2024-03-23 20:08:36"
const (
	Json Format = iota
	Text
)
