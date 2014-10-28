package engine

import "fmt"

type HitKey struct {
	symbol string
	counts int
}

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

// if less than 2, then ZERO chance to win
func calcNormalWins(symbols []*symbol) *Win {
	first := symbols[0]
	if first.isScatter() {
		return nil
	}
	c := 1
	wild := first.isWild()
	for _, v := range symbols[1:] {
		if v.isScatter() {
			break
		} else if *v == *first {
			c = c + 1
		} else if v.isWild() {
			c = c + 1
			wild = true
		} else if first.isWild() {
			c = c + 1
			first = v
		} else {
			break
		}
	}
	if c < 2 {
		return nil
	} else if first.isWild() {
		return nil
	} else {
		return NewWin(first, c, wild)
	}
}

// if less than 2, then ZERO chance to win
func calcWildWins(symbols []*symbol) *Win {
	first := symbols[0]
	if !first.isWild() {
		return nil
	}
	c := 1
	for _, v := range symbols[1:] {
		if *v == *first {
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

// assume just one scatter in the symbols
// if less than 2, then ZERO chance to win
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
