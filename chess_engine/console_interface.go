package chess_engine

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
)

type ConsoleInterface struct {
	CurrGame *Game
}

func NewConsoleInterface() *ConsoleInterface {
	iface := &ConsoleInterface{}
	return iface
}

func (iface *ConsoleInterface) Start() {
	// Main loop
	for {
		iface.DisplayMainMenu()
		iface.ProcessMainUserCmd()
	}
}

func (iface *ConsoleInterface) DisplayMainMenu() {
	fmt.Print("1) Start New Game\n")
	fmt.Print("2) Exit\n")
	fmt.Print("Selection: ")
}

func (iface *ConsoleInterface) ProcessMainUserCmd() {
	var value int
	n, err := fmt.Scanln(&value)
	if n != 1 {
		fmt.Printf("Error reading selection (%s).  Please try again.", err)
		return
	}
	if value == 1 {
		iface.RunGame()
	} else {
		os.Exit(0)
	}
}

func (iface *ConsoleInterface) RunGame() {
	fmt.Print("What color is the user?\n")
	fmt.Print("1) White\n")
	fmt.Print("2) Black\n")
	fmt.Print("Selection: ")
	var value int
	n, err := fmt.Scanln(&value)
	if n != 1 {
		fmt.Printf("Error scanning selection (%s).  Please try again.\n", err)
		return
	}
	if value == 1 {
		iface.CurrGame = NewGame(White)
	} else {
		iface.CurrGame = NewGame(Black)
	}

	for {
		if iface.CurrGame.Status == GAME_STATUS_USER_MOVE {
			var startStr, endStr string
			fmt.Print("Please enter your move in the format 'Start Square End Square (i.e. A1 A3): ")
			fmt.Scanln(&startStr, &endStr)
			move := GetMoveFromUserString(startStr, endStr)

			if move == nil {
				fmt.Print("Invalid move.  Please try again.")
			} else {
				if iface.CurrGame.Board.Squares[move.StartPos].CurrPiece != nil &&
					iface.CurrGame.Board.Squares[move.StartPos].CurrPiece.GetType() == PAWN {
					_, endFile := GetRankFile(move.EndPos)
					if (iface.CurrGame.UserColor == White && endFile == 7) ||
						(iface.CurrGame.UserColor == Black && endFile == 0) {
						fmt.Print("Pawn Promotion! Please enter 1)Rook 2)Knight 3)Bishop 4)Queen: ")
						var promotion string
						fmt.Scanln(&promotion)
						promotion = strings.TrimSpace(promotion)
						if promotion == "1" {
							iface.CurrGame.Board.Squares[move.StartPos].CurrPiece = NewRook(move.StartPos, iface.CurrGame.UserColor)
						} else if promotion == "2" {
							iface.CurrGame.Board.Squares[move.StartPos].CurrPiece = NewKnight(move.StartPos, iface.CurrGame.UserColor)
						} else if promotion == "3" {
							iface.CurrGame.Board.Squares[move.StartPos].CurrPiece = NewBishop(move.StartPos, iface.CurrGame.UserColor)
						} else if promotion == "4" {
							iface.CurrGame.Board.Squares[move.StartPos].CurrPiece = NewQueen(move.StartPos, iface.CurrGame.UserColor)
						}
					}

				}
				res := iface.CurrGame.UserMove(move)
				if res.Result {
					iface.PrintBoard()
				} else {
					fmt.Printf("Inavlid move.  %s", res.ResStr)
				}
			}
		} else if iface.CurrGame.Status == GAME_STATUS_ENGINE_MOVE {
			res, move := iface.CurrGame.EngineMove()
			if res.Result {
				iface.PrintEngineMove(move)
				iface.PrintBoard()
			} else {
				fmt.Printf("Error getting engine move: %s", res.ResStr)
			}

		} else if iface.CurrGame.Status == GAME_STATUS_ENGINE_CHECKMATE {
			fmt.Print("You have checkmated me!  You Win! Game Over\n")
			return
		} else {
			fmt.Print("I have checkmated you!  You Lose! Game Over\n")
			return
		}
	}

}

func (iface *ConsoleInterface) PrintEngineMove(move *Move) {
	fmt.Printf("Engine Move: %s to %s\n", GetPositionString(move.StartPos), GetPositionString(move.EndPos))
}

func (iface *ConsoleInterface) PrintBoard() {
	fmt.Print("\n")
	iface.PrintRow(7)
	iface.PrintRow(6)
	iface.PrintRow(5)
	iface.PrintRow(4)
	iface.PrintRow(3)
	iface.PrintRow(2)
	iface.PrintRow(1)
	iface.PrintRow(0)
	pen := color.New(color.FgHiYellow).Add(color.BgHiBlack).Add(color.Bold)
	pen.Printf("\tA\tB\tC\tD\tE\tF\tG\tH\n")

}

func (iface *ConsoleInterface) PrintRow(rank int) {
	pen := color.New(color.FgHiYellow).Add(color.BgHiBlack).Add(color.Bold)

	pen.Printf("%d\t", rank+1)
	iface.PrintBoardPos(rank, 0)
	iface.PrintBoardPos(rank, 1)
	iface.PrintBoardPos(rank, 2)
	iface.PrintBoardPos(rank, 3)
	iface.PrintBoardPos(rank, 4)
	iface.PrintBoardPos(rank, 5)
	iface.PrintBoardPos(rank, 6)
	iface.PrintBoardPos(rank, 7)
	fmt.Print("\n")
}

func (iface *ConsoleInterface) PrintBoardPos(rank int, file int) {
	pos := GetBoardPos(rank, file)
	str := " "
	fgColor := color.FgHiWhite

	if iface.CurrGame.Board.Squares[pos].Occupied {
		if iface.CurrGame.Board.Squares[pos].CurrPiece.GetType() == BISHOP {
			str = "B"
		} else if iface.CurrGame.Board.Squares[pos].CurrPiece.GetType() == KING {
			str = "K"
		} else if iface.CurrGame.Board.Squares[pos].CurrPiece.GetType() == QUEEN {
			str = "Q"
		} else if iface.CurrGame.Board.Squares[pos].CurrPiece.GetType() == ROOK {
			str = "R"
		} else if iface.CurrGame.Board.Squares[pos].CurrPiece.GetType() == KNIGHT {
			str = "N"
		} else if iface.CurrGame.Board.Squares[pos].CurrPiece.GetType() == PAWN {
			str = "P"
		}

		if iface.CurrGame.Board.Squares[pos].CurrPiece.GetColor() == Black {
			fgColor = color.FgHiMagenta
		}
	}

	bgColor := color.BgHiGreen
	if (rank+1)%2 == 0 {
		if (file+1)%2 != 0 {
			bgColor = color.BgHiCyan
		}
	} else if (file+1)%2 == 0 {
		bgColor = color.BgHiCyan
	}

	pen := color.New(fgColor).Add(bgColor).Add(color.Bold)
	pen.Printf("%s\t", str)
}
