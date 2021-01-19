package driver

type Logger interface {
	Debug(format string, v ...interface{})
	Info(format string, v ...interface{})
	Warning(format string, v ...interface{})
	Error(format string, v ...interface{})
}

type Messager interface {
	OnMessage(b Buffer) error
	SendMessage(b Buffer) error
}

type Buffer interface {
	Read() (b []byte, err error)
	Write(b ...[]byte) (err error)
}
