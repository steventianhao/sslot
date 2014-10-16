package sslot

import (
	"math/rand"
	"time"
)

/*
	ranges: length of all symbols for each reel
	return: random index of the symbol for each reel
*/
func RandomGen(ranges []int) []int {
	l := len(ranges)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	reels := make([]int, l)
	for i := 0; i < l; i++ {
		reels[i] = r.Intn(ranges[i])
	}
	return reels
}

/*
	ranges : length of all symbols for earch reel
	nRows : rows of each reel in the game window
	return : [cols][rows] of symbol indexes
*/
func RandomSeqs(ranges []int, nRows int) [][]int {
	startSeqs := RandomGen(ranges)
	l := len(ranges)
	result := make([][]int, l)
	for i, v := range startSeqs {
		limit := ranges[i]
		rows := make([]int, nRows)
		for j := 0; j < nRows; j++ {
			rows[j] = (v + j) % limit
		}
		result[i] = rows
	}
	return result
}
