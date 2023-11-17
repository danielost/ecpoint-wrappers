package ecwrapper

import (
	"fmt"
	"math/big"
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
