package ecwrapper

import (
	"crypto/elliptic"
	"math/big"
)

// ECWrapper wraps the elliptic.Curve interface to provide more convenient functionality.
type ECWrapper struct {
	curve elliptic.Curve
}

// NewECWrapper creates a new ECWrapper with the specified elliptic curve.
func NewECWrapper(curve elliptic.Curve) *ECWrapper {
	return &ECWrapper{
		curve: curve,
	}
}

// Params returns the parameters of the underlying elliptic curve.
func (ec *ECWrapper) Params() *elliptic.CurveParams {
	return ec.curve.Params()
}

// GetBasePointG returns the base point (generator) of the elliptic curve as an ECPoint.
func (ec *ECWrapper) GetBasePointG() *ECPoint {
	params := ec.Params()
	Gx, Gy := params.Gx, params.Gy
	return NewECPoint(Gx, Gy)
}

// IsOnCurve checks if the given ECPoint lies on the elliptic curve.
func (ec *ECWrapper) IsOnCurve(point *ECPoint) bool {
	params := ec.Params()
	return params.IsOnCurve(point.Params())
}

// Add performs point addition on the elliptic curve and returns the result as an ECPoint.
func (ec *ECWrapper) Add(point1, point2 *ECPoint) *ECPoint {
	params := ec.Params()
	x1, y1 := point1.Params()
	x2, y2 := point2.Params()
	return NewECPoint(params.Add(x1, y1, x2, y2))
}

// Double performs point doubling on the elliptic curve and returns the result as an ECPoint.
func (ec *ECWrapper) Double(point *ECPoint) *ECPoint {
	params := ec.Params()
	return NewECPoint(params.Double(point.Params()))
}

// ScalarMult performs scalar multiplication of an ECPoint with a scalar value and returns the result as an ECPoint.
func (ec *ECWrapper) ScalarMult(k *big.Int, point *ECPoint) *ECPoint {
	params := ec.Params()
	x, y := point.Params()
	return NewECPoint(params.ScalarMult(x, y, k.Bytes()))
}
