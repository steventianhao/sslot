package engine

import (
	"fmt"
	"math/rand"
	"time"
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

type Reel []*Symbol

type spinEngine struct {
	rows  int
	reels []Reel
}

func createEngine(rows int, reels ...Reel) *spinEngine {
	return &spinEngine{rows, reels}
}

// spin the reels, randomly pick index for each reel, then from that symbol, get consecutive [rows] number of symbols
func (self *spinEngine) spin() []Reel {
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
