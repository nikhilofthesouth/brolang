package sieve

import (
	"testing"
	"fmt"
	"math"
)

func TestSieve(t *testing.T) {
	fmt.Println(Sieve(math.MaxUint16).View())
}
