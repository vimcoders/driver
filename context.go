package driver

type Context interface {
	Header() int32
	Unmarshal(v interface{}) error
}
