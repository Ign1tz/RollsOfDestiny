package GameLogic

import (
	"RollsOfDestiny/Components"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GameLoop() {
	reader := bufio.NewReader(os.Stdin)
	playfield := Components.Playfield{Host: "Moritz", Guest: "Tester", HostGrid: Components.Grid{},
		GuestGrid: Components.Grid{}, GameID: "12345", ActivePlayer: "Moritz"}
	for !playfield.HostGrid.IsFull() && !playfield.GuestGrid.IsFull() {
		die := Die{[]int{1, 2, 3, 4, 5, 6}}
		roll := die.Throw()
		fmt.Println(playfield.ActivePlayer + " rolled a " + strconv.Itoa(roll))
		fmt.Println("Pick a column")
		if playfield.ActivePlayer == playfield.Host {
			playfield.PrettyPrint()
			col, _ := reader.ReadString('\n')
			col = strings.Replace(col, "\r\n", "", -1)
			pickColumn(col, &playfield.HostGrid, &playfield.GuestGrid, roll)
			playfield.ActivePlayer = playfield.Guest
		} else {
			playfield.PrettyPrint()
			col, _ := reader.ReadString('\n')
			col = strings.Replace(col, "\r\n", "", -1)
			pickColumn(col, &playfield.GuestGrid, &playfield.HostGrid, roll)
			playfield.ActivePlayer = playfield.Host
		}
	}
	result := playfield.Results()

	fmt.Println("The winner is " + result.Winner + "!")
	fmt.Println("With " + strconv.Itoa(result.WinnerScore) + " against " + strconv.Itoa(result.LoserScore))
}

func pickColumn(col string, grid *Components.Grid, enemyGrid *Components.Grid, roll int) {
	if col == "1" {
		grid.Left.Add(roll)
		enemyGrid.Left.Remove(roll)
	} else if col == "2" {
		grid.Middle.Add(roll)
		enemyGrid.Middle.Remove(roll)
	} else if col == "3" {
		grid.Right.Add(roll)
		enemyGrid.Right.Remove(roll)
	}
}

func main() {
	GameLoop()
}
