package engine

import "fmt"

const (
	Normal  = 0
	Scatter = 1
	Wild    = 2
)

type symbol struct {
	id   int
	kind int
	name string
}

func Ns(id int, name string) *symbol {
	return create(id, name, Normal)
}

func Ss(id int, name string) *symbol {
	return create(id, name, Scatter)
}

func Ws(id int, name string) *symbol {
	return create(id, name, Wild)
}

func create(id int, name string, kind int) *symbol {
	return &symbol{id, kind, name}
}

func (s symbol) String() string {
	var k string
	switch s.kind {
	default:
		k = "Normal"
	case 1:
		k = "Scatter"
	case 2:
		k = "Wild"
	}
	return fmt.Sprint("id:", s.id, ",kind:", k, ",name:", s.name)
}

func (s symbol) isWild() bool {
	return s.kind == Wild
}

func (s symbol) isScatter() bool {
	return s.kind == Scatter
}
