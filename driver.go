package driver

type Logger interface {
	Debug(format string, v ...interface{})
	Info(format string, v ...interface{})
	Warning(format string, v ...interface{})
	Error(format string, v ...interface{})
	Close() (err error)
}

type Header interface {
	Version() (version uint8)
	Length() (length uint32)
	Protocol() (protocol uint16)
	CheckSum() (code uint32)
	ToBytes() (header []byte)
}

type Message interface {
	Header() (header Header)
	Body() (body []byte)
}

type Messager interface {
	Identify() uint64
	OnMessage(message Message) (err error)
	PushMessage(message Message) (err error)
}
