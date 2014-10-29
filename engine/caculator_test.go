package engine

import (
	"fmt"
	"testing"
)

var testSymbols = []*symbol{
	ns(0, "Nine"),
	ns(1, "Ten"),
	ns(2, "Jack"),
	ns(3, "Queen"),
	ns(4, "King"),
	ns(5, "Ace"),
	ns(6, "Clam"),
	ns(7, "Starfish"),
	ns(8, "Nemo"),
	ns(9, "Green"),
	ns(10, "Octopus"),
	ss(11, "Mermaid"),
	ws(12, "Shark"),
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
	expect(t, w1, NewWin(ns(0, "Nine"), 2, false))

	sstr = []string{"Nine", "Shark", "Mermaid", "Nine", "Nine"}
	ss = strings2Symbols(testSymbolsMap, sstr)
	w1 = calcNormalWins(ss)
	fmt.Println("TestCalcNormalWins", w1)
	expect(t, w1, NewWin(ns(0, "Nine"), 2, true))

	sstr = []string{"Shark", "Nine", "Mermaid", "Nine", "Nine"}
	ss = strings2Symbols(testSymbolsMap, sstr)
	w1 = calcNormalWins(ss)
	fmt.Println("TestCalcNormalWins", w1)
	expect(t, w1, NewWin(ns(0, "Nine"), 2, true))
}
