package ecwrapper

import (
	"fmt"
	"math/big"
	"strings"
)

// ECPoint represents a point on an elliptic curve with coordinates (x, y).
type ECPoint struct {
	x, y *big.Int
}

// NewECPoint creates a new ECPoint with the given x and y coordinates.
func NewECPoint(x, y *big.Int) *ECPoint {
	return &ECPoint{
		x: x,
		y: y,
	}
}

// Params returns the x and y coordinates of the ECPoint.
func (ecp *ECPoint) Params() (*big.Int, *big.Int) {
	return ecp.x, ecp.y
}

// Print prints the x and y coordinates of the ECPoint in the specified base.
func (ecp *ECPoint) Print(base int) {
	fmt.Printf("X:%s\nY:%s\n", ecp.x.Text(base), ecp.y.Text(base))
}

// IsEqual checks if two ECPoints are equal by comparing their x and y coordinates.
func (ecp *ECPoint) IsEqual(other *ECPoint) bool {
	x1, y1 := ecp.Params()
	x2, y2 := other.Params()
	return x1.Cmp(x2) == 0 && y1.Cmp(y2) == 0
}

// ECPointToString converts an ECPoint to a string representation in the specified base.
func ECPointToString(point *ECPoint, base int) string {
	x, y := point.Params()
	return x.Text(base) + ":" + y.Text(base)
}

// StringToECPoint converts a string representation to an ECPoint in the specified base.
func StringToECPoint(s string, base int) (*ECPoint, error) {
	params := strings.Split(s, ":")
	if len(params) != 2 {
		return nil, fmt.Errorf("wrong number of parameters expected=%d got=%d", 2, len(params))
	}

	x, y := &big.Int{}, &big.Int{}

	// Convert string values to big.Int for x coordinate.
	x, ok := x.SetString(params[0], base)
	if !ok {
		return nil, fmt.Errorf("invalid value for x")
	}

	// Convert string values to big.Int for y coordinate.
	y, ok = y.SetString(params[1], base)
	if !ok {
		return nil, fmt.Errorf("invalid value for y")
	}

	return NewECPoint(x, y), nil
}
