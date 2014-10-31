package engine

import (
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
)

func TestGameBuilderSuccess(t *testing.T) {
	Symbols := [13]*Symbol{Ns(0, "Nine"), Ns(1, "Ten"), Ns(2, "Jack"), Ns(3, "Queen"), Ns(4, "King"), Ns(5, "Ace"), Ns(6, "Clam"), Ns(7, "Starfish"), Ns(8, "Nemo"), Ns(9, "Green"), Ns(10, "Octopus"), Ss(11, "Mermaid"), Ws(12, "Shark")}

	reels0 := [34]string{"Clam", "Nine", "Nemo", "Mermaid", "Queen", "Jack", "Green", "Ace", "Shark", "Starfish", "Ace", "King", "Clam", "Ten", "Queen", "Nemo", "Nine", "Starfish", "Mermaid", "Ten", "Starfish", "Queen", "Clam", "Octopus", "Ace", "Jack", "Mermaid", "King", "Queen", "Clam", "Ten", "Clam", "Ace", "Octopus"}
	reels1 := [33]string{"Green", "Ace", "Clam", "Mermaid", "King", "Octopus", "Ace", "Queen", "Starfish", "Shark", "Nemo", "King", "Starfish", "Mermaid", "Queen", "Nemo", "Jack", "Green", "Queen", "Ace", "Octopus", "King", "Green", "Ten", "Ace", "Starfish", "Nine", "King", "Green", "Ten", "Clam", "Nine", "Ten"}
	reels2 := [34]string{"King", "Jack", "Nemo", "Queen", "Octopus", "Ten", "Mermaid", "Ace", "Nemo", "Queen", "Starfish", "Ace", "King", "Starfish", "Ace", "Jack", "Octopus", "King", "Clam", "Shark", "Ten", "Nine", "Green", "Ace", "Nine", "Green", "Queen", "Mermaid", "Queen", "Octopus", "Ten", "Ace", "Clam", "Green"}
	reels3 := [46]string{"Jack", "Ten", "Starfish", "Shark", "Ten", "Green", "Nine", "Jack", "Octopus", "Mermaid", "Ace", "Jack", "Green", "Ace", "Nemo", "King", "Queen", "Clam", "King", "Jack", "Octopus", "Ace", "Nemo", "King", "Ace", "Starfish", "Nemo", "Queen", "King", "Nemo", "Nine", "Clam", "Ace", "King", "Nemo", "Ten", "Nine", "Octopus", "Jack", "King", "Starfish", "Ten", "Octopus", "Ace", "Nemo", "Ten"}
	reels4 := [35]string{"Nemo", "Starfish", "King", "Nine", "Mermaid", "Starfish", "Ace", "Nemo", "Octopus", "King", "Nine", "Starfish", "Mermaid", "Clam", "Queen", "Jack", "Octopus", "Ten", "Clam", "Ace", "Jack", "Green", "Octopus", "Nine", "Ten", "Shark", "King", "Octopus", "Queen", "Ace", "Clam", "Jack", "Octopus", "Queen", "Jack"}

	gc, err := createGameCore("main", 3, Symbols[:], reels0[:], reels1[:], reels2[:], reels3[:], reels4[:])

	assert.NoError(t, err, "create game core failed")
	assert.Equal(t, gc.nSymbols(), 13)
	assert.Equal(t, gc.nReels(), 5)
	assert.Equal(t, gc.rows, 3)
	assert.Equal(t, gc.mode, "main")
	matrix := gc.spin()
	assert.Len(t, matrix, 5, "matrix should have 5 reels")
	for _, l := range matrix {
		assert.Len(t, l, 3, "each reel should have 3 Symbols")
	}
}

func TestGameBuilderFailure(t *testing.T) {
	Symbols := [13]*Symbol{Ns(0, "Nine"), Ns(1, "Ten"), Ns(2, "Jack"), Ns(3, "Queen"), Ns(4, "King"), Ns(5, "Ace"), Ns(6, "Clam"), Ns(7, "Starfish"), Ns(8, "Nemo"), Ns(9, "Green"), Ns(10, "Octopus"), Ss(11, "Mermaid"), Ws(12, "Shark")}

	reels0 := [34]string{"Clam", "Ninex", "Nemo", "Mermaid", "Queen", "Jack", "Green", "Ace", "Shark", "Starfish", "Ace", "King", "Clam", "Ten", "Queen", "Nemo", "Nine", "Starfish", "Mermaid", "Ten", "Starfish", "Queen", "Clam", "Octopus", "Ace", "Jack", "Mermaid", "King", "Queen", "Clam", "Ten", "Clam", "Ace", "Octopus"}
	reels1 := [33]string{"Green", "Ace", "Clam", "Mermaid", "King", "Octopus", "Ace", "Queen", "Starfish", "Shark", "Nemo", "King", "Starfish", "Mermaid", "Queen", "Nemo", "Jack", "Green", "Queen", "Ace", "Octopus", "King", "Green", "Ten", "Ace", "Starfish", "Nine", "King", "Green", "Ten", "Clam", "Nine", "Ten"}
	reels2 := [34]string{"King", "Jack", "Nemo", "Queen", "Octopus", "Ten", "Mermaid", "Ace", "Nemo", "Queen", "Starfish", "Ace", "King", "Starfish", "Ace", "Jack", "Octopus", "King", "Clam", "Shark", "Ten", "Nine", "Green", "Ace", "Nine", "Green", "Queen", "Mermaid", "Queen", "Octopus", "Ten", "Ace", "Clam", "Green"}
	reels3 := [46]string{"Jack", "Ten", "Starfish", "Shark", "Ten", "Green", "Nine", "Jack", "Octopus", "Mermaid", "Ace", "Jack", "Green", "Ace", "Nemo", "King", "Queen", "Clam", "King", "Jack", "Octopus", "Ace", "Nemo", "King", "Ace", "Starfish", "Nemo", "Queen", "King", "Nemo", "Nine", "Clam", "Ace", "King", "Nemo", "Ten", "Nine", "Octopus", "Jack", "King", "Starfish", "Ten", "Octopus", "Ace", "Nemo", "Ten"}
	reels4 := [35]string{"Nemo", "Starfish", "King", "Nine", "Mermaid", "Starfish", "Ace", "Nemo", "Octopus", "King", "Nine", "Starfish", "Mermaid", "Clam", "Queen", "Jack", "Octopus", "Ten", "Clam", "Ace", "Jack", "Green", "Octopus", "Nine", "Ten", "Shark", "King", "Octopus", "Queen", "Ace", "Clam", "Jack", "Octopus", "Queen", "Jack"}
	_, err := createGameCore("main", 3, Symbols[:], reels0[:], reels1[:], reels2[:], reels3[:], reels4[:])

	assert.Error(t, err, "create game core should succeed")
}

type TestSpinResult struct {
	nStarfish, nOctupus, nShark int
}

func TestSpin1x1(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	starfish := Ns(7, "Starfish")
	octopus := Ns(10, "Octopus")
	shark := Ws(12, "Shark")

	engine := createEngine(1, Reel{starfish, octopus, shark})

	workers := 100
	jobs := 1000

	rchan := make(chan *TestSpinResult, workers)
	for i := 0; i < workers; i++ {
		go func() {
			var nStarfish, nOctupus, nShark int
			for j := 0; j < jobs; j++ {
				result := engine.spin()
				assert.Len(t, result, 1)
				assert.Len(t, result[0], 1)
				r := result[0][0]
				assert.True(t, r == starfish || r == octopus || r == shark)
				if r == starfish {
					nStarfish++
				} else if r == octopus {
					nOctupus++
				} else {
					nShark++
				}
			}
			rchan <- &TestSpinResult{nStarfish, nOctupus, nShark}
		}()
	}

	var nStarfish, nOctupus, nShark int
	for i := 0; i < workers; i++ {
		tsr := <-rchan
		nStarfish += tsr.nStarfish
		nOctupus += tsr.nOctupus
		nShark += tsr.nShark
	}
	close(rchan)
	assert.True(t, nStarfish > 33000)
	assert.True(t, nOctupus > 33000)
	assert.True(t, nShark > 33000)
	assert.True(t, nShark+nStarfish+nOctupus == jobs*workers)
}

func TestSpin2x1(t *testing.T) {
	starfish := Ns(7, "Starfish")
	octopus := Ns(10, "Octopus")
	shark := Ws(12, "Shark")
	engine := createEngine(2, Reel{starfish, octopus, shark})
	for i := 0; i < 10000; i++ {
		result := engine.spin()
		assert.Len(t, result, 1)
		assert.Len(t, result[0], 2)
		r1 := result[0][0]
		r2 := result[0][1]
		ok := (r1 == starfish && r2 == octopus) || (r1 == octopus && r2 == shark) || (r1 == shark && r2 == starfish)
		assert.True(t, ok)
	}
}

func TestSpin2x3(t *testing.T) {
	engine := createEngine(2,
		Reel{Ns(7, "Starfish"), Ns(10, "Octopus"), Ws(12, "Shark")},
		Reel{Ns(7, "Starfish"), Ns(10, "Octopus"), Ws(12, "Shark")},
		Reel{Ns(10, "Octopus"), Ns(11, "Mermaid"), Ws(12, "Shark")})
	result := engine.spin()
	assert.Len(t, result, 3)
	assert.Len(t, result[0], 2)
	assert.Len(t, result[1], 2)
	assert.Len(t, result[2], 2)
}

func BenchmarkSpin2x3(b *testing.B) {
	engine := createEngine(2,
		Reel{Ns(7, "Starfish"), Ns(10, "Octopus"), Ws(12, "Shark")},
		Reel{Ns(7, "Starfish"), Ns(10, "Octopus"), Ws(12, "Shark")},
		Reel{Ns(10, "Octopus"), Ns(11, "Mermaid"), Ws(12, "Shark")})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		engine.spin()
	}
}
