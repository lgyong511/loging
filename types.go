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

type Format int

const (
	Json Format = iota
	Text
)
