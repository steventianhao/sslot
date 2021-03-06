package engine

import "fmt"

type Win struct {
	Key
	Substitute bool
}

func (w Win) String() string {
	return fmt.Sprint("symbol:", w.Key.Symbol, ",counts:", w.Key.Counts, ",substitude:", w.Substitute)
}

func NewNormalWin(sym string, counts int, wild bool) *Win {
	return &Win{Key{sym, counts}, wild}
}

func NewOtherWin(sym string, counts int) *Win {
	return &Win{Key{sym, counts}, false}
}

//deprecated, later romove this function
func NewWin(sym string, counts int, wild bool) *Win {
	return &Win{Key{sym, counts}, wild}
}

// if less than 2, then ZERO chance to win
func calcNormalWins(symbols SLine) *Win {
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
		return NewNormalWin(first.Name, c, wild)
	}
}

// if less than 2, then ZERO chance to win
func calcWildWins(symbols SLine) *Win {
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
		return NewOtherWin(first.Name, c)
	}
}

// assume just one scatter in the symbols
// if less than 2, then ZERO chance to win
func caclScatterWins(reels []Reel) *Win {
	hasScatter := func(strip []*Symbol) (bool, *Symbol) {
		for _, v := range strip {
			if v.isScatter() {
				return true, v
			}
		}
		return false, nil
	}
	isScatter, first := hasScatter(reels[0])
	if !isScatter {
		return nil
	}
	c := 1
	for _, strip := range reels[1:] {
		if ok, _ := hasScatter(strip); ok {
			c = c + 1
		} else {
			break
		}
	}
	if c < 2 {
		return nil
	} else {
		return NewOtherWin(first.Name, c)
	}
}
