package driver

type Logger interface {
	Debug(format string, v ...interface{})
	Info(format string, v ...interface{})
	Warning(format string, v ...interface{})
	Error(format string, v ...interface{})
}

type Messager interface {
	OnMessage(message Message) error
	SendMessage(message Message) error
}

type Message interface {
	Encoding() (header, message []byte, err error)
	Decoding() (header, message []byte, err error)
}
