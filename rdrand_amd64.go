package rdrand

import (
	"fmt"
	"github.com/intel-go/cpuid"
)

//go:noescape
func rdrandU64() uint64

//go:noescape
func rdseedU64() uint64

var rdseedFn = rdseedU64

var (
	ErrAMDBroken = fmt.Errorf("this AMD processor has broken RDRAND and RDSEED")
	ErrNoRdRand  = fmt.Errorf("this processor does not support RDRAND")
	ErrNoRdSeed  = fmt.Errorf("this processor does not support RDSEED")
)

func Check() error {
	// With AMD processor families < 0x17 (before Zen),
	// RDRAND could return non-random data (0) while also reporting a success.
	if cpuid.VendorIdentificatorString == "AuthenticAMD" && cpuid.DisplayFamily < 0x17 {
		return ErrAMDBroken
	}
	if !cpuid.HasFeature(cpuid.RDRND) {
		return ErrNoRdRand
	}
	if !cpuid.HasFeature(cpuid.RDSEED) {
		return ErrNoRdSeed
	}
	return nil
}

func init() {
	if !cpuid.HasFeature(cpuid.RDSEED) {
		rdseedFn = rdrandU64
	}
}

type RdRand struct{}

func (r RdRand) Int63() int64 {
	return int64(rdrandU64())
}

func (r RdRand) Seed(int64) {}

func (r RdRand) Uint64() uint64 {
	return rdrandU64()
}

type RdSeed struct{}

func (r RdSeed) Int63() int64 {
	return int64(rdseedFn())
}

func (r RdSeed) Seed(int64) {}

func (r RdSeed) Uint64() uint64 {
	return rdseedFn()
}
