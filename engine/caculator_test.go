package engine

import (
	"fmt"
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

func expect(t *testing.T, a *Win, b *Win) {
	if b != nil && a != nil {
		if *(a.Symbol) != *(b.Symbol) || a.Counts != b.Counts || a.Substitute != b.Substitute {
			es := fmt.Sprint("[", *a, "] should equal to [", *b, "]")
			t.Error(es)
		}
	} else if b == nil && a == nil {
		return
	} else {
		es := fmt.Sprint("[", a, "] should equal to [", b, "]")
		t.Error(es)
	}
}

func TestCalcNormalWins(t *testing.T) {
	sstr := []string{"Nine", "Ten", "Jack", "Queen", "King"}
	ss := strings2Symbols(testSymbolsMap, sstr)
	w1 := calcNormalWins(ss)
	expect(t, w1, nil)

	sstr = []string{"Mermaid", "Nine", "Nine", "Nine", "Nine"}
	ss = strings2Symbols(testSymbolsMap, sstr)
	w1 = calcNormalWins(ss)
	expect(t, w1, nil)

	sstr = []string{"Mermaid", "Mermaid", "Nine", "Nine", "Nine"}
	ss = strings2Symbols(testSymbolsMap, sstr)
	w1 = calcNormalWins(ss)
	expect(t, w1, nil)

	sstr = []string{"Shark", "Mermaid", "Nine", "Nine", "Nine"}
	ss = strings2Symbols(testSymbolsMap, sstr)
	w1 = calcNormalWins(ss)
	expect(t, w1, nil)

	sstr = []string{"Mermaid", "Shark", "Nine", "Nine", "Nine"}
	ss = strings2Symbols(testSymbolsMap, sstr)
	w1 = calcNormalWins(ss)
	expect(t, w1, nil)

	sstr = []string{"Nine", "Nine", "Mermaid", "Nine", "Nine"}
	ss = strings2Symbols(testSymbolsMap, sstr)
	w1 = calcNormalWins(ss)
	fmt.Println("TestCalcNormalWins", w1)
	expect(t, w1, NewWin(Ns(0, "Nine"), 2, false))

	sstr = []string{"Nine", "Shark", "Mermaid", "Nine", "Nine"}
	ss = strings2Symbols(testSymbolsMap, sstr)
	w1 = calcNormalWins(ss)
	fmt.Println("TestCalcNormalWins", w1)
	expect(t, w1, NewWin(Ns(0, "Nine"), 2, true))

	sstr = []string{"Shark", "Nine", "Mermaid", "Nine", "Nine"}
	ss = strings2Symbols(testSymbolsMap, sstr)
	w1 = calcNormalWins(ss)
	fmt.Println("TestCalcNormalWins", w1)
	expect(t, w1, NewWin(Ns(0, "Nine"), 2, true))
}
