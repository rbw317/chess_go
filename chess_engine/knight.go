package chess_engine

type Knight struct {
	Type         PieceType
	Color        PieceColor
	CurrPosition BoardPosition
	PrevPosition BoardPosition
	MoveCount    int
}

func NewKnight(pos BoardPosition, color PieceColor) *Knight {
	knight := &Knight{KNIGHT, color, pos,
		INVALID_BOARD_POSITION, 0}
	return knight
}

func (knight *Knight) Copy() ChessPiece {
	copyKnight := &Knight{}
	copyKnight.Type = knight.Type
	copyKnight.Color = knight.Color
	copyKnight.CurrPosition = knight.CurrPosition
	copyKnight.PrevPosition = knight.PrevPosition
	copyKnight.MoveCount = knight.MoveCount
	return copyKnight
}

func (knight *Knight) GetMoveCount() int {
	return knight.MoveCount
}

func (knight *Knight) GetMoves(board *Board) []*Move {
	moves := []*Move{}
	rank, file := GetRankFile(knight.CurrPosition)
	moves = knight.CheckSquare(board, moves, rank+2, file+1)
	moves = knight.CheckSquare(board, moves, rank+1, file+2)
	moves = knight.CheckSquare(board, moves, rank+2, file-1)
	moves = knight.CheckSquare(board, moves, rank+1, file-2)
	moves = knight.CheckSquare(board, moves, rank-2, file+1)
	moves = knight.CheckSquare(board, moves, rank-1, file+2)
	moves = knight.CheckSquare(board, moves, rank-2, file-1)
	moves = knight.CheckSquare(board, moves, rank-1, file-2)
	return moves
}

func (knight *Knight) CheckSquare(board *Board, moves []*Move, rank int, file int) []*Move {
	if (rank >= 0 && rank < 8) && (file >= 0 && file < 8) {
		currPos := GetBoardPos(rank, file)
		if board.Squares[currPos].Occupied {
			if board.Squares[currPos].CurrPiece.GetColor() != knight.Color {
				moves = append(moves, NewMove(knight.CurrPosition, currPos))
			}
		} else {
			moves = append(moves, NewMove(knight.CurrPosition, currPos))
		}
	}

	return moves
}

func (knight *Knight) GetCurrPosition() BoardPosition {
	return knight.CurrPosition
}

func (knight *Knight) GetPrevPosition() BoardPosition {
	return knight.PrevPosition
}

func (knight *Knight) GetColor() PieceColor {
	return knight.Color
}

func (knight *Knight) GetType() PieceType {
	return knight.Type
}

func (knight *Knight) Move(newPos BoardPosition) {
	knight.CurrPosition = newPos
	knight.MoveCount++
}

func (knight *Knight) GetValue() int {
	return 3
}
