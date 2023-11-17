package elliptic

import "crypto/elliptic"

type ECWrapper struct {
	curve elliptic.Curve
}

func (ec *ECWrapper) Params() *elliptic.CurveParams {
	return ec.curve.Params()
}
