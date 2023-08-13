package chess_engine

type Rook struct {
	Type         PieceType
	Color        PieceColor
	CurrPosition BoardPosition
	PrevPosition BoardPosition
	MoveCount    int
}

func NewRook(pos BoardPosition, color PieceColor) *Rook {
	rook := &Rook{ROOK, color, pos,
		INVALID_BOARD_POSITION, 0}
	return rook
}

func (rook *Rook) GetMoveCount() int {
	return rook.MoveCount
}

func (rook *Rook) GetMoves(board *Board) []*Move {
	moves := []*Move{}

	moves = rook.addIterMoves(board, moves, -1, 0)
	moves = rook.addIterMoves(board, moves, 1, 0)
	moves = rook.addIterMoves(board, moves, 0, 1)
	moves = rook.addIterMoves(board, moves, 0, -1)

	return moves
}

func (rook *Rook) addIterMoves(board *Board, moves []*Move, rInc int, fInc int) []*Move {
	startRank, startFile := GetRankFile(rook.CurrPosition)
	startRank += rInc
	startFile += fInc
	for (startRank < 8 && startRank >= 0) && (startFile < 8 && startFile >= 0) {
		currPos := GetBoardPos(startRank, startFile)
		if board.Squares[currPos].Occupied {
			if board.Squares[currPos].CurrPiece.GetColor() != rook.Color {
				moves = append(moves, NewMove(rook.CurrPosition, currPos))
				return moves // Hit another piece so no more moves
			} else {
				return moves // We've hit a blocking piece along the diagonal.
			}
		} else {
			moves = append(moves, NewMove(rook.CurrPosition, currPos))
		}
		startRank += rInc
		startFile += fInc
	}

	return moves
}

func (rook *Rook) GetCurrPosition() BoardPosition {
	return rook.CurrPosition
}

func (rook *Rook) GetPrevPosition() BoardPosition {
	return rook.PrevPosition
}

func (rook *Rook) GetColor() PieceColor {
	return rook.Color
}

func (rook *Rook) GetType() PieceType {
	return rook.Type
}

func (rook *Rook) Move(newPos BoardPosition) {
	rook.CurrPosition = newPos
	rook.MoveCount++
}
