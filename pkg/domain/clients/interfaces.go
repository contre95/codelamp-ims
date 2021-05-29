package clients

type Logger interface {
	Info(format string, i ...interface{})
	Warn(format string, i ...interface{})
	Err(format string, i ...interface{})
	Debug(format string, i ...interface{})
}
