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
	normalHits  map[Key]*Hit
	wildHits    map[Key]*Hit
	scatterHits map[Key]*Hit
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

type Key struct {
	Symbol string
	Counts int
}

type Hit struct {
	Key
	ratio      int
	features   int
	multiplier int
}

type HitResult struct {
	win *Win
	hit *Hit
}

// if there's subsitute in line, then ratio multiply 2
func (hr HitResult) ratio() int {
	ratio := hr.hit.ratio
	if hr.win.Substitute {
		return ratio * 2
	}
	return ratio
}

func NewHit(symbol string, counts int, ratio int) *Hit {
	return &Hit{Key{symbol, counts}, ratio, 0, 0}
}

func NewFeatureHit(symbol string, counts, ratio, features, multiplier int) *Hit {
	return &Hit{Key{symbol, counts}, ratio, features, multiplier}
}

func (h Hit) key() Key {
	return h.Key
}

func makeHitMap(hits []*Hit) map[Key]*Hit {
	m := make(map[Key]*Hit)
	for _, v := range hits {
		m[v.key()] = v
	}
	return m
}

func caclHitResult(win *Win, hits map[Key]*Hit) *HitResult {
	if h, found := hits[win.key()]; found {
		return &HitResult{win, h}
	}
	return nil
}

type ScatterWin struct {
	Ratio, Features, Multiplier int
}

type LineWin struct {
	LineId int
	Ratio  int
	Win    *Win
}

func (lw LineWin) String() string {
	return fmt.Sprint("{lineId:", lw.LineId+1, ",", lw.Win, ",ratio:", lw.Ratio, "}")
}

type SpinResult struct {
	reels      []Reel
	lineWins   []*LineWin
	scatterWin *ScatterWin
}

func (sr SpinResult) String() string {
	return fmt.Sprint("reels:", sr.reels, "lineWins:", sr.lineWins, "scatterWin:", sr.scatterWin)
}

func (g SlotGame) SpinResult(mode string) (*SpinResult, error) {
	reels, err := g.Spin(mode)
	if err != nil {
		return nil, err
	}
	slines := symbolOnLines(reels, g.lines)
	var lineWins []*LineWin
	for i, sl := range slines {
		var nhr, whr *HitResult
		if nw := calcNormalWins(sl); nw != nil {
			nhr = caclHitResult(nw, g.normalHits)
		}
		if ww := calcWildWins(sl); ww != nil {
			whr = caclHitResult(ww, g.wildHits)
		}
		if nhr == nil && whr == nil {
			continue
		}
		var hr *HitResult
		if nhr != nil && whr != nil {
			if nhr.ratio() > whr.ratio() {
				hr = nhr
			} else {
				hr = whr
			}
		} else if nhr != nil {
			hr = nhr
		} else {
			hr = whr
		}
		lineWins = append(lineWins, &LineWin{i, hr.ratio(), hr.win})
	}
	if sw := caclScatterWins(reels); sw != nil {
		if shr := caclHitResult(sw, g.scatterHits); shr != nil {
			r, f, m := shr.hit.ratio, shr.hit.features, shr.hit.multiplier
			return &SpinResult{reels, lineWins, &ScatterWin{r, f, m}}, nil
		}
	}
	return &SpinResult{reels, lineWins, nil}, nil
}
