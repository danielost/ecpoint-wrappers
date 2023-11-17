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
