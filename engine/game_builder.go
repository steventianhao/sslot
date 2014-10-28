package engine

import (
	"errors"
	"fmt"
)

type GameBuilder struct {
	nRows, nCols, nLines int
	symbols              map[string]*symbol
	reels                []Reel
	lines                []Line
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

func (g *GameBuilder) build() *engine {
	return NewEngine(g.nRows, g.reels...)
}

type Hit struct {
	HitKey
	ratio      int
	features   int
	multiplier int
}

type HitResult struct {
	win *Win
	hit *Hit
}

func NewHit(symbol string, counts int, ratio int) *Hit {
	return &Hit{HitKey{symbol, counts}, ratio, 0, 0}
}

func NewFeatureHit(symbol string, counts, ratio, features, multiplier int) *Hit {
	return &Hit{HitKey{symbol, counts}, ratio, features, multiplier}
}

func (nH Hit) key() HitKey {
	return nH.HitKey
}

func caclHitResult(win *Win, hits map[HitKey]*Hit) *HitResult {
	if h, found := hits[win.key()]; found {
		return &HitResult{win, h}
	}
	return nil
}
