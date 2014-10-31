package engine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSymbolsEqual(t *testing.T) {
	s1 := Symbol{Id: 0, Name: "s1", kind: Normal}
	s2 := Symbol{Id: 0, Name: "s1", kind: Normal}
	assert.Equal(t, s1, s2, "two symbols should be the same")
}

func TestSymbolsNotEqual(t *testing.T) {
	s1 := Symbol{Id: 0, Name: "s1", kind: Normal}
	s2 := Symbol{Id: 0, Name: "s2", kind: Normal}
	assert.NotEqual(t, s1, s2, "two symbols should be the same")
}

func TestWsNsSs(t *testing.T) {
	ws := Ws(1, "aa")
	assert.True(t, ws.isWild())
	assert.False(t, ws.isScatter())

	ns := Ns(2, "bb")
	assert.False(t, ns.isWild() && ns.isScatter())

	ss := Ss(3, "cc")
	assert.True(t, ss.isScatter())
	assert.False(t, ss.isWild())
}

func TestSymbols2Map(t *testing.T) {
	symbols := []*Symbol{Ns(1, "aa"), Ns(2, "bb")}
	m := symbols2Map(symbols)
	assert.Equal(t, 2, len(m))

	symbols = []*Symbol{Ns(1, "aa"), Ns(2, "aa")}
	m = symbols2Map(symbols)
	assert.Equal(t, 1, len(m))
	assert.Equal(t, Ns(2, "aa"), m["aa"])
}

func TestCheckSymbolNames(t *testing.T) {
	symbols := []*Symbol{Ns(1, "aa"), Ns(2, "bb")}
	m := symbols2Map(symbols)

	yes := checkSymbolNames(m, []string{"aa", "bb"})
	assert.True(t, yes)

	yes = checkSymbolNames(m, []string{"aa"})
	assert.True(t, yes)

	yes = checkSymbolNames(m, []string{"bb"})
	assert.True(t, yes)

	yes = checkSymbolNames(m, []string{"cc"})
	assert.False(t, yes)

	yes = checkSymbolNames(m, []string{})
	assert.True(t, yes, "empty symbol list return true")

	yes = checkSymbolNames(m, nil)
	assert.True(t, yes, "nil symbol list return true")
}

func TestStrings2Symbols(t *testing.T) {
	symbols := []*Symbol{Ns(1, "aa"), Ns(2, "bb")}
	m := symbols2Map(symbols)

	s := strings2Symbols(m, []string{"aa"})
	assert.Len(t, s, 1)

	s = strings2Symbols(m, []string{"aa", "bb"})
	assert.Len(t, s, 2)

	s = strings2Symbols(m, []string{"aa", "aa"})
	assert.Len(t, s, 2)

	s = strings2Symbols(m, []string{"aa", "acca"})
	assert.Len(t, s, 2)
	assert.Nil(t, s[1])
}
