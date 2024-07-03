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
	FlipClockwise()
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

func (g *Grid) FlipClocwise() {
	tempGrid := g
	g.Left.First = tempGrid.Left.Third
	g.Left.Second = tempGrid.Middle.Third
	g.Left.Third = tempGrid.Right.Third
	g.Middle.First = tempGrid.Left.Second
	g.Middle.Third = tempGrid.Right.Second
	g.Right.First = tempGrid.Left.First
	g.Right.Second = tempGrid.Middle.First
	g.Right.Third = tempGrid.Right.First
}
