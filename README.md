# go-rdrand

This library is to access CPU's `RDRAND` and `RDSEED`
random number generator for some reason like you love
Intel's (or AMD) black box instead of Linux's open source
CSPRNG.

## RDRAND and RDSEED

See [Wikipedia](https://en.wikipedia.org/wiki/RDRAND).

TL;DR: It's an RNG(Random Number Generator) built in
your processor.

RDRAND is broken on AMD processors prior to Zen.

RDSEED is supported after Broadwell (5th Gen and Xeon v4) 
and AMD Zen series. Using RDSEED with this library will 
be automatically fallback to RDRAND if not supported.

Run `Check` before use and fallback to other RNG 
solutions if it's not supported on your platform.

## Usage

This library implements `math/rand.Source64` interface, 
just use it with `math/rand` to get features like read
to a byte slice.

```go
package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/fakeboboliu/go-rdrand"
)

func main() {
	// Check before use
	if err := rdrand.Check(); err != nil {
		if err == rdrand.ErrNoRdSeed {
			fmt.Println("Warning: RdSeed is not supported")
		} else {
			fmt.Println("RdRand cannot use:", err)
			os.Exit(1)
		}
	}

	// Init a rng with RDRAND
	rng := rand.New(rdrand.RdRand{})

	// use it as math/rand
	fmt.Println(rng.Intn(10))
	// or read bytes
	buf := make([]byte, 16)
	rng.Read(buf)
	fmt.Println(buf)

	// Or use RDSEED
	rng = rand.New(rdrand.RdSeed{})
}
```