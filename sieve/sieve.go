package sieve

import (
	"github.com/nikhilofthesouth/brolang/bitarray"
	"math"
)

func Sieve(max uint) (bitarray.BitArray) {
	ba := bitarray.CreateFull(max+1)
	ba.Remove(0)
	ba.Remove(1)

	sqrt := math.Sqrt(float64(max))
	// Use math.Ceil not math.Floor in case of rounding error
	limit := uint(math.Ceil(sqrt))
	for i := uint(2); i <= limit && ba.Contains(i); i++ {
		for j := i*i; j <= max; j += i {
			ba.Remove(j)
		}
	}

	return ba
}
