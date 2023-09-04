package main

import (
	"chess_go/chess_engine"
)

func main() {
	console := chess_engine.NewConsoleInterface()
	console.Start()
}
