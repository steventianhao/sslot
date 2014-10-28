package engine

import "fmt"

type Win struct {
	Symbol     *symbol
	Counts     int
	Substitute bool
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
