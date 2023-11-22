package main

import (
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/danielost/ecpoint-wrappers/pkg/ecwrapper"
)

// Is k*(d*G) = d*(k*G)?
func main() {
	// Define and set the elliptic curve to P256
	curve := elliptic.P256()

	// Dependency injection: Create an elliptic curve wrapper object using the defined curve
	curveWrapper := ecwrapper.NewECWrapper(curve)

	// Retrieve the base point G of the curve
	G := curveWrapper.GetBasePointG()
	k, d := &big.Int{}, &big.Int{}

	// Generate random scalar values k and d
	k, _ = k.SetString(randToken(32), 16)
	d, _ = d.SetString(randToken(32), 16)

	// Calculate H1 = d * G
	H1 := curveWrapper.ScalarMult(d, G)

	// Calculate H2 = k * H1
	H2 := curveWrapper.ScalarMult(k, H1)

	// Calculate H3 = k * G
	H3 := curveWrapper.ScalarMult(k, G)

	// Calculate H4 = d * H3
	H4 := curveWrapper.ScalarMult(d, H3)

	H2.Print(16)
	H4.Print(16)

	// Check if H2 and H4 are equal (commutative property of scalar multiplication)
	fmt.Println(H2.IsEqual(H4))
}

// Generates a random hexadecimal token of the specified length
func randToken(n int) string {
	bytes := make([]byte, n)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
