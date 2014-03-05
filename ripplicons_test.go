package ripplicons

import (
	"fmt"
	"testing"
)

/* This is a table of common ripple addresses, and the md5 representation
   as given by singpolyma's implementation:
   http://userscripts.org/scripts/show/350423
*/

var decodeMD5Tests = []struct {
	address  string // ripple address
	expected string // expected md5 hex
}{
	// Zero and One addresses
	{"rrrrrrrrrrrrrrrrrrrrrhoLvTp", "2319ac34f4848755a639fd524038dfd3"},
	{"rrrrrrrrrrrrrrrrrrrrBZbvji", "0f7173e963614f142641d4699cc329d8"},

	// Common gateway issuing addresses
	{"rNPRNzBB92BVpAhhZr4iXDTveCgV5Pofm9", "90f13b28688fedbaeb0c2a87e60b0618"},
	{"rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B", "2d6e14d2929bd0389784c269660f18d7"},
	{"rfYv1TXnwgDDK4WQNbFALykYuEBnrR4pDX", "bbcf11f51a123d9a673f4d49bb0e56fc"},
	{"rGwUWgN5BEg3QGNY3RX2HfYowjUTZdid3E", "c516884bc476752c209dd827fc7df6f9"},
	{"rnuF96W4SZoCJmbHYBFoJZpR8eCaxNvekK", "be57f17c6ceb71834563f955bc848982"},
	{"r3ADD8kXSUKHd6zTCKfnKT3zV9EZHjzp1S", "7584715a37f688d590027bf8b9db07fe"},
	{"rLEsXccBGNR3UPuPu2hUXPjziKC3qKSBun", "1f3fee5f9b82181a4413c30cf7fc8aa5"},
}

func TestDecodeMD5(t *testing.T) {
	for _, r := range decodeMD5Tests {
		md5 := Decode(r.address)
		if fmt.Sprintf("%x", md5) != r.expected {
			t.Errorf(
				"For %s got %x, expected %s",
				r.address, md5, r.expected,
			)
		}
	}
}
