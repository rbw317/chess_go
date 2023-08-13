package chess_engine

import "testing"

func TestNewBishop(t *testing.T) {
	bishop := NewBishop(C1, Black)

	if bishop == nil {
		t.Errorf("Error! NewBishop function returned nil for bishop!")
	}

	if bishop.Type != BISHOP {
		t.Errorf("Error! NewBishop function returned wrong Type!")
	}

	if bishop.CurrPosition != C1 {
		t.Errorf("Error! NewBishop function returned wrong CurrPosition!")
	}

	if bishop.Color != Black {
		t.Errorf("Error! NewBishop function returned wrong Color!")
	}

	if bishop.GetType() != BISHOP {
		t.Errorf("Error! NewBishop function returned wrong Type!")
	}

	if bishop.GetCurrPosition() != C1 {
		t.Errorf("Error! NewBishop function returned wrong CurrPosition!")
	}

	if bishop.GetColor() != Black {
		t.Errorf("Error! NewBishop function returned wrong Color!")
	}
}

func TestBishopMove(t *testing.T) {
	bishop := NewBishop(C1, Black)
	bishop.Move(C2)
	if bishop.GetCurrPosition() != C2 {
		t.Errorf("Error! Bishop.Move function returned wrong CurrPosition!")
	}
	if bishop.MoveCount != 1 {
		t.Errorf("Error! Bishop.MoveCount wrong after move!")
	}
}

func TestWhiteBishopCenterMoves(t *testing.T) {
	board := &Board{}
	// Test white square Bishop moving from center square to outer corners
	board.Squares[E4] = Square{true, NewBishop(E4, White)}
	moves := board.Squares[E4].CurrPiece.GetMoves(board)
	if len(moves) != 13 {
		t.Errorf("Error! Bishop.GetMoves function did not return 4 moves for center square on empty board!")
	}
	if !MovesContainMove(Move{E4, D3, false}, moves) &&
		!MovesContainMove(Move{E4, C2, false}, moves) &&
		!MovesContainMove(Move{E4, D1, false}, moves) &&
		!MovesContainMove(Move{E4, F5, false}, moves) &&
		!MovesContainMove(Move{E4, G6, false}, moves) &&
		!MovesContainMove(Move{E4, H7, false}, moves) &&
		!MovesContainMove(Move{E4, F3, false}, moves) &&
		!MovesContainMove(Move{E4, G2, false}, moves) &&
		!MovesContainMove(Move{E4, H1, false}, moves) &&
		!MovesContainMove(Move{E4, D5, false}, moves) &&
		!MovesContainMove(Move{E4, C6, false}, moves) &&
		!MovesContainMove(Move{E4, B7, false}, moves) &&
		!MovesContainMove(Move{E4, A8, false}, moves) {
		t.Errorf("Error! Bishop.GetMoves for white square bishop missing moves!")
	}
}

func TestBlackBishopCenterMoves(t *testing.T) {
	board := &Board{}
	// Test white square Bishop moving from center square to outer corners
	board.Squares[D4] = Square{true, NewBishop(D4, White)}
	moves := board.Squares[D4].CurrPiece.GetMoves(board)
	if len(moves) != 13 {
		t.Errorf("Error! Bishop.GetMoves function did not return 4 moves for center square on empty board!")
	}
	if !MovesContainMove(Move{D4, C3, false}, moves) &&
		!MovesContainMove(Move{D4, B2, false}, moves) &&
		!MovesContainMove(Move{D4, A1, false}, moves) &&
		!MovesContainMove(Move{D4, E5, false}, moves) &&
		!MovesContainMove(Move{D4, F6, false}, moves) &&
		!MovesContainMove(Move{D4, G7, false}, moves) &&
		!MovesContainMove(Move{D4, H8, false}, moves) &&
		!MovesContainMove(Move{D4, E3, false}, moves) &&
		!MovesContainMove(Move{D4, F2, false}, moves) &&
		!MovesContainMove(Move{D4, G1, false}, moves) &&
		!MovesContainMove(Move{D4, C5, false}, moves) &&
		!MovesContainMove(Move{D4, B6, false}, moves) &&
		!MovesContainMove(Move{D4, A7, false}, moves) {
		t.Errorf("Error! Bishop.GetMoves for white square bishop missing moves!")
	}
}

func TestWhiteBishopNoMoves(t *testing.T) {
	board := &Board{}
	// Test white square Bishop moving from center square to outer corners
	board.Squares[D4] = Square{true, NewBishop(D4, White)}
	board.Squares[E5] = Square{true, NewPawn(E5, White)}
	board.Squares[E3] = Square{true, NewPawn(E3, White)}
	board.Squares[C5] = Square{true, NewPawn(C5, White)}
	board.Squares[C3] = Square{true, NewPawn(C3, White)}

	moves := board.Squares[D4].CurrPiece.GetMoves(board)
	if len(moves) != 0 {
		t.Errorf("Error! Bishop.GetMoves function did not return 0 moves for white center square when surrounded by white!")
	}
}

func TestBlackBishopNoMoves(t *testing.T) {
	board := &Board{}
	// Test white square Bishop moving from center square to outer corners
	board.Squares[E4] = Square{true, NewBishop(E4, Black)}
	board.Squares[F5] = Square{true, NewPawn(F5, Black)}
	board.Squares[F3] = Square{true, NewPawn(F3, Black)}
	board.Squares[D5] = Square{true, NewPawn(D5, Black)}
	board.Squares[D3] = Square{true, NewPawn(D3, Black)}

	moves := board.Squares[E4].CurrPiece.GetMoves(board)
	if len(moves) != 0 {
		t.Errorf("Error! Bishop.GetMoves function did not return 0 moves for black center square when surrounded by black!")
	}
}

func TestWhiteBishopAttackMoves(t *testing.T) {
	board := &Board{}
	// Test white square Bishop moving from center square to outer corners
	board.Squares[D4] = Square{true, NewBishop(D4, White)}
	board.Squares[E5] = Square{true, NewPawn(E5, Black)}
	board.Squares[E3] = Square{true, NewPawn(E3, Black)}
	board.Squares[C5] = Square{true, NewPawn(C5, Black)}
	board.Squares[C3] = Square{true, NewPawn(C3, Black)}

	moves := board.Squares[D4].CurrPiece.GetMoves(board)
	if len(moves) != 4 {
		t.Errorf("Error! Bishop.GetMoves function did not return 4 moves for white bishop on center square surrounded by black!")
	}
}

func TestBlackBishopAttackMoves(t *testing.T) {
	board := &Board{}
	// Test white square Bishop moving from center square to outer corners
	board.Squares[E4] = Square{true, NewBishop(E4, Black)}
	board.Squares[F5] = Square{true, NewPawn(F5, Black)}
	board.Squares[F3] = Square{true, NewPawn(F3, Black)}
	board.Squares[D5] = Square{true, NewPawn(D5, Black)}
	board.Squares[D3] = Square{true, NewPawn(D3, Black)}

	moves := board.Squares[E4].CurrPiece.GetMoves(board)
	if len(moves) != 0 {
		t.Errorf("Error! Bishop.GetMoves function did not return 4 moves for black center square when surrounded by white!")
	}
}
