package games

import (
	"fmt"
	"github.com/landjur/go-decimal"
	"testing"
)

func TestUnderWaterMainSpin(t *testing.T) {
	bet, _ := decimal.Parse("3.12")
	if r, err := PlayerMainSpin("underwater", "simon", 15, bet); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(r)
	}
}

func TestUnderWaterFreeSpin(t *testing.T) {
	if r, err := PlayerFreeSpin("underwater", "simon"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(r)
	}
}
