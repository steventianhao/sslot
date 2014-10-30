package engine

import (
	"fmt"
)

const (
	Normal  = 0
	Scatter = 1
	Wild    = 2
)

type Symbol struct {
	id   int
	name string
	kind int
}

func Ns(id int, name string) *Symbol {
	return &Symbol{id, name, Normal}
}

func Ss(id int, name string) *Symbol {
	return &Symbol{id, name, Scatter}
}

func Ws(id int, name string) *Symbol {
	return &Symbol{id, name, Wild}
}

func (s Symbol) String() string {
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

func (s Symbol) isWild() bool {
	return s.kind == Wild
}

func (s Symbol) isScatter() bool {
	return s.kind == Scatter
}

func symbols2Map(symbols []*Symbol) map[string]*Symbol {
	m := make(map[string]*Symbol)
	for _, s := range symbols {
		m[s.name] = s
	}
	return m
}

func checkSymbolNames(symbolsMap map[string]*Symbol, symbolNames []string) bool {
	for _, n := range symbolNames {
		_, ok := symbolsMap[n]
		if !ok {
			return false
		}
	}
	return true
}

func strings2Symbols(symbolsMap map[string]*Symbol, symbolNames []string) []*Symbol {
	result := make([]*Symbol, len(symbolNames))
	for i, n := range symbolNames {
		result[i] = symbolsMap[n]
	}
	return result
}
