package rdrand

import "testing"

func TestRdRand(t *testing.T) {
	RdRand{}.Uint64()
	RdSeed{}.Uint64()
	// no panic means good
}
