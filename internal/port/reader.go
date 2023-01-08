package port

type Reader interface {
	Read() ([]string, error)
}
