package Types

import (
	"fmt"
	"strconv"
)

type Game struct {
	HostId       string
	GuestId      string
	HostGrid     string
	GuestGrid    string
	GameID       string
	ActivePlayer string
	LastRoll     string
}

type Playfield struct {
	Host         Player
	Guest        Player
	HostGrid     Grid
	GuestGrid    Grid
	GameID       string
	ActivePlayer Player
	LastRoll     string
}

type PlayfieldLogic interface {
	Clear()
	Results()
	PrettyPrint()
	EnemyPlayer()
}

func (p *Playfield) EnemyPlayer() Player {
	if p.ActivePlayer.UserID != p.Host.UserID {
		return p.Host
	} else {
		return p.Guest
	}
}

func (p *Playfield) Clear() {
	p.HostGrid.Clear()
	p.GuestGrid.Clear()
}

func (p Playfield) Results() Results {
	results := Results{}
	if p.HostGrid.Value() > p.GuestGrid.Value() {
		results.Winner = p.Host.Username
		results.Loser = p.Guest.Username
	} else if p.HostGrid.Value() < p.GuestGrid.Value() {
		results.Winner = p.Guest.Username
		results.Loser = p.Host.Username
	} else {
		results.Winner = p.Host.Username
		results.Loser = p.Guest.Username
		results.Draw = true
	}
	results.WinnerScore = max(p.HostGrid.Value(), p.GuestGrid.Value())
	results.LoserScore = min(p.HostGrid.Value(), p.GuestGrid.Value())
	return results
}
func (p Playfield) PrettyPrint() {
	fmt.Println(p.Guest.Username)
	fmt.Println("+-+-+-+")
	fmt.Println("|" + strconv.Itoa(p.GuestGrid.Left.Third) + "|" + strconv.Itoa(p.GuestGrid.Middle.Third) + "|" + strconv.Itoa(p.GuestGrid.Right.Third) + "|")
	fmt.Println("+-+-+-+")
	fmt.Println("|" + strconv.Itoa(p.GuestGrid.Left.Second) + "|" + strconv.Itoa(p.GuestGrid.Middle.Second) + "|" + strconv.Itoa(p.GuestGrid.Right.Second) + "|")
	fmt.Println("+-+-+-+")
	fmt.Println("|" + strconv.Itoa(p.GuestGrid.Left.First) + "|" + strconv.Itoa(p.GuestGrid.Middle.First) + "|" + strconv.Itoa(p.GuestGrid.Right.First) + "|")
	fmt.Println("+-+-+-+")
	fmt.Println()
	fmt.Println(p.Host.Username)
	fmt.Println("+-+-+-+")
	fmt.Println("|" + strconv.Itoa(p.HostGrid.Left.First) + "|" + strconv.Itoa(p.HostGrid.Middle.First) + "|" + strconv.Itoa(p.HostGrid.Right.First) + "|")
	fmt.Println("+-+-+-+")
	fmt.Println("|" + strconv.Itoa(p.HostGrid.Left.Second) + "|" + strconv.Itoa(p.HostGrid.Middle.Second) + "|" + strconv.Itoa(p.HostGrid.Right.Second) + "|")
	fmt.Println("+-+-+-+")
	fmt.Println("|" + strconv.Itoa(p.HostGrid.Left.Third) + "|" + strconv.Itoa(p.HostGrid.Middle.Third) + "|" + strconv.Itoa(p.HostGrid.Right.Third) + "|")
	fmt.Println("+-+-+-+")
}
