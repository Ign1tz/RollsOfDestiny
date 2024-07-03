package Types

import (
	"math/rand"
	"strconv"
)

type Die struct {
	PossibleThrows []int
}

type DieLogic interface {
	Throw()
}

func (d Die) Throw() string {
	res := rand.Intn(6)
	return strconv.Itoa([]int{1, 2, 3, 4, 5, 6}[res])
}
