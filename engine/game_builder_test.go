package engine

import (
	"fmt"
	"testing"
)

func TestSymbols(t *testing.T) {
	s := symbol{id: 0, kind: Normal}
	s2 := symbol{id: 0, kind: Normal}
	if !(s == s2) {
		t.Error(s, s2, "should be the same")
	}
}

func TestGameBuilderStep1(t *testing.T) {
	g := NewGameBuilder(3, 5, 15)

	symbols := []*symbol{Ns(0, "Nine"), Ns(1, "Ten"), Ns(2, "Jack"), Ns(3, "Queen"), Ns(4, "King"), Ns(5, "Ace"), Ns(6, "Clam"), Ns(7, "Starfish"), Ns(8, "Nemo"), Ns(9, "Green"), Ns(10, "Octopus"), Ss(11, "Mermaid"), Ws(12, "Shark")}
	g.SetSymbols(symbols)
	g2, err := g.str2symbol("Nine")
	if err != nil {
		t.Error("Nine should be existed")
	}
	nine := *Ns(0, "Nine")
	if *g2 != nine {
		t.Error("Nine should be", nine)
	}
	g2, err = g.str2symbol("Ninex")
	if err == nil || g2 != nil {
		t.Error("Ninex should NOT be existed")
	}
}

func TestGameBuilder(t *testing.T) {
	g := NewGameBuilder(3, 5, 15)

	symbols := []*symbol{Ns(0, "Nine"), Ns(1, "Ten"), Ns(2, "Jack"), Ns(3, "Queen"), Ns(4, "King"), Ns(5, "Ace"), Ns(6, "Clam"), Ns(7, "Starfish"), Ns(8, "Nemo"), Ns(9, "Green"), Ns(10, "Octopus"), Ss(11, "Mermaid"), Ws(12, "Shark")}
	g.SetSymbols(symbols)
	if !(len(g.symbols) == 13) {
		t.Error("length should be 13")
	}

	reels0 := [34]string{"Clam", "Nine", "Nemo", "Mermaid", "Queen", "Jack", "Green", "Ace", "Shark", "Starfish", "Ace", "King", "Clam", "Ten", "Queen", "Nemo", "Nine", "Starfish", "Mermaid", "Ten", "Starfish", "Queen", "Clam", "Octopus", "Ace", "Jack", "Mermaid", "King", "Queen", "Clam", "Ten", "Clam", "Ace", "Octopus"}
	reels1 := [33]string{"Green", "Ace", "Clam", "Mermaid", "King", "Octopus", "Ace", "Queen", "Starfish", "Shark", "Nemo", "King", "Starfish", "Mermaid", "Queen", "Nemo", "Jack", "Green", "Queen", "Ace", "Octopus", "King", "Green", "Ten", "Ace", "Starfish", "Nine", "King", "Green", "Ten", "Clam", "Nine", "Ten"}
	reels2 := [34]string{"King", "Jack", "Nemo", "Queen", "Octopus", "Ten", "Mermaid", "Ace", "Nemo", "Queen", "Starfish", "Ace", "King", "Starfish", "Ace", "Jack", "Octopus", "King", "Clam", "Shark", "Ten", "Nine", "Green", "Ace", "Nine", "Green", "Queen", "Mermaid", "Queen", "Octopus", "Ten", "Ace", "Clam", "Green"}
	reels3 := [46]string{"Jack", "Ten", "Starfish", "Shark", "Ten", "Green", "Nine", "Jack", "Octopus", "Mermaid", "Ace", "Jack", "Green", "Ace", "Nemo", "King", "Queen", "Clam", "King", "Jack", "Octopus", "Ace", "Nemo", "King", "Ace", "Starfish", "Nemo", "Queen", "King", "Nemo", "Nine", "Clam", "Ace", "King", "Nemo", "Ten", "Nine", "Octopus", "Jack", "King", "Starfish", "Ten", "Octopus", "Ace", "Nemo", "Ten"}
	reels4 := [35]string{"Nemo", "Starfish", "King", "Nine", "Mermaid", "Starfish", "Ace", "Nemo", "Octopus", "King", "Nine", "Starfish", "Mermaid", "Clam", "Queen", "Jack", "Octopus", "Ten", "Clam", "Ace", "Jack", "Green", "Octopus", "Nine", "Ten", "Shark", "King", "Octopus", "Queen", "Ace", "Clam", "Jack", "Octopus", "Queen", "Jack"}

	g2, err := g.SetReels(reels0[:], reels1[:], reels2[:], reels3[:], reels4[:])
	if err != nil {
		t.Error(err)
	}

	screenshot := g2.build().MainSpin()
	fmt.Println("scrrenshot:", screenshot)

	symbolLines := SymbolOnLines(screenshot, lines())
	for _, lines := range symbolLines {
		for _, l := range lines {
			fmt.Print(l.name, ",")
		}
		fmt.Println()
	}

}

func expect(t *testing.T, a *Win, b *Win) {
	if b == nil && a == nil {
		return
	} else if *(a.Symbol) != *(b.Symbol) && a.Counts != b.Counts && a.Substitute != b.Substitute {
		es := fmt.Sprint(*a, " should equal to ", *b)
		t.Error(es)
	}
}

func TestCalcWin(t *testing.T) {
	ss := []*symbol{Ns(0, "Nine"), Ns(1, "Ten"), Ns(2, "Jack"), Ns(3, "Queen"), Ns(4, "King")}
	w1 := calcNormalWins(ss)
	expect(t, w1, nil)

	w2 := calcWildWins(ss)
	expect(t, w2, nil)

	ss = []*symbol{Ns(0, "Nine"), Ns(0, "Nine"), Ns(2, "Jack"), Ns(3, "Queen"), Ns(4, "King")}
	w1 = calcNormalWins(ss)
	expect(t, w1, NewWin(Ns(0, "Nine"), 2, false))
	xxx := caclHitResult(w1, normalHits())
	if xxx != nil {
		fmt.Println(xxx.win, xxx.hit)
	}
	w2 = calcWildWins(ss)
	expect(t, w2, nil)

	ss = []*symbol{Ns(0, "Nine"), Ws(12, "Shark"), Ns(0, "Nine"), Ns(2, "Jack"), Ns(3, "Queen"), Ns(4, "King")}
	w1 = calcNormalWins(ss)
	expect(t, w1, NewWin(Ns(0, "Nine"), 3, true))
	xxx = caclHitResult(w1, normalHits())
	if xxx != nil {
		fmt.Println(xxx.win, xxx.hit)
	}
	w2 = calcWildWins(ss)
	expect(t, w2, nil)

	ss = []*symbol{Ws(12, "Shark"), Ns(0, "Nine"), Ns(0, "Nine"), Ns(2, "Jack"), Ns(3, "Queen"), Ns(4, "King")}
	w1 = calcNormalWins(ss)
	expect(t, w1, NewWin(Ns(0, "Nine"), 3, true))
	xxx = caclHitResult(w1, normalHits())
	if xxx != nil {
		fmt.Println(xxx.win, xxx.hit)
	}
	w2 = calcWildWins(ss)
	expect(t, w2, nil)

	ss = []*symbol{Ws(12, "Shark"), Ws(12, "Shark"), Ns(0, "Nine"), Ns(2, "Jack"), Ns(3, "Queen"), Ns(4, "King")}
	w1 = calcNormalWins(ss)
	expect(t, w1, NewWin(Ns(0, "Nine"), 3, true))
	xxx = caclHitResult(w1, normalHits())
	if xxx != nil {
		fmt.Println(xxx.win, xxx.hit)
	}
	w2 = calcWildWins(ss)
	expect(t, w2, NewWin(Ws(12, "Shark"), 2, false))
	xxx = caclHitResult(w2, wildHits())
	if xxx != nil {
		fmt.Println(xxx.win, xxx.hit)
	}

	ss = []*symbol{Ws(12, "Shark"), Ws(12, "Shark"), Ns(0, "Nine"), Ns(2, "Jack"), Ns(0, "Nine"), Ns(4, "King")}
	w1 = calcNormalWins(ss)
	expect(t, w1, NewWin(Ns(0, "Nine"), 3, true))
	xxx = caclHitResult(w1, normalHits())
	if xxx != nil {
		fmt.Println(xxx.win, xxx.hit)
	}
	w2 = calcWildWins(ss)
	expect(t, w2, NewWin(Ws(12, "Shark"), 2, false))
	xxx = caclHitResult(w2, wildHits())
	if xxx != nil {
		fmt.Println(xxx.win, xxx.hit)
	}

	ss = []*symbol{Ws(12, "Shark"), Ws(12, "Shark"), Ns(0, "Nine"), Ns(0, "Nine"), Ws(12, "Shark")}
	w1 = calcNormalWins(ss)
	expect(t, w1, NewWin(Ns(0, "Nine"), 5, true))
	xxx = caclHitResult(w1, normalHits())
	if xxx != nil {
		fmt.Println(xxx.win, xxx.hit)
	}
	w2 = calcWildWins(ss)
	expect(t, w2, NewWin(Ws(12, "Shark"), 2, false))
	xxx = caclHitResult(w2, wildHits())
	if xxx != nil {
		fmt.Println(xxx.win, xxx.hit)
	}

	ss = []*symbol{Ws(12, "Shark"), Ws(12, "Shark"), Ws(12, "Shark"), Ws(12, "Shark"), Ns(0, "Nine")}
	w1 = calcNormalWins(ss)
	expect(t, w1, NewWin(Ns(0, "Nine"), 5, true))
	xxx = caclHitResult(w1, normalHits())
	if xxx != nil {
		fmt.Println(xxx.win, xxx.hit)
	}
	w2 = calcWildWins(ss)
	expect(t, w2, NewWin(Ws(12, "Shark"), 4, false))
	xxx = caclHitResult(w2, wildHits())
	if xxx != nil {
		fmt.Println(xxx.win, xxx.hit)
	}

	ss = []*symbol{Ws(12, "Shark"), Ws(12, "Shark"), Ws(12, "Shark"), Ws(12, "Shark"), Ws(12, "Shark")}
	w1 = calcNormalWins(ss)
	expect(t, w1, nil)
	w2 = calcWildWins(ss)
	expect(t, w2, NewWin(Ws(12, "Shark"), 5, false))
	xxx = caclHitResult(w2, wildHits())
	if xxx != nil {
		fmt.Println(xxx.win, xxx.hit)
	}
}

func TestCalcScatter1(t *testing.T) {
	matrix := [][]*symbol{
		[]*symbol{Ns(0, "Nine"), Ns(1, "Ten"), Ss(11, "Mermaid")},
		[]*symbol{Ns(3, "Queen"), Ss(11, "Mermaid"), Ns(4, "King")},
		[]*symbol{Ns(5, "Ace"), Ns(6, "Clam"), Ns(7, "Starfish")},
	}
	w := caclScatterWins(matrix)
	expect(t, w, NewWin(Ss(11, "Mermaid"), 2, false))

	matrix = [][]*symbol{
		[]*symbol{Ss(11, "Mermaid"), Ns(0, "Nine"), Ns(1, "Ten")},
		[]*symbol{Ns(3, "Queen"), Ss(11, "Mermaid"), Ns(4, "King")},
		[]*symbol{Ns(5, "Ace"), Ns(6, "Clam"), Ns(7, "Starfish")},
	}
	w = caclScatterWins(matrix)
	expect(t, w, NewWin(Ss(11, "Mermaid"), 2, false))

}

func TestCalcScatter2(t *testing.T) {
	matrix := [][]*symbol{
		[]*symbol{Ns(0, "Nine"), Ns(1, "Ten"), Ss(11, "Mermaid")},
		[]*symbol{Ns(5, "Ace"), Ns(6, "Clam"), Ns(7, "Starfish")},
		[]*symbol{Ns(3, "Queen"), Ss(11, "Mermaid"), Ns(4, "King")},
	}
	w := caclScatterWins(matrix)
	expect(t, w, nil)

	matrix = [][]*symbol{
		[]*symbol{Ss(11, "Mermaid"), Ns(0, "Nine"), Ns(1, "Ten")},
		[]*symbol{Ns(5, "Ace"), Ns(6, "Clam"), Ns(7, "Starfish")},
		[]*symbol{Ns(3, "Queen"), Ss(11, "Mermaid"), Ns(4, "King")},
	}
	w = caclScatterWins(matrix)
	expect(t, w, nil)
}

func TestCalcScatter3(t *testing.T) {
	matrix := [][]*symbol{
		[]*symbol{Ns(5, "Ace"), Ns(6, "Clam"), Ns(7, "Starfish")},
		[]*symbol{Ns(0, "Nine"), Ns(1, "Ten"), Ss(11, "Mermaid")},
		[]*symbol{Ns(3, "Queen"), Ss(11, "Mermaid"), Ns(4, "King")},
	}
	w := caclScatterWins(matrix)
	expect(t, w, nil)

	matrix = [][]*symbol{
		[]*symbol{Ns(5, "Ace"), Ns(6, "Clam"), Ns(7, "Starfish")},
		[]*symbol{Ss(11, "Mermaid"), Ns(0, "Nine"), Ns(1, "Ten")},
		[]*symbol{Ss(11, "Mermaid"), Ns(3, "Queen"), Ns(4, "King")},
	}
	w = caclScatterWins(matrix)
	expect(t, w, nil)
}
