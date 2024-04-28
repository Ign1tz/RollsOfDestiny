package Components

import (
	"RollsOfDestiny/Components/Types"
	"fmt"
	"strconv"
)

type Playfield struct {
	Host         string
	Guest        string
	HostGrid     Grid
	GuestGrid    Grid
	GameID       string
	ActivePlayer string
}

type PlayfieldLogic interface {
	Clear()
	Results()
	PrettyPrint()
}

func (p *Playfield) Clear() {
	p.HostGrid.Clear()
	p.GuestGrid.Clear()
}

func (p Playfield) Results() Types.Results {
	results := Types.Results{}
	if p.HostGrid.Value() > p.GuestGrid.Value() {
		results.Winner = p.Host
		results.Loser = p.Guest
	} else if p.HostGrid.Value() < p.GuestGrid.Value() {
		results.Winner = p.Guest
		results.Loser = p.Host
	} else {
		results.Winner = p.Host
		results.Loser = p.Guest
		results.Draw = true
	}
	results.WinnerScore = max(p.HostGrid.Value(), p.GuestGrid.Value())
	results.LoserScore = min(p.HostGrid.Value(), p.GuestGrid.Value())
	return results
}
func (p *Playfield) PrettyPrint() {
	fmt.Println(p.Guest)
	fmt.Println("+-+-+-+")
	fmt.Println("|" + strconv.Itoa(p.GuestGrid.Left.Third) + "|" + strconv.Itoa(p.GuestGrid.Middle.Third) + "|" + strconv.Itoa(p.GuestGrid.Right.Third) + "|")
	fmt.Println("+-+-+-+")
	fmt.Println("|" + strconv.Itoa(p.GuestGrid.Left.Second) + "|" + strconv.Itoa(p.GuestGrid.Middle.Second) + "|" + strconv.Itoa(p.GuestGrid.Right.Second) + "|")
	fmt.Println("+-+-+-+")
	fmt.Println("|" + strconv.Itoa(p.GuestGrid.Left.First) + "|" + strconv.Itoa(p.GuestGrid.Middle.First) + "|" + strconv.Itoa(p.GuestGrid.Right.First) + "|")
	fmt.Println("+-+-+-+")
	fmt.Println()
	fmt.Println(p.Host)
	fmt.Println("+-+-+-+")
	fmt.Println("|" + strconv.Itoa(p.HostGrid.Left.First) + "|" + strconv.Itoa(p.HostGrid.Middle.First) + "|" + strconv.Itoa(p.HostGrid.Right.First) + "|")
	fmt.Println("+-+-+-+")
	fmt.Println("|" + strconv.Itoa(p.HostGrid.Left.Second) + "|" + strconv.Itoa(p.HostGrid.Middle.Second) + "|" + strconv.Itoa(p.HostGrid.Right.Second) + "|")
	fmt.Println("+-+-+-+")
	fmt.Println("|" + strconv.Itoa(p.HostGrid.Left.Third) + "|" + strconv.Itoa(p.HostGrid.Middle.Third) + "|" + strconv.Itoa(p.HostGrid.Right.Third) + "|")
	fmt.Println("+-+-+-+")
}
