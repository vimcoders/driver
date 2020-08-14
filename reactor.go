package driver

import "io"

type Listener interface {
	Listen(addr string) error
}

type Processor interface {
	Read() error
	io.WriteCloser
}
