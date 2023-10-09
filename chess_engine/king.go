package chess_engine

type King struct {
	Type         PieceType
	Color        PieceColor
	CurrPosition BoardPosition
	PrevPosition BoardPosition
	MoveCount    int
	GenCastle    bool
}

func NewKing(pos BoardPosition, color PieceColor) *King {
	king := &King{KING, color, pos,
		INVALID_BOARD_POSITION, 0, true}
	return king
}

func (king *King) Copy() ChessPiece {
	copyKing := &King{}
	copyKing.Type = king.Type
	copyKing.Color = king.Color
	copyKing.CurrPosition = king.CurrPosition
	copyKing.PrevPosition = king.PrevPosition
	copyKing.MoveCount = king.MoveCount
	return copyKing
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
	moves = king.GetCastleMoves(board, moves)
	return moves
}

func (king *King) GetAttackMoves(board *Board) []*Move {
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

func (king *King) GetValue() int {
	return 0
}

func (king *King) GetCastleMoves(board *Board, moves []*Move) []*Move {
	king.GenCastle = false
	if king.CurrPosition == E1 && king.Color == White && king.MoveCount == 0 &&
		!board.Squares[F1].Occupied && !board.Squares[G1].Occupied && !board.WhiteInCheck {
		opMoves := board.GetAttackMoves(Black)

		if !MovesContainEndPos(F1, opMoves) && !MovesContainEndPos(G1, opMoves) {
			moves = append(moves, NewCastleMove(king.CurrPosition, G1))
		}
	} else if king.CurrPosition == E1 && king.Color == White && king.MoveCount == 0 &&
		!board.Squares[D1].Occupied && !board.Squares[C1].Occupied &&
		!board.WhiteInCheck {
		opMoves := board.GetAttackMoves(Black)
		if !MovesContainEndPos(D1, opMoves) && !MovesContainEndPos(C1, opMoves) {
			moves = append(moves, NewCastleMove(king.CurrPosition, C1))
		}
	} else if king.CurrPosition == E8 && king.Color == Black && king.MoveCount == 0 &&
		!board.Squares[F8].Occupied && !board.Squares[G8].Occupied && !board.BlackInCheck {
		opMoves := board.GetAttackMoves(White)

		if !MovesContainEndPos(F8, opMoves) && !MovesContainEndPos(G8, opMoves) {
			moves = append(moves, NewCastleMove(king.CurrPosition, G8))
		}
	} else if king.CurrPosition == E8 && king.Color == Black && king.MoveCount == 0 &&
		!board.Squares[D8].Occupied && !board.Squares[C8].Occupied &&
		!board.BlackInCheck {
		opMoves := board.GetAttackMoves(White)
		if !MovesContainEndPos(D8, opMoves) && !MovesContainEndPos(C8, opMoves) {
			moves = append(moves, NewCastleMove(king.CurrPosition, C8))
		}
	}
	king.GenCastle = true
	return moves
}
