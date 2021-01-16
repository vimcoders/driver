package driver

type Buffer interface {
	Take(lenth int) (buffer []byte)
}
