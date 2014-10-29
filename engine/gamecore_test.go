package engine

import (
	"fmt"
	"testing"
)

func TestSymbols(t *testing.T) {
	s := Symbol{id: 0, kind: Normal}
	s2 := Symbol{id: 0, kind: Normal}
	if !(s == s2) {
		t.Error(s, s2, "should be the same")
	}
}

func TestGameBuilder(t *testing.T) {
	Symbols := [13]*Symbol{Ns(0, "Nine"), Ns(1, "Ten"), Ns(2, "Jack"), Ns(3, "Queen"), Ns(4, "King"), Ns(5, "Ace"), Ns(6, "Clam"), Ns(7, "Starfish"), Ns(8, "Nemo"), Ns(9, "Green"), Ns(10, "Octopus"), Ss(11, "Mermaid"), Ws(12, "Shark")}

	reels0 := [34]string{"Clam", "Nine", "Nemo", "Mermaid", "Queen", "Jack", "Green", "Ace", "Shark", "Starfish", "Ace", "King", "Clam", "Ten", "Queen", "Nemo", "Nine", "Starfish", "Mermaid", "Ten", "Starfish", "Queen", "Clam", "Octopus", "Ace", "Jack", "Mermaid", "King", "Queen", "Clam", "Ten", "Clam", "Ace", "Octopus"}
	reels1 := [33]string{"Green", "Ace", "Clam", "Mermaid", "King", "Octopus", "Ace", "Queen", "Starfish", "Shark", "Nemo", "King", "Starfish", "Mermaid", "Queen", "Nemo", "Jack", "Green", "Queen", "Ace", "Octopus", "King", "Green", "Ten", "Ace", "Starfish", "Nine", "King", "Green", "Ten", "Clam", "Nine", "Ten"}
	reels2 := [34]string{"King", "Jack", "Nemo", "Queen", "Octopus", "Ten", "Mermaid", "Ace", "Nemo", "Queen", "Starfish", "Ace", "King", "Starfish", "Ace", "Jack", "Octopus", "King", "Clam", "Shark", "Ten", "Nine", "Green", "Ace", "Nine", "Green", "Queen", "Mermaid", "Queen", "Octopus", "Ten", "Ace", "Clam", "Green"}
	reels3 := [46]string{"Jack", "Ten", "Starfish", "Shark", "Ten", "Green", "Nine", "Jack", "Octopus", "Mermaid", "Ace", "Jack", "Green", "Ace", "Nemo", "King", "Queen", "Clam", "King", "Jack", "Octopus", "Ace", "Nemo", "King", "Ace", "Starfish", "Nemo", "Queen", "King", "Nemo", "Nine", "Clam", "Ace", "King", "Nemo", "Ten", "Nine", "Octopus", "Jack", "King", "Starfish", "Ten", "Octopus", "Ace", "Nemo", "Ten"}
	reels4 := [35]string{"Nemo", "Starfish", "King", "Nine", "Mermaid", "Starfish", "Ace", "Nemo", "Octopus", "King", "Nine", "Starfish", "Mermaid", "Clam", "Queen", "Jack", "Octopus", "Ten", "Clam", "Ace", "Jack", "Green", "Octopus", "Nine", "Ten", "Shark", "King", "Octopus", "Queen", "Ace", "Clam", "Jack", "Octopus", "Queen", "Jack"}

	gc, err := createGameCore("main", 3, Symbols[:], reels0[:], reels1[:], reels2[:], reels3[:], reels4[:])
	if err != nil {
		t.Error("create game core failed", err.Error())
	} else {
		if gc.nSymbols() != 13 || gc.nReels() != 5 || gc.rows != 3 || gc.mode != "main" {
			t.Error("create game core wrong result")
		}
	}

	matrix := gc.spin()
	fmt.Println("scrreNshot:", matrix)
	if len(matrix) != 5 {
		t.Error("matrix should have 5 reels")
	} else {
		for _, l := range matrix {
			if len(l) != 3 {
				t.Error("each reel should have 3 Symbols")
			}
		}
	}
}

func TestCalcScatter3(t *testing.T) {
	matrix := []Reel{
		Reel{Ns(5, "Ace"), Ns(6, "Clam"), Ns(7, "Starfish")},
		Reel{Ns(0, "Nine"), Ns(1, "Ten"), Ss(11, "Mermaid")},
		Reel{Ns(3, "Queen"), Ss(11, "Mermaid"), Ns(4, "King")},
	}
	w := caclScatterWins(matrix)
	expect(t, w, nil)

	matrix = []Reel{
		Reel{Ns(5, "Ace"), Ns(6, "Clam"), Ns(7, "Starfish")},
		Reel{Ss(11, "Mermaid"), Ns(0, "Nine"), Ns(1, "Ten")},
		Reel{Ss(11, "Mermaid"), Ns(3, "Queen"), Ns(4, "King")},
	}
	w = caclScatterWins(matrix)
	expect(t, w, nil)
}
