package engine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testSymbols = []*Symbol{
	Ns(0, "Nine"),
	Ns(1, "Ten"),
	Ns(2, "Jack"),
	Ns(3, "Queen"),
	Ns(4, "King"),
	Ns(5, "Ace"),
	Ns(6, "Clam"),
	Ns(7, "Starfish"),
	Ns(8, "Nemo"),
	Ns(9, "Green"),
	Ns(10, "Octopus"),
	Ss(11, "Mermaid"),
	Ws(12, "Shark"),
}

var testSymbolsMap = symbols2Map(testSymbols)

func TestCalcNormalWins(t *testing.T) {
	sstr := []string{"Nine", "Ten", "Jack", "Queen", "King"}
	ss := strings2Symbols(testSymbolsMap, sstr)
	w1 := calcNormalWins(ss)
	assert.Nil(t, w1)

	sstr = []string{"Mermaid", "Nine", "Nine", "Nine", "Nine"}
	ss = strings2Symbols(testSymbolsMap, sstr)
	w1 = calcNormalWins(ss)
	assert.Nil(t, w1)

	sstr = []string{"Mermaid", "Mermaid", "Nine", "Nine", "Nine"}
	ss = strings2Symbols(testSymbolsMap, sstr)
	w1 = calcNormalWins(ss)
	assert.Nil(t, w1)

	sstr = []string{"Shark", "Mermaid", "Nine", "Nine", "Nine"}
	ss = strings2Symbols(testSymbolsMap, sstr)
	w1 = calcNormalWins(ss)
	assert.Nil(t, w1)

	sstr = []string{"Mermaid", "Shark", "Nine", "Nine", "Nine"}
	ss = strings2Symbols(testSymbolsMap, sstr)
	w1 = calcNormalWins(ss)
	assert.Nil(t, w1)

	sstr = []string{"Shark", "Shark", "Shark", "Shark", "Shark"}
	ss = strings2Symbols(testSymbolsMap, sstr)
	w1 = calcNormalWins(ss)
	assert.Nil(t, w1)

	sstr = []string{"Shark", "Mermaid", "Mermaid", "Mermaid", "Shark"}
	ss = strings2Symbols(testSymbolsMap, sstr)
	w1 = calcNormalWins(ss)
	assert.Nil(t, w1)

	sstr = []string{"Shark", "Shark", "Mermaid", "Mermaid", "Shark"}
	ss = strings2Symbols(testSymbolsMap, sstr)
	w1 = calcNormalWins(ss)
	assert.Nil(t, w1)

	sstr = []string{"Nine", "Nine", "Mermaid", "Nine", "Nine"}
	ss = strings2Symbols(testSymbolsMap, sstr)
	w1 = calcNormalWins(ss)
	assert.Equal(t, w1, NewWin("Nine", 2, false))

	sstr = []string{"Nine", "Shark", "Mermaid", "Nine", "Nine"}
	ss = strings2Symbols(testSymbolsMap, sstr)
	w1 = calcNormalWins(ss)
	assert.Equal(t, w1, NewWin("Nine", 2, true))

	sstr = []string{"Shark", "Nine", "Mermaid", "Nine", "Nine"}
	ss = strings2Symbols(testSymbolsMap, sstr)
	w1 = calcNormalWins(ss)
	assert.Equal(t, w1, NewWin("Nine", 2, true))
}

func TestCalcWildWins(t *testing.T) {

	sstr := []string{"Shark", "Nine", "Mermaid", "Nine", "Nine"}
	ss := strings2Symbols(testSymbolsMap, sstr)
	w1 := calcWildWins(ss)
	assert.Nil(t, w1)

	sstr = []string{"Nine", "Shark", "Shark", "Shark", "Shark"}
	ss = strings2Symbols(testSymbolsMap, sstr)
	w1 = calcWildWins(ss)
	assert.Nil(t, w1)

	sstr = []string{"Shark", "Shark", "Mermaid", "Nine", "Nine"}
	ss = strings2Symbols(testSymbolsMap, sstr)
	w1 = calcWildWins(ss)
	assert.Equal(t, w1, NewWin("Shark", 2, false))

	sstr = []string{"Shark", "Shark", "Shark", "Nine", "Nine"}
	ss = strings2Symbols(testSymbolsMap, sstr)
	w1 = calcWildWins(ss)
	assert.Equal(t, w1, NewWin("Shark", 3, false))

	sstr = []string{"Shark", "Shark", "Shark", "Shark", "Nine"}
	ss = strings2Symbols(testSymbolsMap, sstr)
	w1 = calcWildWins(ss)
	assert.Equal(t, w1, NewWin("Shark", 4, false))

	sstr = []string{"Shark", "Shark", "Shark", "Shark", "Shark"}
	ss = strings2Symbols(testSymbolsMap, sstr)
	w1 = calcWildWins(ss)
	assert.Equal(t, w1, NewWin("Shark", 5, false))
}

func TestCalcNormalWins2(t *testing.T) {
	slines := [][]string{
		[]string{"Shark", "Nine", "Nine", "Nine", "Nine"},
		[]string{"Nine", "Nine", "Shark", "Nine", "Nine"},
		[]string{"Nine", "Nine", "Nine", "Nine", "Shark"},

		[]string{"Shark", "Shark", "Shark", "Nine", "Shark"},
		[]string{"Shark", "Shark", "Shark", "Shark", "Nine"},
		[]string{"Nine", "Shark", "Shark", "Shark", "Shark"},
	}

	for _, sline := range slines {
		ss := strings2Symbols(testSymbolsMap, sline)
		w1 := calcNormalWins(ss)
		assert.Equal(t, w1, NewWin("Nine", 5, true))
	}
}

func TestCalcScatter3(t *testing.T) {
	ss := Ss(11, "Mermaid")

	matrix := []Reel{
		Reel{Ns(5, "Ace"), Ns(6, "Clam"), Ns(7, "Starfish")},
		Reel{Ns(0, "Nine"), Ns(1, "Ten"), ss},
		Reel{Ns(3, "Queen"), ss, Ns(4, "King")},
	}
	w := caclScatterWins(matrix)
	assert.Nil(t, w)

	matrix = []Reel{
		Reel{Ns(5, "Ace"), Ns(6, "Clam"), Ns(7, "Starfish")},
		Reel{Ns(5, "Ace"), Ns(0, "Nine"), Ns(1, "Ten")},
		Reel{ss, Ns(3, "Queen"), Ns(4, "King")},
	}
	w = caclScatterWins(matrix)
	assert.Nil(t, w)

	matrix = []Reel{
		Reel{ss, Ns(3, "Queen"), Ns(4, "King")},
		Reel{Ns(5, "Ace"), Ns(6, "Clam"), Ns(7, "Starfish")},
		Reel{Ns(5, "Ace"), Ns(0, "Nine"), Ns(1, "Ten")},
	}
	w = caclScatterWins(matrix)
	assert.Nil(t, w)

	matrix = []Reel{
		Reel{Ns(5, "Ace"), Ns(6, "Clam"), Ns(7, "Starfish")},
		Reel{ss, Ns(0, "Nine"), Ns(1, "Ten")},
		Reel{ss, Ns(3, "Queen"), Ns(4, "King")},
	}
	w = caclScatterWins(matrix)
	assert.Nil(t, w)

	matrix = []Reel{
		Reel{Ns(5, "Ace"), Ns(6, "Clam"), ss},
		Reel{ss, Ns(0, "Nine"), Ns(1, "Ten")},
		Reel{ss, Ns(3, "Queen"), Ns(4, "King")},
	}
	w = caclScatterWins(matrix)
	assert.Equal(t, w, NewWin("Mermaid", 3, false))

	matrix = []Reel{
		Reel{Ns(5, "Ace"), Ns(6, "Clam"), ss},
		Reel{ss, Ns(0, "Nine"), Ns(1, "Ten")},
		Reel{Ns(5, "Ace"), Ns(3, "Queen"), Ns(4, "King")},
	}
	w = caclScatterWins(matrix)
	assert.Equal(t, w, NewWin("Mermaid", 2, false))
}
