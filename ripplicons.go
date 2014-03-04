/*
Ripplicons decodes a human-readable ripple address into salient bytes suitable for creating
a Sigil-based identicon (https://github.com/cupcake/sigil). The idea was originally
implemented by singpolyma as a Greasemonkey script (http://userscripts.org/scripts/show/350423).
*/
package ripplicons

import (
	"bytes"
	"math"
)

// toBytes accepts a []int of arbitary base smallBase,
// converts it to base-256, and returns it as a []byte
func toBytes(smallBase int, digits []int) []byte {

	n := make([]int, len(digits)) // least-sig first

	for i := 0; i < len(digits); i++ {
		for j := 0; j < len(n); j++ {
			n[j] *= smallBase
		}

		n[0] += digits[i]

		normalizeBase(256, n)
	}

	// Reverse n and convert to bytes
	out := make([]byte, len(digits)+1)
	for i := 0; i < len(n); i++ {
		out[i+1] = byte(n[len(n)-i-1])
	}

	return out
}

func normalizeBase(b int, digits []int) {
	for i := 0; i < len(digits); i++ {
		if digits[i] > b {
			digits[i+1] = digits[i+1] + int(math.Floor(float64(digits[i])/float64(b)))
			digits[i] = digits[i] % b
		}
	}
}

// Decode accepts a ripple address and returns a []byte suitable for
// use with https://github.com/cupcake/sigil.
func Decode(address string) []byte {
	alphabet := []byte("rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz")
	base58digits := []int{}

	for _, r := range address {
		idx := bytes.IndexByte(alphabet, byte(r))
		base58digits = append(base58digits, idx)
	}

	bytes := toBytes(58, base58digits)

	// Trim off the last 4 (checksum) bytes and extra leading zeros
	// to end up with exactly 21 bytes
	return bytes[len(bytes)-25 : len(bytes)-4]
}
