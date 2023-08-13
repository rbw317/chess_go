package main

import (
	"fmt"
	"main/chess_engine"
)

func main() {
	var board = chess_engine.Board{}
	board.InitBoard()
	fmt.Println("Hello World!")
}
