package games

import (
	"errors"
	"fmt"
	"github.com/fzzy/radix/redis"
	"github.com/landjur/go-decimal"
	"sslot/engine"
	"strconv"
	"time"
)

type SpinHistory struct {
	Game     string
	Lines    int
	Bet      string
	Features int
	Fresh    bool
}

func UserHash(username string) string {
	return fmt.Sprint("user:", username)
}

func GameFieldLines(gamename string) string {
	return fmt.Sprint("game_", gamename, "_lines")
}
func GameFieldBet(gamename string) string {
	return fmt.Sprint("game_", gamename, "_bet")
}

func GameFieldFeatures(gamename string) string {
	return fmt.Sprint("game_", gamename, "_features")
}

func NewSpin(game string) *SpinHistory {
	return &SpinHistory{game, 0, "0", 0, true}
}

func OldSpin(game string, lines int, bet string, featrues int) *SpinHistory {
	return &SpinHistory{game, lines, bet, featrues, false}
}

func RestoreSpinHistory(conn *redis.Client, username, gamename string) (*SpinHistory, error) {
	key := UserHash(username)
	res, err := conn.Cmd("HMGET", key, GameFieldLines(gamename), GameFieldBet(gamename), GameFieldFeatures(gamename)).List()
	if err != nil {
		return nil, err
	}
	strLine, strBet, strFeatures := res[0], res[1], res[2]
	if strLine != "" && strBet != "" && strFeatures != "" {
		lines, err := strconv.Atoi(strLine)
		if err != nil {
			return nil, err
		}
		featrues, err := strconv.Atoi(strFeatures)
		if err != nil {
			return nil, err
		}
		return OldSpin(gamename, lines, strBet, featrues), nil
	} else {
		return NewSpin(gamename), nil
	}
}

type InternalError struct {
	cause error
}

func (self InternalError) Error() string {
	return fmt.Sprint("internal state is invalid, root cause is:", self.cause.Error())
}

func PlayerMainSpin(game, player string, lines int, bet *decimal.Decimal) (*engine.SpinResult, error) {
	conn, err := redis.DialTimeout("tcp", "127.0.0.1:6379", time.Duration(2)*time.Second)
	if err != nil {
		return nil, &InternalError{err}
	}
	defer conn.Close()

	history, err := RestoreSpinHistory(conn, player, game)
	if err != nil {
		return nil, &InternalError{err}
	}

	if history.Features > 0 {
		return nil, fmt.Errorf("spin status is invalid, there are still %d free spins left", history.Features)
	}

	if result, err := Spin(game, MODE_MAIN); err != nil {
		return nil, err
	} else {
		conn.Cmd("HMSET", UserHash(player), GameFieldLines(game), lines, GameFieldBet(game), bet.String(), GameFieldFeatures(game), history.Features)
		return result, nil
	}
}

func PlayerFreeSpin(game, player string) (*engine.SpinResult, error) {
	conn, err := redis.DialTimeout("tcp", "127.0.0.1:6379", time.Duration(2)*time.Second)
	if err != nil {
		return nil, &InternalError{err}
	}
	defer conn.Close()

	history, err := RestoreSpinHistory(conn, player, game)
	if err != nil {
		return nil, &InternalError{err}
	}
	if history.Features < 1 {
		return nil, errors.New("spin status is invalid, no more free spins left")
	}

	luaScript := "local v = redis.call('HINCRBY',KEYS[1],KEYS[2],-1) if v>=0 then return v else redis.call('HSET',KEYS[1],KEYS[2],0) return -1 end"
	n, err := conn.Cmd("eval", luaScript, 2, UserHash(player), GameFieldFeatures(game)).Int()
	if err != nil {
		return nil, &InternalError{err}
	}

	if n >= 0 {
		if result, err := Spin(game, MODE_FEATURE); err != nil {
			return nil, err
		} else {
			if result.ScatterWin != nil {
				_, err := conn.Cmd("HINCRBY", UserHash(player), GameFieldFeatures(game), result.ScatterWin.Features).Int()
				if err != nil {
					return nil, &InternalError{err}
				}
			}
			return result, nil
		}
	} else {
		return nil, errors.New("spin status is invalid, no more free spins left")
	}
}
