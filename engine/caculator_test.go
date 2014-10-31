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
	matrix := []Reel{
		Reel{Ns(5, "Ace"), Ns(6, "Clam"), Ns(7, "Starfish")},
		Reel{Ns(0, "Nine"), Ns(1, "Ten"), Ss(11, "Mermaid")},
		Reel{Ns(3, "Queen"), Ss(11, "Mermaid"), Ns(4, "King")},
	}
	w := caclScatterWins(matrix)
	assert.Nil(t, w)

	matrix = []Reel{
		Reel{Ns(5, "Ace"), Ns(6, "Clam"), Ns(7, "Starfish")},
		Reel{Ss(11, "Mermaid"), Ns(0, "Nine"), Ns(1, "Ten")},
		Reel{Ss(11, "Mermaid"), Ns(3, "Queen"), Ns(4, "King")},
	}
	w = caclScatterWins(matrix)
	assert.Nil(t, w)
}
