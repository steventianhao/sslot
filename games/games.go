package games

import (
	"fmt"
	"sslot/engine"
	"sslot/games/underwater"
)

var AllGames = InitGames()

const (
	MODE_MAIN    = "main"
	MODE_FEATURE = "feature"
)

func InitGames() map[string]*engine.SlotGame {
	m := make(map[string]*engine.SlotGame)
	g1 := createUnderWater()
	m[g1.Name()] = g1
	return m
}

func createUnderWater() *engine.SlotGame {
	var err error
	g1 := engine.CreateGame(1, "underwater", 3, 5, 15)
	err = g1.SetLines(underwater.Lines)
	if err != nil {
		panic("game create failed")
	}
	err = g1.AddGameCore(MODE_MAIN, underwater.Symbols, underwater.MainReels...)
	if err != nil {
		panic("game create failed")
	}
	err = g1.AddGameCore(MODE_FEATURE, underwater.Symbols, underwater.FeatureReels...)
	if err != nil {
		panic("game create failed")
	}
	g1.AddHits(underwater.NormalHits, underwater.WildHits, underwater.ScatterHits)
	return g1
}

func Spin(game, mode string) ([]engine.Reel, error) {
	g, ok := AllGames[game]
	if !ok {
		return nil, fmt.Errorf("game [%s] not found", game)
	}
	return g.Spin(mode)
}
