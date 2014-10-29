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

func TestGameBuilder(t *testing.T) {
	symbols := [13]*symbol{ns(0, "Nine"), ns(1, "Ten"), ns(2, "Jack"), ns(3, "Queen"), ns(4, "King"), ns(5, "Ace"), ns(6, "Clam"), ns(7, "Starfish"), ns(8, "Nemo"), ns(9, "Green"), ns(10, "Octopus"), ss(11, "Mermaid"), ws(12, "Shark")}

	reels0 := [34]string{"Clam", "Nine", "Nemo", "Mermaid", "Queen", "Jack", "Green", "Ace", "Shark", "Starfish", "Ace", "King", "Clam", "Ten", "Queen", "Nemo", "Nine", "Starfish", "Mermaid", "Ten", "Starfish", "Queen", "Clam", "Octopus", "Ace", "Jack", "Mermaid", "King", "Queen", "Clam", "Ten", "Clam", "Ace", "Octopus"}
	reels1 := [33]string{"Green", "Ace", "Clam", "Mermaid", "King", "Octopus", "Ace", "Queen", "Starfish", "Shark", "Nemo", "King", "Starfish", "Mermaid", "Queen", "Nemo", "Jack", "Green", "Queen", "Ace", "Octopus", "King", "Green", "Ten", "Ace", "Starfish", "Nine", "King", "Green", "Ten", "Clam", "Nine", "Ten"}
	reels2 := [34]string{"King", "Jack", "Nemo", "Queen", "Octopus", "Ten", "Mermaid", "Ace", "Nemo", "Queen", "Starfish", "Ace", "King", "Starfish", "Ace", "Jack", "Octopus", "King", "Clam", "Shark", "Ten", "Nine", "Green", "Ace", "Nine", "Green", "Queen", "Mermaid", "Queen", "Octopus", "Ten", "Ace", "Clam", "Green"}
	reels3 := [46]string{"Jack", "Ten", "Starfish", "Shark", "Ten", "Green", "Nine", "Jack", "Octopus", "Mermaid", "Ace", "Jack", "Green", "Ace", "Nemo", "King", "Queen", "Clam", "King", "Jack", "Octopus", "Ace", "Nemo", "King", "Ace", "Starfish", "Nemo", "Queen", "King", "Nemo", "Nine", "Clam", "Ace", "King", "Nemo", "Ten", "Nine", "Octopus", "Jack", "King", "Starfish", "Ten", "Octopus", "Ace", "Nemo", "Ten"}
	reels4 := [35]string{"Nemo", "Starfish", "King", "Nine", "Mermaid", "Starfish", "Ace", "Nemo", "Octopus", "King", "Nine", "Starfish", "Mermaid", "Clam", "Queen", "Jack", "Octopus", "Ten", "Clam", "Ace", "Jack", "Green", "Octopus", "Nine", "Ten", "Shark", "King", "Octopus", "Queen", "Ace", "Clam", "Jack", "Octopus", "Queen", "Jack"}

	gc, err := createGameCore(MODE_NORMAL, 3, symbols[:], reels0[:], reels1[:], reels2[:], reels3[:], reels4[:])
	if err != nil {
		t.Error("create game core failed", err.Error())
	} else {
		if gc.nSymbols() != 13 || gc.nReels() != 5 || gc.rows != 3 || gc.mode != MODE_NORMAL {
			t.Error("create game core wrong result")
		}
	}

	screenshot := gc.spin()
	fmt.Println("scrrenshot:", screenshot)

	symbolLines := symbolOnLines(screenshot, lines())
	for _, lines := range symbolLines {
		for _, l := range lines {
			fmt.Print(l.name, ",")
		}
		fmt.Println()
	}

}

func TestCalcWin(t *testing.T) {
	ss := []*symbol{ns(0, "Nine"), ns(1, "Ten"), ns(2, "Jack"), ns(3, "Queen"), ns(4, "King")}
	w1 := calcNormalWins(ss)
	expect(t, w1, nil)

	w2 := calcWildWins(ss)
	expect(t, w2, nil)

	ss = []*symbol{ns(0, "Nine"), ns(0, "Nine"), ns(2, "Jack"), ns(3, "Queen"), ns(4, "King")}
	w1 = calcNormalWins(ss)
	expect(t, w1, NewWin(ns(0, "Nine"), 2, false))
	xxx := caclHitResult(w1, normalHits())
	if xxx != nil {
		fmt.Println(xxx.win, xxx.hit)
	}
	w2 = calcWildWins(ss)
	expect(t, w2, nil)

	ss = []*symbol{ns(0, "Nine"), ws(12, "Shark"), ns(0, "Nine"), ns(2, "Jack"), ns(3, "Queen"), ns(4, "King")}
	w1 = calcNormalWins(ss)
	expect(t, w1, NewWin(ns(0, "Nine"), 3, true))
	xxx = caclHitResult(w1, normalHits())
	if xxx != nil {
		fmt.Println(xxx.win, xxx.hit)
	}
	w2 = calcWildWins(ss)
	expect(t, w2, nil)

	ss = []*symbol{ws(12, "Shark"), ns(0, "Nine"), ns(0, "Nine"), ns(2, "Jack"), ns(3, "Queen"), ns(4, "King")}
	w1 = calcNormalWins(ss)
	expect(t, w1, NewWin(ns(0, "Nine"), 3, true))
	xxx = caclHitResult(w1, normalHits())
	if xxx != nil {
		fmt.Println(xxx.win, xxx.hit)
	}
	w2 = calcWildWins(ss)
	expect(t, w2, nil)

	ss = []*symbol{ws(12, "Shark"), ws(12, "Shark"), ns(0, "Nine"), ns(2, "Jack"), ns(3, "Queen"), ns(4, "King")}
	w1 = calcNormalWins(ss)
	expect(t, w1, NewWin(ns(0, "Nine"), 3, true))
	xxx = caclHitResult(w1, normalHits())
	if xxx != nil {
		fmt.Println(xxx.win, xxx.hit)
	}
	w2 = calcWildWins(ss)
	expect(t, w2, NewWin(ws(12, "Shark"), 2, false))
	xxx = caclHitResult(w2, wildHits())
	if xxx != nil {
		fmt.Println(xxx.win, xxx.hit)
	}

	ss = []*symbol{ws(12, "Shark"), ws(12, "Shark"), ns(0, "Nine"), ns(2, "Jack"), ns(0, "Nine"), ns(4, "King")}
	w1 = calcNormalWins(ss)
	expect(t, w1, NewWin(ns(0, "Nine"), 3, true))
	xxx = caclHitResult(w1, normalHits())
	if xxx != nil {
		fmt.Println(xxx.win, xxx.hit)
	}
	w2 = calcWildWins(ss)
	expect(t, w2, NewWin(ws(12, "Shark"), 2, false))
	xxx = caclHitResult(w2, wildHits())
	if xxx != nil {
		fmt.Println(xxx.win, xxx.hit)
	}

	ss = []*symbol{ws(12, "Shark"), ws(12, "Shark"), ns(0, "Nine"), ns(0, "Nine"), ws(12, "Shark")}
	w1 = calcNormalWins(ss)
	expect(t, w1, NewWin(ns(0, "Nine"), 5, true))
	xxx = caclHitResult(w1, normalHits())
	if xxx != nil {
		fmt.Println(xxx.win, xxx.hit)
	}
	w2 = calcWildWins(ss)
	expect(t, w2, NewWin(ws(12, "Shark"), 2, false))
	xxx = caclHitResult(w2, wildHits())
	if xxx != nil {
		fmt.Println(xxx.win, xxx.hit)
	}

	ss = []*symbol{ws(12, "Shark"), ws(12, "Shark"), ws(12, "Shark"), ws(12, "Shark"), ns(0, "Nine")}
	w1 = calcNormalWins(ss)
	expect(t, w1, NewWin(ns(0, "Nine"), 5, true))
	xxx = caclHitResult(w1, normalHits())
	if xxx != nil {
		fmt.Println(xxx.win, xxx.hit)
	}
	w2 = calcWildWins(ss)
	expect(t, w2, NewWin(ws(12, "Shark"), 4, false))
	xxx = caclHitResult(w2, wildHits())
	if xxx != nil {
		fmt.Println(xxx.win, xxx.hit)
	}

	ss = []*symbol{ws(12, "Shark"), ws(12, "Shark"), ws(12, "Shark"), ws(12, "Shark"), ws(12, "Shark")}
	w1 = calcNormalWins(ss)
	expect(t, w1, nil)
	w2 = calcWildWins(ss)
	expect(t, w2, NewWin(ws(12, "Shark"), 5, false))
	xxx = caclHitResult(w2, wildHits())
	if xxx != nil {
		fmt.Println(xxx.win, xxx.hit)
	}
}

func TestCalcScatter1(t *testing.T) {
	matrix := []Reel{
		Reel{ns(0, "Nine"), ns(1, "Ten"), ss(11, "Mermaid")},
		Reel{ns(3, "Queen"), ss(11, "Mermaid"), ns(4, "King")},
		Reel{ns(5, "Ace"), ns(6, "Clam"), ns(7, "Starfish")},
	}
	w := caclScatterWins(matrix)
	expect(t, w, NewWin(ss(11, "Mermaid"), 2, false))

	matrix = []Reel{
		Reel{ss(11, "Mermaid"), ns(0, "Nine"), ns(1, "Ten")},
		Reel{ns(3, "Queen"), ss(11, "Mermaid"), ns(4, "King")},
		Reel{ns(5, "Ace"), ns(6, "Clam"), ns(7, "Starfish")},
	}
	w = caclScatterWins(matrix)
	expect(t, w, NewWin(ss(11, "Mermaid"), 2, false))

}

func TestCalcScatter2(t *testing.T) {
	matrix := []Reel{
		Reel{ns(0, "Nine"), ns(1, "Ten"), ss(11, "Mermaid")},
		Reel{ns(5, "Ace"), ns(6, "Clam"), ns(7, "Starfish")},
		Reel{ns(3, "Queen"), ss(11, "Mermaid"), ns(4, "King")},
	}
	w := caclScatterWins(matrix)
	expect(t, w, nil)

	matrix = []Reel{
		Reel{ss(11, "Mermaid"), ns(0, "Nine"), ns(1, "Ten")},
		Reel{ns(5, "Ace"), ns(6, "Clam"), ns(7, "Starfish")},
		Reel{ns(3, "Queen"), ss(11, "Mermaid"), ns(4, "King")},
	}
	w = caclScatterWins(matrix)
	expect(t, w, nil)
}

func TestCalcScatter3(t *testing.T) {
	matrix := []Reel{
		Reel{ns(5, "Ace"), ns(6, "Clam"), ns(7, "Starfish")},
		Reel{ns(0, "Nine"), ns(1, "Ten"), ss(11, "Mermaid")},
		Reel{ns(3, "Queen"), ss(11, "Mermaid"), ns(4, "King")},
	}
	w := caclScatterWins(matrix)
	expect(t, w, nil)

	matrix = []Reel{
		Reel{ns(5, "Ace"), ns(6, "Clam"), ns(7, "Starfish")},
		Reel{ss(11, "Mermaid"), ns(0, "Nine"), ns(1, "Ten")},
		Reel{ss(11, "Mermaid"), ns(3, "Queen"), ns(4, "King")},
	}
	w = caclScatterWins(matrix)
	expect(t, w, nil)
}
