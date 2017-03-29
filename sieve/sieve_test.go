package sieve

import (
	"testing"
	"math"
)

func BenchmarkSieve(b *testing.B) {
	limit := uint(math.MaxUint32)
	for i := 0; i < b.N; i++ {
		Sieve(limit)
	}
}
