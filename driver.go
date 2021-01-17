package driver

type Logger interface {
	Debug(format string, v ...interface{})
	Info(format string, v ...interface{})
	Warning(format string, v ...interface{})
	Error(format string, v ...interface{})
}

type Messager interface {
	OnMessage(packet Packet) error
	SendMessage(packet Packet) error
}

type Packet interface {
	Encoding() (header, pkg []byte)
	Decoding() (header, pkg []byte)
}
