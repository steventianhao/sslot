package engine

import (
	"fmt"
	"math/rand"
	"time"
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

type Reel []*Symbol

type engine struct {
	rows  int
	reels []Reel
}

// spin the reels, randomly pick index for each reel, then from that symbol, get consecutive [rows] number of symbols
func (self *engine) spin() []Reel {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	fOneReel := func(reel Reel) Reel {
		length := len(reel)
		ri := random.Intn(length)
		r := make(Reel, self.rows)
		for i := 0; i < self.rows; i++ {
			r[i] = reel[(ri+i)%length]
		}
		return r
	}
	result := make([]Reel, len(self.reels))
	for i, reel := range self.reels {
		cs := fOneReel(reel)
		result[i] = cs
	}
	return result
}

func createEngine(rows int, reels ...Reel) *engine {
	return &engine{rows, reels}
}

type Line []int
type SLine []*Symbol

func symbolOnLines(reels []Reel, lines []Line) []SLine {
	oneLine := func(line Line) SLine {
		r := make(SLine, len(line))
		for i, idx := range line {
			r[i] = reels[i][idx]
		}
		return r
	}
	r := make([]SLine, len(lines))
	for i, line := range lines {
		r[i] = oneLine(line)
	}
	return r
}
