package bitarray

import (
	"strconv"
)

const IntSize = strconv.IntSize

type bitarray []uint

func Create(size uint) (bitarray) {
	return CreateEmpty(size)
}

func CreateEmpty(size uint) (bitarray) {
	sliceSize := size / IntSize
	if size % IntSize != 0 {
		sliceSize += 1
	}
	return make([]uint, sliceSize)
}


func CreateFull(size uint) (bitarray) {
	ba := CreateEmpty(size)
	for i := range ba {
		ba[i] ^= 0
	}
	return ba
}

func GetIndex(location uint) uint {
	return location / IntSize
}

func GetOffset(location uint) uint {
	return location % IntSize
}

func (b bitarray) Add(loc uint) {
	index := GetIndex(loc)
	b[index] |= 1 << (loc % IntSize)
}

func (b bitarray) Remove(loc uint) {
	index := GetIndex(loc)
	b[index] &^= 1 << (loc % IntSize)
}

func (b bitarray) Contains(loc uint) (bool) {
	index := GetIndex(loc)
	return b[index] & (1 << (loc % IntSize)) > 0
}

func (b bitarray) View() (set []uint) {
	set = make([]uint, IntSize)
	length := uint(len(b)) * IntSize
	for j := uint(2); j < length; j++ { // Don't allocate exabyte sized slices plz
		if b.Contains(j) {
			set = append(set, j)
		}
	}
	return
}