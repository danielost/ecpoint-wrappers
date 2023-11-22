package test

import (
	"crypto/elliptic"
	"testing"

	"github.com/danielost/ecpoint-wrappers/pkg/ecwrapper"
)

func TestECPoint_Serialization(t *testing.T) {
	curve := elliptic.P256()
	curveWrapper := ecwrapper.NewECWrapper(curve)
	G := curveWrapper.GetBasePointG()
	serializedG := ecwrapper.ECPointToString(G, 16)
	deserializedG, err := ecwrapper.StringToECPoint(serializedG, 16)

	if err != nil {
		t.Errorf(err.Error())
	} else if !G.IsEqual(deserializedG) {
		t.Errorf("deserialization returned wrong point")
	}
}
