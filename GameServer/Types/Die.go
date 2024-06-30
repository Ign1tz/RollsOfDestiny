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
	res := rand.Intn(len(d.PossibleThrows))
	return strconv.Itoa(d.PossibleThrows[res])
}
