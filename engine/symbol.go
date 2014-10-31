package engine

const (
	Normal  = 0
	Scatter = 1
	Wild    = 2
)

type Symbol struct {
	Id   int
	Name string
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
	return s.Name
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
		m[s.Name] = s
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

//if two symbols with same name, return two symbols back
//if not exist, then nil will be in the result slice
func strings2Symbols(symbolsMap map[string]*Symbol, symbolNames []string) []*Symbol {
	result := make([]*Symbol, len(symbolNames))
	for i, n := range symbolNames {
		result[i] = symbolsMap[n]
	}
	return result
}
