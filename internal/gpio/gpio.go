package gpio

type Pin interface {
	Read() bool
	Write(value bool) error
}
