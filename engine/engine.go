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

type symbol struct {
	id   int
	name string
	kind int
}

func Ns(id int, name string) *symbol {
	return &symbol{id, name, Normal}
}

func Ss(id int, name string) *symbol {
	return &symbol{id, name, Scatter}
}

func Ws(id int, name string) *symbol {
	return &symbol{id, name, Wild}
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

func Symbols2Map(symbols []*symbol) map[string]*symbol {
	m := make(map[string]*symbol)
	for _, s := range symbols {
		m[s.name] = s
	}
	return m
}

func Strings2Symbols(symbolsMap map[string]*symbol, symbolNames []string) []*symbol {
	result := make([]*symbol, len(symbolNames))
	for i, n := range symbolNames {
		result[i] = symbolsMap[n]
	}
	return result
}

type Reel []*symbol

type engine struct {
	rows  int
	reels []Reel
}

// spin the reels, randomly pick index for each reel, then from that symbol, get consecutive [rows] number of symbols
func (self *engine) Spin() []Reel {
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

func NewEngine(rows int, reels ...Reel) *engine {
	return &engine{rows, reels}
}

type Line []int
type SLine []*symbol

func SymbolOnLines(reels []Reel, lines [][]int) []SLine {
	oneLine := func(line []int) []*symbol {
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
