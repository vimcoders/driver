package driver

type Logger interface {
	Debug(format string, v ...interface{})
	Info(format string, v ...interface{})
	Warning(format string, v ...interface{})
	Error(format string, v ...interface{})
	Close() (err error)
}

type Header interface {
	Version() uint8
	Length() uint32
	Protocol() uint16
	ToBytes(payload []byte) []byte
}

type Message interface {
	Header() Header
	Payload() []byte
}
