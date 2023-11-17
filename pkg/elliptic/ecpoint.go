package elliptic

import (
	"fmt"
	"math/big"
)

type ECPoint struct {
	X, Y *big.Int
}

func NewECPoint(x, y *big.Int) *ECPoint {
	return &ECPoint{
		X: x,
		Y: y,
	}
}

func (ecp *ECPoint) Print(base int) {
	fmt.Printf("X:%s\nY:%s\n", ecp.X.Text(base), ecp.Y.Text(base))
}
