package loging

var (
	std = Default()
)

func UpdateConfig(config *Config) {
	std.UpdateConfig(config)
}

// 自定义一个k/v日志消息
// 如果k和预设的k相同（file,func,level,msg,time），会被覆盖。
func WithField(key, value string) *Loging {
	return std.WithField(key, value)
}

// 自定义多个k/v日志消息
// 如果k和预设的k相同（file,func,level,msg,time），会被覆盖。
func WithFields(m map[string]string) *Loging {
	return std.WithFields(m)
}

// Trace级别日志
func Trace(msg string) {
	std.Trace(msg)
}

// Debug级别日志
func Debug(msg string) {
	std.Debug(msg)
}

// Info级别日志
func Info(msg string) {
	std.Info(msg)
}

// Warn级别日志
func Warn(msg string) {
	std.Warn(msg)
}

// Error级别日志
func Error(msg string) {
	std.Error(msg)
}

// Fatal级别日志，程序退出返回状态码1
func Fatal(msg string) {
	std.Fatal(msg)
}
