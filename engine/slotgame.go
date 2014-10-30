package engine

import (
	"fmt"
)

type SlotGame struct {
	id          int
	name        string
	nRows       int
	nReels      int
	nLines      int
	gameCores   map[string]*GameCore
	lines       []Line
	normalHits  map[HitKey]*Hit
	wildHits    map[HitKey]*Hit
	scatterHits map[HitKey]*Hit
}

func (g SlotGame) Name() string {
	return g.name
}

func (g SlotGame) Spin(mode string) ([]Reel, error) {
	gc, ok := g.gameCores[mode]
	if !ok {
		return nil, fmt.Errorf("game mode [%s] not found", mode)
	}
	return gc.spin(), nil
}

func CreateGame(id int, name string, nRows, nReels, nLines int) *SlotGame {
	return &SlotGame{id: id, name: name, nRows: nRows, nReels: nReels, nLines: nLines, gameCores: make(map[string]*GameCore)}
}

func (g *SlotGame) SetLines(lines []Line) error {
	nLines := len(lines)
	if nLines != g.nLines {
		return fmt.Errorf("lines specified as %d not match the lines given %d", g.nLines, nLines)
	}
	for _, line := range lines {
		if len(line) != g.nReels {
			return fmt.Errorf("each line's length should be %d", g.nRows)
		}
		for _, v := range line {
			if v >= g.nRows {
				return fmt.Errorf("the index in one line should be less than %d", g.nRows)
			}
		}
	}
	g.lines = lines
	return nil
}

func (g *SlotGame) AddGameCore(mode string, symbols []*Symbol, reels ...[]string) error {
	if len(reels) != g.nReels {
		return fmt.Errorf("reels length should be %d instead of %d", g.nReels, len(reels))
	}
	if gc, err := createGameCore(mode, g.nRows, symbols, reels...); err != nil {
		return err
	} else {
		g.gameCores[mode] = gc
	}
	return nil
}

func (g *SlotGame) AddHits(normalHits, wildHits, scatterHits []*Hit) {
	g.normalHits = makeHitMap(normalHits)
	g.wildHits = makeHitMap(wildHits)
	g.scatterHits = makeHitMap(scatterHits)
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
