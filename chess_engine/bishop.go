package chess_engine

type Bishop struct {
	Type         PieceType
	Color        PieceColor
	CurrPosition BoardPosition
	PrevPosition BoardPosition
	MoveCount    int
}

func NewBishop(pos BoardPosition, color PieceColor) *Bishop {
	bishop := &Bishop{BISHOP, color, pos,
		INVALID_BOARD_POSITION, 0}
	return bishop
}

func (bishop *Bishop) GetMoveCount() int {
	return bishop.MoveCount
}

func (bishop *Bishop) GetMoves(board *Board) []*Move {
	moves := []*Move{}

	moves = bishop.addIterMoves(board, moves, 1, 1)
	moves = bishop.addIterMoves(board, moves, -1, 1)
	moves = bishop.addIterMoves(board, moves, 1, -1)
	moves = bishop.addIterMoves(board, moves, -1, -1)

	return moves
}

func (bishop *Bishop) addIterMoves(board *Board, moves []*Move, rInc int, fInc int) []*Move {
	startRank, startFile := GetRankFile(bishop.CurrPosition)
	startRank += rInc
	startFile += fInc
	for (startRank < 8 && startRank >= 0) && (startFile < 8 && startFile >= 0) {
		currPos := GetBoardPos(startRank, startFile)
		if board.Squares[currPos].Occupied {
			if board.Squares[currPos].CurrPiece.GetColor() != bishop.Color {
				moves = append(moves, NewMove(bishop.CurrPosition, currPos))
				return moves // Hit another piece so no more moves
			} else {
				return moves // We've hit a blocking piece along the diagonal.
			}
		} else {
			moves = append(moves, NewMove(bishop.CurrPosition, currPos))
		}
		startRank += rInc
		startFile += fInc
	}

	return moves
}

func (bishop *Bishop) GetCurrPosition() BoardPosition {
	return bishop.CurrPosition
}

func (bishop *Bishop) GetPrevPosition() BoardPosition {
	return bishop.PrevPosition
}

func (bishop *Bishop) GetColor() PieceColor {
	return bishop.Color
}

func (bishop *Bishop) GetType() PieceType {
	return bishop.Type
}

func (bishop *Bishop) Move(newPos BoardPosition) {
	bishop.CurrPosition = newPos
	bishop.MoveCount++
}
