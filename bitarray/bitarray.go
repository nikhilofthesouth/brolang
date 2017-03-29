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
	sliceSize := size / intSize
	if size % intSize != 0 {
		sliceSize += 1
	}
	return make([]uint, sliceSize)
}


func CreateFull(size uint) (bitarray) {
	ba := CreateEmpty(size)
	for i := range ba {
		ba[i] = ^uint(0)
	}
	setSize := ba.setSize()
	for i := size; i < setSize; i++ {
		ba.Remove(i)
	}
	return ba
}

func GetIndex(location uint) uint {
	return location / intSize
}

func GetOffset(location uint) uint {
	return location % intSize
}

func (b bitarray) Add(loc uint) {
	index := GetIndex(loc)
	b[index] |= 1 << (loc % intSize)
}

func (b bitarray) Remove(loc uint) {
	index := GetIndex(loc)
	b[index] &^= 1 << (loc % intSize)
}

func (b bitarray) Contains(loc uint) (bool) {
	index := GetIndex(loc)
	return b[index] & (1 << (loc % intSize)) > 0
}

func (b bitarray) View() ([]uint) {
	set := make([]uint, 0)
	length := b.setSize()
	for j := uint(0); j < length; j++ {
		if b.Contains(j) {
			set = append(set, j)
		}
	}
	return set
}

func (b bitarray) setSize() (uint) {
	return uint(len(b)) * intSize
}