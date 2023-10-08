package main

import (
	"chess_go/chess_engine"
	"chess_go/chess_web_service"
	"os"
)

func main() {
	if len(os.Args) >= 2 && os.Args[1] == "-c" {
		console := chess_engine.NewConsoleInterface()
		console.Start()
	} else {
		ws := chess_web_service.GetWebServiceInstance(":8080")
		ws.Run()
	}

}
