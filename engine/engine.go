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

type engine struct {
	rows  int
	reels [][]*symbol
}

// spin the reels, randomly pick index for each reel, then from that symbol, get consecutive [rows] number of symbols
func (self *engine) Spin() [][]*symbol {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	fOneReel := func(symbols []*symbol) []*symbol {
		l := len(symbols)
		ri := random.Intn(l)
		r := make([]*symbol, self.rows)
		for i := 0; i < self.rows; i++ {
			r[i] = symbols[(ri+i)%l]
		}
		return r
	}
	result := make([][]*symbol, len(self.reels))
	for i, reel := range self.reels {
		cs := fOneReel(reel)
		result[i] = cs
	}
	return result
}

func NewEngine(rows int, reels ...[]*symbol) *engine {
	return &engine{rows, reels}
}

func SymbolOnLines(screenshot [][]*symbol, lines [][]int) [][]*symbol {
	oneLine := func(line []int) []*symbol {
		r := make([]*symbol, len(line))
		for i, idx := range line {
			r[i] = screenshot[i][idx]
		}
		return r
	}
	r := make([][]*symbol, len(lines))
	for i, line := range lines {
		r[i] = oneLine(line)
	}
	return r
}
