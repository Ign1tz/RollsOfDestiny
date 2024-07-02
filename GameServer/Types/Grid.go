package Types

import "fmt"

type Grid struct {
	Left   Column
	Middle Column
	Right  Column
	GridId string
}

type GridLogic interface {
	Clear()
	Value()
	IsFull()
	PrettyPrint()
}

func (g *Grid) Clear() {
	g.Left.Clear()
	g.Middle.Clear()
	g.Right.Clear()
}

func (g Grid) Value() int {
	return g.Left.Value() + g.Middle.Value() + g.Right.Value()
}

func (g Grid) IsFull() bool {
	return g.Left.IsFull() && g.Middle.IsFull() && g.Right.IsFull()
}

func (g Grid) PrettyPrint() {
	fmt.Println(g.Left.First, g.Middle.First, g.Right.First)
	fmt.Println(g.Left.Second, g.Middle.Second, g.Right.Second)
	fmt.Println(g.Left.Third, g.Middle.Third, g.Right.Third)
	fmt.Println()
}
