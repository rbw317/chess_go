package main

import (
	"chess_go/chess_engine"
	"fmt"
)

func main() {
	board := chess_engine.NewBoard()
	board.InitBoard()
	fmt.Println("Hello World!")
}
