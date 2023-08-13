package chess_engine

type King struct {
	Type         PieceType
	Color        PieceColor
	CurrPosition BoardPosition
	PrevPosition BoardPosition
	MoveCount    int
}

func NewKing(pos BoardPosition, color PieceColor) *King {
	king := &King{KING, color, pos,
		INVALID_BOARD_POSITION, 0}
	return king
}

func (king *King) GetMoveCount() int {
	return king.MoveCount
}

func (king *King) GetMoves(board *Board) []*Move {
	moves := []*Move{}
	rank, file := GetRankFile(king.CurrPosition)
	moves = king.CheckSquare(board, moves, rank+1, file+1)
	moves = king.CheckSquare(board, moves, rank+1, file)
	moves = king.CheckSquare(board, moves, rank+1, file-1)
	moves = king.CheckSquare(board, moves, rank, file-1)
	moves = king.CheckSquare(board, moves, rank-1, file-1)
	moves = king.CheckSquare(board, moves, rank-1, file)
	moves = king.CheckSquare(board, moves, rank-1, file+1)
	moves = king.CheckSquare(board, moves, rank, file+1)
	return moves
}

func (king *King) CheckSquare(board *Board, moves []*Move, rank int, file int) []*Move {
	if (rank >= 0 && rank < 8) && (file >= 0 && file < 8) {
		currPos := GetBoardPos(rank, file)
		if board.Squares[currPos].Occupied {
			if board.Squares[currPos].CurrPiece.GetColor() != king.Color {
				moves = append(moves, NewMove(king.CurrPosition, currPos))
			}
		} else {
			moves = append(moves, NewMove(king.CurrPosition, currPos))
		}
	}

	return moves
}

func (king *King) GetCurrPosition() BoardPosition {
	return king.CurrPosition
}

func (king *King) GetPrevPosition() BoardPosition {
	return king.PrevPosition
}

func (king *King) GetColor() PieceColor {
	return king.Color
}

func (king *King) GetType() PieceType {
	return king.Type
}

func (king *King) Move(newPos BoardPosition) {
	king.CurrPosition = newPos
	king.MoveCount++
}
