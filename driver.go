package driver

type Logger interface {
	Debug(format string, v ...interface{})
	Info(format string, v ...interface{})
	Warning(format string, v ...interface{})
	Error(format string, v ...interface{})
	Close() (err error)
}

type Header interface {
	Protocol() uint16
}

type Message interface {
	Header
	ToBytes() (header, payload []byte)
}
