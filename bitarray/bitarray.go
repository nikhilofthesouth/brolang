package bitarray

import (
	"strconv"
)

const intSize = strconv.IntSize

type bitarray []uint

func Create(size uint) (bitarray) {
	return CreateEmpty(size)
}

func CreateEmpty(size uint) (bitarray) {
	sliceSize := indexOf(size)
	if offsetOf(size) != 0 {
		sliceSize += 1
	}
	return make([]uint, sliceSize)
}

func CreateFull(size uint) (bitarray) {
	ba := CreateEmpty(size)
	for i := range ba {
		ba[i] = ^uint(0)
	}
	setSize := ba.arraySize()
	for i := size; i < setSize; i++ {
		ba.Remove(i)
	}
	return ba
}

func (b bitarray) Add(loc uint) {
	index := indexOf(loc)
	b[index] |= bitPosition(loc)
}

func (b bitarray) Remove(loc uint) {
	index := indexOf(loc)
	b[index] &^= bitPosition(loc)
}

func (b bitarray) Contains(loc uint) (bool) {
	index := indexOf(loc)
	return b[index] & bitPosition(loc) > 0
}

func (b bitarray) View() ([]uint) {
	set := make([]uint, 0)
	length := b.arraySize()
	for j := uint(0); j < length; j++ {
		if b.Contains(j) {
			set = append(set, j)
		}
	}
	return set
}

func (b bitarray) arraySize() (uint) {
	return uint(len(b)) * intSize
}

func indexOf(location uint) uint {
	return location / intSize
}

func offsetOf(location uint) uint {
	return location % intSize
}

func bitPosition(location uint) uint {
	return 1 << offsetOf(location)
}
