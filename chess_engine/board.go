package chess_engine

type PieceType uint8

const (
	NoPiece PieceType = iota
	PAWN
	ROOK
	KNIGHT
	BISHOP
	QUEEN
	KING
)

type PieceColor uint8

const (
	White = 0
	Black = 1
)

type BoardPosition uint8

const (
	A1 BoardPosition = iota
	A2
	A3
	A4
	A5
	A6
	A7
	A8
	B1
	B2
	B3
	B4
	B5
	B6
	B7
	B8
	C1
	C2
	C3
	C4
	C5
	C6
	C7
	C8
	D1
	D2
	D3
	D4
	D5
	D6
	D7
	D8
	E1
	E2
	E3
	E4
	E5
	E6
	E7
	E8
	F1
	F2
	F3
	F4
	F5
	F6
	F7
	F8
	G1
	G2
	G3
	G4
	G5
	G6
	G7
	G8
	H1
	H2
	H3
	H4
	H5
	H6
	H7
	H8
	INVALID_BOARD_POSITION
)

type Square struct {
	Occupied  bool
	CurrPiece ChessPiece
}

type Board struct {
	Squares [64]Square
}

func NewBoard() *Board {
	b := Board{}
	b.InitBoard()
	return &b
}

func (board *Board) InitBoard() {
	board.Squares[A1] = Square{Occupied: true, CurrPiece: NewRook(A1, White)}
	board.Squares[B1] = Square{Occupied: true, CurrPiece: NewKnight(B1, White)}
	board.Squares[C1] = Square{Occupied: true, CurrPiece: NewBishop(C1, White)}
	board.Squares[D1] = Square{Occupied: true, CurrPiece: NewQueen(D1, White)}
	board.Squares[E1] = Square{Occupied: true, CurrPiece: NewKing(E1, White)}
	board.Squares[F1] = Square{Occupied: true, CurrPiece: NewBishop(F1, White)}
	board.Squares[G1] = Square{Occupied: true, CurrPiece: NewKnight(G1, White)}
	board.Squares[H1] = Square{Occupied: true, CurrPiece: NewRook(H1, White)}
	board.Squares[A2] = Square{Occupied: true, CurrPiece: NewPawn(A2, White)}
	board.Squares[B2] = Square{Occupied: true, CurrPiece: NewPawn(B2, White)}
	board.Squares[C2] = Square{Occupied: true, CurrPiece: NewPawn(C2, White)}
	board.Squares[D2] = Square{Occupied: true, CurrPiece: NewPawn(D2, White)}
	board.Squares[E2] = Square{Occupied: true, CurrPiece: NewPawn(E2, White)}
	board.Squares[F2] = Square{Occupied: true, CurrPiece: NewPawn(F2, White)}
	board.Squares[G2] = Square{Occupied: true, CurrPiece: NewPawn(G2, White)}
	board.Squares[H2] = Square{Occupied: true, CurrPiece: NewPawn(H2, White)}
	board.Squares[A8] = Square{Occupied: true, CurrPiece: NewRook(A8, Black)}
	board.Squares[B8] = Square{Occupied: true, CurrPiece: NewKnight(B8, Black)}
	board.Squares[C8] = Square{Occupied: true, CurrPiece: NewBishop(C8, Black)}
	board.Squares[D8] = Square{Occupied: true, CurrPiece: NewQueen(D8, Black)}
	board.Squares[E8] = Square{Occupied: true, CurrPiece: NewKing(E8, Black)}
	board.Squares[F8] = Square{Occupied: true, CurrPiece: NewBishop(F8, Black)}
	board.Squares[G8] = Square{Occupied: true, CurrPiece: NewKnight(G8, Black)}
	board.Squares[H8] = Square{Occupied: true, CurrPiece: NewRook(H8, Black)}
	board.Squares[A7] = Square{Occupied: true, CurrPiece: NewPawn(A7, Black)}
	board.Squares[B7] = Square{Occupied: true, CurrPiece: NewPawn(B7, Black)}
	board.Squares[C7] = Square{Occupied: true, CurrPiece: NewPawn(C7, Black)}
	board.Squares[D7] = Square{Occupied: true, CurrPiece: NewPawn(D7, Black)}
	board.Squares[E7] = Square{Occupied: true, CurrPiece: NewPawn(E7, Black)}
	board.Squares[F7] = Square{Occupied: true, CurrPiece: NewPawn(F7, Black)}
	board.Squares[G7] = Square{Occupied: true, CurrPiece: NewPawn(G7, Black)}
	board.Squares[H7] = Square{Occupied: true, CurrPiece: NewPawn(H7, Black)}
}

func (board *Board) CopyBoard() *Board {
	copyBoard := &Board{}
	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			currPos := GetBoardPos(rank, file)
			if board.Squares[currPos].Occupied {
				copyBoard.Squares[currPos] = Square{Occupied: true, CurrPiece: board.Squares[currPos].CurrPiece.Copy()}
			}
		}
	}
	return copyBoard
}

func (board *Board) GetScore(color PieceColor) int {
	score := 0
	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			currPos := GetBoardPos(rank, file)
			if board.Squares[currPos].Occupied && board.Squares[currPos].CurrPiece.GetColor() == color {
				score += board.Squares[currPos].CurrPiece.GetValue()
			}
		}
	}
	return score
}

func (board *Board) GetKingPosition(color PieceColor) BoardPosition {
	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			currPos := GetBoardPos(rank, file)
			if board.Squares[currPos].Occupied &&
				board.Squares[currPos].CurrPiece.GetColor() == color &&
				board.Squares[currPos].CurrPiece.GetType() == KING {
				return currPos
			}
		}
	}
	return INVALID_BOARD_POSITION
}

func (board *Board) GetMoves(color PieceColor) []*Move {
	moves := []*Move{}
	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			currPos := GetBoardPos(rank, file)
			if board.Squares[currPos].Occupied &&
				board.Squares[currPos].CurrPiece.GetColor() == color {
				newMoves := board.Squares[currPos].CurrPiece.GetMoves(board)
				if newMoves != nil {
					moves = append(moves, newMoves...)
				}
			}
		}
	}

	return moves
}

func (board *Board) KingInCheck(color PieceColor) bool {
	retVal := false
	oppositeColor := OppositeColor(color)
	moves := board.GetMoves(oppositeColor)
	kingPos := board.GetKingPosition(color)
	if kingPos != INVALID_BOARD_POSITION {
		for _, move := range moves {
			if move.EndPos == kingPos {
				retVal = true
			}
		}
	}

	return retVal
}

func (board *Board) CheckMate(color PieceColor) bool {
	retVal := true

	if board.KingInCheck(color) { // King must be in check in it's current position
		kingPos := board.GetKingPosition(color)
		moves := board.Squares[kingPos].CurrPiece.GetMoves(board)
		for _, m := range moves {
			tempBoard := board.CopyBoard()
			tempBoard.MovePiece(m)
			if !tempBoard.KingInCheck(color) {
				retVal = false
			}
		}
	} else {
		retVal = false
	}

	return retVal
}

func (board *Board) MovePiece(move *Move) Result {
	res := Result{true, NO_ERROR, ""}
	if move == nil ||
		move.StartPos == INVALID_BOARD_POSITION ||
		move.EndPos == INVALID_BOARD_POSITION {
		res.Result = false
		res.ResCode = INVALID_MOVE
	} else if !board.Squares[move.StartPos].Occupied {
		res.Result = false
		res.ResCode = INVALID_MOVE
	} else {
		if (board.Squares[move.EndPos].Occupied) &&
			(board.Squares[move.StartPos].CurrPiece.GetColor() == board.Squares[move.EndPos].CurrPiece.GetColor()) {
			res.Result = false
			res.ResCode = INVALID_MOVE
		} else {
			board.Squares[move.EndPos].CurrPiece = board.Squares[move.StartPos].CurrPiece
			board.Squares[move.EndPos].Occupied = true
			board.Squares[move.StartPos].CurrPiece = nil
			board.Squares[move.StartPos].Occupied = false
			board.Squares[move.EndPos].CurrPiece.Move(move.EndPos)
			if move.Castle {
				res = board.CastleMove(move)
			}
		}
	}
	return res
}

func (board *Board) CastleMove(move *Move) Result {
	res := Result{true, NO_ERROR, ""}
	if (move.StartPos != E1 && (move.EndPos != G1 && move.EndPos != C1)) &&
		(move.StartPos != E8 && (move.EndPos != G8 && move.EndPos != C8)) {
		res = Result{false, INVALID_MOVE, "Invalid castle move"}
	} else {
		rookStart := H1
		rookEnd := F1
		if move.StartPos == E1 && move.EndPos == C1 {
			rookStart = A1
			rookEnd = D1
		} else if move.StartPos == E8 && move.EndPos == G8 {
			rookStart = H8
			rookEnd = F8
		} else if move.StartPos == E8 && move.EndPos == C8 {
			rookStart = A8
			rookEnd = D8
		}
		board.Squares[rookEnd].CurrPiece = board.Squares[rookStart].CurrPiece
		board.Squares[rookEnd].Occupied = true
		board.Squares[rookStart].CurrPiece = nil
		board.Squares[rookStart].Occupied = false
		board.Squares[rookEnd].CurrPiece.Move(rookEnd)
	}

	return res

}
