package commons

import (
	"errors"
	"fmt"
	"infra/game/message"
)

type Communication struct {
	Peer []chan message.Message
}

func SaturatingSub(x uint, y uint) uint {
	res := x - y
	var val uint
	if res <= x {
		val = 1
	}
	res &= -val
	return res
}

func DeleteElFromSlice(s []uint, i int) ([]uint, error) {
	if i < cap(s) && i >= 0 {
		s[i] = s[len(s)-1]
		return s[:len(s)-1], nil
	} else {
		return s, errors.New(fmt.Sprintf("Out of bounds error, attempted to access index %d in slice %v\n", i, s))
	}

}
