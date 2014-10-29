package engine

import (
	"fmt"
)

type Game struct {
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

func (g Game) Name() string {
	return g.name
}

func (g Game) Spin(mode string) ([]Reel, error) {
	gc, ok := g.gameCores[mode]
	if !ok {
		return nil, fmt.Errorf("game mode [%s] not found", mode)
	}
	return gc.spin(), nil
}

func CreateGame(id int, name string, nRows, nReels, nLines int) *Game {
	return &Game{id: id, name: name, nRows: nRows, nReels: nReels, nLines: nLines, gameCores: make(map[string]*GameCore)}
}

func (g *Game) SetLines(lines []Line) error {
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

func (g *Game) AddGameCore(mode string, symbols []*Symbol, reels ...[]string) error {
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

func (g *Game) AddHits(normalHits, wildHits, scatterHits []*Hit) {
	g.normalHits = MakeHitMap(normalHits)
	g.wildHits = MakeHitMap(wildHits)
	g.scatterHits = MakeHitMap(scatterHits)
}
