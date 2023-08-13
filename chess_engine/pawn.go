package chess_engine

type Pawn struct {
	Type         PieceType
	Color        PieceColor
	CurrPosition BoardPosition
	PrevPosition BoardPosition
	MoveCount    int
}

func NewPawn(pos BoardPosition, color PieceColor) *Pawn {
	pawn := &Pawn{PAWN, color, pos,
		INVALID_BOARD_POSITION, 0}
	return pawn
}

func (pawn *Pawn) GetMoveCount() int {
	return pawn.MoveCount
}

func (pawn *Pawn) GetMoves(board *Board) []*Move {
	moves := []*Move{}
	rank, file := GetRankFile(pawn.CurrPosition)
	increment := 1
	newPos := INVALID_BOARD_POSITION
	if pawn.Color == Black {
		increment = -1
	}
	newRank := rank + increment
	if newRank >= 0 && newRank < 8 {
		newPos = GetBoardPos(newRank, file)
		if !board.Squares[newPos].Occupied {
			moves = append(moves, NewMove(pawn.CurrPosition, newPos))
		}
	}

	// Add the two square first move
	if pawn.MoveCount == 0 && !board.Squares[newPos].Occupied {
		newRank = newRank + increment
		if newRank >= 0 && newRank < 8 {
			newPos = GetBoardPos(newRank, file)
			if !board.Squares[newPos].Occupied {
				moves = append(moves, NewMove(pawn.CurrPosition, newPos))
			}
		}
	}

	// Check for attack move to left
	newRank = rank + increment
	newFile := file - 1
	if (newRank >= 0 && newRank < 8) && (newFile >= 0 && newFile < 8) {
		newPos = GetBoardPos(newRank, newFile)
		if board.Squares[newPos].Occupied &&
			board.Squares[newPos].CurrPiece.GetColor() != pawn.Color {
			moves = append(moves, NewMove(pawn.CurrPosition, newPos))
		}
	}

	// Check for attack move to right
	newRank = rank + increment
	newFile = file + 1
	if (newRank >= 0 && newRank < 8) && (newFile >= 0 && newFile < 8) {
		newPos = GetBoardPos(newRank, newFile)
		if board.Squares[newPos].Occupied &&
			board.Squares[newPos].CurrPiece.GetColor() != pawn.Color {
			moves = append(moves, NewMove(pawn.CurrPosition, newPos))
		}
	}

	// Check for en Passant with white
	if pawn.Color == White && rank == 4 {
		if file-1 > 0 { // Check for enPassant to the left
			startPos := GetBoardPos(6, file-1)
			attackPos := GetBoardPos(rank, file-1)
			movePos := GetBoardPos(5, file-1)
			if board.Squares[attackPos].Occupied &&
				board.Squares[attackPos].CurrPiece.GetType() == PAWN &&
				board.Squares[attackPos].CurrPiece.GetColor() == Black &&
				board.Squares[attackPos].CurrPiece.GetMoveCount() == 1 &&
				board.Squares[attackPos].CurrPiece.GetPrevPosition() == startPos {
				moves = append(moves, NewEnPassantMove(pawn.CurrPosition, movePos))
			}
		}
		if file+1 <= 7 { // Check for enPassant to the right
			startPos := GetBoardPos(6, file+1)
			attackPos := GetBoardPos(rank, file+1)
			movePos := GetBoardPos(5, file+1)
			if board.Squares[attackPos].Occupied &&
				board.Squares[attackPos].CurrPiece.GetType() == PAWN &&
				board.Squares[attackPos].CurrPiece.GetColor() == Black &&
				board.Squares[attackPos].CurrPiece.GetMoveCount() == 1 &&
				board.Squares[attackPos].CurrPiece.GetPrevPosition() == startPos {
				moves = append(moves, NewEnPassantMove(pawn.CurrPosition, movePos))
			}
		}
	}

	if pawn.Color == Black && rank == 3 {
		if file-1 > 0 { // Check for enPassant to the left
			startPos := GetBoardPos(1, file-1)
			attackPos := GetBoardPos(rank, file-1)
			movePos := GetBoardPos(2, file-1)
			if board.Squares[attackPos].Occupied &&
				board.Squares[attackPos].CurrPiece.GetType() == PAWN &&
				board.Squares[attackPos].CurrPiece.GetColor() == White &&
				board.Squares[attackPos].CurrPiece.GetMoveCount() == 1 &&
				board.Squares[attackPos].CurrPiece.GetPrevPosition() == startPos {
				moves = append(moves, NewEnPassantMove(pawn.CurrPosition, movePos))
			}
		}
		if file+1 <= 7 { // Check for enPassant to the right
			startPos := GetBoardPos(1, file+1)
			attackPos := GetBoardPos(rank, file+1)
			movePos := GetBoardPos(2, file+1)
			if board.Squares[attackPos].Occupied &&
				board.Squares[attackPos].CurrPiece.GetType() == PAWN &&
				board.Squares[attackPos].CurrPiece.GetColor() == White &&
				board.Squares[attackPos].CurrPiece.GetMoveCount() == 1 &&
				board.Squares[attackPos].CurrPiece.GetPrevPosition() == startPos {
				moves = append(moves, NewEnPassantMove(pawn.CurrPosition, movePos))
			}
		}
	}
	return moves
}

func (pawn *Pawn) GetCurrPosition() BoardPosition {
	return pawn.CurrPosition
}

func (pawn *Pawn) GetPrevPosition() BoardPosition {
	return pawn.PrevPosition
}

func (pawn *Pawn) GetColor() PieceColor {
	return pawn.Color
}

func (pawn *Pawn) GetType() PieceType {
	return pawn.Type
}

func (pawn *Pawn) Move(newPos BoardPosition) {
	pawn.CurrPosition = newPos
	pawn.MoveCount++
}
