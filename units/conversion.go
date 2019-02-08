package units

import (
	"fmt"
	"math"
	"math/big"
	"strings"
)

type Unit int

const (
	Wei Unit = iota * 3
	Kwei
	Mwei
	Gwei
	Szabo
	Finney
	Ether
	Kether
	Mether
	Gether
	Tether
)

func Atou(s string) (Unit, error) {
	var err error
	var u Unit
	switch strings.ToLower(s) {
	case "wei":
		u = Wei
	case "kwei":
		u = Kwei
	case "mwei":
		u = Mwei
	case "gwei":
		u = Gwei
	case "szabo":
		u = Szabo
	case "finney":
		u = Finney
	case "ether":
		u = Ether
	case "kether":
		u = Kether
	case "mether":
		u = Mether
	case "gether":
		u = Gether
	case "tether":
		u = Tether
	default:
		err = fmt.Errorf("Invalid unit type '%s'", s)
	}
	return u, err
}

func ConvertInt(from *big.Int, fromUnit, toUnit Unit) *big.Float {
	to := big.NewFloat(math.Pow10(int(fromUnit - toUnit)))
	to.Mul(new(big.Float).SetInt(from), to)
	return to
}

func ConvertFloat(from *big.Float, fromUnit, toUnit Unit) *big.Float {
	to := big.NewFloat(math.Pow10(int(fromUnit - toUnit)))
	to.Mul(from, to)
	return to
}
