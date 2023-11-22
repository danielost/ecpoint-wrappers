# Elliptic Curve points wrappers

Wrappers for operations on elliptic curves in Golang. Task №7 for the Cryptography for Developers course.

## How to use

Execute the following command first:
```bash
github.com/danielost/ecpoint-wrappers
```
Then simply add the following import to your Go code:
```bash
import (
  // Importing the ecwrappers package
  "github.com/danielost/ecpoint-wrappers/pkg/ecwrapper"
)
```
### API description
- `NewECPoint(x, y *big.Int) *ECPoint` - creates a new ECPoint with the given x and y coordinates.
- `ECPointToString(point *ECPoint, base int) string` - converts an ECPoint to a string representation in the specified base.
- `StringToECPoint(s string, base int) (*ECPoint, error)` - converts a string representation to an ECPoint in the specified base.
- `NewECWrapper(curve elliptic.Curve) *ECWrapper` - creates a new ECWrapper with the specified elliptic curve.
- `(ec *ECWrapper) GetBasePointG() *ECPoint` - returns the base point (generator) of the elliptic curve as an ECPoint.
- `(ec *ECWrapper) IsOnCurve(point *ECPoint) bool` - checks if the given ECPoint lies on the elliptic curve.
- `(ec *ECWrapper) Add(point1, point2 *ECPoint) *ECPoint` - performs point addition on the elliptic curve and returns the result as an ECPoint.
- `(ec *ECWrapper) Double(point *ECPoint) *ECPoint` - performs point doubling on the elliptic curve and returns the result as an ECPoint.
- `(ec *ECWrapper) ScalarMult(k *big.Int, point *ECPoint) *ECPoint` - performs scalar multiplication of an ECPoint with a scalar value and returns the result as an ECPoint.

## Examples
As an example, we can check `k*(d*G) = d*(k*G)` equality. Any curve can be used for this, but for the demonstration `elliptic.P256()` is used.

Full block of code checking the above equality:
```go
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

```

## Testing
To run the tests, clone the repository and execute the following command:
```bash
go test ./...
```
or (for the detailed view):
```bash
go test -v ./...
```
