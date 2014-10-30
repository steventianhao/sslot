package games

import (
	"fmt"
	"testing"
)

func TestUnderWaterSpin(t *testing.T) {
	if r, err := Spin("underwater", "main"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(r)
	}
}
