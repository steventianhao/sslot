package game

import (
	"github.com/landjur/go-decimal"
)

var games = []string{"underwater", "transformer"}

func ShowGame(name string) bool {
	for _, v := range games {
		if v == name {
			return true
		}
	}
	return false
}

func SpinGame(name string, lines int, bet *decimal.Decimal) (*Result, error) {
	//chcecking the lines, should be less then the games provided 0<lines<max
	//checking the money, bet> min, bet < max
	//caculate the user's balance, if not enough, reject the spin
	return testResult(), nil
}

func FreeSpinGame(name string, lines int, bet *decimal.Decimal) (*Result, error) {
	//chcecking the lines, should be less then the games provided 0<lines<max
	//checking the money, bet> min, bet < max
	//caculate the user's balance, if not enough, reject the spin
	return testResult2(), nil
}

func testResult() *Result {
	return &Result{
		[][]int{{1, 2, 3}, {4, 6, 7}, {2, 3, 4}, {1, 1, 1}, {5, 6, 2}},
		[]*Payout{&Payout{1, 400}, &Payout{2, 500}, &Payout{3, 10}},
		3,
	}
}

func testResult2() *Result {
	return &Result{
		[][]int{{1, 2, 3}, {4, 6, 7}, {2, 3, 4}, {1, 1, 1}, {2, 2, 2}},
		[]*Payout{&Payout{1, 2}, &Payout{2, 5}, &Payout{3, 1}},
		0,
	}
}

type Spin struct {
	Game     string
	Lines    int
	Bet      string
	Features int
	Fresh    bool
}

func FreshSpin(game string) *Spin {
	return &Spin{game, 0, "0", 0, true}
}

func CacheSpin(game string, lines int, bet string, featrues int) *Spin {
	return &Spin{game, lines, bet, featrues, false}
}

type Payout struct {
	Line  int
	Ratio int
}

type Result struct {
	Screen   [][]int
	Payouts  []*Payout
	Features int
}
