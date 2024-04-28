package GameLogic

import (
	"math/rand"
)

type Die struct {
	possibleThrows []int
}

type DieLogic interface {
	Throw()
}

func (d Die) Throw() int {
	res := rand.Intn(len(d.possibleThrows))
	return d.possibleThrows[res]
}
