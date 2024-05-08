package Types

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
	return g.Left.IsFull && g.Middle.IsFull && g.Right.IsFull
}
