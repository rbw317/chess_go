package chess_engine

type Queen struct {
	Type         PieceType
	Color        PieceColor
	CurrPosition BoardPosition
	PrevPosition BoardPosition
	MoveCount    int
}

func NewQueen(pos BoardPosition, color PieceColor) *Queen {
	queen := &Queen{QUEEN, color, pos,
		INVALID_BOARD_POSITION, 0}
	return queen
}

func (queen *Queen) Copy() ChessPiece {
	copyQueen := &Queen{}
	copyQueen.Type = queen.Type
	copyQueen.Color = queen.Color
	copyQueen.CurrPosition = queen.CurrPosition
	copyQueen.PrevPosition = queen.PrevPosition
	copyQueen.MoveCount = queen.MoveCount
	return copyQueen
}

func (queen *Queen) GetMoveCount() int {
	return queen.MoveCount
}

func (queen *Queen) GetMoves(board *Board) []*Move {
	moves := []*Move{}

	moves = queen.addIterMoves(board, moves, 1, 1)
	moves = queen.addIterMoves(board, moves, -1, 1)
	moves = queen.addIterMoves(board, moves, 1, -1)
	moves = queen.addIterMoves(board, moves, -1, -1)
	moves = queen.addIterMoves(board, moves, -1, 0)
	moves = queen.addIterMoves(board, moves, 1, 0)
	moves = queen.addIterMoves(board, moves, 0, 1)
	moves = queen.addIterMoves(board, moves, 0, -1)

	return moves
}

func (queen *Queen) addIterMoves(board *Board, moves []*Move, rInc int, fInc int) []*Move {
	startRank, startFile := GetRankFile(queen.CurrPosition)
	startRank += rInc
	startFile += fInc
	for (startRank < 8 && startRank >= 0) && (startFile < 8 && startFile >= 0) {
		currPos := GetBoardPos(startRank, startFile)
		if board.Squares[currPos].Occupied {
			if board.Squares[currPos].CurrPiece.GetColor() != queen.Color {
				moves = append(moves, NewMove(queen.CurrPosition, currPos))
				return moves // Hit another piece so no more moves
			} else {
				return moves // We've hit a blocking piece along the diagonal.
			}
		} else {
			moves = append(moves, NewMove(queen.CurrPosition, currPos))
		}
		startRank += rInc
		startFile += fInc
	}

	return moves
}

func (queen *Queen) GetCurrPosition() BoardPosition {
	return queen.CurrPosition
}

func (queen *Queen) GetPrevPosition() BoardPosition {
	return queen.PrevPosition
}

func (queen *Queen) GetColor() PieceColor {
	return queen.Color
}

func (queen *Queen) GetType() PieceType {
	return queen.Type
}

func (queen *Queen) Move(newPos BoardPosition) {
	queen.CurrPosition = newPos
	queen.MoveCount++
}

func (queen *Queen) GetValue() int {
	return 9
}
