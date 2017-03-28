package bitarray

import (
	"testing"
)

func TestSetBit(t *testing.T) {
	ba := CreateFull(1000)
	toRemove := []uint{10, 100, 123, 500, 600, 750, 999, 1000}
	toAdd := []uint{200, 500, 1000}
	for _, v := range toRemove {
		ba.Remove(v)
	}
	for _, v := range toAdd {
		ba.Add(v)
	}

	for _, v := range []uint{10, 100, 123, 600, 750, 999} {
		if ba.Contains(v) {
			t.Fatal()
		}
	}
}
