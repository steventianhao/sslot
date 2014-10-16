package sslot

import (
	"testing"
)

func TestRandomGen(t *testing.T) {
	reelsLen := []int{34, 23, 45, 56}
	r := RandomGen(reelsLen)
	for i, v := range r {
		if v > reelsLen[i] {
			t.Error("the index on ", i, " reel should be less than ", reelsLen[i])
		}
	}
}

func TestRandomSeqs(t *testing.T) {
	reelsLen := []int{34, 23, 45, 56}
	nRows := 3
	r2 := RandomSeqs(reelsLen, nRows)
	if len(r2) != len(reelsLen) {
		t.Error("result length should equals the total of reels")
	}
	for i, v := range r2 {
		if len(v) != nRows {
			t.Error("the sequence of index should have length of ", nRows)
		}
		for j, idx := range v {
			if idx > reelsLen[i] {
				t.Error("the index on ", j, " reel should be less than ", reelsLen[i])
			}
		}
	}
}
