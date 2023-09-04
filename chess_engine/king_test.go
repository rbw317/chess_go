package chess_engine

import "testing"

func TestNewKing(t *testing.T) {
	king := NewKing(C1, Black)

	if king == nil {
		t.Errorf("Error! NewKing function returned nil!")
	}

	if king.Type != KING {
		t.Errorf("Error! NewKing function returned wrong Type!")
	}

	if king.CurrPosition != C1 {
		t.Errorf("Error! NewKing function returned wrong CurrPosition!")
	}

	if king.Color != Black {
		t.Errorf("Error! NewKing function returned wrong Color!")
	}

	if king.GetType() != KING {
		t.Errorf("Error! NewKing function returned wrong Type!")
	}

	if king.GetCurrPosition() != C1 {
		t.Errorf("Error! NewKing function returned wrong CurrPosition!")
	}

	if king.GetColor() != Black {
		t.Errorf("Error! NewKing function returned wrong Color!")
	}
}

func TestKingMove(t *testing.T) {
	king := NewKing(C1, Black)
	king.Move(C2)
	if king.GetCurrPosition() != C2 {
		t.Errorf("Error! King.Move function returned wrong CurrPosition!")
	}
	if king.MoveCount != 1 {
		t.Errorf("Error! King.MoveCount wrong after move!")
	}
}

func TestKingAllMoves(t *testing.T) {
	board := &Board{}
	// Test white square Bishop moving from center square to outer corners
	board.Squares[E4] = Square{true, NewKing(E4, Black)}

	moves := board.Squares[E4].CurrPiece.GetMoves(board)
	if len(moves) != 8 {
		t.Errorf("Error! King.GetMoves function did not return 8 moves for king on open center square!")
	}
	if !MovesContainMove(Move{E4, E5, false, false, false}, moves) &&
		!MovesContainMove(Move{E4, E3, false, false, false}, moves) &&
		!MovesContainMove(Move{E4, D4, false, false, false}, moves) &&
		!MovesContainMove(Move{E4, F4, false, false, false}, moves) &&
		!MovesContainMove(Move{E4, F5, false, false, false}, moves) &&
		!MovesContainMove(Move{E4, F3, false, false, false}, moves) &&
		!MovesContainMove(Move{E4, D5, false, false, false}, moves) &&
		!MovesContainMove(Move{E4, D3, false, false, false}, moves) {
		t.Errorf("Error! King.GetMoves for king on open center square missing squares!")
	}
}

func TestKingEdgeMoves(t *testing.T) {
	board := &Board{}
	// Test white square Bishop moving from center square to outer corners
	board.Squares[A4] = Square{true, NewKing(A4, Black)}

	moves := board.Squares[A4].CurrPiece.GetMoves(board)
	if len(moves) != 5 {
		t.Errorf("Error! King.GetMoves function did not return 4 moves for king on open edge square!")
	}
	if !MovesContainMove(Move{A4, A5, false, false, false}, moves) &&
		!MovesContainMove(Move{A4, B5, false, false, false}, moves) &&
		!MovesContainMove(Move{A4, B4, false, false, false}, moves) &&
		!MovesContainMove(Move{A4, B3, false, false, false}, moves) &&
		!MovesContainMove(Move{A4, A3, false, false, false}, moves) {
		t.Errorf("Error! King.GetMoves for king on open edge square missing squares!")
	}
}

func TestKingNoMoves(t *testing.T) {
	board := &Board{}
	// Test white square Bishop moving from center square to outer corners
	board.Squares[E4] = Square{true, NewKing(E4, White)}
	board.Squares[E5] = Square{true, NewPawn(E5, White)}
	board.Squares[F5] = Square{true, NewPawn(F5, White)}
	board.Squares[F4] = Square{true, NewPawn(F4, White)}
	board.Squares[F3] = Square{true, NewPawn(F3, White)}
	board.Squares[E3] = Square{true, NewPawn(E3, White)}
	board.Squares[D3] = Square{true, NewPawn(D3, White)}
	board.Squares[D4] = Square{true, NewPawn(D4, White)}
	board.Squares[D5] = Square{true, NewPawn(D5, White)}
	moves := board.Squares[E4].CurrPiece.GetMoves(board)
	if len(moves) != 0 {
		t.Errorf("Error! King.GetMoves function did not return 0 moves for king on totally blocked center square!")
	}
}

func TestKingAttackMoves(t *testing.T) {
	board := &Board{}
	// Test white square Bishop moving from center square to outer corners
	board.Squares[E4] = Square{true, NewKing(E4, White)}
	board.Squares[E5] = Square{true, NewPawn(E5, Black)}
	board.Squares[F5] = Square{true, NewPawn(F5, Black)}
	board.Squares[F4] = Square{true, NewPawn(F4, Black)}
	board.Squares[F3] = Square{true, NewPawn(F3, Black)}
	board.Squares[E3] = Square{true, NewPawn(E3, Black)}
	board.Squares[D3] = Square{true, NewPawn(D3, Black)}
	board.Squares[D4] = Square{true, NewPawn(D4, Black)}
	board.Squares[D5] = Square{true, NewPawn(D5, Black)}
	moves := board.Squares[E4].CurrPiece.GetMoves(board)
	if len(moves) != 8 {
		t.Errorf("Error! King.GetMoves function did not return 8 moves for king on center square surrounded by opponent pieces!")
	}
}

func TestKingCastle(t *testing.T) {
	board := NewBoard()
	board.Squares[F1].Occupied = false
	board.Squares[G1].Occupied = false
	moves := board.Squares[E1].CurrPiece.GetMoves(board)
	if len(moves) != 2 {
		t.Errorf("Error! King.GetMoves function did not return 2 moves for king in castling position!")
	} else if !moves[0].Castle && !moves[1].Castle {
		t.Errorf("Error! King.GetMoves did not return a move flagged Castle when in Castle position!")
	}
	if !MovesContainMove(Move{E1, F1, false, false, false}, moves) &&
		!MovesContainMove(Move{E1, H1, false, false, false}, moves) {
		t.Errorf("Error! King.GetMoves for king on E1 castling square missing squares!")
	}

	board = NewBoard()
	board.Squares[B1].Occupied = false
	board.Squares[C1].Occupied = false
	board.Squares[D1].Occupied = false
	moves = board.Squares[E1].CurrPiece.GetMoves(board)
	if len(moves) != 2 {
		t.Errorf("Error! King.GetMoves function did not return 2 moves for king in castling position!")
	} else if !moves[0].Castle && !moves[1].Castle {
		t.Errorf("Error! King.GetMoves did not return a move flagged Castle when in Castle position!")
	}
	if !MovesContainMove(Move{E1, D1, false, false, false}, moves) &&
		!MovesContainMove(Move{E1, C1, false, false, false}, moves) {
		t.Errorf("Error! King.GetMoves for king on E1 castling square missing squares!")
	}

	board = NewBoard()
	board.Squares[B8].Occupied = false
	board.Squares[C8].Occupied = false
	board.Squares[D8].Occupied = false
	moves = board.Squares[E8].CurrPiece.GetMoves(board)
	if len(moves) != 2 {
		t.Errorf("Error! King.GetMoves function did not return 2 moves for king in castling position!")
	} else if !moves[0].Castle && !moves[1].Castle {
		t.Errorf("Error! King.GetMoves did not return a move flagged Castle when in Castle position!")
	}
	if !MovesContainMove(Move{E8, D8, false, false, false}, moves) &&
		!MovesContainMove(Move{E8, C8, false, false, false}, moves) {
		t.Errorf("Error! King.GetMoves for king on E1 castling square missing squares!")
	}

	board = NewBoard()
	board.Squares[F8].Occupied = false
	board.Squares[G8].Occupied = false
	moves = board.Squares[E8].CurrPiece.GetMoves(board)
	if len(moves) != 2 {
		t.Errorf("Error! King.GetMoves function did not return 2 moves for king in castling position!")
	} else if !moves[0].Castle && !moves[1].Castle {
		t.Errorf("Error! King.GetMoves did not return a move flagged Castle when in Castle position!")
	}
	if !MovesContainMove(Move{E8, F8, false, false, false}, moves) &&
		!MovesContainMove(Move{E8, H8, false, false, false}, moves) {
		t.Errorf("Error! King.GetMoves for king on E1 castling square missing squares!")
	}
}
