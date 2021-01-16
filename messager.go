package driver

type Messager interface {
	OnMessage(pkg []byte) error
	Close()
}
