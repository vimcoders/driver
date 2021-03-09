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
}

type Message interface {
	Header() (header Header)
	ToBytes() (header, pkg []byte)
}
