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
	return params.IsOnCurve(point.Params())
}

func (ec *ECWrapper) Add(point1, point2 *ECPoint) *ECPoint {
	params := ec.Params()
	x1, y1 := point1.Params()
	x2, y2 := point2.Params()
	return NewECPoint(params.Add(x1, y1, x2, y2))
}

func (ec *ECWrapper) Double(point *ECPoint) *ECPoint {
	params := ec.Params()
	return NewECPoint(params.Double(point.Params()))
}
