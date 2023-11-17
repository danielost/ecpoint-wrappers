package ecwrapper

import (
	"fmt"
	"math/big"
	"strings"
)

type ECPoint struct {
	x, y *big.Int
}

func NewECPoint(x, y *big.Int) *ECPoint {
	return &ECPoint{
		x: x,
		y: y,
	}
}

func (ecp *ECPoint) Params() (*big.Int, *big.Int) {
	return ecp.x, ecp.y
}

func (ecp *ECPoint) Print(base int) {
	fmt.Printf("X:%s\nY:%s\n", ecp.x.Text(base), ecp.y.Text(base))
}

func ECPointToString(point *ECPoint, base int) string {
	x, y := point.Params()
	return x.Text(base) + ":" + y.Text(base)
}

func StringToECPoint(s string, base int) (*ECPoint, error) {
	params := strings.Split(s, ":")
	if len(params) != 2 {
		return nil, fmt.Errorf("wrong number of parameters expected=%d got=%d", 2, len(params))
	}

	var x, y *big.Int

	x, ok := x.SetString(params[0], base)
	if !ok {
		return nil, fmt.Errorf("invalid value for x")
	}

	y, ok = y.SetString(params[1], base)
	if !ok {
		return nil, fmt.Errorf("invalid value for y")
	}

	return NewECPoint(x, y), nil
}
