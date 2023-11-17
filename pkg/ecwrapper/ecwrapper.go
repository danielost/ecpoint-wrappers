package ecwrapper

import "crypto/elliptic"

type ECWrapper struct {
	curve elliptic.Curve
}

func NewECWrapper(curve elliptic.Curve) *ECWrapper {
	return &ECWrapper{
		curve: curve,
	}
}

func (ec *ECWrapper) Params() *elliptic.CurveParams {
	return ec.curve.Params()
}

func (ec *ECWrapper) GetBasePointG() *ECPoint {
	params := ec.Params()
	Gx, Gy := params.Gx, params.Gy
	return NewECPoint(Gx, Gy)
}

func (ec *ECWrapper) IsOnCurve(point *ECPoint) bool {
	params := ec.Params()
	return params.IsOnCurve(point.X, point.Y)
}

func (ec *ECWrapper) Add(point1, point2 *ECPoint) *ECPoint {
	params := ec.Params()
	return NewECPoint(params.Add(point1.X, point1.Y, point2.X, point2.Y))
}
