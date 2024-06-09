package Types

import (
	"math/rand"
)

type Die struct {
	PossibleThrows []int
}

type DieLogic interface {
	Throw()
}

func (d Die) Throw() int {
	res := rand.Intn(len(d.PossibleThrows))
	return d.PossibleThrows[res]
}
