package hasher

import "fmt"

type Mock struct {
}

func NewMock() *Mock {
	return &Mock{}
}

func (m *Mock) Hash(s string) string {
	return fmt.Sprintf("<%s>", s)
}
