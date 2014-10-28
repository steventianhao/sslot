package engine

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Reel []*symbol
type Line []int

type GameBuilder struct {
	nRows, nCols, nLines int
	symbols              map[string]*symbol
	reels                []Reel
	lines                []Line
}

type Game interface {
	MainSpin() [][]*symbol
}

func NewGameBuilder(rows, columns, lines int) *GameBuilder {
	return &GameBuilder{nRows: rows, nCols: columns, nLines: lines}
}

func (g *GameBuilder) SetSymbols(symbols []*symbol) *GameBuilder {
	m := make(map[string]*symbol)
	for _, v := range symbols {
		m[v.name] = v
	}
	g.symbols = m
	return g
}

func (g GameBuilder) str2symbol(str string) (*symbol, error) {
	if s, found := g.symbols[str]; found {
		return s, nil
	}
	es := fmt.Sprint("symbols name is not correct ", str)
	return nil, errors.New(es)
}

func (g *GameBuilder) SetReels(reels ...[]string) (*GameBuilder, error) {
	length := len(reels)
	if length != g.nCols {
		es := fmt.Sprint("columns is ", g.nCols, " but reels columns is ", length)
		return nil, errors.New(es)
	}
	g.reels = make([]Reel, g.nCols)
	for i, v := range reels {
		l := len(v)
		one := make(Reel, l)
		for j, v2 := range v {
			s, err := g.str2symbol(v2)
			if err != nil {
				return nil, err
			}
			one[j] = s
		}
		g.reels[i] = one
	}
	return g, nil
}

func (g *GameBuilder) build() Game {
	return g
}

func (g GameBuilder) MainSpin() [][]*symbol {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	fOneReel := func(symbols []*symbol, nRows int) []*symbol {
		l := len(symbols)
		idx := random.Intn(l)
		r := make([]*symbol, nRows)
		for i := 0; i < nRows; i++ {
			r[i] = symbols[(idx+i)%l]
		}
		return r
	}
	result := make([][]*symbol, len(g.reels))
	for i, reel := range g.reels {
		cs := fOneReel(reel, g.nRows)
		result[i] = cs
	}
	return result
}

func hotLines(screenshot [][]*symbol, lines [][]int) [][]*symbol {
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

type HitKey struct {
	symbol string
	counts int
}

type Hit struct {
	HitKey
	ratio int
}

type Win struct {
	Symbol     *symbol
	Counts     int
	Substitute bool
}

type HitResult struct {
	win *Win
	hit *Hit
}

func NewHit(symbol string, counts int, ratio int) *Hit {
	return &Hit{HitKey{symbol, counts}, ratio}
}

func (nH Hit) key() HitKey {
	return HitKey{nH.symbol, nH.counts}
}

func (w Win) String() string {
	return fmt.Sprint(w.Symbol, w.Counts, w.Substitute)
}

func (w Win) key() HitKey {
	return HitKey{w.Symbol.name, w.Counts}
}

func NewWin(sym *symbol, counts int, wild bool) *Win {
	return &Win{sym, counts, wild}
}

func calcNormalWins(symbols []*symbol) *Win {
	c := 1
	first := *symbols[0]
	wild := first.isWild()
	for _, v := range symbols[1:] {
		s := *v
		if s == first {
			c = c + 1
		} else if s.isWild() {
			c = c + 1
			wild = true
		} else if first.isWild() {
			c = c + 1
			first = s
		} else {
			break
		}
	}
	if c < 2 {
		return nil
	} else if first.isWild() {
		return nil
	} else {
		return NewWin(&first, c, wild)
	}
}

func caclHitResult(win *Win, hits map[HitKey]*Hit) *HitResult {
	if h, found := hits[win.key()]; found {
		return &HitResult{win, h}
	}
	return nil
}

func calcWildWins(symbols []*symbol) *Win {
	c := 1
	first := *symbols[0]
	if !first.isWild() {
		return nil
	}
	for _, v := range symbols[1:] {
		if *v == first {
			c = c + 1
		} else {
			break
		}
	}
	if c < 2 {
		return nil
	} else {
		return NewWin(&first, c, false)
	}
}

/*
	screenshot the current result of 3*5 symbols
	here the premise is just one Scatter symbol for one game
*/
func caclScatterWins(screenshot [][]*symbol) *Win {
	hasScatter := func(strip []*symbol) (bool, *symbol) {
		for _, v := range strip {
			if v.isScatter() {
				return true, v
			}
		}
		return false, nil
	}
	isScatter, first := hasScatter(screenshot[0])
	if !isScatter {
		return nil
	}
	c := 1
	for _, strip := range screenshot[1:] {
		if ok, _ := hasScatter(strip); ok {
			c = c + 1
		} else {
			break
		}
	}
	if c < 2 {
		return nil
	} else {
		return NewWin(first, c, false)
	}
}
