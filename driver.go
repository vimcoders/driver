package driver

type Logger interface {
	Debug(format string, v ...interface{})
	Info(format string, v ...interface{})
	Warning(format string, v ...interface{})
	Error(format string, v ...interface{})
}

type Messager interface {
	OnMessage(pkg Package) error
	SendMessage(pkg Package) error
}

type Packer interface {
	Package() (header, body []byte)
}

type Unpacker interface {
	Unpackage() (protocol uint16, body []byte)
}
