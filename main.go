package main

import (
	"fmt"

	"github.com/dkNinjaCoder/chessengine/board"
)

func main() {
	b := board.NewBoard()
	_ = b.StartingPosition()
	fmt.Println(b.GetFen())
}
