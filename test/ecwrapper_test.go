package test

import (
	"crypto/elliptic"
	"fmt"
	"math/big"
	"testing"

	"github.com/danielost/ecpoint-wrappers/pkg/ecwrapper"
)

func TestECWrapper_IsOnCurve(t *testing.T) {
	curve := elliptic.P256()
	ecw := ecwrapper.NewECWrapper(curve)
	pointOnCurve := ecw.GetBasePointG()
	Gx, Gy := pointOnCurve.Params()
	x, y := &big.Int{}, &big.Int{}
	x.SetString(Gx.String(), 10)
	y.SetString(Gy.String(), 10)
	pointNotOnCurve := ecwrapper.NewECPoint(x.Add(x, big.NewInt(1)), y)

	tests := []struct {
		point *ecwrapper.ECPoint
		want  bool
	}{
		{point: pointOnCurve, want: true},
		{point: pointNotOnCurve, want: false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("IsOnCurve() for %v\n", ecwrapper.ECPointToString(tt.point, 16)), func(t *testing.T) {
			if res := ecw.IsOnCurve(tt.point); res != tt.want {
				t.Errorf("IsOnCurve() want=%t got=%t", !tt.want, res)
			}
		})
	}
}

func TestECWrapper_Double(t *testing.T) {
	curve := elliptic.P256()
	ecw := ecwrapper.NewECWrapper(curve)
	x, y := &big.Int{}, &big.Int{}
	x.SetString("56515219790691171413109057904011688695424810155802929973526481321309856242040", 10)
	y.SetString("3377031843712258259223711451491452598088675519751548567112458094635497583569", 10)
	wantPoint := ecwrapper.NewECPoint(x, y)

	tests := []struct {
		point *ecwrapper.ECPoint
		want  *ecwrapper.ECPoint
	}{
		{point: ecw.GetBasePointG(), want: wantPoint},
	}

	for _, tt := range tests {
		t.Run("Double()", func(t *testing.T) {
			if res := ecw.Double(tt.point); !res.IsEqual(wantPoint) {
				t.Errorf("Double() want=%v got=%v", tt.want, res)
			}
		})
	}
}

func TestECWrapper_ScalarMult(t *testing.T) {
	curve := elliptic.P256()
	ecw := ecwrapper.NewECWrapper(curve)
	n := &big.Int{}
	n.SetString("6903864864128780597790631180119990471252548033682705256603083227360459446930", 10)

	x, y := &big.Int{}, &big.Int{}
	x.SetString("93183920130258995916676756816899397813723383078082481087319266791963799753870", 10)
	y.SetString("82318360383842367894787038688154348486923097067905335792545204245318567426864", 10)
	wantPoint := ecwrapper.NewECPoint(x, y)

	tests := []struct {
		point *ecwrapper.ECPoint
		n     *big.Int
		want  *ecwrapper.ECPoint
	}{
		{point: ecw.GetBasePointG(), n: n, want: wantPoint},
	}

	for _, tt := range tests {
		t.Run("ScalarMult()", func(t *testing.T) {
			if res := ecw.ScalarMult(tt.n, tt.point); !res.IsEqual(wantPoint) {
				t.Errorf("ScalarMult() want=%v got=%v", tt.want, res)
			}
		})
	}
}
