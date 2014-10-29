package engine

import (
	//"fmt"
	"testing"
)

func TestSpin1x1(t *testing.T) {
	engine := createEngine(1, Reel{ns(7, "Starfish"), ns(10, "Octopus"), ws(12, "Shark")})
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
	engine := createEngine(2, Reel{ns(7, "Starfish"), ns(10, "Octopus"), ws(12, "Shark")})
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
		Reel{ns(7, "Starfish"), ns(10, "Octopus"), ws(12, "Shark")},
		Reel{ns(7, "Starfish"), ns(10, "Octopus"), ws(12, "Shark")},
		Reel{ns(10, "Octopus"), ss(11, "Mermaid"), ws(12, "Shark")})
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
