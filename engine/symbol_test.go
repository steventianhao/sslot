package engine

import (
	//"fmt"
	"testing"
)

func TestSpin1x1(t *testing.T) {
	engine := createEngine(1, Reel{Ns(7, "Starfish"), Ns(10, "Octopus"), Ws(12, "Shark")})
	result := engine.spin()
	if len(result) != 1 {
		t.Error("1 reel so just give one")
	} else {
		if len(result[0]) != 1 {
			t.Error("1 row so just one symbol")
		}
	}
	//fmt.Println("test spin of engine:", result)
}

func TestSpin2x1(t *testing.T) {
	engine := createEngine(2, Reel{Ns(7, "Starfish"), Ns(10, "Octopus"), Ws(12, "Shark")})
	result := engine.spin()
	if len(result) != 1 {
		t.Error("1 reel so just give one")
	} else {
		if len(result[0]) != 2 {
			t.Error("1 row so just one symbol")
		}
	}
	//fmt.Println("test spin of engine:", result)
}

func TestSpin1x3(t *testing.T) {
	engine := createEngine(1,
		Reel{Ns(7, "Starfish"), Ns(10, "Octopus"), Ws(12, "Shark")},
		Reel{Ns(7, "Starfish"), Ns(10, "Octopus"), Ws(12, "Shark")},
		Reel{Ns(10, "Octopus"), Ns(11, "Mermaid"), Ws(12, "Shark")})
	result := engine.spin()
	if len(result) != 3 {
		t.Error("1 reel so just give one")
	} else {
		if len(result[0]) != 1 || len(result[1]) != 1 || len(result[2]) != 1 {
			t.Error("1 row so just one symbol")
		}
	}
	//fmt.Println("test spin of engine:", result)
}
