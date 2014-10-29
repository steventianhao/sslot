package engine

import (
	"fmt"
)

type GameCore struct {
	mode    string
	rows    int
	symbols map[string]*Symbol
	reels   []Reel
}

func (g GameCore) nReels() int {
	return len(g.reels)
}

func (g GameCore) nSymbols() int {
	return len(g.symbols)
}

func (g *GameCore) setReels(reels [][]string) error {
	g.reels = make([]Reel, len(reels))
	for i, v := range reels {
		if !checkSymbolNames(g.symbols, v) {
			return fmt.Errorf("symbol name {%s} is not correct", v)
		}
		g.reels[i] = strings2Symbols(g.symbols, v)
	}
	return nil
}

func createGameCore(mode string, rows int, symbols []*Symbol, reels ...[]string) (*GameCore, error) {
	gc := &GameCore{mode: mode, rows: rows}
	gc.symbols = symbols2Map(symbols)
	if err := gc.setReels(reels); err != nil {
		return nil, err
	}
	return gc, nil
}

func (g *GameCore) spin() []Reel {
	return createEngine(g.rows, g.reels...).spin()
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

func makeHitMap(hits []*Hit) map[HitKey]*Hit {
	m := make(map[HitKey]*Hit)
	for _, v := range hits {
		m[v.key()] = v
	}
	return m
}

func caclHitResult(win *Win, hits map[HitKey]*Hit) *HitResult {
	if h, found := hits[win.key()]; found {
		return &HitResult{win, h}
	}
	return nil
}
