package main

import (
	"fmt"
	"github.com/nikhilofthesouth/brolang/bitarray"
)

func main() {
	ba := bitarray.CreateFull(70)
	ba.Remove(0)
	ba.Remove(1)
	ba.Remove(2)
	ba.Remove(3)
	ba.Remove(4)
	ba.Remove(64)
	fmt.Println(ba.View())
}
