package main

import (
	"chess_go/chess_engine"
	"chess_go/chess_web_service"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Chess Go is starting!")
	if len(os.Args) >= 2 && os.Args[1] == "-c" {
		console := chess_engine.NewConsoleInterface()
		console.Start()
	} else {
		fmt.Printf("Starting Chess Go Web Service on Port 80.")
		ws := chess_web_service.GetWebServiceInstance(":80")
		ws.Run()
	}

}
